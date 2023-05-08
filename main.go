package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Topic string `json:"topic"`
	Time  int64  `json:"time"`
}

func main() {
	router := mux.NewRouter()

	fmt.Println(router)

	//response := Response{"Jonas", time.Now().UnixMicro()}

	var response Response

	router.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		response = Response{vars["topic"], time.Now().UnixMilli()}
		b, err := json.Marshal(response)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	})
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
