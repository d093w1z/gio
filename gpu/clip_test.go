// SPDX-License-Identifier: Unlicense OR MIT

package gpu

import (
	"testing"

	"github.com/d093w1z/gio/internal/f32"
)

func BenchmarkEncodeQuadTo(b *testing.B) {
	var data [vertStride * 4]byte
	for i := 0; b.Loop(); i++ {
		v := float32(i)
		encodeQuadTo(data[:], 123,
			f32.Point{X: v, Y: v},
			f32.Point{X: v, Y: v},
			f32.Point{X: v, Y: v},
		)
	}
}
