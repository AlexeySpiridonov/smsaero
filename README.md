# smsaero

SMSAero API Client (Go)

## Example

```go
package main

import (
	"fmt"
	"log"

	"gitlab.com/hiteam/smsaero"
)

func main() {
	client := smsaero.NewClient("e-mail", "api-secret", "", "")
	msg := smsaero.MessageRequest{
		Numbers: []string{"79871234567"},
		Sign: "hightech+",
		Text: "Ваш код: 123456",
		Channel: smsaero.ChannelDirect,
	}
	resp, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v\n", resp)
}
```