package redisv1

import (
	"log"
	"strconv"
)

//Append appends the value to the existing value
func (plant *Redis) Append(key, value string) string {
	appendCmd, err := fireCommand(plant, "APPEND", key, value)
	if err != nil {
		log.Println("Error Appending to the Key  : ", err)
	}
	return appendCmd.(string)
}

//Decr decrements the value
func (plant *Redis) Decr(key string) string {
	decrCmd, err := fireCommand(plant, "DECR", key)
	if err != nil {
		log.Println("Error Decrementing the Key value  : ", err)
		return err.Error()
	}
	return decrCmd.(string)
}

//Incr increments the value
func (plant *Redis) Incr(key string) string {
	incrCmd, err := fireCommand(plant, "INCR", key)
	if err != nil {
		log.Println("Error Incrementing  the Key value  : ", err)
		return err.Error()
	}
	return incrCmd.(string)
}

//Set sets the value for the key
func (plant *Redis) Set(key, value string) string {
	setCmd, err := fireCommand(plant, "SET", key, value)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return setCmd.(string)
}

//Get gets the value of the key
func (plant *Redis) Get(key string) interface{} {
	getCmd, err := fireCommand(plant, "GET", key)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return getCmd
}

//StrLen gives the length of value
func (plant *Redis) StrLen(key string) interface{} {
	lengthCmd, err := fireCommand(plant, "STRLEN", key)
	if err != nil {
		log.Println("Error Setting the Key value  : ", err)
		return err.Error()
	}
	return lengthCmd
}

//SetRange sets the exisiting value from the offset
func (plant *Redis) SetRange(key string, rangeVal int, value string) (interface{}, error) {
	setRangeCmd, err := fireCommand(plant, "SETRANGE", key, strconv.Itoa(rangeVal), value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return nil, err
	}
	return setRangeCmd, nil
}

//SetNx sets the value ifkey doesn't exist
func (plant *Redis) SetNx(key, value string) (string, error) {
	setNxCmd, err := fireCommand(plant, "SETNX", key, value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return "0", err
	}
	return setNxCmd.(string), nil
}

//SetEx sets the value with the expiration time
func (plant *Redis) SetEx(key string, expiryInSec int, value string) (string, error) {
	setExCmd, err := fireCommand(plant, "SETEX", key, strconv.Itoa(expiryInSec), value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return "", err
	}
	return setExCmd.(string), nil
}

//SetBit sets the Bit for the key from the offset
func (plant *Redis) SetBit(key string, offset int, value string) (int, error) {
	setBitCmd, err := fireCommand(plant, "SETBIT", key, strconv.Itoa(offset), value)
	if err != nil {
		log.Println("Server Error  : ", err)
		return 0, err
	}
	setBit, err := strconv.Atoi(setBitCmd.(string))
	if err != nil {
		log.Println("Server Error  : ", err)
		return 0, err
	}
	return setBit, nil
}

//MSet sets multiple key values
func (plant *Redis) MSet(args ...string) (string, error) {
	mSetCmd, err := fireCommand(plant, "MSET", args...)
	if err != nil {
		log.Println("Server Error  : ", err)
		return "", err
	}

	return mSetCmd.(string), nil
}

//GetSet gets the key's value and set the key with the new value
func (plant *Redis) GetSet(key, value string) (string, error) {
	getSetCmd, err := fireCommand(plant, "GETSET", key, value)
	if err != nil {
		log.Println("Server Error : ", err)
		return "", err
	}
	return getSetCmd.(string), nil
}

//MGet gets values for multiple keys
func (plant *Redis) MGet(args ...string) ([][]byte, error) {
	mGetCmd, err := fireCommand(plant, "MGET", args...)
	if err != nil {
		log.Println("Server Error  : ", err)
		return nil, err
	}

	return mGetCmd.([][]byte), nil
}
