package main

import (
		"strings"
		"time"
)
 
var (
	APN string = "internet.beeline.ru"            // точка доступа выхода в интернет вашего сотового оператора
	Connect_status bool = false
	com uint8 = 0
 )

func Resp_serial(){
	size := Serial.Buffered()
	data := make([]byte,size)
	for i := 0; i< len(data);i++ {
		data[i],_ = Serial.ReadByte()
	}
	AT2 := string(data)
	SIM800.Write([]byte(AT2+"\r\n"))
	AT2 = ""
}

func Resp_modem(){
	
	size := SIM800.Buffered()
	data := make([]byte,size)
	for i := 0; i< len(data);i++ {
		data[i],_ = SIM800.ReadByte()
	}
	
	AT = string(data)
	Serial.Write([]byte(AT+"\r\n")) 
	

	if strings.Compare(string(AT), "\r\nOK\r\n") == 0 && com == 0 && Connect_status == false {
		comconnect = true
		com = 1
		SIM800.Write([]byte("AT+CSTT=\"internet.beeline.ru\",\"beeline\",\"beeline\"\r\n"))  
		time.Sleep(time.Millisecond * 200)
		println(AT)
		println("send command  1")
        
	} else if strings.Compare(string(AT), "\r\nOK\r\n") == 0 && com == 1 && Connect_status == false {
		com = 2 
		SIM800.Write([]byte("AT+CIICR\r\n"))  //настройка контекста
		time.Sleep(time.Millisecond * 200)
		println("send command  2")

	} else if strings.Compare(string(AT), "\r\nOK\r\n") == 0 && com == 2 && Connect_status == false {
		com = 3
		SIM800.Write([]byte("AT+CIFSR\r\n"))  //запрос IP
		time.Sleep(time.Millisecond * 200)
		println("send command  3") 

	} else if strings.Count(string(AT), "0.0.0.0") != 1 && com == 3 && Connect_status == false {
		com = 4
		SIM800.Write([]byte("AT+CIPSTART=\"TCP\",\"srv2.clusterfly.ru\",\"9991\"\r\n"))
		time.Sleep(time.Millisecond * 1000)
		println("send command  4")
	
	} else if strings.Compare(string(AT), "\r\nCONNECT OK\r\n") == 0 {
		led.High()
		Connect_status = true
		MQTT_CONNECT()

	} else if strings.Compare(string(AT), "\r\nCLOSED\r\n") == 0 {
		comconnect = false
		led.Low()
		Connect_status = false
		SIM800.Write([]byte("AT+CIPSTART=\"TCP\",\"srv2.clusterfly.ru\",\"9991\"\r\n"))

	} else if strings.Compare(string(AT), "\r\nCONNECT FAIL\r\n") == 0 {
		led.Low()
		SIM800.Write([]byte("AT+CFUN=1,1\r\n"))  //посылаем в GSM модуль 

	} else if strings.Compare(string(AT), "\r\nALREADY CONNECT\r\n") == 0 {
		led.Low()
		SIM800.Write([]byte("AT+CFUN=1,1\r\n"))  //посылаем в GSM модуль 

	} else if strings.Count(string(AT), "+CME ERROR:") == 1 {
		com = 0
		SIM800.Write([]byte("AT+CFUN=1,1\r\n"))  //
		time.Sleep(time.Millisecond * 3000)
		SIM800.Write([]byte("AT\r\n"))  //посылаем в
		println(AT)
		println("SIM800 is reset")
	}

	AT = ""
}
