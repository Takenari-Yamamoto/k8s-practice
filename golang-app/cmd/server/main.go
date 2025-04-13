package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// サンプルデータ
	books := []Book{
		{ID: "1", Title: "Go言語プログラミング", Author: "山田太郎", Year: 2023},
		{ID: "2", Title: "実践Kubernetes", Author: "鈴木一郎", Year: 2024},
		{ID: "3", Title: "マイクロサービスアーキテクチャ", Author: "佐藤花子", Year: 2022},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Kubernetes!")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
