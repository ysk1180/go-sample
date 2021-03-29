package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var result string

func main() {
	shuffle()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, result)
}

func shuffle() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(9)
	fmt.Println(r)
	switch r {
	case 0, 1:
		result = "大吉"
	case 2, 3, 4:
		result = "中吉"
	case 5, 6, 7:
		result = "小吉"
	case 8:
		result = "凶"
	default:
		result = "大凶"
	}
}
