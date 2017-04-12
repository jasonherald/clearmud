package objects

type wmap struct {
	limitx, limity int64
}

type coord struct {
	x, y, z, parent int64
}

type wobject struct {
	coords      coord
	description string
	name        string
}

type planet struct {
	*wobject
}
