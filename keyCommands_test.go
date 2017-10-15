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

func TestDelete(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("key1", "Please delete me")
	redisReply, err := client.Del("key1")
	if err != nil || redisReply != "1" {
		t.Error(err)
	}
	keyExists, _ := client.Exists("key1")
	if keyExists != "0" {
		t.Error("Expected the Key to delete but still it exists")
	}
	_ = client.Set("key1", "Please delete me 1")
	_ = client.Set("key2", "Please delete me 2")
	redisReply, err = client.Del("key1", "key2")
	if err != nil || redisReply != "2" {
		t.Error(err)
	}
	keyExists, _ = client.Exists("key1", "key2")
	if keyExists != "0" {
		t.Error("Expected the Key to delete but still it exists")
	}
}

func TestDeleteNegative(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	r, err := client.Del("1234 2323")

	if err == nil && r != "0" {
		t.Error("Expected Error in Delete Cmd but error is nil")
	}
}

func TestExpiry(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("key1", "Expire me in 100 ms")
	redisReply, err := client.Expire("Key1", 100)
	if err != nil {
		t.Errorf("Expected theKey to expire but got error %v", err)
	}
	if redisReply != "1" {
		t.Errorf("Expected 1 while expiring but got %v", redisReply)
	}
}
