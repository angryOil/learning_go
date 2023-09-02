package fir

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func testMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("testFirst uses stuff set up in TestMain time:", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond also uses stuff set up in testMain time:", testTime)
}
