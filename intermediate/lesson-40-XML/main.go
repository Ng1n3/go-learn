package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

func main() {
	book := Book{
		ISBN:       "123213213123123",
		Title:      "Go Bootcamp",
		Author:     "EAZY",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}

	fmt.Println(string(xmlDataAttr))
}
