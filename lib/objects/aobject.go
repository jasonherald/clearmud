package objects

type wmap struct {
	limitx, limity, limitz int64
}

type basecoord struct {
	x, y, z int64
}

type coord struct {
	*basecoord
	parent int64
}

type wobject struct {
	coords      coord
	description string
	name        string
}

type planet struct {
	*wobject
}

type user struct {
	coords basecoord
}
