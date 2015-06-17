package main

import (
	// "encoding/json"
	"flag"
	// "log"
	"os"

	"github.com/wangch/ripple/terminal"
	"github.com/wangch/ripple/websockets"
)

func checkErr(err error, quit bool) {
	if err != nil {
		terminal.Println(err.Error(), terminal.Default)
		if quit {
			os.Exit(1)
		}
	}
}

var (
	host    = flag.String("host", "wss://local.icloud.com:19528", "websockets host to connect to")
	account = flag.String("acc", "iHb9CJAWyB4ij91VRWn96DkukG4bwdtyTh", "account address for monitor")
)

func main() {
	flag.Parse()
	r, err := websockets.NewRemote(*host)
	checkErr(err, true)

	// Subscribe to all streams
	_, err = r.Subscribe(false, false, false, false, []string{*account})
	checkErr(err, true)
	// terminal.Println(fmt.Sprint("Subscribed at: ", confirmation.LedgerSequence), terminal.Default)

	// Consume messages as they arrive
	for {
		msg, ok := <-r.Incoming
		if !ok {
			return
		}

		switch msg := msg.(type) {
		case *websockets.LedgerStreamMsg:
			terminal.Println(msg, terminal.Default)
		case *websockets.TransactionStreamMsg:
			terminal.Println(&msg.Transaction, terminal.Indent)
			for _, path := range msg.Transaction.PathSet() {
				terminal.Println(path, terminal.DoubleIndent)
			}
			trades, err := msg.Transaction.Trades()
			checkErr(err, false)
			for _, trade := range trades {
				terminal.Println(trade, terminal.DoubleIndent)
			}
			balances, err := msg.Transaction.Balances()
			checkErr(err, false)
			for _, balance := range balances {
				terminal.Println(balance, terminal.DoubleIndent)
			}
			// log.Println(msg.Transaction.GetType())
			// b, err := json.MarshalIndent(msg, "", "  ")
			// log.Println(string(b))

		case *websockets.ServerStreamMsg:
			terminal.Println(msg, terminal.Default)
		}
	}
}
