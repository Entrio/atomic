package Game

func (g *Game) Draw2D() {
	(*g.activeScene).draw2D()
}
