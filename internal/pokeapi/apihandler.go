package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Xeninon/pokedexcli/internal/pokecache"
)

func PokeGet(cache *pokecache.Cache, url string) ([]byte, error) {
	body, ok := cache.Get(url)
	if ok {
		return body, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	cache.Add(url, body)
	return body, nil
}
