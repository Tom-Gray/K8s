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

type SentenceSubmission struct {
	Sentence string `json:"sentence"`
}
type SentenceData struct {
	Sentence string  `json:"sentence"`
	Polarity float64 `json:"polarity"`
}

func doTheThings(w http.ResponseWriter, r *http.Request) {
	//receive http request
	log.Printf("Sentence received")
	body, _ := ioutil.ReadAll(r.Body)
	sb := string(body)

	var sentenceSubmission SentenceSubmission
	json.Unmarshal([]byte(sb), &sentenceSubmission)

	score := getSentencePolarity(sb)
	response := SentenceData{
		Sentence: sentenceSubmission.Sentence,
		Polarity: score,
	}

	postBody, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	//write the response back
	w.Write(postBody)

}

type Result struct {
	Polarity float64 `json:"polarity"`
}

// takes a sentence and sends it to downstream service to calculate polarity
func getSentencePolarity(sentence string) float64 {
	SA_LOGIC_API_URL := os.Getenv("SA_LOGIC_API_URL")
	resp, err := http.Post(SA_LOGIC_API_URL+"/analyse/sentiment", "application/json", bytes.NewBuffer([]byte(sentence)))
	if err != nil {
		log.Fatalf("errors: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var result Result
	json.Unmarshal(body, &result)
	fmt.Println("Polarity: ", result.Polarity)

	return result.Polarity //return polarity score somehow.
}

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/sentiment", doTheThings).Methods("POST")
	router.HandleFunc("/health", healthcheck).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Origin, X-Requested-With, Content-Type, content-type, Accept, Authorization"}, //https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
		AllowedMethods:     []string{"GET,PUT,POST,DELETE,PATCH,OPTIONS"},
		OptionsPassthrough: true,
	})

	cors.Default().Handler(router)
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("healthy")
	w.Write([]byte("true"))
}
