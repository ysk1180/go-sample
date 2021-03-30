package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var tmpl = template.Must(template.New("msg").Parse("<html><body>{{if .Name}}{{.Name}}さんの{{end}}運勢は<b>「{{.Result}}」</b>です！</body></html>"))

type Message struct {
	Name   string
	Result string
}

func main() {
	shuffle()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := Message{r.FormValue("p"), omikuji()}
	tmpl.Execute(w, msg)
}

func shuffle() {
	rand.Seed(time.Now().Unix())
}

func omikuji() (result string) {
	switch rand.Intn(9) {
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
	return
}
