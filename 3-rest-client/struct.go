package main

//Movie for importing imdb json
type Movie struct {
	Title      string  `json:"Title"`
	Year       string  `json:"Year"`
	Rated      string  `json:"Rated"`
	Released   string  `json:"Released"`
	Runtime    string  `json:"Runtime"`
	Genre      string  `json:"Genre"`
	Director   string  `json:"Director"`
	Writer     string  `json:"Writer"`
	Actors     string  `json:"Actors"`
	Plot       string  `json:"Plot"`
	Language   string  `json:"Language"`
	Country    string  `json:"Country"`
	Awards     string  `json:"Awards"`
	Poster     string  `json:"Poster"`
	Metascore  string  `json:"Metascore"`
	IMDBRating float64 `json:"imdbRating,string"`
	IMDBVotes  string  `json:"imdbVotes"`
	IMDBID     string  `json:"imdbID"`
	Type       string  `json:"Type"`
	Response   string  `json:"Response"`
}
