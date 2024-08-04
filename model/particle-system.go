package model

type ParticleSystem struct {
	SizeX     int
	SizeY     int
	Particles []Particle
}

func NewParticleSystem(particleCount, sizeX, sizeY int, maxSpeed int) *ParticleSystem {

	particles := make([]Particle, particleCount)
	for idx := range particleCount {
		particles[idx] = *NewParticle(sizeX, sizeY, maxSpeed)
	}

	return &ParticleSystem{
		SizeX:     sizeX,
		SizeY:     sizeY,
		Particles: particles,
	}
}
