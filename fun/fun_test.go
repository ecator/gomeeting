package fun_test

import (
	"fmt"
	"testing"

	"github.com/ecator/gomeeting/fun"
)

type people struct {
	Name string
	Age  uint32
}

type student struct {
	Name  string
	Score int
}

func TestGetByName(t *testing.T) {
	p := new(people)
	p.Age = 123
	p.Name = "martin"
	fmt.Println(fun.GetStrByName(p, "Name"))
	fmt.Println(fun.GetStrByName(p, "Name1"))
	fun.SetStrByName(p, "Name", "mark")
	fun.SetStrByName(p, "Name1", "mark1")
	fmt.Println(fun.GetStrByName(p, "Name"))
	fmt.Println(fun.GetUint32ByName(p, "Age"))
	fmt.Println(fun.GetUint32ByName(p, "Age1"))
	fun.SetUint32ByName(p, "Age", 32)
	fun.SetUint32ByName(p, "Age1", 32)
	fmt.Println(fun.GetUint32ByName(p, "Age"))

}

func TestSetByObj(t *testing.T) {
	p := new(people)
	s := new(people)
	s.Age = 123
	s.Name = "aaaa"
	fmt.Println(p)
	fun.SetByObj(p, s)
	fmt.Println(p)
}
