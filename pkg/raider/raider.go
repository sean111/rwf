package raider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
)

var apiUrl string = "https://raider.io/api/v1"

func GetStaticData(options StaticDataOptions) (StaticDataResponse, error) {
	if options.Expansion < 6 || options.Expansion > 10 {
		return StaticDataResponse{}, errors.New("invalid expansion selected")
	}
	staticDataUrl := fmt.Sprintf("%s/raiding/static-data/?expansion_id=%d", apiUrl, options.Expansion)
	resp, _ := http.NewRequest(http.MethodGet, staticDataUrl, nil)
	resp.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Error(err.Error())
		return StaticDataResponse{}, err
	}
	defer response.Body.Close()
	var raidsResponse StaticDataResponse
	raidsResponse.Status = response.StatusCode
	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&raidsResponse)
		if err != nil {
			log.Error(err.Error())
			return StaticDataResponse{}, err
		}
	}
	return raidsResponse, nil
}

func GetRaidRankings(options RaidRankingsOptions) (RaidRankingsResponse, error) {
	dataUrl := fmt.Sprintf("%s/raiding/raid-rankings?raid=%s&difficulty=%s&region=%s&limit=%d&page=%d", apiUrl, options.Raid, options.Difficulty, options.Region, options.Limit, options.Page)
	resp, _ := http.NewRequest(http.MethodGet, dataUrl, nil)

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return RaidRankingsResponse{}, err
	}
	defer response.Body.Close()

	var rankingResponse RaidRankingsResponse
	rankingResponse.Status = response.StatusCode
	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&rankingResponse)
		if err != nil {
			return RaidRankingsResponse{}, err
		}
	}
	return rankingResponse, nil
}
