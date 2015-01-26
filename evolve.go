package main 

import (
	"fmt"
	"log"
	"os"
	
)

func main () () {
	var (
	)


}


const BodyFormat string = "%e,%e,%e,%e,%e,%e,%e"

type Body struct {
	Mass float64
	Pos [3]float64 // should they be a slice instead of an array?
	Vel [3]float64
}

func (b *Body) ReadFromLine (line string) (err error) {
	// Read data from line
	if _, err = fmt.Sscanf(line, BodyFormat,
		&(b.Mass), &(b.Pos[0]), &(b.Pos[1]), &(b.Pos[2]), &(b.Vel[0]), &(b.Vel[1]), &(b.Vel[2])); err != nil {
		return err
	}
	return nil
}

func (b *Body) Format () (body string) {
	// Read data from line
	body = fmt.Sprintf(BodyFormat, b.Mass, b.Pos[0], b.Pos[1], b.Pos[2], b.Vel[0], b.Vel[1], b.Vel[2])
	return body
}

func (b *Body) Print () () {
	fmt.Println(b.Format())
}


type System struct {
	Bodies []*Body
}

func (s *System) LoadFromFile (inFileName string) (err error) {
	// Read data from file
	var (
		bodies = []*Body{}
		b *Body
		inFile *os.File
	)
	
	if inFile, err = os.Open("test.dat"); err != nil {
		log.Fatalf("Error while opening %v file: %v \n", inFileName, err)
	}
	
	for {
		b = &Body{}
		_, err = fmt.Fscanf(inFile, BodyFormat+"\n",
			&(b.Mass), 
			&(b.Pos[0]), &(b.Pos[1]), &(b.Pos[2]), 
			&(b.Vel[0]), &(b.Vel[1]), &(b.Vel[2]))
		
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}
			break
		}
		bodies = append(bodies, b)
	}
	s.Bodies = bodies
	return nil
}

func (s *System) Print () () {
	for _, p := range(s.Bodies) {
		p.Print()
	}
}



