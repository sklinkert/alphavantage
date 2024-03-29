package alphavantage

import (
	"github.com/AMekss/assert"
	"testing"
)

func TestToCompanyOverview(t *testing.T) {
	var buf = `
{
    "Symbol": "IBM",
    "AssetType": "Common Stock",
    "Name": "International Business Machines",
    "Description": "International Business Machines Corporation (IBM) is an American multinational technology company headquartered in Armonk, New York, with operations in over 170 countries. The company began in 1911, founded in Endicott, New York, as the Computing-Tabulating-Recording Company (CTR) and was renamed International Business Machines in 1924. IBM is incorporated in New York. IBM produces and sells computer hardware, middleware and software, and provides hosting and consulting services in areas ranging from mainframe computers to nanotechnology. IBM is also a major research organization, holding the record for most annual U.S. patents generated by a business (as of 2020) for 28 consecutive years. Inventions by IBM include the automated teller machine (ATM), the floppy disk, the hard disk drive, the magnetic stripe card, the relational database, the SQL programming language, the UPC barcode, and dynamic random-access memory (DRAM). The IBM mainframe, exemplified by the System/360, was the dominant computing platform during the 1960s and 1970s.",
    "CIK": "51143",
    "Exchange": "NYSE",
    "Currency": "USD",
    "Country": "USA",
    "Sector": "TECHNOLOGY",
    "Industry": "COMPUTER & OFFICE EQUIPMENT",
    "Address": "1 NEW ORCHARD ROAD, ARMONK, NY, US",
    "FiscalYearEnd": "December",
    "LatestQuarter": "2023-03-31",
    "MarketCapitalization": "112577577000",
    "EBITDA": "12644000000",
    "PERatio": "55.33",
    "PEGRatio": "1.276",
    "BookValue": "23.79",
    "DividendPerShare": "6.6",
    "DividendYield": "0.0544",
    "EPS": "2.22",
    "RevenuePerShareTTM": "66.97",
    "ProfitMargin": "0.0303",
    "OperatingMarginTTM": "0.132",
    "ReturnOnAssetsTTM": "0.0376",
    "ReturnOnEquityTTM": "0.101",
    "RevenueTTM": "60585001000",
    "GrossProfitTTM": "32688000000",
    "DilutedEPSTTM": "2.22",
    "QuarterlyEarningsGrowthYOY": "0.253",
    "QuarterlyRevenueGrowthYOY": "0.004",
    "AnalystTargetPrice": "140.79",
    "TrailingPE": "55.33",
    "ForwardPE": "15.55",
    "PriceToSalesRatioTTM": "2.108",
    "PriceToBookRatio": "6.75",
    "EVToRevenue": "2.969",
    "EVToEBITDA": "25.81",
    "Beta": "0.851",
    "52WeekHigh": "149.31",
    "52WeekLow": "111.29",
    "50DayMovingAverage": "126.63",
    "200DayMovingAverage": "133.09",
    "SharesOutstanding": "908045000",
    "DividendDate": "2023-06-10",
    "ExDividendDate": "2023-05-09"
}
`
	var companyOverview, err = toCompanyOverview([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "IBM", companyOverview.Symbol)
	assert.EqualStrings(t, "Common Stock", companyOverview.AssetType)
	assert.EqualStrings(t, "International Business Machines", companyOverview.Name)
	assert.EqualStrings(t, "51143", companyOverview.CIK)
	assert.EqualStrings(t, "NYSE", companyOverview.Exchange)
	assert.EqualStrings(t, "USD", companyOverview.Currency)
	assert.EqualInt(t, 112577577000, companyOverview.MarketCapitalization)
	assert.EqualInt(t, 12644000000, companyOverview.EBITDA)
	assert.EqualFloat64(t, 55.33, companyOverview.PERatio)
	assert.EqualFloat64(t, 1.276, companyOverview.PEGRatio)
	assert.EqualFloat64(t, 23.79, companyOverview.BookValue)
	assert.EqualFloat64(t, 6.6, companyOverview.DividendPerShare)
	assert.EqualFloat64(t, 0.0544, companyOverview.DividendYield)
	assert.EqualFloat64(t, 2.22, companyOverview.EPS)
	assert.EqualFloat64(t, 66.97, companyOverview.RevenuePerShareTTM)
	assert.EqualFloat64(t, 0.0303, companyOverview.ProfitMargin)
	assert.EqualFloat64(t, 0.132, companyOverview.OperatingMarginTTM)
	assert.EqualFloat64(t, 0.0376, companyOverview.ReturnOnAssetsTTM)
	assert.EqualFloat64(t, 0.101, companyOverview.ReturnOnEquityTTM)
	assert.EqualInt(t, 60585001000, companyOverview.RevenueTTM)
	assert.EqualInt(t, 32688000000, companyOverview.GrossProfitTTM)
	assert.EqualFloat64(t, 2.22, companyOverview.DilutedEPSTTM)
	assert.EqualFloat64(t, 0.253, companyOverview.QuarterlyEarningsGrowthYOY)
	assert.EqualFloat64(t, 0.004, companyOverview.QuarterlyRevenueGrowthYOY)
	assert.EqualFloat64(t, 140.79, companyOverview.AnalystTargetPrice)
	assert.EqualFloat64(t, 55.33, companyOverview.TrailingPE)
	assert.EqualFloat64(t, 15.55, companyOverview.ForwardPE)
	assert.EqualFloat64(t, 2.108, companyOverview.PriceToSalesRatioTTM)
	assert.EqualFloat64(t, 6.75, companyOverview.PriceToBookRatio)
	assert.EqualFloat64(t, 2.969, companyOverview.EVToRevenue)
	assert.EqualFloat64(t, 25.81, companyOverview.EVToEBITDA)
	assert.EqualFloat64(t, 0.851, companyOverview.Beta)
	assert.EqualFloat64(t, 149.31, companyOverview.Week52High)
	assert.EqualFloat64(t, 111.29, companyOverview.Week52Low)
	assert.EqualFloat64(t, 126.63, companyOverview.MovingAverage50Day)
	assert.EqualFloat64(t, 133.09, companyOverview.MovingAverage200Day)
	assert.EqualInt(t, 908045000, companyOverview.SharesOutstanding)
}
