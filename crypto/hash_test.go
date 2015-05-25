package crypto

import (
	"testing"

	. "github.com/wangch/ripple/testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type HashSuite struct{}

var _ = Suite(&HashSuite{})

var testAccounts = map[string]struct {
	Account, Secret string
}{
	"alice":    {"iG1QQv2nh2gi7RCZ1P8YYcBUKCCN633jCn", "alice"},
	"bob":      {"iPMh7Pr9ct699rZUTWaytJUoHcJ7cgyzrK", "bob"},
	"carol":    {"iH4KEcG9dEwGwpn6AyoWK9cZPLL4RLSmWW", "carol"},
	"dan":      {"iJ85Mok8YRNxSo7NnxKGrPuk29uAeZQqwZ", "dan"},
	"bitstamp": {"i4jKmc2nQb5yEU6eycefrNKGHTU5NQJASx", "bitstamp"},
	"mtgox":    {"iGrhwhaqU8g7ahwAvTq6rX5ivsfcbgZw6v", "mtgox"},
	"amazon":   {"ihheXqX7bDnXePJeMHhubDDvw2uUTtenPd", "amazon"},
	"root":     {"iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "masterpassphrase"},
}

var accountTests = TestSlice{
	{accountCheck("0").Value().String(), Equals, "0", "Parse 0"},
	{accountCheck("0").String(), Equals, ACCOUNT_ZERO, "Parse 0 export"},
	{accountCheck("1").Value().String(), Equals, "1", "Parse 1"},
	{accountCheck("1").String(), Equals, ACCOUNT_ONE, "Parse 1 export"},
	{accountCheck(ACCOUNT_ZERO).String(), Equals, ACCOUNT_ZERO, "Parse iiiiiiiiiiiiiiiiiiiiihoLvTp export"},
	{accountCheck(ACCOUNT_ONE).String(), Equals, ACCOUNT_ONE, "Parse iiiiiiiiiiiiiiiiiiiiBZbvjr export"},
	{accountCheck(testAccounts["mtgox"].Account).String(), Equals, testAccounts["mtgox"].Account, "Parse mtgox export"},
	{accountCheck(ACCOUNT_ZERO), Not(Equals), nil, "IsValid iiiiiiiiiiiiiiiiiiiiihoLvTp"},
	{ErrorCheck(NewRippleHash("iiiiiiiiiiiiiiiiiiiiihoLvT")), ErrorMatches, "Bad Base58 checksum:.*", "IsValid iiiiiiiiiiiiiiiiiiiiihoLvT"},
}

func accountCheck(v interface{}) Hash {
	if a, err := NewRippleHash(v.(string)); err != nil {
		println(v.(string))
		panic(err)
	} else {
		return a
	}
}

func (s *HashSuite) TestHashes(c *C) {
	accountTests.Test(c)
}
