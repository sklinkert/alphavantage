package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToIncomeStatements(t *testing.T) {
	var buf = `
	{
		"symbol": "IBM",
		"annualReports": [
			{
				"fiscalDateEnding": "2020-12-31",
				"reportedCurrency": "USD",
				"grossProfit": "35575000000",
				"totalRevenue": "73620000000",
				"costOfRevenue": "38046000000",
				"costofGoodsAndServicesSold": "439000000",
				"operatingIncome": "4609000000",
				"sellingGeneralAndAdministrative": "23082000000",
				"researchAndDevelopment": "6333000000",
				"operatingExpenses": "30966000000",
				"investmentIncomeNet": "None",
				"netInterestIncome": "-1288000000",
				"interestIncome": "105000000",
				"interestExpense": "1288000000",
				"nonInterestIncome": "None",
				"otherNonOperatingIncome": "-861000000",
				"depreciation": "4227000000",
				"depreciationAndAmortization": "2468000000",
				"incomeBeforeTax": "4726000000",
				"incomeTaxExpense": "-864000000",
				"interestAndDebtExpense": "1288000000",
				"netIncomeFromContinuingOperations": "5501000000",
				"comprehensiveIncomeNetOfTax": "4850000000",
				"ebit": "6014000000",
				"ebitda": "8482000000",
				"netIncome": "5590000000"
			}
		],
		"quarterlyReports": [
			{
				"fiscalDateEnding": "2021-06-30",
				"reportedCurrency": "USD",
				"grossProfit": "9004000000",
				"totalRevenue": "18745000000",
				"costOfRevenue": "9741000000",
				"costofGoodsAndServicesSold": "103000000",
				"operatingIncome": "2304000000",
				"sellingGeneralAndAdministrative": "5334000000",
				"researchAndDevelopment": "1657000000",
				"operatingExpenses": "6700000000",
				"investmentIncomeNet": "None",
				"netInterestIncome": "-281000000",
				"interestIncome": "11000000",
				"interestExpense": "281000000",
				"nonInterestIncome": "None",
				"otherNonOperatingIncome": "-315000000",
				"depreciation": "1050000000",
				"depreciationAndAmortization": "630000000",
				"incomeBeforeTax": "1552000000",
				"incomeTaxExpense": "227000000",
				"interestAndDebtExpense": "281000000",
				"netIncomeFromContinuingOperations": "1325000000",
				"comprehensiveIncomeNetOfTax": "1930000000",
				"ebit": "1833000000",
				"ebitda": "2463000000",
				"netIncome": "1325000000"
			}
		]
	}
`
	incomeStatements, err := toIncomeStatements([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "IBM", incomeStatements.Symbol)

	assert.EqualStrings(t, "2020-12-31", incomeStatements.AnnualReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", incomeStatements.AnnualReports[0].ReportedCurrency)
	assert.EqualInt(t, 35575000000, int(incomeStatements.AnnualReports[0].GrossProfit))
	assert.EqualInt(t, 73620000000, int(incomeStatements.AnnualReports[0].TotalRevenue))
	assert.EqualInt(t, 38046000000, int(incomeStatements.AnnualReports[0].CostOfRevenue))
	assert.EqualInt(t, 439000000, int(incomeStatements.AnnualReports[0].CostOfGoodsAndServicesSold))
	assert.EqualInt(t, 4609000000, int(incomeStatements.AnnualReports[0].OperatingIncome))
	assert.EqualInt(t, 23082000000, int(incomeStatements.AnnualReports[0].SellingGeneralAndAdministrative))
	assert.EqualInt(t, 6333000000, int(incomeStatements.AnnualReports[0].ResearchAndDevelopment))
	assert.EqualInt(t, 30966000000, int(incomeStatements.AnnualReports[0].OperatingExpenses))
	assert.EqualInt(t, 0, int(incomeStatements.AnnualReports[0].InvestmentIncomeNet))
	assert.EqualInt(t, -1288000000, int(incomeStatements.AnnualReports[0].NetInterestIncome))
	assert.EqualInt(t, 105000000, int(incomeStatements.AnnualReports[0].InterestIncome))
	assert.EqualInt(t, 1288000000, int(incomeStatements.AnnualReports[0].InterestExpense))
	assert.EqualInt(t, 0, int(incomeStatements.AnnualReports[0].NonInterestIncome))
	assert.EqualInt(t, -861000000, int(incomeStatements.AnnualReports[0].OtherNonOperatingIncome))
	assert.EqualInt(t, 4227000000, int(incomeStatements.AnnualReports[0].Depreciation))
	assert.EqualInt(t, 2468000000, int(incomeStatements.AnnualReports[0].DepreciationAndAmortization))
	assert.EqualInt(t, 4726000000, int(incomeStatements.AnnualReports[0].IncomeBeforeTax))
	assert.EqualInt(t, -864000000, int(incomeStatements.AnnualReports[0].IncomeTaxExpense))
	assert.EqualInt(t, 1288000000, int(incomeStatements.AnnualReports[0].InterestAndDebtExpense))
	assert.EqualInt(t, 5501000000, int(incomeStatements.AnnualReports[0].NetIncomeFromContinuingOperations))
	assert.EqualInt(t, 4850000000, int(incomeStatements.AnnualReports[0].ComprehensiveIncomeNetOfTax))
	assert.EqualInt(t, 6014000000, int(incomeStatements.AnnualReports[0].Ebit))
	assert.EqualInt(t, 8482000000, int(incomeStatements.AnnualReports[0].Ebitda))
	assert.EqualInt(t, 5590000000, int(incomeStatements.AnnualReports[0].NetIncome))

	assert.EqualStrings(t, "2021-06-30", incomeStatements.QuarterlyReports[0].FiscalDateEnding)
	assert.EqualStrings(t, "USD", incomeStatements.QuarterlyReports[0].ReportedCurrency)
	assert.EqualInt(t, 9004000000, int(incomeStatements.QuarterlyReports[0].GrossProfit))
	assert.EqualInt(t, 18745000000, int(incomeStatements.QuarterlyReports[0].TotalRevenue))
	assert.EqualInt(t, 9741000000, int(incomeStatements.QuarterlyReports[0].CostOfRevenue))
	assert.EqualInt(t, 103000000, int(incomeStatements.QuarterlyReports[0].CostOfGoodsAndServicesSold))
	assert.EqualInt(t, 2304000000, int(incomeStatements.QuarterlyReports[0].OperatingIncome))
	assert.EqualInt(t, 5334000000, int(incomeStatements.QuarterlyReports[0].SellingGeneralAndAdministrative))
	assert.EqualInt(t, 1657000000, int(incomeStatements.QuarterlyReports[0].ResearchAndDevelopment))
	assert.EqualInt(t, 6700000000, int(incomeStatements.QuarterlyReports[0].OperatingExpenses))
	assert.EqualInt(t, 0, int(incomeStatements.QuarterlyReports[0].InvestmentIncomeNet))
	assert.EqualInt(t, -281000000, int(incomeStatements.QuarterlyReports[0].NetInterestIncome))
	assert.EqualInt(t, 11000000, int(incomeStatements.QuarterlyReports[0].InterestIncome))
	assert.EqualInt(t, 281000000, int(incomeStatements.QuarterlyReports[0].InterestExpense))
	assert.EqualInt(t, 0, int(incomeStatements.QuarterlyReports[0].NonInterestIncome))
	assert.EqualInt(t, -315000000, int(incomeStatements.QuarterlyReports[0].OtherNonOperatingIncome))
	assert.EqualInt(t, 1050000000, int(incomeStatements.QuarterlyReports[0].Depreciation))
	assert.EqualInt(t, 630000000, int(incomeStatements.QuarterlyReports[0].DepreciationAndAmortization))
	assert.EqualInt(t, 1552000000, int(incomeStatements.QuarterlyReports[0].IncomeBeforeTax))
	assert.EqualInt(t, 227000000, int(incomeStatements.QuarterlyReports[0].IncomeTaxExpense))
	assert.EqualInt(t, 281000000, int(incomeStatements.QuarterlyReports[0].InterestAndDebtExpense))
	assert.EqualInt(t, 1325000000, int(incomeStatements.QuarterlyReports[0].NetIncomeFromContinuingOperations))
	assert.EqualInt(t, 1930000000, int(incomeStatements.QuarterlyReports[0].ComprehensiveIncomeNetOfTax))
	assert.EqualInt(t, 1833000000, int(incomeStatements.QuarterlyReports[0].Ebit))
	assert.EqualInt(t, 2463000000, int(incomeStatements.QuarterlyReports[0].Ebitda))
	assert.EqualInt(t, 1325000000, int(incomeStatements.QuarterlyReports[0].NetIncome))
}
