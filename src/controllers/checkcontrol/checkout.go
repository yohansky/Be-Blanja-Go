package checkcontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/checkmodel"
	"encoding/json"
	"net/http"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/checkout/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(checkmodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "PUT" {
		var input checkmodel.Checkout
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newCheckout := checkmodel.Checkout{
			Userid: input.Userid,
			Bagid: input.Bagid,
			Addressid: input.Addressid,
			Total: input.Total,
		}

		checkmodel.Updates(id, &newCheckout)
		msg := map[string]string{
			"Message": "Checkout Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		checkmodel.Deletes(id)
		msg := map[string]string{
			"Message": "Checkout Deleted",
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

func Checkouts(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res := checkmodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "POST" {
		var input checkmodel.Checkout
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := checkmodel.Checkout{
			Userid: input.Userid,
			Bagid: input.Bagid,
			Addressid: input.Addressid,
			Total: input.Total,
		}
		checkmodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Checkout Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}