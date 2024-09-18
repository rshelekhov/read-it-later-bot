package telegram

const msgHelp = `I can save and keep your pages. Also I can offer you them to read.

In order to save the page, just send me a link to it.

In order to get a random page from your list, just send me /rnd command.
Caution! After that, this page will be removed from your list!`

const msgHello = "Hello! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You have no saved pages"
	msgSaved          = "Saved"
	msgAlreadyExists  = "You have already saved this page"
)
