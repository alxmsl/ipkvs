package ipkvs

type (
	block struct {
		number uint

		data [blockSize * wordSize]byte
		len  int
	}
)
