package data

import (
	"fmt"
	"testing"

	. "github.com/wangch/ripple/testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type AmountSuite struct{}

var _ = Suite(&AmountSuite{})

var amountTests = TestSlice{
	// {amountCheck("0").Add(amountCheck("-1")).ToHuman(), Equals, "-1", "Negatives"},
	{amountCheck("1").IsPositive(), Equals, true, "Positives"},
	// {amountCheck(int64(1)).String(), Equals, "1/1/rrrrrrrrrrrrrrrrrrrrBZbvji", "FromNumber"}, //WHY?
	{amountCheck(int64(1)).String(), Equals, "0.000001/ICC", "int64(1) String"},
	{amountCheck("1/ICC").String(), Equals, "1/ICC", "Parse 1/ICC"},
	{amountCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL").String(), Equals, "1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "Parse 1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"},
	{amountCheck("10/015841551A748AD2C1F76FF6ECB0CCCD00000000/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Not(Equals), "10/XAU/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Demurrage"},
	{amountCheck("0").String(), Equals, "0/ICC", "Parse native 0"},
	{amountCheck("0.0").String(), Equals, "0/ICC", "Parse native 0.0"},
	{amountCheck("-0").String(), Equals, "0/ICC", "Parse native -0"},
	{amountCheck("-0.0").String(), Equals, "0/ICC", "Parse native -0.0"},
	{amountCheck("1000").String(), Equals, "0.001/ICC", "Parse native 1000"},
	{amountCheck("1234").String(), Equals, "0.001234/ICC", "Parse native 1234"},
	{amountCheck("12.3").String(), Equals, "12.3/ICC", "Parse native 12.3"},
	{amountCheck("-12.3").String(), Equals, "-12.3/ICC", "Parse native -12.3"},
	{amountCheck("123./USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse 123./USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("12300/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "12300/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse 12300/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("12.3/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "12.3/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse 12.3/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("1.2300/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "1.23/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse 1.2300/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("-0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse -0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("-0.0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Parse -0.0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("123").Negate().String(), Equals, "-0.000123/ICC", "Negate native 123"},
	{amountCheck("-123").Negate().String(), Equals, "0.000123/ICC", "Negate native -123"},
	{amountCheck("123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").Negate().String(), Equals, "-123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Negate 123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("-123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").Negate().String(), Equals, "123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Negate -123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{amountCheck("-123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").Clone().String(), Equals, "-123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Clone -123/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh"},
	{addCheck("150", "50").String(), Equals, "0.0002/ICC", "Add ICC to ICC"},
	{addCheck("150.02/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "50.5/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "200.52/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Add USD to USD"},
	{addCheck("0/USD", "1/USD").String(), Equals, "1/USD", "Add 0 USD to 1 USD"},
	{ErrorCheck(amountCheck("1/ICC").Add(amountCheck("1/USD"))), ErrorMatches, "Cannot add.*", "Add 1 ICC to 1 USD"},
	{subCheck("150.02/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "50.5/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "99.52/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Subtract USD from USD"},
	{mulCheck("0", "0").String(), Equals, "0/ICC", "Multiply 0 ICC with 0 ICC"},
	{mulCheck("0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "0").String(), Equals, "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply 0 USD with 0 ICC"},
	{mulCheck("0", "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0/ICC", "Multiply 0 ICC with 0 USD"},
	{mulCheck("1", "0").String(), Equals, "0/ICC", "Multiply 1 ICC with 0 ICC"},
	{mulCheck("1/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "0").String(), Equals, "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply 1 USD with 0 ICC"},
	{mulCheck("1", "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0/ICC", "Multiply 1 ICC with 0 USD"},
	{mulCheck("0", "1").String(), Equals, "0/ICC", "Multiply 0 ICC with 1 ICC"},
	{mulCheck("0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "1").String(), Equals, "0/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply 0 USD with 1 ICC"},
	{mulCheck("0", "1/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0/ICC", "Multiply 0 ICC with 1 USD"},
	{mulCheck("200", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.002/ICC", "Multiply ICC with USD"},
	{mulCheck("20000", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.2/ICC", "Multiply ICC with USD"},
	{mulCheck("2000000", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "20/ICC", "Multiply ICC with USD"},
	{mulCheck("200", "-10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-0.002/ICC", "Multiply ICC with USD, neg"},
	{mulCheck("-6000", "37/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-0.222/ICC", "Multiply ICC with USD, neg, frac"},
	{mulCheck("2000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "20000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply USD with USD"},
	{mulCheck("2000000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "100000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "2e11/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply USD with USD"},
	{mulCheck("100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "1000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "100000/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with USD, result < 1"},
	{mulCheck("-24000/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-48000000/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with USD, neg"},
	{mulCheck("0.1/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "-1000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with USD, neg, <1"},
	{mulCheck("0.05/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000").String(), Equals, "100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with ICC, factor < 1"},
	{mulCheck("-100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "5").String(), Equals, "-500/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with ICC, neg"},
	{mulCheck("-0.05/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000").String(), Equals, "-100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply EUR with ICC, neg, <1"},
	{mulCheck("10", "10").String(), Equals, "0.0001/ICC", "Multiply ICC with ICC"},
	{mulCheck("2000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "10/015841551A748AD2C1F76FF6ECB0CCCD00000000/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Not(Equals), "20000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Multiply USD with XAU (demurred)"},
	{divCheck("200", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.00002/ICC", "Divide ICC by USD"},
	{divCheck("20000", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.002/ICC", "Divide ICC by USD"},
	{divCheck("2000000", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.2/ICC", "Divide ICC by USD"},
	{divCheck("200", "-10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-0.00002/ICC", "Divide ICC by USD, neg"},
	{divCheck("-6000", "37/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-0.000162/ICC", "Divide ICC by USD, neg, frac"},
	{divCheck("2000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "10/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "200/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide USD by USD"},
	{divCheck("2000000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "35/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "57142.85714285714/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide USD by USD, fractional"},
	{divCheck("2000000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "100000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "20/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide USD by USD"},
	{divCheck("100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "1000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "0.1/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by USD, factor < 1"},
	{divCheck("-24000/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-12/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by USD, neg"},
	{divCheck("100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "-1000/USD/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh").String(), Equals, "-0.1/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by USD, neg, <1"},
	{divCheck("100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000").String(), Equals, "0.05/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by ICC, result < 1"},
	{divCheck("-100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "5").String(), Equals, "-20/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by ICC, neg"},
	{divCheck("-100/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "2000").String(), Equals, "-0.05/EUR/iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "Divide EUR by ICC, neg, <1"},
	{equalCheck("0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, true, "0 USD == 0 USD"},
	{equalCheck("0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, true, "0 USD == 0 USD"},
	{equalCheck("0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "-0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, true, "0 USD == -0 USD"},
	{equalCheck("0", "0.0"), Equals, true, "0 ICC == 0 ICC"},
	{equalCheck("0", "-0"), Equals, true, "0 ICC == -0 ICC"},
	{equalCheck("10/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "10/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, true, "10 USD == 10 USD"},
	{equalCheck("123.4567/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "123.4567/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, true, "123.4567 USD == 123.4567 USD"},
	{equalCheck("10", "10"), Equals, true, "10 ICC == 10 ICC"},
	// {equalCheck("1.1", "11.0").ratio_human(10,false),Equals,true, "1.1 ICC == 1.1 ICC"},
	{amountCheck("0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL").SameValue(amountCheck("0/USD/iH5aWQJ4R7v4Mpyf4kDBUvDFT5cbpFq3XP")), Equals, true, "0 USD == 0 USD (ignore issuer)"},
	{amountCheck("1.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL").SameValue(amountCheck("1.10/USD/iH5aWQJ4R7v4Mpyf4kDBUvDFT5cbpFq3XP")), Equals, true, "1.1 USD == 1.10 USD (ignore issuer)"},
	{equalCheck("10/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "100/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "10 USD != 100 USD"},
	{equalCheck("10", "100"), Equals, false, "10 ICC != 100 ICC"},
	{equalCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "2/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "1 USD != 2 USD"},
	{equalCheck("1", "2"), Equals, false, "1 ICC != 2 ICC"},
	{equalCheck("0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "0.2/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "0.1 USD != 0.2 USD"},
	{equalCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "-1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "1 USD != -1 USD"},
	{equalCheck("1", "-1"), Equals, false, "1 ICC != -1 ICC"},
	{equalCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "1/USD/iH5aWQJ4R7v4Mpyf4kDBUvDFT5cbpFq3XP"), Equals, false, "1 USD != 1 USD (issuer mismatch)"},
	{equalCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "1/EUR/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "1 USD != 1 EUR"},
	{equalCheck("1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "1"), Equals, false, "1 USD != 1 ICC"},
	{equalCheck("1", "1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"), Equals, false, "1 ICC != 1 USD"},
	{ErrorCheck(amountCheck("1").Divide(amountCheck("0"))), ErrorMatches, "Division by zero", "Divide one by zero"},
	{amountCheck("-1/ICC").Abs().String(), Equals, "1/ICC", "Abs -1"},
	// {ErrorCheck(NewAmount("xx")), ErrorMatches, "Bad amount:.*", "IsValid xx"},
	{ErrorCheck(NewAmount(nil)), ErrorMatches, "Bad type:.*", "IsValid nil"},
	{ErrorCheck(NewAmount(int(1))), ErrorMatches, "Bad type:.*", "IsValid int(0)"},

	{checkBinaryMarshal(amountCheck("0/ICC")).String(), Equals, "0/ICC", "Binary Marshal 0/ICC"},
	{checkBinaryMarshal(amountCheck("0.1/ICC")).String(), Equals, "0.1/ICC", "Binary Marshal 0.1/ICC"},
	{checkBinaryMarshal(amountCheck("-0.1/ICC")).String(), Equals, "-0.1/ICC", "Binary Marshal -0.1/ICC"},
	{checkBinaryMarshal(amountCheck("0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL")).String(), Equals, "0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "Binary Marshal 0/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"},
	{checkBinaryMarshal(amountCheck("0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL")).String(), Equals, "0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "Binary Marshal 0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"},
	{checkBinaryMarshal(amountCheck("-0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL")).String(), Equals, "-0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL", "Binary Marshal -0.1/USD/iNDKeo9RiCrRdfsMG8AdoZvNZxHASGzbZL"},
}

func subCheck(a, b string) *Amount {
	if sum, err := amountCheck(a).Subtract(amountCheck(b)); err != nil {
		panic(err)
	} else {
		return sum
	}
}

func addCheck(a, b string) *Amount {
	if sum, err := amountCheck(a).Add(amountCheck(b)); err != nil {
		panic(err)
	} else {
		return sum
	}
}

func mulCheck(a, b string) *Amount {
	if product, err := amountCheck(a).Multiply(amountCheck(b)); err != nil {
		panic(err)
	} else {
		return product
	}
}

func divCheck(a, b string) *Amount {
	if quotient, err := amountCheck(a).Divide(amountCheck(b)); err != nil {
		panic(err)
	} else {
		return quotient
	}
}

func amountCheck(v interface{}) *Amount {
	if a, err := NewAmount(v); err != nil {
		panic(err)
	} else {
		return a
	}
}

func equalCheck(a, b string) bool {
	return amountCheck(a).Equals(*amountCheck(b))
}

func (s *AmountSuite) TestAmount(c *C) {
	amountTests.Test(c)
}

func ExampleValue_Add() {
	v1, _ := NewValue("100", false)
	v2, _ := NewValue("200.199", false)
	sum, _ := v1.Add(*v2)
	fmt.Println(v1.String())
	fmt.Println(v2.String())
	fmt.Println(sum.String())
	// Output:
	// 100
	// 200.199
	// 300.199
}

func checkBinaryMarshal(v1 *Amount) *Amount {
	var b []byte
	var err error

	if b, err = v1.MarshalBinary(); err != nil {
		panic(err)
	}

	v2 := &Amount{}
	if err = v2.UnmarshalBinary(b); err != nil {
		panic(err)
	}

	return v2
}
