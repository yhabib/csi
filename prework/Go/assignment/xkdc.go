// The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
// produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an
// offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that
// matches a search term provided on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	pathToComicCount = "https://xkcd.com/info.0.json"
	pathToComic      = "https://xkcd.com/%d/info.0.json"
	pathToIndex      = "./index.json"
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
	createIndex()
}

func createIndex() bool {
	comics, err := getComics()
	if err != nil {
		//  ! Where to handle errors? here or at the fetch level?? What error do I throw here
		// fmt.Fprintf(os.Stderr, "page cannot be fetched: %s\n", resp.Status)
		return false
	}
	fmt.Println(comics)

	file, _ := json.MarshalIndent(comics, "", " ")
	err = ioutil.WriteFile(pathToIndex, file, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to file: %v\n", err)
		return false
	}
	return true
}

func getComics() ([]Comic, error) {
	count, _ := getCountOfComics()
	fmt.Println(count)
	var comics []Comic
	for i := 1; i < 5; i++ {
		var comic Comic
		path := fmt.Sprintf(pathToComic, i)
		fmt.Println(path)
		comic, err := getComic(path)
		if err != nil {
			return comics, err
		}
		comics = append(comics, comic)
	}
	return comics, nil
}

func getComic(url string) (Comic, error) {
	var comic Comic
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching comic: %v\n", err)
		return comic, err
	}

	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		fmt.Fprintf(os.Stderr, "error decoding the body: %v\n", err)
		return comic, err
	}

	return comic, nil
}

func getCountOfComics() (int, error) {
	comic, _ := getComic(pathToComicCount)
	return comic.Num, nil
}
