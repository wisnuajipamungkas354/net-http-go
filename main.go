package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	type Product struct {
		Nama  string `json:"nama"`
		Harga int    `json:"harga"`
	}

	type APIResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "Halo, Wisnu!",
			"status":  "success",
		}

		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"nama":      "Wisnu",
			"pekerjaan": "Backend Developer",
			"bahasa":    "Golang",
		}

		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method tidak didukung", http.StatusMethodNotAllowed)
			return
		}

		var products Product

		err := json.NewDecoder(r.Body).Decode(&products)

		if err != nil {
			http.Error(w, "JSON tidak valid", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	http.HandleFunc("/api/products/save", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(APIResponse{
				Success: false,
				Message: "Method tidak didukung!",
			})
			return
		}

		var products Product

		err := json.NewDecoder(r.Body).Decode(&products)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{
				Success: false,
				Message: "JSON tidak valid!",
			})
			return
		}

		var response APIResponse = APIResponse{
			Success: true,
			Message: "Data berhasil dibuat",
			Data:    products,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server berjalan di http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
