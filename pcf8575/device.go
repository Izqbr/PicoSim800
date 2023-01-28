package pcf8575

import (
	"machine"
)


type Device interface {
	Configure(bus machine.I2C,sda machine.Pin, scl machine.Pin )
	ReadRegister(addr uint8, r uint8) 
	WriteRegister(addr uint8, r uint8) 
}
type Pins uint16

// // Set sets the value for the given pin.
// func (p *Pins) Set(pin int, value bool) {
// 	if value {
// 		p.High(pin)
// 	} else {
// 		p.Low(pin)
// 	}
// }  

type device struct {
	bus  machine.I2C
	addr uint8
	pins Pins
}




