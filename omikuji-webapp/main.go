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
	var msg string
	if name := r.FormValue("p"); name != "" {
		msg = fmt.Sprintf("%sさんの運勢は「%s」です！", name, result)
	} else {
		msg = fmt.Sprintf("運勢は「%s」です！", result)
	}
	fmt.Fprint(w, msg)
}

func shuffle() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(9)
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
