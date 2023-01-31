package pcf8575

import (
	"machine"
)
var dataOut uint8 = 0

const(
	P0 uint8 = iota
	P1
	P2
	P3
	P4
	P5
	P6
	P7
)	
const (
	P10 uint8 = iota
	P11
	P12
	P13
	P14
	P15
	P16
	P17
)
	 


type Pcf8575 interface {
	Configure(i2c machine.I2C, sda machine.Pin, scl machine.Pin)
	Setpin(pin uint8,value bool) 
	Writeport(dataOut uint8)
 
}
  

type Port struct {
	bus  machine.I2C
	addr uint8
	
}
func NewI2C(bus machine.I2C, addr uint8) Pcf8575 {
	return &Port{
		bus: bus,
		addr: addr,
	}
}

func(device Port) Configure(i2c machine.I2C, sda machine.Pin, scl machine.Pin)  {
	
	device.bus = i2c
	device.bus.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		SDA: sda,
		SCL: scl,
	})
}

func(device Port) Writeport(dataOut uint8) {
	
	//device.Writeport(dataOut)
}


func(device Port) Setpin(pin uint8,value bool) {
	if value == false {
		dataOut &= ^(1<<pin)
	} else {
		dataOut |= (1<<pin)
	}
	//device.Writeport(dataOut)
}

// Setpin(P17,true)