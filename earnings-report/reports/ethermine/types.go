package ethermine

// Payout is a single payout as returned by the get payouts call
type Payout struct {
	PaidOn int64  `json:"paidOn"`
	Start  int64  `json:"start"`
	End    int64  `json:"end"`
	Amount int64  `json:"amount"`
	TxHash string `json:"txHash"`
}
