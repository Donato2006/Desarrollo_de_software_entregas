package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func EnviarCorreo(destino string, concierto string) error {

	remitente := os.Getenv("MAIL_USER")
	password := os.Getenv("MAIL_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		remitente,
		password,
		"smtp.gmail.com",
	)

	asunto := "Subject: Entrada asignada automaticamente\r\n"

	mensaje := []byte(
		asunto +
			"\r\n" +
			"Hola.\r\n\r\n" +
			"Se te asignó automáticamente una entrada para:\r\n\r\n" +
			concierto +
			"\r\n\r\n" +
			"Ingresá a la plataforma para verla.\r\n\r\n" +
			"Ticket Conciertos",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		remitente,
		[]string{destino},
		mensaje,
	)

	if err != nil {
		fmt.Println("Error enviando correo:", err)
		return err
	}

	fmt.Println("Correo enviado a:", destino)

	return nil
}
