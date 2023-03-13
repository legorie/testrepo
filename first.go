package main

// trest file
import (
	"fmt"

	"github.com/buger/jsonparser"
)

func main() {

	// raw_data, err := http.Get("https://raw.githubusercontent.com/cunningr/testrepo/main/test-data.json")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(raw_data)
	// data := raw_data

	raw_data_bytes := []byte(`
	{
		"random": "61",
		"random float": "70.113",
		"bool": "true",
		"date": "1996-08-15",
		"regEx": "helloooooooooooooooooooooooooooooooooooooooooo to you",
		"enum": "json",
		"firstname": "Cassondra",
		"lastname": "Justinn",
		"city": "Tampere",
		"country": "Bahrain",
		"countryCode": "MM",
		"email uses current data": "Cassondra.Justinn@gmail.com",
		"email from expression": "Cassondra.Justinn@yopmail.com",
		"array": [
			"Fawne",
			"Ofilia",
			"Joy",
			"Gale",
			"Teddie"
		],
		"array of objects": [
			{
				"index": "0",
				"index start at 5": "5"
			},
			{
				"index": "1",
				"index start at 5": "6"
			},
			{
				"index": "2",
				"index start at 5": "7"
			}
		],
		"Celestyna": {
			"age": "44"
		}
	}
	`)
	value, a, b, c := jsonparser.Get(raw_data_bytes, "firstname")
	fmt.Println(a, b, c)
	fmt.Println(string(value))
}
