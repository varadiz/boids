package model

import (
	"math"
	"math/rand"
)

type Position struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Particle struct {
	Position
	Velocity
}

func NewParticle(maxX int, maxY int, maxVelocityComponent int) *Particle {
	sign := int(math.Round(rand.Float64()))*2 - 1 // -1 or 1
	return &Particle{
		Position: Position{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		},
		Velocity: Velocity{
			X: sign * rand.Intn(maxVelocityComponent),
			Y: sign * rand.Intn(maxVelocityComponent),
		},
	}
}

func (p *Particle) Update() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}
