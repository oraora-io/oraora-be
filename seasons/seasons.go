package seasons

import (
	"encoding/json"
	"fmt"
	"net/http"

	jikan "github.com/darenliang/jikan-go"
)

// Name of the anime
type Name struct {
	Title string
}

// AnimeHandler reports the data in raw JSON
var AnimeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Gets all the anime information data based on year and season
	season, _ := jikan.Season{Year: 2019, Season: "Fall"}.Get()

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
