package main

import (
	"github.com/own3dh2so4/telegram/client"
)

func main()  {
	hc := client.NewClient()
	/*hc.SendMessage(message.Message{
		ChatID:                "-1001357498410",
		Text:                  "Hola",
	});*/
	hc.GetUpdate(432399265)
}

