package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
)

// Service adalah struct untuk layanan Cloudinary
type Service struct {
	cld *cloudinary.Cloudinary
}

// NewService membuat instance baru dari Service Cloudinary
func NewService(config *configs.CloudinaryConfig) (*Service, error) {
	// Inisialisasi Cloudinary dari konfigurasi
	cld, err := cloudinary.NewFromParams(config.CloudName, config.APIKey, config.APISecret)
	if err != nil {
		return nil, err
	}

	// Mengatur konfigurasi untuk menggunakan HTTPS
	cld.Config.URL.Secure = true

	return &Service{cld: cld}, nil
}

// UploadFile mengunggah file ke Cloudinary
// Menerima file multipart, dan folder tujuan
// Mengembalikan URL gambar dan public ID jika berhasil
func (s *Service) UploadFile(fileHeader *multipart.FileHeader, folder string) (string, string, error) {
	// Buka file
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	// Set folder default jika tidak disediakan
	if folder == "" {
		folder = "darahconnect"
	}

	ctx := context.Background()

	// Upload file ke Cloudinary
	result, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:         "darahconnect/" + folder,
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(true),
		Transformation: "q_auto:low,",
	})

	if err != nil {
		return "", "", err
	}

	return result.SecureURL, result.PublicID, nil
}

// DeleteFile menghapus file dari Cloudinary berdasarkan public ID
func (s *Service) DeleteFile(publicID string) error {
	ctx := context.Background()

	// Hapus file dari Cloudinary
	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})

	return err
}
