package Game

func (g *Game) Draw2D() {
	(*g.activeScene).Draw2D()
}
