package reports

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/config"
	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

/*
This is used to handle automated mining reports for cryptocurrency mining farms
The idea is to create an easy to use system that can be used by farm operators to create accurate book reports for the tax man
*/

const (
	USDAPI = "https://free.currencyconverterapi.com/api/v5/convert?q=USD_CAD&compact=y"
)

var methodList = []string{"24hour_credit", "credit"}

type Manager struct {
	Config         *config.Config   `json:"config"`
	EthUSD         float64          `json:"eth_usd"` // keeps track of the ETH->USD conversion ratio
	UsdCad         float64          `json:"usd_cad"` // keeps track of the USD -> USD conversion ratio
	SendgridClient *sendgrid.Client `json:"sendgrid_client"`
}

func GenerateReportManagerFromFile(path string) (*Manager, error) {
	cfg, err := config.LoadConfigFromFile(path)
	if err != nil {
		return nil, err
	}
	usd, err := ParseUSDCAD()
	if err != nil {
		return nil, err
	}
	eth, err := ParseETHUSD()
	if err != nil {
		return nil, err
	}
	return &Manager{Config: cfg, EthUSD: eth, UsdCad: usd, SendgridClient: sendgrid.NewSendClient(cfg.SendgridAPIKey)}, nil

}

func GenerateReportManager(coin, apikey string) (*Manager, error) {
	cfg, err := config.LoadConfig(coin, apikey)
	if err != nil {
		return nil, err
	}
	return &Manager{Config: cfg}, nil
}

// CreateReportAndSend is used to create and send a mining report
func (m *Manager) CreateReportAndSend(method string) error {
	switch method {
	case "24hour_credit":
		credit, err := m.GetRecentCredits24Hours()
		if err != nil {
			return err
		}
		usdValue := credit.Amount * m.EthUSD
		resp, err := m.SendEmail(credit.Amount, usdValue)
		if err != nil {
			return err
		}
		if resp != 202 {
			return fmt.Errorf("unacceptable return code, expected 200 got %v", resp)
		}
	case "credit":
		break
	default:
		return errors.New(fmt.Sprint("invalid method must be one of ", methodList))
	}
	return nil
}

func (m *Manager) GetRecentCredits24Hours() (*types.RecentCredits, error) {
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
	marshaled, err = json.Marshal(data.Data["recent_credits_24hours"])
	if err != nil {
		return nil, err
	}
	var credits types.RecentCredits
	err = json.Unmarshal(marshaled, &credits)
	if err != nil {
		return nil, err
	}
	return &credits, nil
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

func (m *Manager) SendEmail(ethMined, usdValue float64) (int, error) {
	content := fmt.Sprintf("<br>Eth Mined: %v<br>USD Value: %v", ethMined, usdValue)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "Ethereum Mining Report"
	to := mail.NewEmail("Postables", "postables@rtradetechnologies.com")

	mContent := mail.NewContent("text/html", content)
	mail := mail.NewV3MailInit(from, subject, to, mContent)

	response, err := m.SendgridClient.Send(mail)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}
