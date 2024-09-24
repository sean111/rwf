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

// Private Methods

// Public Methods

func GetStaticData(expansionId int) (StaticDataResponse, error) {
	staticDataUrl := fmt.Sprintf("%s/raiding/static-data/?expansion_id=%d", apiUrl, expansionId)
	resp, err := http.NewRequest(http.MethodGet, staticDataUrl, nil)
	resp.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Error(err.Error())
		return StaticDataResponse{}, err
	}
	defer response.Body.Close()
	var raidsResponse StaticDataResponse
	log.Info(response.Header)
	err = json.NewDecoder(response.Body).Decode(&raidsResponse)
	if err != nil {
		log.Error(err.Error())
		return StaticDataResponse{}, err
	}
	return raidsResponse, nil
}
