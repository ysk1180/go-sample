package main

import "fmt"

type Stringer interface {
	String() string
}

type MyStr string

func (s MyStr) String() string {
	return "This is MyStr"
}

type MyInt int

func (i MyInt) String() string {
	return "This is MyInt"
}

type MyFloat float64

func (f MyFloat) String() string {
	return "This is MyFloat"
}

func main() {
	var str MyStr = "Cat"
	PutTypeAndValue(str)

	var i MyInt = 123
	PutTypeAndValue(i)

	var f MyFloat = 4.32
	PutTypeAndValue(f)
}

func PutTypeAndValue(s Stringer) {
	fmt.Println(s)

	switch v := s.(type) {
	case MyStr:
		fmt.Printf("type: %v, value: %s\n", "MyString", string(v))
	case MyInt:
		fmt.Printf("type: %v, value: %d\n", "MyInt", int(v))
	case MyFloat:
		fmt.Printf("type: %v, value: %g\n", "MyFloat", float64(v))
	default:
		fmt.Println("Nothing")
	}

	fmt.Println("---")
}
