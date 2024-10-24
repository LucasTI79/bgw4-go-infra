package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func sendMail() {
	from := "from@email.com"
	to := []string{"to1@email.com", "to2@email.com"}

	message := []byte(strings.Join([]string{
		"From: " + from,
		"To: " + strings.Join(to, ","),
		"ReplyTo: " + to[0],
		"Subject: Teste de envio de email",
		"",
		"Mensagem de teste",
	}, "\r\n"))

	err := smtp.SendMail("mail:1025", nil, from, to, message)
	if err != nil {
		fmt.Println("erro ao enviar o email")
		return
	}

	fmt.Println("email enviado com sucesso")
}

func main() {

	http.HandleFunc("/sendMail", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		sendMail()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
