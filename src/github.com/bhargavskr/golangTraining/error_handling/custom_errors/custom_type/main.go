package main

import (
	"fmt"
	"log"
)

type MathError struct {
	lat, lon string
	err      error
}

func (n *MathError) Error() string {
	return fmt.Sprintf("Math Error Occured %v %v %v \n", n.lat, n.lon, n.err)
}

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		var ErrMsg = fmt.Errorf("square root of negative number %v", f)
		return 0, &MathError{"8397498.8N", "87398.44S", ErrMsg}
	}
	return 42, nil
}
