package coscontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/cosmodel"
	"encoding/json"
	"net/http"
)

func Costumer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/costumer/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(cosmodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
		
	} else if r.Method == "PUT" {
		var input cosmodel.Costumer
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newCostumer := cosmodel.Costumer{
			CName :input.CName,
			Email : input.Email,
			Password : input.Password,
		}
		cosmodel.Updates(id, &newCostumer)
		msg := map[string]string{
			"Message": "Costumer Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		cosmodel.Deletes(id)
		msg := map[string]string{
			"Message": "Costumer Updated",
		}
		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Costumers(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res := cosmodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "POST" {
		var input cosmodel.Costumer
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := cosmodel.Costumer{
			CName :input.CName,
			Email : input.Email,
			Password : input.Password,
		}
		cosmodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Costumer Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	
}
