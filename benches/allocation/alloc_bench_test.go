package main

import "testing"

const (
	size6  = 1 << 6
	size7  = 1 << 7
	size8  = 1 << 8
	size9  = 1 << 9
	size10 = 1 << 10
	size11 = 1 << 11
	size12 = 1 << 12
	size13 = 1 << 13
	size14 = 1 << 14
	size15 = 1 << 15
	size16 = 1 << 16
	size17 = 1 << 17
	size18 = 1 << 18
	size19 = 1 << 19
	size20 = 1 << 20
	size21 = 1 << 21
	size22 = 1 << 22
	size23 = 1 << 23
	size24 = 1 << 24
	size25 = 1 << 25
	size26 = 1 << 26
	size27 = 1 << 27
	size28 = 1 << 28
	size29 = 1 << 29
	size30 = 1 << 30
)

var (
	array6  *[size6]byte
	array7  *[size7]byte
	array8  *[size8]byte
	array9  *[size9]byte
	array10 *[size10]byte
	array11 *[size11]byte
	array12 *[size12]byte
	array13 *[size13]byte
	array14 *[size14]byte
	array15 *[size15]byte
	array16 *[size16]byte
	array17 *[size17]byte
	array18 *[size18]byte
	array19 *[size19]byte
	array20 *[size20]byte
	array21 *[size21]byte
	array22 *[size22]byte
	array23 *[size23]byte
	array24 *[size24]byte
	array25 *[size25]byte
	array26 *[size26]byte
	array27 *[size27]byte
	array28 *[size28]byte
	array29 *[size29]byte
	array30 *[size30]byte
)

func BenchmarkAllocate6(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array6 = &[size6]byte{}
	}
}

func BenchmarkAllocate7(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array7 = &[size7]byte{}
	}
}

func BenchmarkAllocate8(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array8 = &[size8]byte{}
	}
}

func BenchmarkAllocate9(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array9 = &[size9]byte{}
	}
}

func BenchmarkAllocate10(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array10 = &[size10]byte{}
	}
}

func BenchmarkAllocate11(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array11 = &[size11]byte{}
	}
}

func BenchmarkAllocate12(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array12 = &[size12]byte{}
	}
}

func BenchmarkAllocate13(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array13 = &[size13]byte{}
	}
}

func BenchmarkAllocate14(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array14 = &[size14]byte{}
	}
}

func BenchmarkAllocate15(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array15 = &[size15]byte{}
	}
}

func BenchmarkAllocate16(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array16 = &[size16]byte{}
	}
}

func BenchmarkAllocate17(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array17 = &[size17]byte{}
	}
}

func BenchmarkAllocate18(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array18 = &[size18]byte{}
	}
}

func BenchmarkAllocate19(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array19 = &[size19]byte{}
	}
}

func BenchmarkAllocate20(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array20 = &[size20]byte{}
	}
}

func BenchmarkAllocate21(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array21 = &[size21]byte{}
	}
}

func BenchmarkAllocate22(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array22 = &[size22]byte{}
	}
}

func BenchmarkAllocate23(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array23 = &[size23]byte{}
	}
}

func BenchmarkAllocate24(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array24 = &[size24]byte{}
	}
}

func BenchmarkAllocate25(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array25 = &[size25]byte{}
	}
}

func BenchmarkAllocate26(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array26 = &[size26]byte{}
	}
}

func BenchmarkAllocate27(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array27 = &[size27]byte{}
	}
}

func BenchmarkAllocate28(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array28 = &[size28]byte{}
	}
}

func BenchmarkAllocate29(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array29 = &[size29]byte{}
	}
}

func BenchmarkAllocate30(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		array30 = &[size30]byte{}
	}
}
