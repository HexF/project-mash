package main

import (
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	myNode.Uptime = uptime()
	vars := mux.Vars(r)
	mash := vars["mash"]

	mashNo := new(big.Int)
	mashNo.SetString(mash, 36)

	change := new(big.Int)
	change.SetInt64(1)

	prevMashNo := new(big.Int)
	prevMashNo.Sub(mashNo, change)

	nextMashNo := new(big.Int)
	nextMashNo.Add(mashNo, change)

	factText := getFact(mash)

	page := Page{
		Fact:        factText,
		Mash:        mash,
		NextURL:     nextMashNo.Text(36),
		PreviousURL: prevMashNo.Text(36),
	}
	renderTemplate(myNode, page, w)
	myNode.PreviousRequest = page.Mash
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{mash:[a-z0-9]+}", handler).Methods("GET")

	http.Handle("/", rtr)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
