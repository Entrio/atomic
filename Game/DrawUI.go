package Game

func (g *Game) DrawUI() {
	(*g.activeScene).DrawUI()
}
