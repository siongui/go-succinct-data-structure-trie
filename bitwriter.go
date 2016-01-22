package bits

import "strings"

/**
  The BitWriter will create a stream of bytes, letting you write a certain
  number of bits at a time. This is part of the encoder, so it is not
  optimized for memory or speed.
*/
type BitWriter struct {
	bits []uint
}

/**
  Write some data to the bit string. The number of bits must be 32 or
  fewer.
*/
func (bw *BitWriter) Write(data, numBits uint) {
	//for i := (numBits-1); i >= 0; i-- {
	//FIXME: the above commented line will cause infinite loop, why???
	for i := numBits; i > 0; i-- {
		j := i - 1
		if (data & (1 << j)) != 0 {
			bw.bits = append(bw.bits, 1)
		} else {
			bw.bits = append(bw.bits, 0)
		}
	}
}

/**
  Get the bitstring represented as a javascript string of bytes
*/
func (bw *BitWriter) GetData() string {
	var chars []string
	var b, i uint = 0, 0

	for j := 0; j < len(bw.bits); j++ {
		b = (b << 1) | bw.bits[j]
		i += 1
		if i == W {
			chars = append(chars, CHR(b))
			i = 0
			b = 0
		}
	}

	if i != 0 {
		chars = append(chars, CHR(b<<(W-i)))
	}

	return strings.Join(chars, "")
}

/**
  Returns the bits as a human readable binary string for debugging
*/
func (bw *BitWriter) GetDebugString(group uint) string {
	var chars []string
	var i uint = 0

	for j := 0; j < len(bw.bits); j++ {
		if bw.bits[j] == 1 {
			chars = append(chars, "1")
		} else {
			chars = append(chars, "0")
		}
		i++
		if i == group {
			chars = append(chars, " ")
			i = 0
		}
	}

	return strings.Join(chars, "")
}
