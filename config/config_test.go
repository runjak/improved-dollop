package config

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	//Testing if FromJson(…ToJson()) works:
	config1 := EmptyConfig()
	config2, err := FromJson(config1.ToJson())
	if err != nil {
		t.Errorf("TestJson case 1.1 had an error:\n%s\n", err)
	}
	if !reflect.DeepEqual(config1, config2) {
		t.Errorf("TestJson case 1.2 had an error:\n\t%v\n\t%v\n", config1, config2)
	}
}

func TestReadWrite(t *testing.T) {
	//Building a random path:
	//https://stackoverflow.com/a/31832326/448591
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randStringRunes := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}
	path := "/tmp/" + randStringRunes(16) + ".json"
	//Testing if ReadFile(WriteFile(…)) works:
	fmt.Printf("Creating random file: %s\n", path)
	config1 := EmptyConfig()
	err := config1.WriteFile(path)
	if err != nil {
		t.Errorf("TestReadWrite case 1.1 had an error:\n%s\n", err)
	}
	config2, err := ReadFile(path)
	if err != nil {
		t.Errorf("TestReadWrite case 1.2 had an error:\n%s\n", err)
	}
	if !reflect.DeepEqual(config1, config2) {
		t.Errorf("TestReadWrite case 1.3 had an error:\n\t%v\n\t%v\n", config1, config2)
	}
	fmt.Printf("Removing random file: %s\n", path)
	err = os.Remove(path)
	if err != nil {
		t.Errorf("TestReadWrite case 1.4 had an error:\n%s\n", err)
	}
}
