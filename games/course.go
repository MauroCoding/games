package games

type Game struct {
	name, publisher string
	year, price     float64
}

func (g Game) getInfo() string {
	txt := g.name + " " + g.publisher
	return txt
}
