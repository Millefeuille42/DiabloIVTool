package fetchers

import (
	"encoding/json"
	"fetcher/pkg/globals"
	"io"
	"log"
	"net/http"
)

const uniquesEndpoint = "/uniques/page-data.json"
const uniquesUrl = globals.BuildApiUrl + uniquesEndpoint

type uniqueItemRawResponse struct {
	ComponentChunkName string `json:"componentChunkName"`
	Path               string `json:"path"`
	Result             struct {
		PageContext struct {
			Uniques []struct {
				Name           string   `json:"name"`
				Class          string   `json:"class,omitempty"`
				Type           string   `json:"type"`
				SecondaryStats []string `json:"secondary_stats,omitempty"`
				TerciaryStats  []string `json:"terciary_stats,omitempty"`
				Effect         string   `json:"effect"`
				Flavor         string   `json:"flavor"`
				Filters        []string `json:"filters,omitempty"`
			} `json:"uniques"`
		} `json:"pageContext"`
	} `json:"result"`
	StaticQueryHashes []string `json:"staticQueryHashes"`
}

func getUniqueItemsRaw() ([]byte, error) {
	resp, err := http.Get(uniquesUrl)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	jsonData := uniqueItemRawResponse{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	data, err = json.Marshal(jsonData.Result.PageContext)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}

	return data, nil
}
