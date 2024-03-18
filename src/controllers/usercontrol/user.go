package usercontrol

import (
	"Backend-Golang/src/helper"
	"Backend-Golang/src/middleware"
	"Backend-Golang/src/models/usermodel"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func RegisterSeller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input usermodel.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid request body")
			return
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashPassword)

		item := usermodel.User{
			Name:        input.Name,
			Email:       input.Email,
			Password:    Password,
			Phonenumber: input.Phonenumber,
			Storename:   input.Storename,
			Role:        "Seller",
		}
		usermodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Seller Registered",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input usermodel.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := usermodel.User{
			Name:        input.Name,
			Email:       input.Email,
			Password:    Password,
			Phonenumber: "-",
			Storename:   "-",
			Role:        "Customer",
		}
		usermodel.Post(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Customer Registered",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	// middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input usermodel.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		//testing validate role
		// ValidateRole := usermodel.FindRole(&input)
		// if len(ValidateRole) == 0 {

		// }
		// var RoleSecond string
		// for _ ,user := range ValidateRole {
		// 	RoleSecond = user.Role
		// }

		ValidateEmail := usermodel.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Email is not Found")
			return
		}
		var PasswordSecond string
		for _, user := range ValidateEmail {
			PasswordSecond = user.Password
		}

		if err := bcrypt.CompareHashAndPassword([]byte(PasswordSecond), []byte(input.Password)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Password not Found")
			return
		}
		jwtKey := os.Getenv("SECRETKEY")
		token, err := helper.GenerateToken(jwtKey, input.Email, input.Role)
		if err != nil {
			http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
			return
		}

		// if input.Role == "Seller" {
		// 	input.Role = "Seller"
		// } else if input.Role == "Customer"{
		// 	input.Role = "Customer"
		// }
		item := map[string]string{
			"Email": input.Email,
			"Role":  "Seller", //User.Role
			"Token": token,
		}
		res, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Gagal konversi ke Json", http.StatusInternalServerError)
			return
		}
		// w.WriteHeader(http.StatusOK)
		w.Write(res)
		// debug
		// fmt.Fprintf(w, input.Email)
		// fmt.Fprintf(w, input.Role)
		// fmt.Fprintf(w, token)
		return

	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res := usermodel.SelectAll()
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/user/"):]
	if r.Method == "GET" {
		result, err := json.Marshal(usermodel.Select(id).Value)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "DELETE" {
		usermodel.Deletes(id)
		msg := map[string]string{
			"Message": "User Deleted",
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

func UpdateSeller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-seller/"):]
	if r.Method == "PUT" {
		var input usermodel.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		//hashing
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newSeller := usermodel.User{
			Name:        input.Name,
			Email:       input.Email,
			Password:    Password,
			Phonenumber: input.Phonenumber,
			Storename:   input.Storename,
		}

		usermodel.UpdatesSeller(id, &newSeller)
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
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-customer/"):]
	if r.Method == "PUT" {
		var input usermodel.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newCustomer := usermodel.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: Password,
		}

		usermodel.UpdatesCustomer(id, &newCustomer)
		msg := map[string]string{
			"Message": "Customer Updated",
		}

		result, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := middleware.ExtractToken(r)

	if refreshToken == "" {
		http.Error(w, "Refresh token tidak tersedia", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if err != nil || !token.Valid { // jika ada error / bukan token valid
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	email := claims["email"].(string) //di konversi menjadi string
	role := claims["role"].(string)

	newToken, err := helper.GenerateToken(os.Getenv("SECRETKEY"), email, role) //tambahkan role
	if err != nil {
		http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
		return
	}

	msg := map[string]string{ // tampilkan hasil token yang di generate
		"token": newToken,
	}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Failed convert to Json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Typee", "application/json")
	w.Write(res)
}
