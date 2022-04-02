package helper

import (
	"net/smtp"

	"github.com/pkg/errors"
)

const (
	// server we are authorized to send email through
	host     = "smtp.gmail.com"
	hostPort = ":587"

	// user we are authorizing as
	from     string = "upm.udevs.io@gmail.com"
	password string = "gehwhgelispgqoql"
)

func SendEmail(subject, to, link, token string) error {
	message := `
		You can update your password using the following url
   
	   ` + link + "?token=" + token

	auth := smtp.PlainAuth("", from, password, host)

	//  // // NOTE: Using the backtick here ` works like a heredoc, which is why all the
	//  // // rest of the lines are forced to the beginning of the line, otherwise the
	//  // // formatting is wrong for the RFC 822 style
	//  msg := `To: "` + to + `" <` + to + `>
	// From: "` + from + `" <` + from + `>
	// Subject: ` + subject + `
	// ` + message
	msg := "To: \"" + to + "\" <" + to + ">\n" +
		"From: \"" + from + "\" <" + from + ">\n" +
		"Subject: " + subject + "\n" +
		message + "\n"

	if err := smtp.SendMail(host+hostPort, auth, from, []string{to}, []byte(msg)); err != nil {
		return errors.Wrap(err, "error while sending message to email")
	}

	return nil
}
