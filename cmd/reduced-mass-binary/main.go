package main

import (
	"log"
	"math"
	"github.com/brunetto/evolve"
)

func main () () {
	var (
		p *evolve.Particle
		err error
		dt float64 = 0.01
		nTS float64 = 100000
		ts float64
		G float64 = 1.
	)
	
	// Read a particle with mass the sum of the masses
	p = &evolve.Particle{}
	err = p.ReadFromLine("1e2,5e0,0e0,0e0,0e0,4.46e0,0e0")
	if err != nil {
		log.Fatal("Can't read particle with err: ", err)
	}
	
	p.Print()
	
	for ts=0.; ts<nTS*dt; ts=ts+dt {
		r2 := p.Pos[0]*p.Pos[0] + p.Pos[1]*p.Pos[1] + p.Pos[2]*p.Pos[2]
		r3 := r2*math.Sqrt(r2)
		for idx:=0; idx<3; idx++ {
			p.Acc[idx] = - G * p.Mass * p.Pos[idx] / r3
			p.Pos[idx] += p.Vel[idx] * dt
			p.Vel[idx] += p.Acc[idx] * dt
		}
		p.Print()
	}
}


