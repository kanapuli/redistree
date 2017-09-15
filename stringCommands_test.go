package redisv1

import (
	"testing"
)

func TestAppend(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
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
	defer client.Close()
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
	defer client.Close()
	redisReply := client.Set("Number", "10")
	if redisReply != "OK" {
		t.Error("Error setting the Value for the Key Number")
	}
}

func TestGet(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("b", "qwerty")
	redisReply := client.Get("b")

	if redisReply != "qwerty" {
		t.Error("Error getting the Value for the Key ")
	}

}

func TestIncr(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("b", "1")
	redisReply := client.Incr("b")

	if redisReply != "2" {
		t.Error("Error getting the Value for the Key ")
	}

}

func TestStrLen(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("b", "Writing length")
	redisReply := client.StrLen("b")

	if redisReply != "14" {
		t.Error("Error getting the Value for the Key ")
	}

}

func TestSetRange(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("key1", "Go is written by Jon Skeet")
	_, err = client.SetRange("key1", 17, "Rob Pike")
	if err != nil {
		t.Errorf("Expected to change Jon Skeet but %v", err)
	}

}

func TestSetNx(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	redisReply, err := client.SetNx("key1", "This Key exists already")
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if redisReply != "0" {
		//Since the key1 exists , no set operation is performed . Hence 0 should be the reply
		t.Errorf("SetNx Expected 0  but got %v", redisReply)
	}

}

func TestSetEx(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	redisReply, err := client.SetEx("key1", 10, "This Key should expire in 10 seconds")
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	if redisReply != "OK" {
		//Since the key1 exists , no set operation is performed . Hence 0 should be the reply
		t.Errorf("Expected OK  but got %v", redisReply)
	}

}

func TestSetBit(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	redisReply, err := client.SetBit("key1", 10, "1")
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	if redisReply != 1 {
		//Since the key1 exists , no set operation is performed . Hence 0 should be the reply
		t.Errorf("Expected OK  but got %v", redisReply)
	}
}

func TestMSet(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	redisReply, err := client.MSet("key1", "1", "key2", "2", "key3", "3")
	if err != nil {
		t.Errorf("Unexpected Error %v\n", err)
	}

	if redisReply != "OK" {
		//Since the key1 exists , no set operation is performed . Hence 0 should be the reply
		t.Errorf("Expected OK  but got %v", redisReply)
	}
}

func TestGetSet(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_ = client.Set("MyKey", "2")
	redisReply, err := client.GetSet("MyKey", "5")
	if err != nil {
		t.Errorf("Unexpected Error %v\n", err)
	}
	if redisReply != "2" {
		t.Errorf("Expected 2 but got %s\n", redisReply)
	}
	newSetValue := client.Get("MyKey")
	if newSetValue != "5" {
		t.Errorf("Expected 5 in GetSet but got %v\n", newSetValue)
	}
}

func TestMGet(t *testing.T) {
	client, err := Seed("127.0.0.1", "6379", "letmein", 3, 0)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	_, _ = client.MSet("key1", "1", "key2", "2", "key3", "3")
	_, err = client.MGet("key1", "key2", "key3")
	if err != nil {
		t.Errorf("Unexpected Error %v\n", err)
	}
	// for _, v := range redisReply {
	// 	fmt.Println(string(v))
	// }
}
