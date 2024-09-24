package raider

import (
	"encoding/json"
	"fmt"
	logger "github.com/charmbracelet/log"
	"net/http"
	"os"
)

var apiUrl string = "https://raider.io/api/v1"
var log = logger.NewWithOptions(os.Stderr, logger.Options{
	ReportTimestamp: true,
	ReportCaller:    true,
})

func GetStaticData(expansionId int) (StaticDataResponse, error) {
	staticDataUrl := fmt.Sprintf("%s/raiding/static-data/?expansion_id=%d", apiUrl, expansionId)
	resp, _ := http.NewRequest(http.MethodGet, staticDataUrl, nil)
	resp.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Error(err.Error())
		return StaticDataResponse{}, err
	}
	defer response.Body.Close()
	var raidsResponse StaticDataResponse
	err = json.NewDecoder(response.Body).Decode(&raidsResponse)
	if err != nil {
		log.Error(err.Error())
		return StaticDataResponse{}, err
	}
	return raidsResponse, nil
}

func GetRaidRankings(raid string, difficulty string, region string, limit int, page int) (RaidRankingsResponse, error) {
	dataUrl := fmt.Sprintf("%s/raiding/raid-rankings?raid=%s&difficulty=%s&region=%s&limit=%d&page=%d", apiUrl, raid, difficulty, region, limit, page)
	log.Info("dataUrl", "url", dataUrl)
	resp, _ := http.NewRequest(http.MethodGet, dataUrl, nil)

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return RaidRankingsResponse{}, err
	}

	defer response.Body.Close()

	var rankingResponse RaidRankingsResponse
	err = json.NewDecoder(response.Body).Decode(&rankingResponse)
	if err != nil {
		return RaidRankingsResponse{}, err
	}
	return rankingResponse, nil
}
