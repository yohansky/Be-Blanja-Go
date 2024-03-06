package ordcontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/ordmodel"
	"encoding/json"
	"net/http"
)

func Order(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/order/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(ordmodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "PUT" {
		var input ordmodel.Order
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newOrder := ordmodel.Order{
			Productid: input.Productid,
    		Name: input.Name,
    		Price: input.Price,
    		Sum: input.Sum,
    		Delivery: input.Delivery,
    		Total: input.Total,
    		Costumerid: input.Costumerid,
    		CName: input.CName,
		}

		ordmodel.Updates(id, &newOrder)
		msg := map[string]string{
			"Message": "Order Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		ordmodel.Deletes(id)
		msg := map[string]string{
			"Message": "Order Deleted",
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

func Orders(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res := ordmodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "POST" {
		var input ordmodel.Order
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := ordmodel.Order{
			Productid: input.Productid,
    		Name: input.Name,
    		Price: input.Price,
    		Sum: input.Sum,
    		Delivery: input.Delivery,
    		Total: input.Total,
    		Costumerid: input.Costumerid,
    		CName: input.CName,
		}
		ordmodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Order Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

