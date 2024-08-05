package model

import "math/rand"

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}

type Particle struct {
	Position
	Velocity
}

func NewParticle(maxX int, maxY int, maxVelocityComponent int) *Particle {
	return &Particle{
		Position: Position{
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		},
		Velocity: Velocity{
			X: rand.Intn(maxVelocityComponent),
			Y: rand.Intn(maxVelocityComponent),
		},
	}
}

func (p *Particle) Update() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}
