package catecontrol

import (
	models "Backend-Golang/src/models/catemodel"
	"encoding/json"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/category/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(models.SelectCategory(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "PUT" {
		var input models.Category
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newcate := models.Category{
			Name:   input.Name,
			Imgurl: input.Imgurl,
		}

		models.UpdatesCategory(id, &newcate)
		msg := map[string]string{
			"Message": "Category Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		models.DeletesCategory(id)
		msg := map[string]string{
			"Message": "Category Deleted",
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

func Categories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res := models.SelectAllCategory()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var input models.Category
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := models.Category{
			Name:   input.Name,
			Imgurl: input.Imgurl,
		}
		models.PostCategory(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Category Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
