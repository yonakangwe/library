package mail

import (
	"library/config"
	"library/package/log"
	"bytes"
	"net"
	"net/smtp"

	"text/template"
)

var cfg *config.Config

func init() {
	var err error
	cfg, err = config.New()
	if err != nil {
		log.Fatalf("error initializing config: %v", err)
	}
}

type MailTemplate struct {
	SystemName   string
	SystemHost   string
	Title        string
	Url          string
	Domain       string
	Uri          string
	Description  string
	ReceiverName string
}

// MailRequest struct
type MailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewMailRequest(to []string, subject string) *MailRequest {
	return &MailRequest{
		from:    cfg.Mail.UserName,
		to:      to,
		subject: subject,
	}
}

func (r *MailRequest) SendEmail() (bool, error) {

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	secret := cfg.Mail.Secret
	// Join all recipients for the To header
	toHeader := "To: " + joinAddresses(r.to) + "\n"
	subject := "Subject: " + r.subject + "\n"
	msg := []byte(toHeader + subject + mime + "\n" + r.body)

	servername := cfg.Mail.ServerName
	host, _, err := net.SplitHostPort(servername)
	if err != nil {
		log.Errorf("error occurred splitting host port %v", err)
	}

	auth := smtp.PlainAuth("", r.from, secret, host)

	if err := smtp.SendMail(servername, auth, r.from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil

}

// joinAddresses joins a slice of email addresses into a comma-separated string
func joinAddresses(addresses []string) string {
	if len(addresses) == 0 {
		return ""
	}
	result := addresses[0]
	for i := 1; i < len(addresses); i++ {
		result += ", " + addresses[i]
	}
	return result
}

func (r *MailRequest) ParseTemplate(templateFileName string, data interface{}) *MailRequest {

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Errorf("error occurred parsing file %v", err)
		return nil
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Errorf("error occurred executing file %v", err)
		return nil
	}
	r.body = buf.String()
	return r
}
