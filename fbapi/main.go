package main

import (
	"fmt"

	fb "github.com/huandu/facebook/v2"
)

func main() {
	res, _ := fb.Get("/3261885224075207", fb.Params{
		"fields":       "first_name",
		"access_token": "EAAWcwdyDtvMBAMK97LxZAbQ1R2hDXaTfAKHzCskhXVPeiNabZBBGnIMnaVItZByjyBFRQGdHITAhgMuent0Lg9M7x6Y9KKNKKRWoFDXc9xsctbgywGr5YWfv9bla6h8o6NblMQUlSDDeZARNT5SRoCyEYg73ieGxGm1LOTQDAFgoaX4ghqkeglXZAV6G8Pi7GljzQkAkaA6QWrplbZAc83Cy5oYPAl2W9CYOHnFNxNASBr5OmUqXPlCcZAEvWwM4UQZD",
	})
	fmt.Println("Here is my Facebook first name:", res["first_name"])
	// Decode "first_name" to a Go string.
	var first_name string
	res.DecodeField("first_name", &first_name)
	fmt.Println("Here's an alternative way to get first_name:", first_name)

	//It's also possible to decode the whole result into a predefined struct.
	type User struct {
		FirstName string
	}

	var user User

	res.Decode(&user)

	fmt.Println("print first_name in struct:", user.FirstName)
	// res, err := fb.Get("/me/feed", fb.Params{
	// 	"access_token": "EAAWcwdyDtvMBANqdLVTjtJsCt9Izj66g6CdFbrqZCY4ZAiO47dbHQfbVcZAJexTIlFszPnV7ZBAcnF1pRSCi31ZCSh3pRfebY9khcmU6QJca7hfypZCT76kCscjp6aZC083aUJtJBGG8pGzNmaFTGWK3zSRfrKuqnqAZA396mS7249sZAk3tKuf6tRskA9vKFZA8qWp7UkZA8lp6Q484EWevybGhlyas0svWCJ7SsmSawYaECbhdmiZC8Bp7UWV74ZA4qJd8ZD",
	// })

	// if err != nil {
	// 	// err can be a Facebook API error.
	// 	// if so, the Error struct contains error details.
	// 	if e, ok := err.(*Error); ok {
	// 		fmt.Printf("facebook error. [message:%v] [type:%v] [code:%v] [subcode:%v] [trace:%v]",
	// 			e.Message, e.Type, e.Code, e.ErrorSubcode, e.TraceID)
	// 		return
	// 	}

	// err can be an unmarshal error when Facebook API returns a message which is not JSON.

	// 	if e, ok := err.(*UnmarshalError); ok {
	// 		fmt.Printf("facebook error. [message:%v] [err:%v] [payload:%v]",
	// 			e.Message, e.Err, string(e.Payload))
	// 		return
	// 	}

	// 	return
	// }

	// read my last feed story.
	//fmt.Println("My latest feed story is:", res.Get("data.1.story"))
}
