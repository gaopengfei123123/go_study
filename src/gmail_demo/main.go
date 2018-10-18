package main

import (
	"fmt"
	"gmail_demo/mailer"
)

func main() {
	subject := "Hello!"
	body := "Hello <b>Bob</b> and <i>Cora</i>!"

	err := mailer.SendToMail(subject, body)

	// Send the email to Bob, Cora and Dan.
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
