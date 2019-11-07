package email

import (
	"gopkg.in/gomail.v2"
    "crypto/tls"
    "log"
)

// Send PlainText.
func SendText(host string, port int, username, password string, to []string, from, cc, subject, body string) error {
    addresses := make([]string, len(to))

    m := gomail.NewMessage()
    m.SetHeader("From", from)
    for i := range addresses {
        addresses[i] = m.FormatAddress(to[i], "")
    }
    m.SetHeader("To", addresses...)
    m.SetAddressHeader("Cc", cc, "")
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)
    //m.Attach()

    d := gomail.Dialer{
        Host: host,
        Port: port,
        Username: username,
        Password: password,
    }
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(m); err != nil {
        log.Println("Send email failed.")
        return err
    }
    return nil
}