package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// strings splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple /orange /banana"

	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "/")

	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"Ghana", "Cote D'Ivore", "Benin", "Cameroon"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(fruits, "apple"))

	replaced := strings.Replace(str, "Go", "World", 2)
	fmt.Println(replaced)

	strwspace := "  Hello Everyone!"
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))
	fmt.Println(strings.ToTitle(strwspace))

	fmt.Println(strings.Repeat("foo", 3))

	fmt.Println(strings.Count("Hello", "1"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str5 := "Hello, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, "
	fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// Convert builder to a string
	resultBuilder := builder.String()
	fmt.Println(resultBuilder)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	resultBuilder = builder.String()
	fmt.Println(resultBuilder)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	resultBuilder = builder.String()
	fmt.Println(resultBuilder)
}
