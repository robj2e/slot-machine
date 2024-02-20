package main

import gp "slot-machine/internal/gameparameters"

func main() {
	gameParams := gp.NewGameParams()
	for {
		gameParams.Begin()
		gameParams.Spin()
		gameParams.OutputReels()
		gameParams.DetermineOutcome()
		gameParams.Cleanup()
	}
}
