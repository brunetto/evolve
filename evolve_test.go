package evolve

import (
	"testing"
)

func Test_ParticleReadFromLine(t *testing.T){
	var (
		p = &Particle{}
		check = &Particle{
					Mass: 1.000000e+01,
					Pos: [3]float64{2.000000e+02, 3.000000e+03, 4.000000e+04},
					Vel: [3]float64{5.000000e+05, 6.000000e+06, 7.000000e+07},
				}
	)
	
	err := p.ReadFromLine("1.000000e+01,2.000000e+02,3.000000e+03,4.000000e+04,5.000000e+05,6.000000e+06,7.000000e+07")
	
	if err != nil {
		t.Error("Error reading from line.")
	}
		
	if *p != *check {
		t.Errorf("Found difference between: \n%v and \n%v\n", 
					 p.Format(), check.Format())
	}
	
}

func Test_SystemLoadFromFile(t *testing.T){
	
	var (
		err error
		system = &System{}
		check = &System{
			Particles: []*Particle{
				&Particle{
					Mass: 1.000000e+01,
					Pos: [3]float64{2.000000e+02, 3.000000e+03, 4.000000e+04},
					Vel: [3]float64{5.000000e+05, 6.000000e+06, 7.000000e+07},
				},
				&Particle{
					Mass: 1.000000e+01,
					Pos: [3]float64{2.000000e+02, -3.000000e+03, 4.000000e+04},
					Vel: [3]float64{5.000000e+05, 6.000000e+06, 7.000000e+07},
				},
			},
		}
	)
	
	if err = system.LoadFromFile("test-data.test"); err != nil {
		t.Error("Error loading bodies: ", err)
	}
	
	for idx,p := range system.Particles {
		if *p != *(check.Particles[idx]) {
			t.Errorf("At idx %v found difference between: \n%+v and \n%+v\n", 
					 idx, p.Format(), check.Particles[idx].Format())
		}
	}

}


