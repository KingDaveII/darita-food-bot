package telegram

const msgHelp = `I can save and keep your pages. Also I can offer you them to read

In order to save page just send me its URL. To get random page just send /rnd command`

const msgHello = "Hi there \n\n" + msgHelp

const (
	msgUnkownCommand = "Unkown command. "
	msgNoSavedPages   = "You don't have saved pages yet. Just send me URL of page you want to save"
	msgSaved = "Page is saved. I will offer it to you later"
	msgAlreadyExists = "You have already saved this page. I will offer it to you later"
)