package alphavantage

import (
	"encoding/json"
	"fmt"
)

type IncomeStatements struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []AnnualReport    `json:"annualReports"`
	QuarterlyReports []QuarterlyReport `json:"quarterlyReports"`
}
type AnnualReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"` // Have to handle case where value is "None"
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

type QuarterlyReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       AVInt  `json:"grossProfit"`
	TotalRevenue                      AVInt  `json:"totalRevenue"`
	CostOfRevenue                     AVInt  `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        AVInt  `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   AVInt  `json:"operatingIncome"`
	SellingGeneralAndAdministrative   AVInt  `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            AVInt  `json:"researchAndDevelopment"`
	OperatingExpenses                 AVInt  `json:"operatingExpenses"`
	InvestmentIncomeNet               AVInt  `json:"investmentIncomeNet"`
	NetInterestIncome                 AVInt  `json:"netInterestIncome"`
	InterestIncome                    AVInt  `json:"interestIncome"`
	InterestExpense                   AVInt  `json:"interestExpense"`
	NonInterestIncome                 AVInt  `json:"nonInterestIncome"`
	OtherNonOperatingIncome           AVInt  `json:"otherNonOperatingIncome"`
	Depreciation                      AVInt  `json:"depreciation"`
	DepreciationAndAmortization       AVInt  `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   AVInt  `json:"incomeBeforeTax"`
	IncomeTaxExpense                  AVInt  `json:"incomeTaxExpense"`
	InterestAndDebtExpense            AVInt  `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations AVInt  `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       AVInt  `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              AVInt  `json:"ebit"`
	Ebitda                            AVInt  `json:"ebitda"`
	NetIncome                         AVInt  `json:"netIncome"`
}

func toIncomeStatements(buf []byte) (*IncomeStatements, error) {
	incomeStatements := &IncomeStatements{}
	if err := json.Unmarshal(buf, incomeStatements); err != nil {
		return nil, err
	}
	return incomeStatements, nil
}

func (c *Client) IncomeStatements(symbol string) (*IncomeStatements, error) {
	const function = "INCOME_STATEMENT"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s", baseURL, function, symbol, c.apiKey)
	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toIncomeStatements(body)
}
