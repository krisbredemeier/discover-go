package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		return
	}
	fmt.Println("status code is", resp.Status)

	var m Movie
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return
	}
	rating := int(m.IMDBRating * 10)

	fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.\n", m.Title, m.Released, rating, m.IMDBVotes)
}
