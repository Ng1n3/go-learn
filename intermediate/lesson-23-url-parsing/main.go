package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("Raw Query: ", parsedURL.RawQuery)
	fmt.Println("Fragment: ", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=john&age=30"

	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	// Building URL
	baseURL := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path",
		Fragment: "dafduu314",
	}

	query := baseURL.Query()
	query.Set("hame", "john")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL: ", baseURL.String())

	values := url.Values{}

	values.Add("name", "Sam")
	values.Add("age", "27")
	values.Add("city", "Ikeja")
	values.Add("country", "Nigeria")

	// Encode
	encodedQuery := values.Encode()

	fmt.Println(encodedQuery)

	//  Build a URL
	baseURL2 := "https://example.com/search"
	fullURL := baseURL2 + "?" + encodedQuery

	fmt.Println(fullURL)
}
