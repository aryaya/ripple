package crypto

import (
	"encoding/hex"
	"fmt"

	. "gopkg.in/check.v1"
	"testing"
)

type KeySuite struct{}

var _ = Suite(&KeySuite{})

func checkHash(h Hash, err error) string {
	if err != nil {
		panic(err)
	}
	return h.String()
}

func checkSignature(c *C, privateKey, publicKey, hash, msg []byte) bool {
	sig, err := Sign(privateKey, hash, msg)
	c.Assert(err, IsNil)
	ok, err := Verify(publicKey, hash, msg, sig)
	c.Assert(err, IsNil)
	return ok
}

func b2h(b []byte) string {
	return fmt.Sprintf("%X", b)
}

func h2b(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return h
}

// Examples from https://ripple.com/wiki/Account_Family
func (s *KeySuite) TestWikiVectors(c *C) {
	zero, err := NewRippleHash("0")
	c.Check(err, IsNil)
	c.Check(zero.String(), Equals, ACCOUNT_ZERO)
	c.Check(b2h(Sha512Half(zero.PayloadTrimmed())), Equals, "B8244D028981D693AF7B456AF8EFA4CAD63D282E19FF14942C246E50D9351D22")

	seed := h2b("71ED064155FFADFA38782C5E0158CB26")
	key, err := NewECDSAKey(seed)
	c.Check(err, IsNil)
	var sequenceZero uint32
	c.Check(b2h(key.Private(nil)), Equals, "7CFBA64F771E93E817E15039215430B53F7401C34931D111EAB3510B22DBB0D8")
	c.Check(checkHash(AccountId(key, &sequenceZero)), Equals, "ihcfR9Cg98qCxHpCcPBmMonbDBXo84wyTn")
	c.Check(checkHash(NodePublicKey(key)), Equals, "n9MXXueo837zYH36DvMc13BwHcqtfAWNJY5czWVbp7uYTj7x17TH")
	c.Check(checkHash(NodePrivateKey(key)), Equals, "pa91wmE8V8K63SAMGMpdFprk8wGAcbUdSmHABccV9jFfqhTrjH1")
	c.Check(checkHash(AccountPublicKey(key, &sequenceZero)), Equals, "aBRoQrbr2jpDofohooFuzZr9nEzKw9Zdfc4ExVNmuXHaJpSPh8uJ")
	c.Check(checkHash(AccountPrivateKey(key, &sequenceZero)), Equals, "pwMPbuE25inajrgDPBEh9Pwv8bMV2ebN9gVPTWTh4c3DtB14rGL")
}

// Examples from https://github.com/ripple/rippled/blob/develop/src/ripple_data/protocol/RippleAddress.cpp
func (s *KeySuite) TestRippledVectors(c *C) {
	seed, err := GenerateFamilySeed("masterpassphrase")
	c.Check(err, IsNil)
	c.Check(seed.String(), Equals, "snoPBiXtMeMyMHUVTgbuqAfg1SUTb")
	key, err := NewECDSAKey(seed.Payload())
	c.Check(err, IsNil)
	sequenceZero, sequenceOne := uint32(0), uint32(1)
	c.Check(checkHash(NodePublicKey(key)), Equals, "n94a1u4jAz288pZLtw6yFWVbr89YamrC6JBXPVUj5zmExe5fTVg9")
	c.Check(checkHash(NodePrivateKey(key)), Equals, "pnen77YEeUd4fFKG7rycBWcwKpTaeFRkW2WFostaATy1DSupwXe")
	c.Check(checkHash(AccountId(key, &sequenceZero)), Equals, "iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh")
	c.Check(checkHash(AccountPublicKey(key, &sequenceZero)), Equals, "aBQG8RQAzjs1eTKFEAQXi2gS4utcDrEC9wmr7pfUPTr27VCahwgw")
	c.Check(checkHash(AccountPrivateKey(key, &sequenceZero)), Equals, "p9JfM6HHr64m6mvB6v5k7G2b1cXzGmYrCNJf6GHPKvFTWdeRVjh")
	c.Check(checkHash(AccountId(key, &sequenceOne)), Equals, "i4bYF7SLUMD7QgSLLpgJx38WJSY12VrRjP")
	c.Check(checkHash(AccountPublicKey(key, &sequenceOne)), Equals, "aBPXpTfuLy1Bhk3HnGTTAqnovpKWQ23NpFMNkAF6F1Atg5vDyPiw")
	c.Check(checkHash(AccountPrivateKey(key, &sequenceOne)), Equals, "p9JEm822LMizJrr1k7TvdphfENTp6G5ji253Xa5ikzUWVi8ogQt")

	msg := []byte("Hello, nurse!")
	hash := Sha512Half(msg)
	c.Check(checkSignature(c, key.Private(nil), key.Public(nil), hash, msg), Equals, true)
	c.Check(checkSignature(c, key.Private(&sequenceZero), key.Public(&sequenceZero), hash, msg), Equals, true)
	c.Check(checkSignature(c, key.Private(&sequenceOne), key.Public(&sequenceOne), hash, msg), Equals, true)
	c.Check(checkSignature(c, key.Private(&sequenceOne), key.Public(&sequenceZero), hash, msg), Equals, false)
	c.Check(checkSignature(c, key.Private(&sequenceZero), key.Public(&sequenceOne), hash, msg), Equals, false)

}

func (s *KeySuite) TestEd25119(c *C) {
	seed, err := GenerateFamilySeed("masterpassphrase")
	c.Check(err, IsNil)
	c.Check(seed.String(), Equals, "snoPBiXtMeMyMHUVTgbuqAfg1SUTb")
	key, err := NewEd25519Key(seed.Payload())
	c.Check(err, IsNil)
	c.Check(checkHash(NodePublicKey(key)), Equals, "nHUeeJCSY2dM71oxM8Cgjouf5ekTuev2mwDpc374aLMxzDLXNmjf")
	// c.Check(checkHash(NodePrivateKey(key)), Equals, "pnen77YEeUd4fFKG7iycBWcwKpTaeFRkW2WFostaATy1DSupwXe") // Needs a new version encoding
	c.Check(checkHash(AccountId(key, nil)), Equals, "iGWiZyQqhTp9Xu7G5Pkayo7bXjH4k4QYpf")
	c.Check(checkHash(AccountPublicKey(key, nil)), Equals, "aKGheSBjmCsKJVuLNKRAKpZXT6wpk2FCuEZAXJupXgdAxX5THCqR")
	// c.Check(checkHash(AccountPrivateKey(key, nil)), Equals, "p9JfM6HHi64m6mvB6v5k7G2b1cXzGmYiCNJf6GHPKvFTWdeRVjh") //Needs a new version encoding

	other, err := NewEd25519Key(nil)
	c.Check(err, IsNil)

	msg := []byte("Hello, nurse!")
	hash := Sha512Half(msg)

	c.Check(checkSignature(c, key.Private(nil), key.Public(nil), hash, msg), Equals, true)
	c.Check(checkSignature(c, other.Private(nil), other.Public(nil), hash, msg), Equals, true)
	c.Check(checkSignature(c, key.Private(nil), other.Public(nil), hash, msg), Equals, false)
	c.Check(checkSignature(c, other.Private(nil), key.Public(nil), hash, msg), Equals, false)
}

func Test0(t *testing.T) {
	seed, err := GenerateFamilySeed("masterpassphrase")
	if err != nil {
		t.Error(err)
	}
	if seed.String() != "snoPBiXtMeMyMHUVTgbuqAfg1SUTb" {
		t.Error("!= snoPBiXtMeMyMHUVTgbuqAfg1SUTb")
	}
}
