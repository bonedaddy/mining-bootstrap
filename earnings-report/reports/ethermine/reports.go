package ethermine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/config"
	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/database"
	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/utils"
	"github.com/jinzhu/gorm"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

/*
Utilizing this package one can easily generate reports about their cryptocurrency mining income when utilizing miningpoolhub.
that can parse this data for non solo miners. Utilizing MPHs API, you can easily generate mining income reports for any of the supported currencies.
*/

const (
	// USDAPI is the URL We use to query for USD->CAD conversion
	USDAPI       = "https://free.currencyconverterapi.com/api/v5/convert?q=USD_CAD&compact=y"
	ethermineAPI = "https://api.ethermine.org"
)

var methodList = []string{"24hour_credit", "credit"}

// Manager is a helper struct used for report generation
type Manager struct {
	Config         *config.Config   `json:"config"`
	EthUSD         float64          `json:"eth_usd"` // keeps track of the ETH->USD conversion ratio
	UsdCad         float64          `json:"usd_cad"` // keeps track of the USD -> USD conversion ratio
	SendgridClient *sendgrid.Client `json:"sendgrid_client"`
	DB             *gorm.DB         `json:"db"`
}

// GenerateReportManagerFromFile is used to generate our helper struct from the config file
func GenerateReportManagerFromFile(path string, initializeDatabase bool) (*Manager, error) {
	cfg, err := config.LoadConfigFromFile(path)
	if err != nil {
		return nil, err
	}
	usd, err := utils.ParseUSDCAD()
	if err != nil {
		return nil, err
	}
	eth, err := utils.ParseETHUSD()
	if err != nil {
		return nil, err
	}
	if initializeDatabase {
		dbm, err := database.Initialize(cfg.DBPass, cfg.DBURL, cfg.DBUser)
		if err != nil {
			return nil, err
		}
		dbm.RunMigrations()
		return &Manager{Config: cfg, EthUSD: eth, UsdCad: usd, SendgridClient: sendgrid.NewSendClient(cfg.SendgridAPIKey), DB: dbm.DB}, nil
	}
	return &Manager{Config: cfg, EthUSD: eth, UsdCad: usd, SendgridClient: sendgrid.NewSendClient(cfg.SendgridAPIKey)}, nil

}

// GetPayouts is used to get payouts from ethermine
func (m *Manager) GetPayouts(minerAddress string) (*[]Payout, error) {
	url := fmt.Sprintf("%s/miner/%s/payouts", m.Config.URL, minerAddress)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var intf map[string]interface{}
	var data []Payout
	if err = json.Unmarshal(bodyBytes, &intf); err != nil {
		return nil, err
	}
	//fmt.Println(intf["data"])
	marshaled, err := json.Marshal(intf["data"])
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(marshaled, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
