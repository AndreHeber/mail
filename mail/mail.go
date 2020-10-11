package mail

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Email struct {
	From        string   `json:"from" xml:"from" form:"from"`
	To          string   `json:"to" xml:"to" form:"to"`
	Subject     string   `json:"subject" xml:"subject" form:"subject"`
	Content     string   `json:"content" xml:"content" form:"content"`
	Attachments []string `json:"attachments" xml:"attachments" form:"attachments"`
}

// Init initailzes mail module
func Init(app *fiber.App) {
	// send endpoint
	app.Get("/mail/send/:id", sendMail)
	app.Post("/mail/create", createMail)
}

func sendMail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("wrong id: ", c.Params("id"))
		return c.SendStatus(404)
	}

	// // from config
	// mail := mailyak.New("mail.host.com:25", smtp.PlainAuth("", "user", "pass", "mail.host.com"))

	// mail.To("dom@itsallbroken.com")
	// mail.From("oops@itsallbroken.com")
	// mail.Subject("I am a teapot")
	// mail.HTML().Set("Don't panic")

	// // input can be a bytes.Buffer, os.File, os.Stdin, etc.
	// // call multiple times to attach multiple files
	// mail.Attach("filename.txt", &input)

	// if err := mail.Send(); err != nil {
	// 	panic(" 💣 ")
	// }

	log.Println("Send Mail with id: ", id)
	return c.SendStatus(200)
}

type CreateEmailRequest struct {
	Name        string   `json:"name" xml:"name" form:"name"`
	From        string   `json:"from" xml:"from" form:"from"`
	To          string   `json:"to" xml:"to" form:"to"`
	Subject     string   `json:"subject" xml:"subject" form:"subject"`
	Content     string   `json:"content" xml:"content" form:"content"`
	Attachments []string `json:"attachments" xml:"attachments" form:"attachments"`
}

func createMail(c *fiber.Ctx) error {
	e := new(CreateEmailRequest)

	err := c.BodyParser(e)
	if err != nil {
		return err
	}

	log.Printf("Email Name: %s, From: %s, To: %s, Subject %s, Content: %s\n", e.Name, e.From, e.To, e.Subject, e.Content)

	return c.SendStatus(200)
}
