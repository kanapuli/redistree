package redisv1

import "log"

//Exists check for the existence of a key
func (plant *Redis) Exists(keys ...string) (string, error) {
	existCmd, err := fireCommand(plant, "EXISTS", keys...)
	if err != nil {
		log.Println("Server Error : ", err)
		return "0", err
	}
	return existCmd.(string), nil

}
