package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/markbates/goth"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/entity"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/http/dto"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/internal/repository"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cache"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/cloudinary"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/mailer"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/timezone"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/pkg/token"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]entity.User, int64, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
	Login(ctx context.Context, email, password string) (string, bool, error)
	Register(ctx context.Context, req dto.UserRegisterRequest) error
	CheckGoogleOAuth(ctx context.Context, email string, user *goth.User) (*entity.User, bool, error)
	Update(ctx context.Context, req dto.UpdateUserRequest) error
	Delete(ctx context.Context, user *entity.User) error
	ResendTokenVerifyEmail(ctx context.Context, email string) (string, error)
	VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error
	RequestResetPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	WalletAddress(ctx context.Context, user *entity.User, req dto.WalletAddressRequest) error
}

type userService struct {
	userRepository    repository.UserRepository
	tokenUseCase      token.TokenUseCase
	cacheable         cache.Cacheable
	mailer            *mailer.Mailer
	cfg               *configs.Config
	cloudinaryService *cloudinary.Service
}

func NewUserService(
	userRepository repository.UserRepository,
	tokenUseCase token.TokenUseCase,
	cacheable cache.Cacheable,
	cfg *configs.Config,
	mailer *mailer.Mailer,
	cloudinaryService *cloudinary.Service,
) UserService {
	return &userService{userRepository, tokenUseCase, cacheable, mailer, cfg, cloudinaryService}
}

func (s *userService) Login(ctx context.Context, email string, password string) (string, bool, error) {
	metamask := true
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", false, errors.New("Email atau password salah")
	}
	if user.WalletAddress == "" {
		metamask = false
	}

	if bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); bcryptErr != nil {
		return "", false, errors.New("Email atau password salah")
	}

	if !user.IsVerified {
		TokenExpiresAt := user.TokenExpiresAt.Format(time.RFC3339)
		return TokenExpiresAt, false, nil
	}

	expiredTime := time.Now().Add(time.Hour * 12)

	claims := token.JwtCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
		Metamask: metamask,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Darah Connect",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", false, errors.New("ada kesalahan di server")
	}

	return token, true, nil
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	exist, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err == nil && exist != nil {
		return errors.New("Email sudah digunakan")
	}

	user := new(entity.User)
	user.Email = req.Email
	user.Name = req.Name
	user.Gender = req.Gender
	user.Phone = req.Phone
	user.BloodType = req.BloodType
	user.BirthDate = req.BirthDate
	user.Address = req.Address
	user.Role = "User"
	user.VerifyEmailToken = utils.RandomString(16)
	user.IsVerified = false
	user.TokenExpiresAt = time.Now().In(timezone.JakartaLocation).Add(1 * time.Hour)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ada kesalahan di server")
	}

	user.Password = string(hashedPassword)

	// Prepare email data
	emailData := mailer.EmailData{
		To:       user.Email,
		Subject:  "Darah Connect : Verifikasi Email!",
		Template: "verify-email.html",
		Data: struct {
			Token string
		}{
			Token: user.VerifyEmailToken,
		},
	}

	// Gunakan path relatif terhadap root project
	templatePath := "./templates/email/verify-email.html"
	if Senderr := s.mailer.SendEmail(templatePath, emailData); Senderr != nil {
		return errors.New("gagal mengirim email")
	}

	// Create user in database
	if err = s.userRepository.Create(ctx, user); err != nil {
		return errors.New("gagal membuat user")
	}

	return nil
}

func (s *userService) GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]entity.User, int64, error) {
	users, total, err := s.userRepository.GetAll(ctx, req)
	if err != nil {
		return nil, 0, errors.New("gagal mendapatkan data user")
	}
	return users, total, nil
}

func (s *userService) GetById(ctx context.Context, id int64) (*entity.User, error) {
	return s.userRepository.GetById(ctx, id)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) error {
	var oldPublicId string
	var newPublicId string

	user, err := s.userRepository.GetById(ctx, req.Id)
	if err != nil {
		return errors.New("User tidak ditemukan")
	}

	if req.Image != nil {
		// Simpan publicId lama sebelum mengubahnya
		oldPublicId = user.PublicId

		UrlFile, publicId, err := s.cloudinaryService.UploadFile(req.Image, "Users")
		if err != nil {
			return errors.New("Gagal mengupload gambar")

		}
		newPublicId = publicId
		user.UrlFile = UrlFile
		user.PublicId = publicId
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("ada kesalahan di server")
		}
		user.Password = string(hashedPassword)
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.BloodType != "" {
		user.BloodType = req.BloodType
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			return err
		}
		user.BirthDate = birthDate
	}

	if err := s.userRepository.Update(ctx, user); err != nil {
		// Jika database update gagal dan ada gambar baru yang diunggah, hapus gambar baru
		if req.Image != nil {
			if err := s.cloudinaryService.DeleteFile(newPublicId); err != nil {
				return errors.New("Gagal menghapus gambar baru")
			}
		}
		return errors.New("Gagal mengupdate user")
	}

	// Jika berhasil dan ada gambar lama, hapus gambar lama
	if req.Image != nil && oldPublicId != "" {
		if err := s.cloudinaryService.DeleteFile(oldPublicId); err != nil {
			return errors.New("Gagal menghapus gambar lama")
		}
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, user *entity.User) error {
	return s.userRepository.Delete(ctx, user)
}

func (s *userService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	user, err := s.userRepository.GetByResetPasswordToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token reset password salah")
	}
	if req.Password == "" {
		return errors.New("Password tidak boleh kosong")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.ResetPasswordToken = "expired"
	return s.userRepository.Update(ctx, user)
}

func (s *userService) RequestResetPassword(ctx context.Context, email string) error {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("Email tersebut tidak ditemukan")
	}

	expiredTime := time.Now().Add(10 * time.Minute)

	claims := token.ResetPasswordClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			Issuer:    "Reset Password",
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return errors.New("ada kesalahan di server")
	}

	user.ResetPasswordToken = token
	err = s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	// Prepare email data
	emailData := mailer.EmailData{
		To:       user.Email,
		Subject:  "Darah Connect : Reset Password!",
		Template: "reset-password.html",
		Data: struct {
			Token string
		}{
			Token: user.ResetPasswordToken,
		},
	}

	// Send reset password email
	if err := s.mailer.SendEmail("./templates/email/reset-password.html", emailData); err != nil {
		log.Printf("Gagal mengirim email reset password: %v", err)
		return errors.New("gagal mengirim reset password")
	}

	return nil
}

func (s *userService) ResendTokenVerifyEmail(ctx context.Context, email string) (string, error) {

	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Email tidak ditemukan")
	}
	
	if user.IsVerified {
		return "", errors.New("Email sudah diverifikasi")
	}

	if time.Now().In(timezone.JakartaLocation).Before(user.TokenExpiresAt) {
		return user.TokenExpiresAt.Format(time.RFC3339), nil
	}

	user.VerifyEmailToken = utils.RandomString(16)
	user.TokenExpiresAt = time.Now().In(timezone.JakartaLocation).Add(1 * time.Hour)

	// Prepare email data
	emailData := mailer.EmailData{
		To:       user.Email,
		Subject:  "Darah Connect : Verifikasi Email!",
		Template: "verify-email.html",
		Data: struct {
			Token string
		}{
			Token: user.VerifyEmailToken,
		},
	}

	// Gunakan path relatif terhadap root project
	templatePath := "./templates/email/verify-email.html"
	if Senderr := s.mailer.SendEmail(templatePath, emailData); Senderr != nil {
		return "", errors.New("gagal mengirim email")

	}

	if err := s.userRepository.Update(ctx, user); err != nil {
		return "", errors.New("gagal mengupdate user")
	}
	TokenExpiresAt := user.TokenExpiresAt.Format(time.RFC3339)
	return TokenExpiresAt, nil
}

func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) error {
	user, err := s.userRepository.GetByVerifyEmailToken(ctx, req.Token)
	if err != nil {
		return errors.New("Token verifikasi email salah")
	}
	user.IsVerified = true
	return s.userRepository.Update(ctx, user)
}

func (s *userService) CheckGoogleOAuth(ctx context.Context, email string, user *goth.User) (*entity.User, bool, error) {
	existingUser, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		newUser := new(entity.User)
		newUser.Email = user.Email
		newUser.Name = user.Name
		newUser.Role = "User"
		newUser.IsVerified = true
		if err = s.userRepository.Create(ctx, newUser); err != nil {
			log.Printf("Error creating new user: %v", err)
			return nil, false, err
		}
		return newUser, true, nil
	}
	if existingUser.Name == "" {
		return existingUser, true, nil
	}
	return existingUser, false, nil
}

func (s *userService) WalletAddress(ctx context.Context, user *entity.User, req dto.WalletAddressRequest) error {
	user.WalletAddress = req.WalletAddress
	if err := s.userRepository.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

// func (s *userService) GetAll(ctx context.Context) (result []entity.User, err error) {
// 	keyFindAll := "github.com/mhusainh/DarahConnect/DarahConnectAPI-api:users:find-all"
// 	data := s.cacheable.Get(keyFindAll)
// 	if data == "" {
// 		result, err = s.userRepository.GetAll(ctx)
// 		if err != nil {
// 			return nil, err
// 		}

// 		marshalledData, err := json.Marshal(result)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = s.cacheable.Set(keyFindAll, marshalledData, 5*time.Minute)
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		err = json.Unmarshal([]byte(data), &result)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return result, nil
// }
