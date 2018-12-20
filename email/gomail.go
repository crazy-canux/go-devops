package main

import (
	"gopkg.in/gomail.v2"
    "crypto/tls"
)



func main() {
    to := []string{
        "wcheng@sonicwall.com",
        "hyma@sonicwall.com",
    }

    addresses := make([]string, len(to))

    m := gomail.NewMessage()
    m.SetHeader("From", "from@sonicwall.com")
    for i := range addresses {
        addresses[i] = m.FormatAddress(to[i], "")
    }
    m.SetHeader("To", addresses...)
    m.SetAddressHeader("Cc", "cc@sonicwall.com", "")
    m.SetHeader("Subject", "test subject")
    m.SetBody("text/plain", "Hello baby")
    //m.Attach()

    d := gomail.Dialer{
        Host: "mail.sonicwall.com",
        Port: 25,
        Username: "",
        Password: "",
    }
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}