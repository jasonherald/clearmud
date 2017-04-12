package actions

import (
	objects "../objects"
)

func moveNorth(u *objects.User) {
	u.Coords.Y++
}

func moveSouth(u *objects.User) {
	u.Coords.Y--
}

func moveEast(u *objects.User) {
	u.Coords.X++
}

func movWest(u *objects.User) {
	u.Coords.X--
}
