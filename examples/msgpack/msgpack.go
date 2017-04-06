package main

import (
	"fmt"
	"math/rand"

	"github.com/alxmsl/ipkvs"
	"github.com/mgutz/logxi/v1"
)

type (
	Data struct {
		V1 int16
	}
)

func main() {
	c := ipkvs.NewIpKvs()
	c.Logger = log.Warn

	v1 := Data{int16(rand.Intn(265))}
	err := c.CreateMsgPack("aaa", v1)
	fmt.Println(err)

	v2 := Data{}
	err = c.ReadMsgPack("aaa", &v2)
	fmt.Println(err)
	fmt.Println(v1)
	fmt.Println(v2)
}
