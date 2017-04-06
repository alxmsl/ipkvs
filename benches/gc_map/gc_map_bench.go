package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
	"unsafe"

	"github.com/dustin/go-humanize"
	"github.com/mkideal/cli"
)

const (
	boolMap          = "boolmap"
	intMap           = "intmap"
	pointerToIntMap  = "pintmap"
	uintPtrMap       = "uintptrmap"
	unsafePointerMap = "unsafepointermap"
	emptyStructMap   = "emptystructmap"
	staticStructMap  = "staticstructmap"
	dynamicStructMap = "dynamicstructmap"
	intStringMap     = "intstringmap"
)

type arg struct {
	ShowTypes bool `cli:"show-types" usage:"show available types"`

	Size int    `cli:"s,size" usage:"map size"`
	Type string `cli:"t,type" usage:"map type"`
}

var (
	argv *arg

	types = map[string]string{
		boolMap:          "map[int]bool",
		intMap:           "map[int]int",
		pointerToIntMap:  "map[int]*int",
		uintPtrMap:       "map[int]unsafe.Pointer",
		unsafePointerMap: "map[int]uintptr",
		emptyStructMap:   "map[int]struct{} (struct is empty here)",
		staticStructMap:  "map[int]struct{} (struct has fixed length)",
		dynamicStructMap: "map[int]struct{} (struct has pointer)",
		intStringMap:     "map[string]int",
	}
)

func init() {
	debug.SetGCPercent(-1)
}

func main() {
	argv = &arg{}
	cli.Run(argv, func(ctx *cli.Context) error {
		argv := ctx.Argv().(*arg)
		if argv.ShowTypes {
			for t, h := range types {
				fmt.Println(t, "is", h)
			}
			return nil
		}

		switch argv.Type {
		case boolMap:
			boolmap := make(map[int]bool, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				boolmap[i] = rand.Int31n(100) > 50
			}
		case intMap:
			intmap := make(map[int]int, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				intmap[i] = rand.Int()
			}
		case intStringMap:
			intmap := make(map[string]int, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				intmap[strconv.Itoa(i)] = rand.Int()
			}
		case pointerToIntMap:
			pintmap := make(map[int]*int, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				v := rand.Int()
				pintmap[i] = &v
			}
		case unsafePointerMap:
			unsafepointermap := make(map[int]unsafe.Pointer, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				v := rand.Int()
				unsafepointermap[i] = unsafe.Pointer(&v)
			}
		case uintPtrMap:
			uintptrmap := make(map[int]uintptr, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				v := rand.Int()
				uintptrmap[i] = uintptr(unsafe.Pointer(&v))
			}
		case staticStructMap:
			type test struct {
				a int
				b [2]int
			}
			structmap := make(map[int]test, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				v := rand.Int()
				structmap[i] = test{
					a: v,
					b: [2]int{v, v},
				}
			}
		case dynamicStructMap:
			type test struct {
				a *int
			}
			structmap := make(map[int]test, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				v := rand.Int()
				structmap[i] = test{
					a: &v,
				}
			}
		case emptyStructMap:
			structmap := make(map[int]struct{}, argv.Size)
			for i := 0; i < argv.Size; i += 1 {
				structmap[i] = struct{}{}
			}
		default:
			fmt.Println(ctx.Usage())
			return nil
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
