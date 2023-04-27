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
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received search call")
		query := r.URL.Query().Get("q")
		items := searchItems(query)
		json.NewEncoder(w).Encode(items)
	})

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
func searchItems(query string) []Item {
	// TODO: Perform search logic and return a list of items
	return []Item{
		{Id: "64471f8ad9469e85b6238d9e",
			Name:     "BEHR MARQUEE Keystone Satin Enamel",
			Category: "Exterior Paint",
			Sku:      "945305"},
	}
}
