package main

import (
	"net/smtp"
	"net/mail"
	"io/ioutil"
	"strings"
	"bytes"
	"flag"
	"os"
	"fmt"
)

func die(s string, msg ...interface{}) {
	fmt.Fprintf(os.Stderr, s + "\n" , msg...)
	os.Exit(1)
}

var MtaServer string

func mailf() error {
	rb, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	ri := bytes.NewReader(rb)
	eml, err := mail.ReadMessage(ri)
	if err != nil {
		return err
	}

	from := eml.Header.Get("from")
	if from == "" {
		fmt.Printf("empty from header.\n")
		return nil
	}

	stos := eml.Header.Get("to")
	tos := strings.Split(stos, ",")
	if len(tos) < 1 {
		fmt.Printf("empty To header.\n")
		return nil
	}

	if err := send(from, tos, rb); err != nil {
		return err
	}
	fmt.Printf("file sended.\n")
	return nil
}

func send (from string, to []string, body []byte) error {
	if err := smtp.SendMail(MtaServer, nil,
		from, to, body); err != nil {
			return err
	}
	return nil
}

func init() {
	var mta_port int
	var mta_host string
	flag.StringVar(&mta_host, "s", "", "The address of the mail server.")
	flag.IntVar(&mta_port, "p", 25, "The Port of the mail server.")
	flag.Parse()

	if mta_host == "" {
		die("mail server address is blank.")
	}
	if mta_port < 1 || mta_port > 65535 {
		die("port is out of range.")
	}
	MtaServer = fmt.Sprintf("%s:%v", mta_host, mta_port)
}

func main() {
	if err := mailf(); err != nil {
		die("mailf error : %s", err)
	}
}
