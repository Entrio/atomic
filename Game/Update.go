package Game

import "time"

/**
Main update cycle
*/
func (g *Game) update() {
	//TODO: implement time lapsed since last frame
	var delta time.Duration
	last := time.Now()

	for !g.shouldExit {
		delta = last.Sub(time.Now())
		//TODO: Update
		(*g.activeScene).Update(delta)
		last = time.Now()
	}
}
