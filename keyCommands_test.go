package redisv1

import (
	"testing"
)

func TestExistence(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	//Set a Key
	_ = client.Set("Key0", "I Exist")
	redisReply, err := client.Exists("Key0")
	if err != nil || redisReply == "0" {
		t.Error(err)
	}

	redisReply, err = client.Exists("NosuchKey")
	if err != nil || redisReply != "0" {
		t.Error(err)
	}
	_ = client.Set("Key1", "I too Exist")
	redisReply, err = client.Exists("Key0", "Key1", "NoSuchKey")
	//fmt.Println(redisReply)
	if err != nil || redisReply == "0" {
		t.Error(err)
	}

}
