package comparision

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/alxmsl/ipkvs"
	"github.com/coocood/freecache"
	"github.com/patrickmn/go-cache"
)

type (
	Data struct {
		V int8
	}
)

func BenchmarkGoCacheSet(b *testing.B) {
	b.StopTimer()
	c := cache.New(20*time.Minute, 0)

	b.StartTimer()
	for i := 0; i < b.N; i += 1 {
		c.Set(strconv.Itoa(rand.Intn(b.N)), Data{int8(rand.Intn(265))}, time.Minute)
	}
}

func BenchmarkFreeCacheSet(b *testing.B) {
	b.StopTimer()
	c := freecache.NewCache(512)
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Set([]byte(keys[i]), []byte{byte(rand.Intn(265))}, 20*60)
	}
}

func BenchmarkBigCacheSet(b *testing.B) {
	b.StopTimer()
	c, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:       256,
		LifeWindow:   10 * time.Minute,
		MaxEntrySize: 256,
		Verbose:      false,
	})
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Set(keys[i], []byte{byte(rand.Intn(265))})
	}
}

func BenchmarkCacheSetMsgPack(b *testing.B) {
	b.StopTimer()
	c := ipkvs.NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.StartTimer()
	for i := 0; i < b.N; i += 1 {
		c.CreateMsgPack(keys[i], Data{int8(rand.Intn(265))})
	}
}

func BenchmarkCacheSet(b *testing.B) {
	b.StopTimer()
	c := ipkvs.NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.StartTimer()
	for i := 0; i < b.N; i += 1 {
		c.Create(keys[i], []byte{byte(rand.Intn(265))})
	}
}
