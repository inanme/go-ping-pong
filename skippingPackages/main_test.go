package jojo

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if val, found := os.LookupEnv("HOME1"); found {
		fmt.Println(val, "skipping all")
	} else {
		m.Run()
	}
}
