package redisv1

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Redis struct {
	connection net.Conn
}

//Seed - Create a new connection of redis
func Seed(address, port, password string, timeout int, db int) (*Redis, error) {

	defaultTimeout, _ := time.ParseDuration(fmt.Sprintf("%ds", timeout))

	hostAddress := fmt.Sprintf("%s:%s", address, port)
	con, err := net.DialTimeout("tcp", hostAddress, defaultTimeout)
	if err != nil {
		log.Println("Error Connecting to Redis : ", err)
		return nil, err
	}
	if password != "" {
		cmd := fmt.Sprintf("AUTH %s\r\n", password)

		_, err := sendCo2(con, []byte(cmd))
		if err != nil {
			log.Println("Error Authenticating to Redis : ", err)
			return nil, err
		}

	}

	if db != 0 {
		defaultDb := fmt.Sprintf("SELECT %d\r\n", db)
		_, err := sendCo2(con, []byte(defaultDb))
		if err != nil {
			log.Println("Error Selecting Database to Redis : ", err)
			return nil, err
		}

	}
	client := new(Redis)
	client.connection = con
	return client, nil
}

//Ping - pings the Redis server
func (plant *Redis) Ping() string {
	pingCmd := fmt.Sprintf("PING\r\n")
	pong, err := sendCo2(plant.connection, []byte(pingCmd))
	if err != nil {
		log.Println("Error Pinging to Redis : ", err)
		return "-ERR"
	}
	return pong.(string)
}

//Echo - Echoes back the message
func (plant *Redis) Echo(message string) string {
	if message == "" {
		return "Cannot Set Empty String"
	}
	echoCmd := fmt.Sprintf("ECHO \"%s\"\r\n", message)
	echo, err := sendCo2(plant.connection, []byte(echoCmd))
	if err != nil {
		log.Println("Error Echoing From Redis : ", err)
		return ""
	}
	return echo.(string)
}

//Quit - closes the Redis server connection
func (plant *Redis) Quit() string {
	quitCmd := fmt.Sprint("QUIT\r\n")
	quitOk, err := sendCo2(plant.connection, []byte(quitCmd))
	if err != nil {
		log.Println("Error Quitting the Connection: ", err)
		return ""
	}
	if plant.connection != nil {
		plant.connection.Close()
		plant.connection = nil
	}
	return quitOk.(string)
}
