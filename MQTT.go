package main

import (
//	"strings"
//	"machine"
	// "strconv"
	"time"
)

var (
	MQTT_type string = "MQIsdp"
	MQTT_CID string = "Lada"
	MQTT_user string = "user_xxxxxxxx"
	MQTT_pass string = "pass_xxxxxxxx"
	broker bool = false
)

func MQTT_CONNECT(){
	SIM800.Write([]byte("AT+CIPSEND\r\n")) //Отправить данные
	time.Sleep(time.Millisecond * 100)

	SIM800.WriteByte(0x10)  // маркер пакета на установку соединения

	SIM800.WriteByte(byte(len(MQTT_type)+len(MQTT_CID)+len(MQTT_user)+len(MQTT_pass)+12))

	SIM800.WriteByte(0)
	SIM800.WriteByte(byte(len(MQTT_type)))
	SIM800.Write([]byte(MQTT_type))  // тип протокола
	SIM800.WriteByte(0x03)
	SIM800.WriteByte(0xC2)

	SIM800.WriteByte(0)
	SIM800.WriteByte(0x3C) // просто так нужно

	SIM800.WriteByte(0)
	SIM800.WriteByte(byte(len(MQTT_CID)))
	SIM800.Write([]byte(MQTT_CID))  // MQTT  идентификатор устройства

	SIM800.WriteByte(0)
	SIM800.WriteByte(byte(len(MQTT_user)))
	SIM800.Write([]byte(MQTT_user)) // MQTT логин

	SIM800.WriteByte(0)
	SIM800.WriteByte(byte(len(MQTT_pass)))
	SIM800.Write([]byte(MQTT_pass)) // MQTT пароль
	
  
	MQTT_PUB ("user_f73fd7c4/C5/status", "Подключено");                                            // пакет публикации
	MQTT_SUB ("user_f73fd7c4/C5/comand");                                                          // пакет подписки на присылаемые команды
	MQTT_SUB ("user_f73fd7c4/C5/settimer");                                                        // пакет подписки на присылаемые значения таймера
	
	SIM800.WriteByte(0x1A)
	broker = true
}


func MQTT_SUB(MQTT_topic string){

}



func MQTT_PUB(MQTT_topic string, MQTT_messege string){

}


