package objects

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
	Coords Basecoord
}
