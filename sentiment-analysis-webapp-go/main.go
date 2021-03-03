package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//What we send to the API
type SentenceRequest struct {
	Sentence string
}

//What we get back from the API
type SentenceData struct {
	Sentence string  `json:"sentence"`
	Polarity float64 `json:"polarity"`
}

func doTheThings(w http.ResponseWriter, r *http.Request) {
	//receive http request
	log.Printf("Sentence received")
	body, _ := ioutil.ReadAll(r.Body)
	sb := string(body)
	score := getSentencePolarity(sb)
	response := SentenceData{
		Sentence: sb,
		Polarity: score,
	}

	postBody, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	//write the response back
	fmt.Println(postBody)
	w.Write(postBody)

}

// takes a sentence and sends it to downstream service to calculate polarity
func getSentencePolarity(sentence string) float64 {
	SA_LOGIC_API_URL := os.Getenv("SA_LOGIC_API_URL")
	sentencejson := SentenceRequest{
		Sentence: sentence,
	}
	postBodyinBytes, _ := json.Marshal(sentencejson)

	resp, err := http.Post(SA_LOGIC_API_URL+"/analyse/sentiment", "application/json", bytes.NewBuffer(postBodyinBytes))
	if err != nil {
		log.Fatalf("errors: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return 1 //return polarity score somehow.
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sentiment", doTheThings).Methods("POST")
	router.HandleFunc("/health", healthcheck).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))

}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("healthy")
	w.Write([]byte("true"))
}
