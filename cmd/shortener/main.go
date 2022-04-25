package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var urls = map[string]string{}

// PostUrl — создает короткий урл.
func PostUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST метод")

	body, _ := ioutil.ReadAll(r.Body)
	url := string(body)

	defer r.Body.Close()

	h := sha256.New()
	h.Write(body)

	key := fmt.Sprintf("%x", h.Sum(nil))

	urls[key] = url

	fmt.Println(urls)

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(key))
}

// GetShortUrl — возвращает полный урл по короткому.
func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET метод")

	id := strings.Trim(r.RequestURI, "/")

	url := urls[id]

	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Header().Add("Location", url)

	w.Write([]byte(url))
}

func Short(w http.ResponseWriter, r *http.Request) {

	// в зависимости от метода
	switch r.Method {
	// если методом POST
	case "POST":
		PostUrl(w, r)
	// если методом GET
	case "GET":
		GetShortUrl(w, r)
	default:
		fmt.Println("Неизвестный метод")
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", Short)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on :8080")

	server.ListenAndServe()
}
