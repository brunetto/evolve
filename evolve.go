package evolve 

import (
	"fmt"
	"log"
	"os"
)

const Gmks float64 = 6.67384e-11 // m^3 kg^-1 s^-2

const ParticleFormat string = "%e,%e,%e,%e,%e,%e,%e"

type Particle struct {
	Mass float64
	Pos [3]float64 // should they be a slice instead of an array?
	Vel [3]float64
	Acc [3]float64
}

func (p *Particle) ReadFromLine (line string) (err error) {
	// Read data from line
	if _, err = fmt.Sscanf(line, ParticleFormat,
		&(p.Mass), &(p.Pos[0]), &(p.Pos[1]), &(p.Pos[2]), &(p.Vel[0]), &(p.Vel[1]), &(p.Vel[2])); err != nil {
		return err
	}
	return nil
}

func (p *Particle) Format () (particle string) {
	// Read data from line
	particle = fmt.Sprintf(ParticleFormat, p.Mass, p.Pos[0], p.Pos[1], p.Pos[2], p.Vel[0], p.Vel[1], p.Vel[2])
	return particle
}

func (p *Particle) Print () () {
	fmt.Println(p.Format())
}

func (p *Particle) PrettyPrint () () {
	fmt.Printf("Mass:\t%v\nPos:\t%v\nVel:\t%v\n", p.Mass, p.Pos, p.Vel)
}

type System struct {
	Particles []*Particle
}

func (s *System) LoadFromFile (inFileName string) (err error) {
	// Read data from file
	var (
		particles = []*Particle{}
		p *Particle
		inFile *os.File
	)
	
	if inFile, err = os.Open("test.dat"); err != nil {
		log.Fatalf("Error while opening %v file: %v \n", inFileName, err)
	}
	
	for {
		p = &Particle{}
		_, err = fmt.Fscanf(inFile, ParticleFormat+"\n",
			&(p.Mass), 
			&(p.Pos[0]), &(p.Pos[1]), &(p.Pos[2]), 
			&(p.Vel[0]), &(p.Vel[1]), &(p.Vel[2]))
		
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}
			break
		}
		particles = append(particles, p)
	}
	s.Particles = particles
	return nil
}

func (s *System) Print () () {
	for _, p := range(s.Particles) {
		p.Print()
	}
}



