package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Product ...
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	Product{Id: 1, Name: "Dororo", Slug: "dororo", Description: "Give me back my body"},
	Product{Id: 2, Name: "O Maiden", Slug: "o-maiden", Description: "Literature club becomes sex ed"},
	Product{Id: 3, Name: "Demon Slayer", Slug: "demon-slayer", Description: "Year of the basket loli"},
	Product{Id: 4, Name: "Boku No Hero", Slug: "my-hero", Description: "Imagine being useless until you're not"},
	Product{Id: 5, Name: "Haikyuu", Slug: "haikyuu", Description: "Volleyball is cool"},
	Product{Id: 6, Name: "Boruto", Slug: "boruto", Description: "Boruto's dad needs an anime"},
}

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")

	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

// StatusHandler ...
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

// ProductsHandler ...
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// AddFeedbackHandler ...
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

// NotImplemented ...
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
