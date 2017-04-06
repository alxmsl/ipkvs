package main

import (
	"fmt"

	"github.com/alxmsl/ipkvs"
	"github.com/mgutz/logxi/v1"
)

func main() {
	c := ipkvs.NewIpKvs()
	c.Logger = log.Warn

	fmt.Println(c.Read("hui"))

	err := c.Create("hui", []byte("hui"))
	fmt.Println(err)
	fmt.Println(string(c.Read("hui")))
}
