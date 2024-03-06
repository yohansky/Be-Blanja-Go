package selcontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/selmodel"
	"encoding/json"
	"net/http"
)

func Seller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/seller/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(selmodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "PUT" {
		var input selmodel.Seller
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newSeller := selmodel.Seller{
			SName :input.SName,
			Email : input.Email,
			Phone_Number : input.Phone_Number,
			Store_Name : input.Store_Name,
			Password: input.Password,	
		}

		selmodel.Updates(id, &newSeller)
		msg := map[string]string{
			"Message": "Seller Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		selmodel.Deletes(id)
		msg := map[string]string{
			"Message": "Seller Deleted",
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

func Sellers(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res := selmodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
	} else if r.Method == "POST"{
		var input selmodel.Seller
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := selmodel.Seller{
			SName :input.SName,
			Email : input.Email,
			Phone_Number : input.Phone_Number,
			Store_Name : input.Store_Name,
			Password: input.Password,
		}
		selmodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Seller Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}