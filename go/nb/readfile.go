package nb

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Particles []Particle
type Particle struct {
	// ID         string
	Mass          float64
	X, Y, Z       float64
	Vx, Vy, Vz    float64
	Ax, Ay, Az    float64
	A0x, A0y, A0z float64
}

func ReadFile(fn string) Particles {
	f, err := os.Open(fn)
	DieIf(err)

	s := bufio.NewScanner(f)

	var ps Particles
	for s.Scan() {
		tmp := strings.Fields(s.Text())
		if len(tmp) != 8 {
			continue
		}

		p := Particle{}
		// p.ID =   tmp[0]
		p.Mass, err = strconv.ParseFloat(tmp[1], 64)
		DieIf(err)
		p.X, err = strconv.ParseFloat(tmp[2], 64)
		DieIf(err)
		p.Y, err = strconv.ParseFloat(tmp[3], 64)
		DieIf(err)
		p.Z, err = strconv.ParseFloat(tmp[4], 64)
		DieIf(err)
		p.Vx, err = strconv.ParseFloat(tmp[5], 64)
		DieIf(err)
		p.Vy, err = strconv.ParseFloat(tmp[6], 64)
		DieIf(err)
		p.Vz, err = strconv.ParseFloat(tmp[7], 64)
		DieIf(err)

		ps = append(ps, p)

	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	return ps
}

func (ps Particles) Energy(ek, ep float64) (float64, float64) {
	// kinetic energy
	for _, p := range ps {
		ek += 0.5 * p.Mass * (p.Vx*p.Vx + p.Vy*p.Vy + p.Vz*p.Vz)
	}

	// potential energy
	pi := ps[0]
	for _, pj := range ps[1:] {
		dx := pi.X - pj.X
		dy := pi.Y - pj.Y
		dz := pi.Z - pj.Z

		ep -= pi.Mass * pj.Mass / math.Sqrt(dx*dx+dy*dy+dz*dz)

		pi = pj
	}

	return ek, ep
}

func (ps Particles) Acceleration() {
	for i, pi := range ps {
		pi.Ax, pi.Ay, pi.Az = 0, 0, 0

		// TODO: nothing better?
		for j, pj := range ps {
			if i == j {
				continue
			}

			dx := pi.X - pj.X
			dy := pi.Y - pj.Y
			dz := pi.Z - pj.Z

			r2 := dx*dx + dy*dy + dz*dz

			a := 1.0 / math.Sqrt(r2*r2*r2)

			ps[i].Ax -= pj.Mass * a * dx
			ps[i].Ay -= pj.Mass * a * dy
			ps[i].Az -= pj.Mass * a * dz

		}
	}
}

func (ps Particles) Positions(dt float64) {
	for i, p := range ps {
		// save accelerations
		ps[i].A0x = p.Ax
		ps[i].A0y = p.Ay
		ps[i].A0z = p.Az

		ps[i].X += dt*p.Vx + 0.5*dt*dt*p.Ax
		ps[i].Y += dt*p.Vy + 0.5*dt*dt*p.Ay
		ps[i].Z += dt*p.Vz + 0.5*dt*dt*p.Az
	}
}

func (ps Particles) Velocities(dt float64) {
	for i, p := range ps {
		ps[i].Vx += 0.5 * dt * (p.A0x + p.Ax)
		ps[i].Vy += 0.5 * dt * (p.A0y + p.Ay)
		ps[i].Vz += 0.5 * dt * (p.A0z + p.Az)

		// save accelerations
		ps[i].A0x = p.Ax
		ps[i].A0y = p.Ay
		ps[i].A0z = p.Az
	}
}

func DieIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
