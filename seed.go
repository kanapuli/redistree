package redisv1

import (
	"fmt"
	"log"
	"net"
	"time"
)

//Seed - Create a new connection of redis
func Seed(address, port, password string, timeout int) (net.Conn, error) {

	defaultTimeout, _ := time.ParseDuration(fmt.Sprintf("%ds", timeout))

	hostAddress := fmt.Sprintf("%s:%s", address, port)
	con, err := net.DialTimeout("tcp", hostAddress, defaultTimeout)
	if err != nil {
		log.Println("Error Connecting to Redis : ", err)
		return nil, err
	}
	if password != "" {
		cmd := fmt.Sprintf("AUTH %s\r\n", password)

		isAuthenticated, err := sendCo2(con, []byte(cmd))
		if err != nil {
			log.Println("Error Authenticating to Redis : ", err)
			return nil, err
		}
		fmt.Println("Authentication Status: ", isAuthenticated)
	}
	return con, nil
}
