package main

import (
	"fmt"
	"github.com/zhangzj/tgpl/ch1_tutorial/lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cycle, err := strconv.Atoi(r.Form.Get("cycle"))
	fmt.Println(cycle)
	if err != nil {
		lissajous.Lissajous(w, float64(cycle))
	}
}
