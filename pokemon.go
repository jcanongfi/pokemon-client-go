package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
        "strings"
)

// GetPokemons - Returns list of pokemons (no auth required)
func (c *Client) GetPokemons() ([]Pokemon, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/pokemons", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pokemons := []Pokemon{}
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

// CreatePokemon - Create new pokemon
func (c *Client) CreatePokemon(pokemon_in Pokemon) (*Pokemon, error) {
	rb, err := json.Marshal(pokemon_in)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pokemons", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pokemon_out := Pokemon{}
	err = json.Unmarshal(body, &pokemon_out)
	if err != nil {
		return nil, err
	}

	return &pokemon_out, nil
}
