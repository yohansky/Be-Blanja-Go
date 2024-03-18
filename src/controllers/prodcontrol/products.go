package prodcontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	models "Backend-Golang/src/models/prodmodels"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Product(w http.ResponseWriter, r *http.Request) { // GET & PUT & DELETE
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/product/"):]
	if r.Method == "GET" {
		// result, err := json.Marshal(models.SelectProduct(id).Value)
		result, err := json.Marshal(models.SelectProduct(id))
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	} else if r.Method == "PUT" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		newProduct := models.Product{
			Name:        input.Name,
			Rating:      input.Rating,
			Price:       input.Price,
			Color:       input.Color,
			Size:        input.Size,
			Amount:      input.Amount,
			Condition:   input.Condition,
			Description: input.Description,
			Sellerid:    input.Sellerid,
			CategoryId:  input.CategoryId,
		}

		models.UpdatesProduct(id, &newProduct)
		msg := map[string]string{
			"Message": "Product Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	} else if r.Method == "DELETE" {
		models.DeletesProduct(id)
		msg := map[string]string{
			"Message": "Product Deleted",
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

func SelectProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res := models.SelectAllProduct()
		// result, err := json.Marshal(res.Value)
		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
}

func Products(w http.ResponseWriter, r *http.Request) { // GET & POST
	middleware.GetCleanedInput(r)
	helper.EnableCors(w) //memungkinkan sharing sesama localhost
	if r.Method == "GET" {
		pageOld := r.URL.Query().Get("page") // ini string
		limitOld := r.URL.Query().Get("limit")
		page, _ := strconv.Atoi(pageOld) //ini diubah dari string ke integer
		limit, _ := strconv.Atoi(limitOld)
		offset := (page - 1) * limit
		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = "ASC"
		}
		sortby := r.URL.Query().Get("sortBy")
		if sortby == "" {
			sortby = "name"
		}
		sort = sortby + " " + strings.ToLower(sort)
		respons := models.FindCondProduct(sort, limit, offset)
		totalData := models.CountDataProduct()
		totalPage := math.Ceil(float64(totalData) / float64(limit))
		result := map[string]interface{}{
			"status": "Berhasil",
			// "data":        respons.Value,
			"data":        respons,
			"currentPage": page,
			"limit":       limit, //tadinya limitOld
			"totalData":   totalData,
			"totalPage":   totalPage,
		}

		res, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return

		// //debug
		// fmt.Fprint(w. sort)
		// fmt.Fprint(w. sortby)
		// fmt.Fprint(w. offset)
		// fmt.Fprint(w. page)
		// fmt.Fprint(w. limit)

	} else if r.Method == "POST" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := models.Product{
			Name:        input.Name,
			Rating:      input.Rating,
			Price:       input.Price,
			Color:       input.Color,
			Size:        input.Size,
			Amount:      input.Amount,
			Condition:   input.Condition,
			Description: input.Description,
			Sellerid:    input.Sellerid,
			CategoryId:  input.CategoryId,
		}
		models.PostProduct(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		result, _ := json.Marshal(msg)
		w.Write(result)
		return

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Handle_upload(w http.ResponseWriter, r *http.Request) {
	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20 // 2 MB
	)
	// Memeriksa method request, harus POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mendapatkan file dari form
	file, handler, err := r.FormFile("file") //untuk menangkap request body yang typenya file
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close() //file harus ditutup
	ext := filepath.Ext(handler.Filename)
	ext = strings.ToLower(ext)
	allowedExts := strings.Split(AllowedExtensions, ",") //pecah pakai method split
	validExtension := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			validExtension = true
			break
		}
	}
	if !validExtension {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	// Mengecek ukuran file
	fileSize := handler.Size
	if fileSize > MaxFileSize {
		http.Error(w, "File size exceeds the allowed limit", http.StatusBadRequest)
		return
	}

	// Membuat format waktu dengan detik
	timestamp := time.Now().Format("20060102_150405")

	// Membuat nama unik untuk file
	filename := fmt.Sprintf("src/uploads/%s_%s", timestamp, handler.Filename)

	// Membuat file untuk menyimpan gambar
	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Menyalin isi file yang diupload ke file yang baru dibuat
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// helper.ValidateUpload(w, handler) // harus di gabung sama controller
	// helper.Upload(w, file, handler) // bisa dipakai

	// Menyampaikan respons berhasil
	msg := map[string]string{
		"Message": "File uploaded successfully",
	}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	keyWord := r.URL.Query().Get("search")

	// res, err := json.Marshal(models.FindDataProduct(keyWord).Value)
	res, err := json.Marshal(models.FindDataProduct(keyWord))
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
