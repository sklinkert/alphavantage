package alphavantage

import (
	"encoding/json"
	"fmt"
)

type CashFlows struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}

type AnnualReport struct {
	FiscalDateEnding                      string `json:"fiscalDateEnding"`
	ReportedCurrency                      string `json:"reportedCurrency"`
	OperatingCashflow                     AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities        AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities       AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities          AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets               AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization  AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                   AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                   AVInt  `json:"changeInReceivables"`
	ChangeInInventory                     AVInt  `json:"changeInInventory"`
	ProfitLoss                            AVInt  `json:"profitLoss"`
	CashflowFromInvestment                AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                 AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock    AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity         AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                        AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock             AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock          AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock     AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebt    AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock  AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity        AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock       AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents        AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                  AVInt  `json:"changeInExchangeRate"`
	NetIncome                             AVInt  `json:"netIncome"`
}

type QuarterlyReport struct {
	FiscalDateEnding                      string `json:"fiscalDateEnding"`
	ReportedCurrency                      string `json:"reportedCurrency"`
	OperatingCashflow                     AVInt  `json:"operatingCashflow"`
	PaymentsForOperatingActivities        AVInt  `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities       AVInt  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities          AVInt  `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets               AVInt  `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization  AVInt  `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                   AVInt  `json:"capitalExpenditures"`
	ChangeInReceivables                   AVInt  `json:"changeInReceivables"`
	ChangeInInventory                     AVInt  `json:"changeInInventory"`
	ProfitLoss                            AVInt  `json:"profitLoss"`
	CashflowFromInvestment                AVInt  `json:"cashflowFromInvestment"`
	CashflowFromFinancing                 AVInt  `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt AVInt  `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock    AVInt  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity         AVInt  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock AVInt  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                        AVInt  `json:"dividendPayout"`
	DividendPayoutCommonStock             AVInt  `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock          AVInt  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock     AVInt  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebt    AVInt  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock  AVInt  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity        AVInt  `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock       AVInt  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents        AVInt  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                  AVInt  `json:"changeInExchangeRate"`
	NetIncome                             AVInt  `json:"netIncome"`
}

func toCashFlows(buf []byte) (*CashFlows, error) {
	cashFlows := &CashFlows{}
	if err := json.Unmarshal(buf, cashFlows); err != nil {
		return nil, fmt.Errorf("unable to unmarshal cash flows: %v", err)
	}
	return cashFlows, nil
}

func (c *Client) CashFlows(symbol string) (*CashFlows, error) {
	const function = "CASH_FLOW"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	return toCashFlows(body)
}
