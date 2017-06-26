package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/samitpal/run-ssh/plugins"
	"github.com/samitpal/run-ssh/plugins/ssh"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func RunIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func RunStatus(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
func RunCreate(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("X-Run-Plugin")
	if h == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var run plugins.Runner
	if h == "ssh" {
		run = ssh.NewRunSsh()
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &run); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	vrun, err := RunValidate(run)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		err_json, _ := json.Marshal(err.Error())
		fmt.Fprintf(w, string(err_json))
		return
	}

	go ssh.Hello()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(vrun); err != nil {
		panic(err)
	}

}
