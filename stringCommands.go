package redisv1

import (
	"log"
	"strconv"
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
func (plant *Redis) Incr(key string) string {
	incrCmd, err := fireCommand(plant, "INCR", key)
	if err != nil {
		log.Println("Error Incrementing  the Key value  : ", err)
		return err.Error()
	}
	return incrCmd.(string)
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

func (plant *Redis) StrLen(key string) interface{} {
	lengthCmd, err := fireCommand(plant, "STRLEN", key)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return lengthCmd
}

func (plant *Redis) SetRange(key string, rangeVal int, value string) (interface{}, error) {
	setRangeCmd, err := fireCommand(plant, "SETRANGE", key, strconv.Itoa(rangeVal), value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return nil, err
	}
	return setRangeCmd, nil
}

func (plant *Redis) SetNx(key, value string) (string, error) {
	setNxCmd, err := fireCommand(plant, "SETNX", key, value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return "0", err
	}
	return setNxCmd.(string), nil
}

func (plant *Redis) SetEx(key string, expiryInSec int, value string) (string, error) {
	setExCmd, err := fireCommand(plant, "SETEX", key, strconv.Itoa(expiryInSec), value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return "", err
	}
	return setExCmd.(string), nil
}
