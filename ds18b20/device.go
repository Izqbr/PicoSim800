package ds18b20

import (
	
	"machine"
	 "time"
)

type Device interface{
	
	Init() bool
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


