package objects

import "strconv"

type User struct {
	X, Y, Z int64
}

func (u *User) ToString() string {
	return strconv.FormatInt(u.X, 10) + ", " + strconv.FormatInt(u.Y, 10) + ", " + strconv.FormatInt(u.Z, 10)
}

func (u *User) Move(direction int) {
	switch direction {
	case 0:
		u.Y = u.Y + 1
	case 1:
		u.X = u.X + 1
	case 2:
		u.X = u.X - 1
	case 3:
		u.Y = u.Y - 1
	case 4:
		u.Z = u.Z + 1
	case 5:
		u.Z = u.Z - 1
	}
}
