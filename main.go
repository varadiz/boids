package main

import (
	"boids/model"
	"fmt"
)

func main() {
	particleCount := 3
	sizeX, sizeY := 300, 300
	maxSpeed := 10

	ps := model.NewParticleSystem(particleCount, sizeX, sizeY, maxSpeed)
	fmt.Println(ps)

	for _, p := range ps.Particles {
		p.Update()
	}
}
