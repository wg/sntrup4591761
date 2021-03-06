// These routines are a port of the public domain, C reference implementation.

package vector

import (
	"github.com/companyzero/sntrup4591761/r3/mod3"
)

// Swap swaps x and y if mask is -1. If mask is 0, x and y retain
// their original values.
func Swap(x, y []int8, bytes, mask int) {
	c := int8(mask)
	for i := 0; i < bytes; i++ {
		t := c & (x[i] ^ y[i])
		x[i] ^= t
		y[i] ^= t
	}
}

func Product(z []int8, n int, x []int8, c int8) {
	for i := 0; i < n; i++ {
		z[i] = mod3.Product(x[i], c)
	}
}

func MinusProduct(z []int8, n int, x, y []int8, c int8) {
	for i := 0; i < n; i++ {
		z[i] = mod3.MinusProduct(x[i], y[i], c)
	}
}

func Shift(z []int8, n int) {
	for i := n - 1; i > 0; i-- {
		z[i] = z[i-1]
	}
	z[0] = 0
}
