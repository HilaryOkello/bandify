package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

func FetchArtists() ([]Artist, error) {
	var artists []Artist
	response, err := http.Get(artistsURL)
	if err != nil {
		return artists, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artists, err
	}
	err = json.Unmarshal(bytes, &artists)
	if err != nil {
		return artists, err
	}
	defer response.Body.Close()
	return artists, nil
}

func FetchLocations() (Loc, error) {
	var location Loc
	response, err := http.Get(locationsURL)
	if err != nil {
		return location, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	defer response.Body.Close()
	return location, nil
}

func ArtistID(id int) (Artist, error) {
	var artist Artist
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id))
	if err != nil {
		return artist, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return artist, err
	}

	err = json.Unmarshal(bytes, &artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func LocationID(id int) (Loc, error) {
	var location Loc // see the structure
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id))
	if err != nil {
		return location, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return location, err
	}
	return location, nil
}

func FetchRelation(id int) (Relation, error) {
	var rel Relation
	response, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return rel, err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return rel, err
	}
	err = json.Unmarshal(bytes, &rel)
	if err != nil {
		return rel, err
	}
	return rel, nil
}

func Search(data Everything, searchTerm string) (Everything, error) {
	var output Everything
	ids := make(map[int]int)
	var artists []Artist
	for _, result := range data.Everyone {
		if strings.Contains(strings.ToLower(result.Name), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(result.FirstAlbum), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(strconv.Itoa(result.CreationDate)), strings.ToLower(searchTerm)) {
			if _, ok := ids[result.ID]; ok {
				continue
			} else {
				artists = append(artists, result)
				ids[result.ID] += 1
			}
		}
	}
	output.Everyone = artists
	return output, nil
}
