/*
*	program takes string argument, make http GET request to
*	Hacker News ( Y Combinator ) and print result on stdout.
*
*	NOTE: Go program without goroutines and channels is crap. I know... THIS was just for fun! ;)
 */

package search

import (
	"encoding/json" // need to decode json object
	"errors"        // convert error string to error value
	"fmt"           // format output string
	"net/http"      // to make GET request
	"strings"       // string manipulation functions
)

// structure Response contains array of structure Item
type Response struct {
	Hits []Item
}

// structure Item will hold values of json object
// you can see the format of json object at https://hn.algolia.com/api
type Item struct {
	Title  string
	Url    string
	Author string
	Points int
}

// Get() takes string `search` and returns array of Item and error number
// Yes!.. Go function can return multiple values
func Get(search string) ([]Item, error) {
	url := fmt.Sprintf("https://hn.algolia.com/api/v1/search?query=%s", search)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	// new(T) returns pointer of type T
	r := new(Response)

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	// make(T, size) returns type T
	items := make([]Item, r.Len())

	for i, child := range r.Hits {
		if child.Title != "" {
			items[i] = child
		}
	}

	return items, nil
}

// String() function allows to take receiver `i Item`
// and return string as per body of function
// So, you can format output of custom Type
func (i Item) String() string {
	title := fmt.Sprintf("Title:\t%s", i.Title)
	url := fmt.Sprintf("Url:\t%s", i.Url)
	author := fmt.Sprintf("Author:\t%s", i.Author)
	points := fmt.Sprintf("Points:\t%d", i.Points)
	return strings.Join([]string{title, url, author, points, ""}, "\n")
}

// Len() receives pointer of Response struct and
// return count of non-blank member Title of Item
func (r *Response) Len() (count int) {
	for _, child := range r.Hits {
		if child.Title != "" {
			count++
		}
	}
	return
}
