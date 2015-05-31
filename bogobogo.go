package bogobogo

import (
	"errors"
	"time"
)

func MIPS() (float64, error) {
	t := time.Millisecond * 33
	for l := uint64(1); l > 0; l <<= 1 {
		s := time.Now()
		for i := uint64(0); i < l; i++ {
		}
		e := time.Now().Sub(s)

		if e > t {
			for l2 := l >> 1; l2 < l; l2 = ((l2 + l) / 2) + l2/2 {
				s = time.Now()
				for i := uint64(0); i < l2; i++ {
				}
				e2 := time.Now().Sub(s)
				if e2 > t {
					l = l2
					e = e2
				}
			}
			bm := (float64(l) / (float64(e) / 1e+9) / 1e+6)
			return float64(int(bm*100)) / 100, nil
		}
	}
	return 0, errors.New("failed")
}
