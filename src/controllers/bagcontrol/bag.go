package bagcontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/bagmodel"
	"encoding/json"
	"net/http"
)

func Bag(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/bag/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(bagmodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "PUT" {
		var input bagmodel.Bag
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newBag := bagmodel.Bag{
			Productid: input.Productid,
			Total_Price: input.Total_Price,
		}

		bagmodel.Updates(id, &newBag)
		msg := map[string]string{
			"Message": "Bag Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		
	} else if r.Method == "DELETE" {
		bagmodel.Deletes(id)
		msg := map[string]string{
			"Message": "Bag Deleted",
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

func Bags(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res := bagmodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var input bagmodel.Bag
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := bagmodel.Bag{
			Productid: input.Productid,
			Total_Price: input.Total_Price,
		}
		bagmodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "User Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}