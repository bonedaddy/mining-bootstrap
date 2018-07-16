package types

import "github.com/jinzhu/gorm"

type EthReports struct {
	gorm.Model
	Date     string `gorm:"type:varchar(255)" json:"date"`
	ETHMined string `gorm:"type:varchar(255)" json:"eth_mined"`
	CADValue string `gorm:"type:varchar(255)" json:"cad_value"`
	USDValue string `gorm:"type:varchar(255)" json:"usd_value"`
}

type EthReportsManager struct {
	DB *gorm.DB
}

func NewEthReports(db *gorm.DB) *EthReportsManager {
	return &EthReportsManager{DB: db}
}

func (erm *EthReportsManager) FindByDate(date string) (*EthReports, error) {
	report := &EthReports{}
	if check := erm.DB.Where("date = ?", date).First(report); check.Error != nil {
		return nil, check.Error
	}
	return report, nil
}
func (erm *EthReportsManager) AddNewEntry(date, eth, cad, usd string) error {
	_, err := erm.FindByDate(date)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	report := &EthReports{
		Date:     date,
		ETHMined: eth,
		CADValue: cad,
		USDValue: usd,
	}
	if check := erm.DB.Create(report); check.Error != nil {
		return check.Error
	}
	return nil
}
