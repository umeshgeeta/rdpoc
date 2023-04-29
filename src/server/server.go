package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	Id       string `json:"_id"`
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func main() {
	mime.AddExtensionType(".js", "application/javascript")
	r := mux.NewRouter()

	// Handle the /search URL and return a JSON-encoded list of items
	r.HandleFunc("/search", handleSearch)

	r.HandleFunc("/add-item", handleAddItem).Methods(http.MethodPost)

	// Serve the index.html file
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "public/index.html")
		http.ServeFile(w, r, "/Users/umeshpatil/git/rdpoc/public/index.html")
	})

	// Serve the bundle.js file
	r.HandleFunc("/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "public/bundle.js")
		http.ServeFile(w, r, "/Users/umeshpatil/git/rdpoc/public/bundle.js/bundle.js")
	})

	// Serve static assets from the public directory
	fs := http.FileServer(http.Dir("/Users/umeshpatil/git/rdpoc/public"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", r)
}

// searchItems returns a list of items that match the query
func searchItems2(query string) []Item {
	// TODO: Perform search logic and return a list of items
	return []Item{
		{Id: "64471f8ad9469e85b6238d9e",
			Name:     "BEHR MARQUEE Keystone Satin Enamel",
			Category: "Exterior Paint",
			Sku:      "945305"},
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received search call")
	query := r.URL.Query().Get("q")
	items, err := searchItems(query)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	// Return the item list
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(items)
}

func handleAddItem(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new item data
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the new item data
	if newItem.Sku == "" || newItem.Name == "" || newItem.Category == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	fmt.Println(newItem)

	// Add the new item to the database
	err = insertItem(newItem)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Return the new item data with the generated ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}
