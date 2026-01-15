package main

import "fmt"

type wheel struct {
	radios   int
	material string
}
type car struct {
	model      string
	milage     float64
	brand      string
	doors      int
	frontWheel wheel
	backWheel  wheel
}

func (c car) start() {
	fmt.Println(c.model + " is Starting ")
}
func main() {

	myCar := car{
		model:  "Audi V3",
		milage: 15,
		brand:  "Audi",
		doors:  4,
		frontWheel: wheel{
			radios:   12,
			material: "gold",
		},
		backWheel: wheel{
			radios:   45,
			material: "silver",
		},
	}

	// fmt.Println(myCar)
	myCar.start()

}
