package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client)CatchPokemon(pokemonName string) (PokemonResp, error) {
	endpoint := "/pokemon/"
	fullURL := baseURL+endpoint+pokemonName

	dat, ok := c.cache.Get(fullURL)
	if ok {
		pokemonResp := PokemonResp{}
		if err := json.Unmarshal(dat, &pokemonResp); err != nil {
			return PokemonResp{}, err
		}
	
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer res.Body.Close()

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonResp{}, nil
	}

	pokemonResp := PokemonResp{}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}
	c.cache.Add(fullURL, dat)
	return pokemonResp, nil
}