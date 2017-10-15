package redisv1

import (
	"log"
	"strconv"
)

//Exists check for the existence of a key
func (plant *Redis) Exists(keys ...string) (string, error) {
	existCmd, err := fireCommand(plant, "EXISTS", keys...)
	if err != nil {
		log.Println("Server Error : ", err)
		return "0", err
	}
	return existCmd.(string), nil

}

//Del deletes the list of keys
func (plant *Redis) Del(keys ...string) (string, error) {
	delCmd, err := fireCommand(plant, "DEL", keys...)
	if err != nil {
		log.Println("Server Error : ", err)
		return "", err
	}
	return delCmd.(string), nil
}

//Expire sets a ttl for the Redis Key
func (plant *Redis) Expire(Key string, TTL int) (string, error) {
	expireCmd, err := fireCommand(plant, "EXPIRE", Key, strconv.Itoa(TTL))
	if err != nil {
		log.Println("Server Error : ", err)
		return "", err
	}
	return expireCmd.(string), nil
}
