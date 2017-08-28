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

	if redisReply < 0 {
		t.Error("Expected Integer greater then zero")
	}
}
