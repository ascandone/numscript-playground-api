package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"numscript_playground_api/handlers"
)

const PORT = 3000

func RunHandler(w http.ResponseWriter, r *http.Request) {
	// What method should it be?
	if r.Method != "POST" {
		// TODO return bad verb error
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var runArgs handlers.RunArgs
	err = json.Unmarshal(body, &runArgs)
	if err != nil {
		panic(err)
	}

	ret, err := handlers.Run(runArgs)
	if err != nil {
		panic(err)
	}

	outBytes, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}

func main() {
	http.HandleFunc("/run", RunHandler)
	fmt.Printf("Serving on https://localhost:%d \n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		panic(err)
	}
}
