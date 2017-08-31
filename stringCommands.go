package redisv1

import (
	"log"
)

func (plant *Redis) Append(key, value string) string {
	appendCmd, err := fireCommand(plant, "APPEND", key, value)
	if err != nil {
		log.Println("Error Appending to the Key  : ", err)
	}
	return appendCmd.(string)
}

func (plant *Redis) Decr(key string) string {
	decrCmd, err := fireCommand(plant, "DECR", key)
	if err != nil {
		log.Println("Error Decrementing the Key value  : ", err)
		return err.Error()
	}
	return decrCmd.(string)
}

func (plant *Redis) Set(key, value string) string {
	setCmd, err := fireCommand(plant, "SET", key, value)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return setCmd.(string)
}

func (plant *Redis) Get(key string) interface{} {
	getCmd, err := fireCommand(plant, "GET", key)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return getCmd
}
