package main

import (
	"boids/model"
	"fmt"
	"time"
)

func main() {
	particleCount := 15
	sizeX, sizeY := 10_000, 10_000
	maxSpeed := 10

	ps := model.NewParticleSystem(particleCount, sizeX, sizeY, maxSpeed)

	timer := time.NewTicker(10 * time.Millisecond)
	for {
		<-timer.C
		for i := range ps.Particles {
			ps.Particles[i].Update()
		}
		fmt.Print("\033[H\033[2J")
		ps.Display(35.0, 150.0)
	}
}
