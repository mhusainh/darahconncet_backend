package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"	
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/mhusainh/DarahConnect/DarahConnectAPI/configs"
)

// Mailer adalah struct untuk mengelola pengiriman email menggunakan Mailjet
type Mailer struct {
	client  *mailjet.Client
	smtpCfg configs.SMTPConfig
}

// EmailData berisi data yang diperlukan untuk mengirim email
type EmailData struct {
	To       string
	Subject  string
	Data     interface{}
	Template string // Nama template yang akan digunakan
}
func NewMailer(smtpCfg *configs.SMTPConfig) (*Mailer, error) {
	// Inisialisasi client Mailjet
	mailjetClient := mailjet.NewMailjetClient(smtpCfg.APIKey, smtpCfg.SecretKey)
	return &Mailer{
		client:  mailjetClient,
		smtpCfg: *smtpCfg,
	}, nil
}

func (m *Mailer) SendEmail(templatePath string, emailData EmailData) error {
	// Dapatkan direktori kerja saat ini untuk logging
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Gagal mendapatkan direktori kerja: %v", err)
		cwd = "."
	} else {
		log.Printf("Direktori kerja saat ini: %s", cwd)
	}

	// Coba beberapa path alternatif untuk template
	templatePaths := []string{
		templatePath,
		fmt.Sprintf("templates/email/%s", emailData.Template),
		fmt.Sprintf("../templates/email/%s", emailData.Template),
		fmt.Sprintf("%s/templates/email/%s", cwd, emailData.Template),
	}

	var tmpl *template.Template
	var templateFound bool = false

	// Coba setiap path sampai template ditemukan
	for _, path := range templatePaths {
		log.Printf("Mencoba membaca template dari: %s", path)
		if _, errPath := os.Stat(path); errPath == nil {
			var errParse error
			tmpl, errParse = template.ParseFiles(path)
			if errParse == nil {
				log.Printf("Template ditemukan dan berhasil dibaca dari: %s", path)
				templateContent, errRead := os.ReadFile(path)
				if errRead != nil {
					log.Printf("Peringatan: Gagal membaca konten template untuk logging: %v", errRead)
				} else {
					log.Printf("Ukuran template: %d bytes", len(templateContent))
				}
				templateFound = true
				break
			} else {
				log.Printf("Gagal mem-parse template %s: %v", path, errParse)
			}
		} else {
			log.Printf("Template tidak ditemukan di %s: %v", path, errPath)
		}
	}

	if !templateFound {
		return fmt.Errorf("template tidak ditemukan di path manapun")
	}

	// Eksekusi template dengan data
	var emailBody bytes.Buffer
	if errExecute := tmpl.Execute(&emailBody, emailData.Data); errExecute != nil {
		return fmt.Errorf("gagal mengeksekusi template: %v", errExecute)
	}

	// Siapkan pesan Mailjet menggunakan struktur MessagesV31 dan InfoMessagesV31
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.smtpCfg.Username,
				Name:  "Darah Connect",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: emailData.To,
				},
			},
			Subject:  emailData.Subject,
			TextPart: "Silakan gunakan email client yang mendukung HTML untuk melihat pesan ini.",
			HTMLPart: emailBody.String(),
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	// Kirim email menggunakan SendMailV31
	log.Printf("Mengirim email ke: %s dengan subjek: %s", emailData.To, emailData.Subject)
	res, err := m.client.SendMailV31(&messages)
	if err != nil {
		log.Printf("Error mengirim email: %v", err)
		return fmt.Errorf("gagal mengirim email: %v", err)
	}
	fmt.Printf("Data: %+v\n", res)

	return nil
}