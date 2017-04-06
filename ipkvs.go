package ipkvs

import (
	"errors"
	"sync"

	"gopkg.in/vmihailenco/msgpack.v2"
)

const (
	// Memory word size. Must be 8 for 64bit or 4 for 32bit
	wordSize = 8

	// Block size in words
	blockSize = 4

	// Page size in blocks
	pageSize = 524288
)

type (
	IpKVs struct {
		sync.RWMutex

		entries   map[uint64]uint
		freed     map[uint]struct{}
		pages     []page
		nextBlock uint

		Logger LoggerFunc
	}

	LoggerFunc func(msg string, args ...interface{})
)

var (
	ErrorBigValue = errors.New("value is too big")
	ErrorFound    = errors.New("key already found")
)

func NewIpKvs() *IpKVs {
	return &IpKVs{
		entries:   map[uint64]uint{},
		freed:     map[uint]struct{}{},
		pages:     []page{},
		nextBlock: 0,
	}
}

func (c *IpKVs) block() uint {
	for r, _ := range c.freed {
		delete(c.freed, r)
		return r
	}

	r := c.nextBlock
	c.nextBlock += 1
	return r
}

func (c *IpKVs) address(block uint) (int, int) {
	return int(block / uint(pageSize)),
		int(block % uint(pageSize))
}

func (c *IpKVs) CreateMsgPack(key string, value interface{}) error {
	d, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}
	return c.Create(key, d)
}

func (c *IpKVs) Create(key string, value []byte) error {
	if len(value) > blockSize*wordSize {
		return ErrorBigValue
	}

	hash := FNV1a(key)
	if _, ok := c.entries[hash]; ok {
		return ErrorFound
	}

	c.Lock()
	defer c.Unlock()

	b := c.block()
	p, o := c.address(b)
	for len(c.pages) < p+1 {
		c.pages = append(c.pages, page{})
	}

	c.pages[p][o].number = b
	for i := 0; i < len(value); i += 1 {
		c.pages[p][o].data[i] = value[i]
	}
	c.pages[p][o].len = len(value)
	c.entries[hash] = b
	return nil
}

func (c *IpKVs) ReadMsgPack(key string, value interface{}) error {
	return msgpack.Unmarshal(c.Read(key), value)
}

func (c *IpKVs) Read(key string) []byte {
	hash := FNV1a(key)

	c.RLock()
	defer c.RUnlock()

	b, ok := c.entries[hash]
	if !ok {
		return nil
	}

	p, o := c.address(b)
	return c.pages[p][o].data[:c.pages[p][o].len]
}

func (c *IpKVs) Delete(key string) {
	hash := FNV1a(key)

	c.Lock()
	defer c.Unlock()

	if block, ok := c.entries[hash]; ok {
		delete(c.entries, hash)
		c.freed[block] = struct{}{}
	}
}
