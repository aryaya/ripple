package memdb

import (
	"github.com/wangch/ripple/data"
	"testing"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMemStore(t *testing.T) {
	mem, err := NewMemoryDB([]string{"testdata/mem.gz"})
	checkErr(t, err)
	h1, err := data.NewHash256("CAD2E1FDC45A01998C75A2F50D2DFF3B77CE1451F3F58A328D1323917AC72FD7")
	checkErr(t, err)
	n1, err := mem.Get(*h1)
	checkErr(t, err)
	if _, ok := n1.(data.LedgerEntry); !ok {
		t.Fatalf("Expected LedgerEntry Got:%+v", n1)
	}
	if h1.Compare(*n1.GetHash()) != 0 {
		t.Fatalf("Expected Hash: %s Got:%s", h1, n1.GetHash())
	}
}
