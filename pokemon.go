package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
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

