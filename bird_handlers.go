package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Song struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var songs []Song

func getSongHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "songs" variable to json
	songListBytes, err := json.Marshal(songs)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of songs to the response
	w.Write(songListBytes)
}

func createSongHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Song
	song := Song{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the song from the form info
	song.Species = r.Form.Get("id")
	song.Description = r.Form.Get("title")

	// Append our existing list of songs with a new entry
	songs = append(songs, song)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
