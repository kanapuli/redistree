package redisv1

import (
	"testing"
)

func TestAppend(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	redisReply := client.Append("a", "Jha Athavan")

	if redisReply < "0" {
		t.Error("Expected Integer greater then zero")
	}
}

func TestDecr(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	redisReply := client.Decr("Jack") // A key which doesnt exist. hence the reply should be 0

	if redisReply >= "0" {
		t.Error("Expected Value less than 0 for a Key which doesn't exists")
	}
	_ = client.Set("Number", "9")
	redisReply = client.Decr("Number") // A key which doesnt exist. hence the reply should be 0

	if redisReply != "8" {
		t.Error("Expected Value is 9 for a Key Number")
	}
}

func TestSet(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	redisReply := client.Set("Number", "10")
	if redisReply != "OK" {
		t.Error("Error setting the Value for the Key Number")
	}
}
