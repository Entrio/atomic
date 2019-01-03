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
		delta = time.Now().Sub(last)
		//TODO: Update
		(*g.activeScene).Update(delta)
		ups++
		if time.Now().Sub(lastUPS).Seconds() > 1 {
			aups = ups
			ups = 0
			lastUPS = time.Now()
		}
		last = time.Now()
	}
}
