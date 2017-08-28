package redisv1

import "log"

func (plant *Redis) Append(key, value string) int {
	appendCmd, err := fireCommand(plant, "APPEND", key, value)
	if err != nil {
		log.Println("Error Appending to the Key  : ", err)
	}
	return appendCmd.(int)
}
