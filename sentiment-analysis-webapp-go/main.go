package main

import (
	"bytes"
	"encoding/json"
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
	Sentence string
	Polarity int
}

func doTheThings(w http.ResponseWriter, r *http.Request) {
	//receive http request
	log.Printf("Sentence received")
	body, _ := ioutil.ReadAll(r.Body)
	sb := string(body)
	//extract sentence json and send that to the poster function
	response := post(sb)
	postBody, err := json.Marshal(map[string]interface{}(response))
	if err != nil {
		log.Fatal(err)
	}
	//write the response back
	w.Write([]byte(postBody))

}

//'{"sentence": "i love you"}'
func post(sentence string) map[string]interface{} {

	SA_LOGIC_API_URL := os.Getenv("SA_LOGIC_API_URL")
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"sentence": sentence,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(SA_LOGIC_API_URL+"/analyse/sentiment", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read body to json

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		log.Fatal(err)
	}
	log.Println(obj)
	return obj

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sentiment", doTheThings).Methods("POST")
	// router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/health", healthcheck).Methods("GET")
	// router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

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
