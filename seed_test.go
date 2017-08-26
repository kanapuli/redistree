package redisv1

import (
	"fmt"
	"testing"
)

func TestSeed(t *testing.T) {
	_, err := Seed("127.0.0.1", "6379", "letmein", 3, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestSeedErr(t *testing.T) {
	_, err := Seed("127.0.0.0", "6379", "letmein", 3, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestSeedDbErr(t *testing.T) {
	//default number of Db is 15 but 25 is specified which should result in error
	_, err := Seed("127.0.0.1", "6379", "letmein", 3, 25)
	if err != nil {
		t.Error(err)
	}
}
func TestSeedPasswordErr(t *testing.T) {
	_, err := Seed("127.0.0.1", "6379", "wrongPassword", 3, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestSeedPing(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	pong := client.Ping()
	if pong != "PONG" {
		t.Fail()
	}
	fmt.Println("ping status ", pong)
}
