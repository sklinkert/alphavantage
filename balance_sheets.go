package alphavantage

import (
	"encoding/json"
	"fmt"
)

type BalanceSheets struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}

type AnnualReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}

type QuarterlyReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            AVInt  `json:"totalAssets"`
	TotalCurrentAssets                     AVInt  `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  AVInt  `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            AVInt  `json:"cashAndShortTermInvestments"`
	Inventory                              AVInt  `json:"inventory"`
	CurrentNetReceivables                  AVInt  `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  AVInt  `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 AVInt  `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE AVInt  `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       AVInt  `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      AVInt  `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               AVInt  `json:"goodwill"`
	Investments                            AVInt  `json:"investments"`
	LongTermInvestments                    AVInt  `json:"longTermInvestments"`
	ShortTermInvestments                   AVInt  `json:"shortTermInvestments"`
	OtherCurrentAssets                     AVInt  `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  AVInt  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       AVInt  `json:"totalLiabilities"`
	TotalCurrentLiabilities                AVInt  `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 AVInt  `json:"currentAccountsPayable"`
	DeferredRevenue                        AVInt  `json:"deferredRevenue"`
	CurrentDebt                            AVInt  `json:"currentDebt"`
	ShortTermDebt                          AVInt  `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             AVInt  `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                AVInt  `json:"capitalLeaseObligations"`
	LongTermDebt                           AVInt  `json:"longTermDebt"`
	CurrentLongTermDebt                    AVInt  `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 AVInt  `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 AVInt  `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                AVInt  `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             AVInt  `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 AVInt  `json:"totalShareholderEquity"`
	TreasuryStock                          AVInt  `json:"treasuryStock"`
	RetainedEarnings                       AVInt  `json:"retainedEarnings"`
	CommonStock                            AVInt  `json:"commonStock"`
	CommonStockSharesOutstanding           AVInt  `json:"commonStockSharesOutstanding"`
}

func toBalanceSheets(buf []byte) (*BalanceSheets, error) {
	balanceSheets := &BalanceSheets{}
	if err := json.Unmarshal(buf, balanceSheets); err != nil {
		return nil, err
	}
	return balanceSheets, nil
}

func (c *Client) BalanceSheets(symbol string) (*BalanceSheets, error) {
	const function = "BALANCE_SHEET"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toBalanceSheets(body)
}
