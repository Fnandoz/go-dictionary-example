package main

import (
	"./Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Enter a word: ")
	var input string
	fmt.Scanln(&input)

	searchWord := model.SearchModel{}
	searchWord.Word = input

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	req, err := http.NewRequest("GET", "https://owlbot.info/api/v4/dictionary/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Token f95ba2b2e035519d6ce737d859afb923195e88fa")
	req.URL.Path += searchWord.Word

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	wordModel := model.WordModel{}
	json.Unmarshal([]byte(body), &wordModel)
	fmt.Println("word: ", wordModel.Word)
	fmt.Println("pronunciation: ", wordModel.Pronunciation)
	for i := 0; i < len(wordModel.Definitions); i++ {
		fmt.Println("type: ", wordModel.Definitions[i].WordType)
		fmt.Println("definition: ", wordModel.Definitions[i].Definition)
		fmt.Println("example: ", wordModel.Definitions[i].Example)
		fmt.Println("imageUrl: ", wordModel.Definitions[i].ImageUrl)
		fmt.Println("emoji: ", wordModel.Definitions[i].Emoji)
		fmt.Println("")
	}
}
