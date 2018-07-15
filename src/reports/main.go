package reports

import "github.com/RTradeLtd/mining-bootstrap/src/reports/config"

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
