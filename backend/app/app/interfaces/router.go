package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func Dispatch() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootPage).Methods("GET")
	router.HandleFunc("/user", signup).Methods("POST")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

func rootPage(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "success"); err != nil {
		log.Fatal(err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	fmt.Println(token)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var input SignupInput
	if err := json.Unmarshal(reqBody, &input); err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(User{ID: "id", UserName: "test"}); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID       string
	UserName string
}

type SignupInput struct {
	Token string `json:"token"`
}
