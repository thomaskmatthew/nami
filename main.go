package main

import (
	"fmt"
	database2 "github.com/thomaskmatthew/nami/internal/database"
	"github.com/thomaskmatthew/nami/internal/fetch"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type PokemonDB struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func main() {
	url := "https://pokeapi.co/api/v2/pokemon/%d"
	// var p pokemon.Pokemon
	var p PokemonDB
	var listOfPokemons []PokemonDB
	for i := 1; i <= 25; i++ {
		currentPokemon := fmt.Sprintf(url, i)
		err := fetch.FetchJson(currentPokemon, &p)
		if err != nil {
			fmt.Printf("Can't fetch pokemon: %d\n", i)
			continue
		}

		fmt.Printf("%+v\n", p)
		listOfPokemons = append(listOfPokemons, p)

	}
	//fetcher := fetch.FetchJson(url, &p)
	//
	//fmt.Println("fetching pokemon")
	//fmt.Println(fetcher)
	//fmt.Println("current pokemon")
	//fmt.Println(p)

	database := database2.Connection()
	if database == nil {
		fmt.Println("could not connect to database")
	}
	dbErr := database.DB.Ping()
	fmt.Println(dbErr)

	//if database != nil {
	//	statement := "CREATE TABLE POKEMON (ID INT, "
	//
	//	created := database.CreateTable(statement)
	//	fmt.Println(created)
	//}

}
