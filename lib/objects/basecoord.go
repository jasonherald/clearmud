package objects

import "strconv"

type Basecoord struct {
	X, Y, Z int64
}

func (b *Basecoord) ToString() string {
	return strconv.FormatInt(b.X, 10) + ", " + strconv.FormatInt(b.Y, 10) + ", " + strconv.FormatInt(b.Z, 10)
}
