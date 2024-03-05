package API_GMAIL_Professor

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	professor "sigaa.ufpe/packages/data/professor_data"
	//"gopkg.in/gomail.v2" // Para enviar através de SMTP
)

type Professor_Email struct{
	Professor professor.Professor
	Subject string
	Message string
}

func Send_Professor_email(Professor_Email Professor_Email) {
	// Crie um e-mail
	e := email.NewEmail()
	e.From = string(os.Getenv("EMAIL_SENDER_NAME")) + " <" + string(os.Getenv("EMAIL_SENDER_ADDRESS"))+">" //"Remetente <remetente@example.com>"
	e.To = []string{string(Professor_Email.Professor.Email)}
	e.Subject = Professor_Email.Subject
	e.Text = []byte(string(Professor_Email.Message))

	fmt.Println(e.From, e.To, e.Subject, e.Text)

	// Configuração para enviar e-mail via SMTP
	host := "smtp.gmail.com"
	port := 587
	username := string(os.Getenv("EMAIL_SENDER_ADDRESS"))
	password := string(os.Getenv("GMAIL_KEY"))

	// Enviar e-mail via SMTP
	err := e.Send(fmt.Sprintf("%s:%d", host, port), smtp.PlainAuth("", username, password, host))
	if err != nil {
		log.Fatal("Erro ao enviar e-mail:", err)
	}

	fmt.Println("E-mail enviado com sucesso!")
}