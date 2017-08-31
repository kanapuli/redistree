package redisv1

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

func composeError(message string) error {
	err := errors.New(message)
	return err
}

//sendCO2 - Send RawBytes to the Redis Server
func sendCo2(c net.Conn, cmd []byte) (interface{}, error) {
	_, err := c.Write(cmd)
	if err != nil {
		log.Println("Error in Sending Raw Bytes to the Redis Server: ", err)
		return nil, err
	}
	//Read the Response from the Redis
	reader := bufio.NewReader(c)
	response, err := getOxygen(reader)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//getOxygen - Gets the response back from the Redis server when the sendCo2 method is called
func getOxygen(reader *bufio.Reader) (interface{}, error) {
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		if len(line) == 0 || err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			break
		}
	}
	switch line[0] {
	case '+':
		return line[1:], nil
	case '-':
		//Slice starts from 5 because the first four chars are "-ERR "
		return nil, composeError(line[5:])
	case ':':
		return line[1:], nil
	case '*':
		return line[1:], nil
	case '$':
		byteSize, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			return nil, err
		}
		if byteSize == -1 {
			return nil, composeError("Key has no value")
		}
		lineReader := io.LimitReader(reader, int64(byteSize))
		data, err := ioutil.ReadAll(lineReader)
		if err != nil {
			return nil, err
		}
		//data, err = reader.ReadString('\n')
		return string(data), nil
	}
	return nil, composeError("Redis server did not reply")
}

func fireCommand(plant *Redis, cmd string, args ...string) (data interface{}, err error) {
	var b []byte
	b = composeCommandsBytes(cmd, args...)
	_, err = plant.connection.Write(b)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(plant.connection)
	data, err = getOxygen(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func composeCommandsBytes(cmd string, args ...string) []byte {
	var bufferCmd bytes.Buffer
	fmt.Fprintf(&bufferCmd, "*%d\r\n$%d\r\n%s\r\n", len(args)+1, len(cmd), cmd) // len(args)+1 is used because the cmd is also added to length
	for _, arg := range args {
		fmt.Fprintf(&bufferCmd, "$%d\r\n%s\r\n", len(arg), arg)
	}
	return bufferCmd.Bytes()
}
