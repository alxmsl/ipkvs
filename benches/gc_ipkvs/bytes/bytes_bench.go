package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/alxmsl/ipkvs"
	"github.com/dustin/go-humanize"
	"github.com/mkideal/cli"
)

type arg struct {
	Limit int `cli:"l,limit" usage:"records limit"`
}

var (
	argv *arg
)

func init() {
	debug.SetGCPercent(10)
}

func main() {
	argv = &arg{}
	cli.Run(argv, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*arg)
		c := ipkvs.NewIpKvs()

		for i := 0; i < argv.Limit; i += 1 {
			c.Create(strconv.Itoa(rand.Intn(argv.Limit)), []byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)})
		}
		fmt.Println(gcPause())
		fmt.Println(memStat())
		return nil
	})
}

func gcPause() time.Duration {
	runtime.GC()
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	return stats.Pause[0]
}

func memStat() string {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	return humanize.IBytes(m.TotalAlloc)
}
