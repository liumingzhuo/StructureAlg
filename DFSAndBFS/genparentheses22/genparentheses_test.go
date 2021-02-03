package genparentheses22

import (
	"log"
	"reflect"
	"testing"
)

func TestGenparent(t *testing.T) {
	got := generateParenthesis(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(got, want) {
		log.Fatalf("except, got:%v, want:%v", got, want)
	}
}

func TestGenparent1(t *testing.T) {
	got := generateParenthesis(1)
	want := []string{"()"}
	if !reflect.DeepEqual(got, want) {
		log.Fatalf("except, got:%v, want:%v", got, want)
	}
}
