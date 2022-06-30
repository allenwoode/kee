package main

import (
	"testing"
)

func TestHello(t *tesing.T) {
	t.Log("hello world")
}


func TestWorld(t *testing.T) {
	t.Logf("%s", "hello world")
}
