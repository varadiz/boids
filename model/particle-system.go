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

	for i, p := range ps.Particles {
		// reposition if needed
		if p.Position.X >= ps.Width {
			ps.Particles[i].Position.X -= ps.Width
		}

		if p.Position.X <= 0 {
			ps.Particles[i].Position.X += ps.Width
		}

		if p.Position.Y >= ps.Height {
			ps.Particles[i].Position.Y -= ps.Height
		}

		if p.Position.Y <= 0 {
			ps.Particles[i].Position.Y += ps.Height
		}

		// scale down
		var x, y int
		x = int(math.Ceil(float64(ps.Particles[i].Position.X) * xScale))
		y = int(math.Ceil(float64(ps.Particles[i].Position.Y) * yScale))

		// hack
		if x == 0 {
			x += 1
		}
		if y == 0 {
			y += 1
		}

		screen[x-1][y-1] = '\u2593'
	}

	for _, line := range screen {
		fmt.Println(string(line))
	}
}
