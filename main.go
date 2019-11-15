package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"net/http"
	"os"
	jikan "github.com/darenliang/jikan-go"
	"fmt"
	"encoding/json"

	// seasons "github.com/oraora-io/oraora-be/seasons"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Name struct {
	Title string
}

func main() {
	// Here we are instantiating the gorilla/mux router

	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/anime", AnimeHandler).Methods("GET")

	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

// AnimeHandler ...
var AnimeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Gets all the anime information data based on year and season
	season, _ := jikan.Season{Year: 2019, Season: "fall"}.Get()

	// Starting with an empty slice of strings
	var animeSlice []string

	//TODO: figure out a more elegant way of getting number of anime in a given season
	//Can't natively index type interface{}, so convert to []interface{}, and then index, with output of type map[string]interface{}
	for i := 0; i < 182; i++ {

		anime := season["anime"].([]interface{})[i].(map[string]interface{})

		//Marshal the data into json

		animeData, err := json.Marshal(anime)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// type Name declared as a struct at the top
		var title Name

		// Unmarshal the JSON to only grab the title of the anime
		json.Unmarshal([]byte(string(animeData)), &title)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Append title of the anime to empty slice of strings animeSlice
		animeSlice = append(animeSlice, title.Title)

	}

	// Marshal the entire slice of strings back into JSON
	json.Marshal(animeSlice)
	titleJSON, err := json.Marshal(animeSlice)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(titleJSON))
})

// StatusHandler ...
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

// NotImplemented ...
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
