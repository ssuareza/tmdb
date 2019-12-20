package themoviedb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Results holds the results from the request
type Results struct {
	Results []Movie `json:"results"`
}

// Movie holds the movie structure
type Movie struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}

// SearchMovie search movies based in a query string
func (client *Client) SearchMovie(query string) ([]Movie, error) {
	url := client.URL + "/search/movie?include_adult=false&page=1&query=" + query + "&language=en-US&api_key=" + client.APIKey
	// payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var r Results
	json.Unmarshal(body, &r)

	return r.Results, nil
}
