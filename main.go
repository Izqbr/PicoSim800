package main

import (
	"machine"
	"time"
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
	
	for  {
		//  if Connect_status != true {
		// 	SIM800.Write([]byte("AT\r\n"))  //посылаем в GSM модуль
		// }
		if SIM800.Buffered() > 0 {Resp_modem()}    // если что-то пришло от SIM800 отправляем в Raspberry для разбора                              
		if Serial.Buffered() > 0 {Resp_serial()}
		time.Sleep(time.Millisecond * 1000)
		
	}

} 


		



