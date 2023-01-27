package main

import (
	"PicoSim800/buzzer"
	"PicoSim800/ds18b20"
	"machine"
	"time"
)

var (
	Serial      = machine.UART0
	SIM800      = machine.UART1
	MQTT_SERVER = "srv.clusterfly.ru"
	PORT        = "9991"
	AT          = ""
)

func main() {
	Configure()

	dat := ds18b20.NewDevice(machine.GP2)
	buzzer := buzzer.NewBuzzer(machine.GP15)
	for {
		//  if Connect_status != true {
		// 	SIM800.Write([]byte("AT\r\n"))  //посылаем в GSM модуль
		// }
		if SIM800.Buffered() > 0 {
			Resp_modem()
		} // если что-то пришло от SIM800 отправляем в Raspberry для разбора
		if Serial.Buffered() > 0 {
			Resp_serial()
		}
		time.Sleep(time.Millisecond * 1000)
		buzzer.Beep(200,2)
		 T := dat.GetTemp()
		 println(T)
		// dat.Init()
	// dat.SendCommand(0xcc)
	// dat.SendCommand(0x44)
	// time.Sleep(time.Millisecond * 750)
	// dat.Init()
	// dat.SendCommand(0xcc)
	// dat.SendCommand(0xbe)
	
	// lbt := uint16(dat.Readbyte())
	// 	println(lbt)
	// hbt := uint16(dat.Readbyte())
	// 	println(hbt)
	
		led.High()
		time.Sleep(time.Millisecond * 750)
		led.Low()
		time.Sleep(time.Millisecond * 750)
	}

}
