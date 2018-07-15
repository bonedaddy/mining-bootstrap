package reports

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/config"
	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
)

/*
This is used to handle automated mining reports for cryptocurrency mining farms
The idea is to create an easy to use system that can be used by farm operators to create accurate book reports for the tax man
*/

type Manager struct {
	Config *config.Config `json:"config"`
}

func GenerateReportManagerFromFile() (*Manager, error) {
	cfg, err := config.LoadConfigFromFile("")
	if err != nil {
		return nil, err
	}
	return &Manager{Config: cfg}, nil

}

func GenerateReportManager(coin, apikey string) (*Manager, error) {
	cfg, err := config.LoadConfig(coin, apikey)
	if err != nil {
		return nil, err
	}
	return &Manager{Config: cfg}, nil
}

func (m *Manager) GetRecentCredits() (*[]types.RecentCredits, error) {
	s := "getdashboarddata"
	m.FormatURL(s)
	resp, err := http.Get(m.Config.URL)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var intf map[string]interface{}
	var data types.MiningPoolHubAPIResponse
	err = json.Unmarshal(bodyBytes, &intf)
	if err != nil {
		return nil, err
	}
	marshaled, err := json.Marshal(intf[s])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshaled, &data)
	if err != nil {
		return nil, err
	}
	marshaled, err = json.Marshal(data.Data["recent_credits"])
	if err != nil {
		return nil, err
	}
	var credits []types.RecentCredits
	err = json.Unmarshal(marshaled, &credits)
	if err != nil {
		return nil, err
	}
	return &credits, nil
}

func (m *Manager) FormatURL(action string) {
	m.Config.URL = fmt.Sprintf(m.Config.URL, m.Config.Coin, action, m.Config.APIKey)
}
