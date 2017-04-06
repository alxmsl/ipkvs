package ipkvs

import (
	"math/rand"
	"strconv"
	"testing"
)

type (
	Data struct {
		V int8
	}
)

func BenchmarkSetSerial(b *testing.B) {
	c := NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		c.Create(keys[i], []byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)})
	}
}

func BenchmarkSetMsgPackSerial(b *testing.B) {
	c := NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		c.CreateMsgPack(keys[i], Data{int8(rand.Intn(265))})
	}
}

func BenchmarkGetSerial(b *testing.B) {
	c := NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	for i := 0; i < b.N; i += 1 {
		c.Create(keys[i], []byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		c.Read(keys[i])
	}
}

func BenchmarkGetMsgpackSerial(b *testing.B) {
	c := NewIpKvs()
	keys := make([]string, 0, b.N)
	for i := 0; i < b.N; i += 1 {
		keys = append(keys, strconv.Itoa(rand.Intn(b.N)))
	}

	for i := 0; i < b.N; i += 1 {
		c.CreateMsgPack(keys[i], Data{int8(rand.Intn(265))})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		c.ReadMsgPack(keys[i], &Data{})
	}
}
