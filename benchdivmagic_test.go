// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package benchdivmagic_test

import (
	"math/bits"
	"testing"
	"unsafe"
)

const StartAddr = 8 << 20 // arbitrary

func TestDivMagic(t *testing.T) {
	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}

	for k := uintptr(0); k < 4096; k += 97 {
		for i := range spans {
			size := uintptr(class_to_size[i+1])

			base, index := spans[i].findObject(StartAddr + k)
			if want := StartAddr + index*size; want != base {
				t.Errorf("bad: %v != %v", want, base)
			}
			if StartAddr+k < base || StartAddr+k >= base+size {
				t.Errorf("also bad; %v %v %v", k, base, size)
			}
			if want := spans[i].objIndex(StartAddr + k); want != index {
				t.Errorf("triple bad: %v %v", want, index)
			}
		}
	}
}

func TestLemire(t *testing.T) {
	var spans [_NumSizeClasses - 1]lspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}

	for k := uintptr(0); k < 4096; k += 97 {
		for i := range spans {
			size := uintptr(class_to_size[i+1])

			base, index := spans[i].findObject(StartAddr + k)
			if want := StartAddr + index*size; want != base {
				t.Errorf("bad: %v != %v", want, base)
			}
			if StartAddr+k < base || StartAddr+k >= base+size {
				t.Errorf("also bad; %v %v %v", k, base, size)
			}
			if want := spans[i].objIndex(StartAddr + k); want != index {
				t.Errorf("triple bad: %v %v", want, index)
			}
		}
	}
}

var Output uintptr

func BenchmarkDivMagic_objIndex(b *testing.B) {
	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				sum += spans[i].objIndex(StartAddr + k)
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_objIndex2(b *testing.B) {
	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				sum += spans[i].objIndex2(StartAddr + k)
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_objIndex3(b *testing.B) {
	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				sum += spans[i].objIndex3(StartAddr + k)
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_objIndex4(b *testing.B) {
	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				sum += spans[i].objIndex4(StartAddr + k)
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_findObject(b *testing.B) {

	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				a, b := spans[i].findObject(StartAddr + k)
				sum += a + b
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_findObject2(b *testing.B) {

	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				a, b := spans[i].findObject2(StartAddr + k)
				sum += a + b
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_findObject3(b *testing.B) {

	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				a, b := spans[i].findObject3(StartAddr + k)
				sum += a + b
			}
		}
	}
	Output = sum
}

func BenchmarkDivMagic_findObject4(b *testing.B) {

	var spans [_NumSizeClasses - 1]mspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				a, b := spans[i].findObject4(StartAddr + k)
				sum += a + b
			}
		}
	}
	Output = sum
}

func BenchmarkLemire_objIndex(b *testing.B) {

	var spans [_NumSizeClasses - 1]lspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				sum += spans[i].objIndex(StartAddr + k)
			}
		}
	}
	Output = sum
}

func BenchmarkLemire_findObject(b *testing.B) {

	var spans [_NumSizeClasses - 1]lspan
	for i := range spans {
		spans[i].init(StartAddr, i+1)
	}
	b.ResetTimer()

	var sum uintptr
	for n := b.N; n > 0; n-- {
		for k := uintptr(0); k < 4096; k += 97 {
			for i := range spans {
				a, b := spans[i].findObject(StartAddr + k)
				sum += a + b
			}
		}
	}
	Output = sum
}

type lspan struct {
	startAddr uintptr
	elemsize  uintptr
	magic     lemire
}

type lemire uintptr

func lemireOf(d uintptr) lemire {
	return ^lemire(0)/lemire(d) + 1
}

func (l lemire) div(n uintptr) uintptr {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return uintptr((uint64(n) * uint64(l)) >> 32)
	} else {
		x, _ := bits.Mul64(uint64(n), uint64(l))
		return uintptr(x)
	}
}

func (s *lspan) init(addr uintptr, sizeclass int) {
	s.startAddr = addr
	s.elemsize = uintptr(class_to_size[sizeclass])
	s.magic = lemireOf(s.elemsize)
}

func (s *lspan) base() uintptr { return s.startAddr }

func (s *lspan) objIndex(p uintptr) uintptr {
	return s.magic.div(p - s.base())
}

func (s *lspan) findObject(p uintptr) (base, objIndex uintptr) {
	objIndex = s.objIndex(p)
	base = s.base() + objIndex*s.elemsize
	return
}

type mspan struct {
	startAddr uintptr
	elemsize  uintptr // computed from sizeclass or from npages
	baseMask  uint16  // if non-0, elemsize is a power of 2, & this will get object allocation base
	divMul    uint16  // for divide by elemsize - divMagic.mul
	divShift  uint8   // for divide by elemsize - divMagic.shift
	divShift2 uint8   // for divide by elemsize - divMagic.shift2
}

func (s *mspan) init(addr uintptr, sizeclass int) {
	s.startAddr = addr

	s.elemsize = uintptr(class_to_size[sizeclass])

	m := &class_to_divmagic[sizeclass]
	s.divShift = m.shift
	s.divMul = m.mul
	s.divShift2 = m.shift2
	s.baseMask = m.baseMask
}

func (s *mspan) base() uintptr {
	return s.startAddr
}

func (s *mspan) objIndex(p uintptr) uintptr {
	byteOffset := p - s.base()
	if byteOffset == 0 {
		return 0
	}
	if s.baseMask != 0 {
		// s.baseMask is non-0, elemsize is a power of two, so shift by s.divShift
		return byteOffset >> s.divShift
	}
	return uintptr(((uint64(byteOffset) >> s.divShift) * uint64(s.divMul)) >> s.divShift2)
}

func (s *mspan) objIndex2(p uintptr) uintptr {
	byteOffset := p - s.base()
	if byteOffset == 0 {
		return 0
	}
	if s.baseMask != 0 {
		// s.baseMask is non-0, elemsize is a power of two, so shift by s.divShift
		return byteOffset >> s.divShift
	}
	return uintptr(((uint64(byteOffset) >> (s.divShift & 31)) * uint64(s.divMul)) >> (s.divShift2 & 31))
}

func (s *mspan) objIndex3(p uintptr) uintptr {
	byteOffset := p - s.base()
	if s.baseMask != 0 {
		// s.baseMask is non-0, elemsize is a power of two, so shift by s.divShift
		return byteOffset >> s.divShift
	}
	return uintptr(((uint64(byteOffset) >> (s.divShift & 31)) * uint64(s.divMul)) >> (s.divShift2 & 31))
}

func (s *mspan) objIndex4(p uintptr) uintptr {
	byteOffset := p - s.base()
	return uintptr(((uint64(byteOffset) >> (s.divShift & 31)) * uint64(s.divMul)) >> (s.divShift2 & 31))
}

func (s *mspan) findObject(p uintptr) (base, objIndex uintptr) {
	if s.baseMask != 0 {
		// optimize for power of 2 sized objects.
		base = s.base()
		base = base + (p-base)&uintptr(s.baseMask)
		objIndex = (base - s.base()) >> s.divShift
		// base = p & s.baseMask is faster for small spans,
		// but doesn't work for large spans.
		// Overall, it's faster to use the more general computation above.
	} else {
		base = s.base()
		if p-base >= s.elemsize {
			// n := (p - base) / s.elemsize, using division by multiplication
			objIndex = uintptr(p-base) >> s.divShift * uintptr(s.divMul) >> s.divShift2
			base += objIndex * s.elemsize
		}
	}
	return
}

func (s *mspan) findObject2(p uintptr) (base, objIndex uintptr) {
	if s.baseMask != 0 {
		// optimize for power of 2 sized objects.
		base = s.base()
		base = base + (p-base)&uintptr(s.baseMask)
		objIndex = (base - s.base()) >> (s.divShift & 31)
		// base = p & s.baseMask is faster for small spans,
		// but doesn't work for large spans.
		// Overall, it's faster to use the more general computation above.
	} else {
		base = s.base()
		if p-base >= s.elemsize {
			// n := (p - base) / s.elemsize, using division by multiplication
			objIndex = uintptr(p-base) >> (s.divShift & 31) * uintptr(s.divMul) >> (s.divShift2 & 31)
			base += objIndex * s.elemsize
		}
	}
	return
}

func (s *mspan) findObject3(p uintptr) (base, objIndex uintptr) {
	if s.baseMask != 0 {
		// optimize for power of 2 sized objects.
		base = s.base()
		base = base + (p-base)&uintptr(s.baseMask)
		objIndex = (base - s.base()) >> (s.divShift & 31)
		// base = p & s.baseMask is faster for small spans,
		// but doesn't work for large spans.
		// Overall, it's faster to use the more general computation above.
	} else {
		base = s.base()
		// n := (p - base) / s.elemsize, using division by multiplication
		objIndex = uintptr(p-base) >> (s.divShift & 31) * uintptr(s.divMul) >> (s.divShift2 & 31)
		base += objIndex * s.elemsize
	}
	return
}

func (s *mspan) findObject4(p uintptr) (base, objIndex uintptr) {
	if s.baseMask != 0 {
		// optimize for power of 2 sized objects.
		base = s.base()
		base = base + (p-base)&uintptr(s.baseMask)
		objIndex = (base - s.base()) >> (s.divShift & 31)
		// base = p & s.baseMask is faster for small spans,
		// but doesn't work for large spans.
		// Overall, it's faster to use the more general computation above.
	} else {
		base = s.base()
		// n := (p - base) / s.elemsize, using division by multiplication
		objIndex = uintptr(p-base) >> (s.divShift & 31) * uintptr(s.divMul) >> (s.divShift2 & 31)
		base += objIndex * s.elemsize
	}
	return
}

// From $GOROOT/src/runtime/sizeclasses.go.

const (
	_MaxSmallSize   = 32768
	smallSizeDiv    = 8
	smallSizeMax    = 1024
	largeSizeDiv    = 128
	_NumSizeClasses = 68
	_PageShift      = 13
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432, 19072, 20480, 21760, 24576, 27264, 28672, 32768}
var class_to_allocnpages = [_NumSizeClasses]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 2, 3, 1, 3, 2, 3, 4, 5, 6, 1, 7, 6, 5, 4, 3, 5, 7, 2, 9, 7, 5, 8, 3, 10, 7, 4}

type divMagic struct {
	shift    uint8
	shift2   uint8
	mul      uint16
	baseMask uint16
}

var class_to_divmagic = [_NumSizeClasses]divMagic{{0, 0, 0, 0}, {3, 0, 1, 65528}, {4, 0, 1, 65520}, {3, 11, 683, 0}, {5, 0, 1, 65504}, {4, 11, 683, 0}, {6, 0, 1, 65472}, {4, 10, 205, 0}, {5, 9, 171, 0}, {4, 11, 293, 0}, {7, 0, 1, 65408}, {4, 13, 911, 0}, {5, 10, 205, 0}, {4, 12, 373, 0}, {6, 9, 171, 0}, {4, 13, 631, 0}, {5, 11, 293, 0}, {4, 13, 547, 0}, {8, 0, 1, 65280}, {5, 9, 57, 0}, {6, 9, 103, 0}, {5, 12, 373, 0}, {7, 7, 43, 0}, {5, 10, 79, 0}, {6, 10, 147, 0}, {5, 11, 137, 0}, {9, 0, 1, 65024}, {6, 9, 57, 0}, {7, 9, 103, 0}, {6, 11, 187, 0}, {8, 7, 43, 0}, {7, 8, 37, 0}, {10, 0, 1, 64512}, {7, 9, 57, 0}, {8, 6, 13, 0}, {7, 11, 187, 0}, {9, 5, 11, 0}, {8, 8, 37, 0}, {11, 0, 1, 63488}, {8, 9, 57, 0}, {7, 10, 49, 0}, {10, 5, 11, 0}, {7, 10, 41, 0}, {7, 9, 19, 0}, {12, 0, 1, 61440}, {8, 9, 27, 0}, {8, 10, 49, 0}, {11, 5, 11, 0}, {7, 13, 161, 0}, {7, 13, 155, 0}, {8, 9, 19, 0}, {13, 0, 1, 57344}, {8, 12, 111, 0}, {9, 9, 27, 0}, {11, 6, 13, 0}, {7, 14, 193, 0}, {12, 3, 3, 0}, {8, 13, 155, 0}, {11, 8, 37, 0}, {14, 0, 1, 49152}, {11, 8, 29, 0}, {7, 13, 55, 0}, {12, 5, 7, 0}, {8, 14, 193, 0}, {13, 3, 3, 0}, {7, 14, 77, 0}, {12, 7, 19, 0}, {15, 0, 1, 32768}}
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 19, 19, 20, 20, 20, 20, 21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24, 24, 24, 24, 25, 25, 25, 25, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 28, 28, 28, 28, 28, 28, 28, 28, 29, 29, 29, 29, 29, 29, 29, 29, 30, 30, 30, 30, 30, 30, 30, 30, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{32, 33, 34, 35, 36, 37, 37, 38, 38, 39, 39, 40, 40, 40, 41, 41, 41, 42, 43, 43, 44, 44, 44, 44, 44, 45, 45, 45, 45, 45, 45, 46, 46, 46, 46, 47, 47, 47, 47, 47, 47, 48, 48, 48, 49, 49, 50, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 53, 53, 54, 54, 54, 54, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 61, 61, 61, 61, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67}
