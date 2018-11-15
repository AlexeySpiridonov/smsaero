package main

import (
	"fmt"
	"log"

	"gitlab.com/hiteam/smsaero"
)

func main() {
	client := smsaero.NewClient("dev@hightech.plus", "qFx5Jr3wtu8z24RZbsEyVLCWIp4i", "", "")
	msg := smsaero.MessageRequest{
		Numbers: []string{"79870601510"},
		Sign:    "hightech+",
		Text:    "Ваш код: 123456",
		Channel: smsaero.ChannelDirect,
	}
	resp, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v\n", resp)
}
