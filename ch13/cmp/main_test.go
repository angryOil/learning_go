package cmp

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "JOY",
		Age:  28,
	}
	result := CreatePerson("JOY", 28)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

var comparer = cmp.Comparer(func(x, y Person) bool {
	return x.Name == y.Name && x.Age == y.Age
})

func TestCreatePerson2(t *testing.T) {
	expected := Person{
		Name: "JOY",
		Age:  28,
	}
	result := CreatePerson("JOY", 28)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
