package redisv1

import "testing"

func TestSeed(t *testing.T) {
	_, err := Seed("127.0.0.1", "6379", "athavan", 3)
	if err != nil {
		t.Error(err)
	}
}

func TestSeedErr(t *testing.T) {
	_, err := Seed("127.0.0.0", "6379", "Letmein", 3)
	if err != nil {
		t.Error(err)
	}
}
