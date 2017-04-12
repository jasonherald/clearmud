package objects

import "strconv"

type Wmap struct {
	Limitx, Limity, Limitz int64
}

type Basecoord struct {
	X, Y, Z int64
}

type Coord struct {
	*Basecoord
	Parent int64
}

type Wobject struct {
	Coords      Coord
	Description string
	Name        string
}

type Planet struct {
	*Wobject
}

type User struct {
	X, Y, Z int64
}

func (b *Basecoord) ToString() string {
	return strconv.FormatInt(b.X, 10) + ", " + strconv.FormatInt(b.Y, 10) + ", " + strconv.FormatInt(b.Z, 10)
}

func (u *User) ToString() string {
	return strconv.FormatInt(u.X, 10) + ", " + strconv.FormatInt(u.Y, 10) + ", " + strconv.FormatInt(u.Z, 10)
}

func (u *User) Move(direction int) {
	switch direction {
	case 0:
		u.Y = u.Y + 1
	case 1:
		u.X = u.X - 1
	case 2:
		u.X = u.X + 1
	case 3:
		u.Y = u.Y - 1
	}
}
