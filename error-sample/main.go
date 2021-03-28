package main

import "fmt"

type Stringer interface {
	String() string
}

type MyStr string

func (s MyStr) String() string {
	return "This is MyStr"
}

type MyError string

func (e MyError) Error() string {
	return fmt.Sprintf("エラー発生。エラー内容:%s", string(e))
}

func main() {
	var s1 MyStr = "good"
	if s, err := ToStringer(s1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	var s2 string = "bad"
	if s, err := ToStringer(s2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("キャスト失敗!!")
}
