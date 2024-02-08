package facades

import "github.com/go-unity/framework/contracts/mail"

func Mail() mail.Mail {
	return App().MakeMail()
}
