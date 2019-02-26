package main
import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934


type car struct {
	//unsigned with u.. starts at 0 rather than -32k
	gas_pedal uint16 // 0 to 65535
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64

	//value receivers vs pointer receivers
}

//value receivers just receives values; doesnt change values of car
//first parenthesis attaches the method to car c
// kmh is method name
//value receiver is safer. You can do all kinds of things with it
// and nothing will ever permanently effect it.
//VALUE RECEIVER
func(c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

//can change values and does so with * pointer.
//Makes lasting changes to the struct
//POINTER RECEIVER
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	//var a_car car
	a_car := car{gas_pedal: 2234,
				 brake_pedal: 0,
				 steering_wheel: 0, 
				 top_speed_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(400.0)
	fmt.Println(a_car.kmh())
}