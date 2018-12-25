package Game

func (g *Game) Draw3D() {
	(*g.activeScene).draw3D()
}
