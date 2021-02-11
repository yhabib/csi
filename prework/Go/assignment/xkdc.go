// The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
// produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an
// offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that
// matches a search term provided on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	pathToComicCount = "https://xkcd.com/info.0.json"
	pathToIndex      = "./index"
)

// Comic represetns the XDCD structure
type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Image            string
}

func main() {
	num, _ := getCountOfComics()
	println(num)
}

func getCountOfComics() (int, error) {
	resp, err := http.Get(pathToComicCount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "page cannot be fetched: %s\n", resp.Status)
		return 0, err
	}
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer resp.Body.Close()

	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		fmt.Fprintf(os.Stderr, "error decoding the body: %v\n", err)
		return 0, err
	}

	return comic.Num, nil
}
