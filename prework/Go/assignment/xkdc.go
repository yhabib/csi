// The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json
// produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an
// offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that
// matches a search term provided on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//  ! Where to handle errors? here or at the fetch level?? What error do I throw here. Strategy for error handling
// Can I have some multifield structure? Year, Month, Day       string `json:"year", "month", "day"`
// Comic is the representation of xkdc Comic
type Comic struct {
	Num        int    `json:"num,omitempty"`
	Year       string `json:"year,omitempty"`
	Month      string `json:"month,omitempty"`
	Day        string `json:"day,omitempty"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt,omitempty"`
	Image      string `json:"image,omitempty"`
	URL        string `json:"url"`
}

// StoredComic is the representation of a Comic in file
type StoredComic struct {
	URL        string `json: "url"`
	Transcript string `json: "transcript"`
}

func main() {
	fmt.Print("Enter pattern to search: ")
	pattern, err := readInput()
	if err != nil {
		log.Printf("enter pattern fails: %v", err)
	}

	pathToFile, err := createIndex()
	if err != nil {
		log.Printf("create index fails: %v", err)
	}

	matches, err := searchForPattern(pattern, pathToFile)
	if err != nil {
		log.Printf("search for pattern fails: %v", err)
	}
	if len(matches) == 0 {
		fmt.Println("No matches found")
		os.Exit(1)
	}
	for i, match := range matches {
		fmt.Printf("%d:\t%s \n\t%s\n", i+1, match.URL, match.Transcript)
	}
}

func readInput() (string, error) {
	var input string
	fmt.Scanln(&input)

	return input, nil
}

func searchForPattern(pattern string, pathToFile string) ([]StoredComic, error) {
	var _, comics []Comic
	var matches []StoredComic
	fmt.Printf("Starting search for provided pattern: %s\n", pattern)
	jsonFile, err := os.Open(pathToFile)
	if err != nil {
		return matches, fmt.Errorf("opening file %s: %v", pathToFile, err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		return matches, err
	}

	if err := json.Unmarshal(byteValue, &comics); err != nil {
		return matches, err
	}

	for _, comic := range comics {
		if strings.Contains(comic.Transcript, pattern) {
			matches = append(matches, StoredComic{comic.URL, comic.Transcript})
		}
	}
	return matches, nil
}

func createIndex() (string, error) {
	pathToFile := "./index.json"
	fmt.Println("Checking for existance of index ...")
	if byte, err := ioutil.ReadFile(pathToFile); err != nil && !os.IsNotExist(err) {
		return "", err
	} else if len(byte) > 0 {
		fmt.Println("   File already exists √")
		return pathToFile, nil
	}

	fmt.Println("Generating index ...")
	comics, err := getComics()
	if err != nil {
		return "", fmt.Errorf("createIndex fails: %v", err)
	}

	file, err := json.MarshalIndent(comics, "", "  ")
	if err != nil {
		return "", fmt.Errorf("parsin %s to JSON: %v", pathToFile, err)
	}
	if err := ioutil.WriteFile(pathToFile, file, 0644); err != nil {
		return "", fmt.Errorf("writing to file %s: %v", pathToFile, err)
	}
	fmt.Println("   File generated √")
	return pathToFile, nil
}

func getComics() ([]Comic, error) {
	pathToComic := "https://xkcd.com/%d/info.0.json"
	var comics []Comic
	count, err := getCountOfComics()
	if err != nil {
		return comics, fmt.Errorf("getComics fails: %v", err)
	}
	for i := 1; i < count-2416; i++ {
		var comic Comic
		path := fmt.Sprintf(pathToComic, i)
		comic, err := getComic(path)
		if err != nil {
			return comics, fmt.Errorf("getComics fails: %v", err)
		}
		comics = append(comics, comic)
	}
	return comics, nil
}

func getCountOfComics() (int, error) {
	pathToComicCount := "https://xkcd.com/info.0.json"
	comic, err := getComic(pathToComicCount)
	if err != nil {
		return 0, fmt.Errorf("getCountOfComics fails: %v", err)
	}
	return comic.Num, nil
}

func getComic(url string) (Comic, error) {
	var comic Comic
	resp, err := http.Get(url)
	if err != nil {
		return comic, fmt.Errorf("fetching %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %s: %s", url, resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	comic.URL = url
	return comic, nil
}

// Some final comments:
//  Best way of handling errors? Log them low in the chain, exit the program high in the chain
//  Marshalling and unmarshalling too many times
