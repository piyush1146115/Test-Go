package SOLID

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/smtp"
)

//The Single Responsibility Principle (SRP) states that each software module should have one and only one reason to change.
//As soon as describing the responsibility of some code struct requires the usage of the word "and", it already breaks the Single Responsibility Principle.

//EmailService does not hold the responsibility for storing and sending emails. It delegates them to the structs below.
//Its responsibility is to delegate requests for processing emails to the underlying services.

//There is a difference between holding and delegating responsibility.
//If an adaptation of a particular code can remove the whole purpose of responsibility, we talk about holding.
//If that responsibility still exists even after removing a specific code, then we talk about delegation.

type EmailGorm struct {
	gorm.Model
	From    string
	To      string
	Subject string
	Message string
}

type EmailRepository interface {
	Save(from string, to string, subject string, message string) error
}

type EmailDBRepository struct {
	db *gorm.DB
}

func (r *EmailDBRepository) Save(from string, to string, subject string, message string) error {
	email := EmailGorm{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := r.db.Create(&email).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewEmailRepository(db *gorm.DB) EmailRepository {
	return &EmailDBRepository{
		db: db,
	}
}

type EmailSender interface {
	Send(from string, to string, subject string, message string) error
}

type EmailSMTPSender struct {
	smtpHost     string
	smtpPassword string
	smtpPort     int
}

func (s *EmailSMTPSender) Send(from string, to string, subject string, message string) error {
	auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)

	server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)

	err := smtp.SendMail(server, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewEmailSender(smtpHost string, smtpPassword string, smtpPort int) EmailSender {
	return &EmailSMTPSender{
		smtpHost:     smtpHost,
		smtpPassword: smtpPassword,
		smtpPort:     smtpPort,
	}
}

type EmailService struct {
	repository EmailRepository
	sender     EmailSender
}

func (s *EmailService) Send(from string, to string, subject string, message string) error {
	err := s.repository.Save(from, to, subject, message)
	if err != nil {
		return err
	}

	return s.sender.Send(from, to, subject, message)
}
