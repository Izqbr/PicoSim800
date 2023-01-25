package main

import (
	"machine"
	"time"
	"PicoSim800/ds18b20"
)

var (
		Serial = machine.UART0
		SIM800 = machine.UART1
		MQTT_SERVER = "srv.clusterfly.ru"
		PORT = "9991"
		AT = ""
		)	




func main()  {
	Configure()


	dallas := ds18b20.NewDevice(machine.GP13)
	


	for  {
		//  if Connect_status != true {
		// 	SIM800.Write([]byte("AT\r\n"))  //посылаем в GSM модуль
		// }
		if SIM800.Buffered() > 0 {Resp_modem()}    // если что-то пришло от SIM800 отправляем в Raspberry для разбора                              
		if Serial.Buffered() > 0 {Resp_serial()}
		time.Sleep(time.Millisecond * 4000)
		
		dallas.Init()
		dallas.SendCommand(0x44)
		//dallas.SendCommand(SKIP_ROM)
		led2.High()
		time.Sleep(time.Millisecond * 2000)
		led2.Low()
		time.Sleep(time.Millisecond * 2000)
	}

} 