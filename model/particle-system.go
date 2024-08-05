package model

import (
	"fmt"
	"math"
)

type ParticleSystem struct {
	Width     int
	Height    int
	Particles []Particle
}

func NewParticleSystem(particleCount, width, height int, maxSpeed int) *ParticleSystem {

	particles := make([]Particle, particleCount)
	for idx := range particleCount {
		particles[idx] = *NewParticle(width, height, maxSpeed)
	}

	return &ParticleSystem{
		Width:     width,
		Height:    height,
		Particles: particles,
	}
}

func (ps *ParticleSystem) Display(screenWidth, screenHeight float64) {
	xScale := screenWidth / float64(ps.Width)
	yScale := screenHeight / float64(ps.Height)

	screen := make([][]rune, int(screenWidth))
	for i := range screen {
		screen[i] = make([]rune, int(screenHeight))
		for j := range screen[i] {
			screen[i][j] = ' '
		}
	}

	for _, p := range ps.Particles {
		// scale down
		x := int(math.Round(float64(p.Position.X) * xScale))
		y := int(math.Round(float64(p.Position.Y) * yScale))
		screen[x-1][y-1] = '\u2593'
	}

	for _, line := range screen {
		fmt.Println(string(line))
	}
}
