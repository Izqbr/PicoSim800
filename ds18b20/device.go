package ds18b20

import (
	
	"machine"
	 "time"
)

type Device interface{
	
	Init() bool
	SendCommand(command uint8)
	SendBit(bit uint8)
}

type device struct {
	load machine.Pin
}

func NewDevice(load machine.Pin) Device {
	return &device {
		load: load,
	}
}


func (d *device) Init() bool{
	d.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.load.Set(false) //низкий уровень
	time.Sleep(time.Microsecond * 485)
	d.load.Set(true) // высокий уровень
	time.Sleep(time.Microsecond * 65)
	d.load.Configure(machine.PinConfig{Mode: machine.PinInput})
	status := d.load.Get()
	time.Sleep(time.Microsecond * 500)
	return status
}

func (d *device)SendCommand(command uint8) {
	for i:=0;i<8;i++ {
		d.SendBit(command >> i & 1)
		time.Sleep(time.Microsecond * 5)
} 
	}
func(d *device) SendBit(bit uint8) {
	d.load.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.load.Set(false)
	if bit==1 {
		time.Sleep(time.Microsecond * 2)
	} else {time.Sleep(time.Microsecond * 65) } 
	d.load.Set(true)	  
	if bit == 1 {
		time.Sleep(time.Microsecond * 65)
	} else {time.Sleep(time.Microsecond * 2) }	
}


