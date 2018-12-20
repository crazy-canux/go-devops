package main

import (
	"net/smtp"
	"log"
	//"fmt"
	//"bytes"
	"bytes"
)

func main() {
	//auth := smtp.PlainAuth("", "", "", "mail.sonicwall.com:25")
	//err := smtp.SendMail("mail.sonicwall.com:25", auth, "test@sonicwall.com", []string{"wcheng@sonicwall.com"}, []byte("test"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	client, err := smtp.Dial("mail.company.com:25")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Mail("taf@company.com"); err != nil {
		log.Fatal(err)
	}
	if err := client.Rcpt("wcheng@company.com"); err != nil {
		log.Fatal(err)
	}

	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	buf := bytes.NewBufferString("this is body")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
	//_, err = fmt.Fprintf(wc, "this is email body")
	//if err != nil {
	//	log.Fatal(err)
	//}
}