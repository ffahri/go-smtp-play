package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	c, err := smtp.Dial("localhost:2525")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("test@test.test"); err != nil {
		log.Fatal(err)
	}
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "howdy")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}

}
