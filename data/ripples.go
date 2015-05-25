package data

import (
	"fmt"
	"sort"
)

// Trade
type Trade struct {
	Buyer  Account
	Seller Account
	Paid   Amount
	Got    Amount
}

func (t Trade) Price() *Value {
	price, err := t.Paid.Value.Ratio(*t.Got.Value)
	if err != nil {
		return nil
	}
	return price
}

// Transfer is a directional representation of a RippleState or AccountRoot balance change.
// Payments and OfferCreates lead to the creation of zero or more Transfers.
//
// 	TransitFee is earned by the Issuer
// 	QualityIn and QualityOut are earned by the Liquidity Provider and can be negative.
//
// Four scenarios:
// 	1. ICC -> ICC
// 	2. ICC -> IOU/Issuer 			Requires an orderbook
// 	3. IOU/Issuer -> ICC			Requires an orderbook
// 	4. IOU/IssuerA <-> IOU/IssuerB		Also known as Rippling, requires an account which trusts both currency/issuer pairs
type Transfer struct {
	Source             Account
	Destination        Account
	SourceBalance      Amount
	DestinationBalance Amount
	Change             Value
	TransitFee         *Value // Applies to all transfers except ICC -> ICC
	QualityIn          *Value // Applies to IOU -> IOU transfers
	QualityOut         *Value // Applies to IOU -> IOU transfers
}

type Balance struct {
	Account  Account
	Balance  Value
	Change   Value
	Currency Currency
}

// func (t Trade) String() string {
// 	return fmt.Sprintf("%-34s=>%-34s %18s@%18s", t.Currency, t.Issuer, t.Seller, t.Buyer, t.Bought.Value, t.Price())
// }

func (b Balance) String() string {
	return fmt.Sprintf("Account: %-34s  Currency: %s Balance: %20s Change: %20s", b.Account, b.Currency, b.Balance, b.Change)
}

type TradeSlice []Trade
type BalanceSlice []Balance

func (s TradeSlice) Len() int           { return len(s) }
func (s TradeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TradeSlice) Less(i, j int) bool { return s[i].Price().Less(*s[j].Price()) }

func (s BalanceSlice) Len() int      { return len(s) }
func (s BalanceSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s BalanceSlice) Less(i, j int) bool {
	switch {
	case !s[i].Currency.Equals(s[j].Currency):
		return s[i].Currency.Less(s[j].Currency)
	case s[i].Change.Abs().Equals(*s[j].Change.Abs()):
		return s[i].Change.negative != s[j].Change.negative
	default:
		return s[i].Change.Abs().Less(*s[j].Change.Abs())
	}
}

func (s *TradeSlice) Add(buyer, seller *Account, paid, got *Amount) {
	*s = append(*s, Trade{*buyer, *seller, *paid, *got})
}

func (s TradeSlice) Sum() (*Amount, error) {
	if len(s) == 0 {
		return nil, nil
	}
	return nil, nil
	// sum, err := NewAmount(fmt.Sprintf("0/%s/%s", s[0].Currency, s[0].Issuer))
	// if err != nil {
	// 	return nil, err
	// }
	// for _, trade := range s {
	// 	if sum.Value, err = sum.Value.Add(trade.Amount); err != nil {
	// 		return nil, err
	// 	}
	// }
	// return sum, nil
}

func (s *BalanceSlice) Add(account *Account, balance, change *Value, currency *Currency) {
	*s = append(*s, Balance{*account, *balance, *change, *currency})
}

func (txm *TransactionWithMetaData) Trades() (TradeSlice, error) {
	if txm.GetTransactionType() != OFFER_CREATE && txm.GetTransactionType() != PAYMENT {
		return nil, nil
	}
	var (
		trades  TradeSlice
		account = txm.Transaction.GetBase().Account
	)
	for _, node := range txm.MetaData.AffectedNodes {
		var reason = ""
		switch {
		case node.CreatedNode != nil && node.CreatedNode.LedgerEntryType == OFFER:
			// No actual side effect
		case node.DeletedNode != nil && node.DeletedNode.LedgerEntryType == OFFER:
			// An OfferCreate specifying previous OfferSequence has no side effect
			if node.DeletedNode.PreviousFields == nil {
				// reason = "No side effect"
				break
			}
			// Fully consumed offer
			var (
				previous = node.DeletedNode.PreviousFields.(*Offer)
				final    = node.DeletedNode.FinalFields.(*Offer)
			)
			if previous.TakerPays == nil {
				reason = "Deleted Offer PreviousFields missing TakerPays"
				break
			}
			if final.TakerPays == nil {
				reason = "Deleted Offer FinalFields missing TakerPays"
				break
			}
			if previous.TakerGets == nil {
				reason = "Deleted Offer PreviousFields missing TakerGets"
				break
			}
			if final.TakerGets == nil {
				reason = "Deleted Offer FinalFields missing TakerGets"
				break
			}
			trades.Add(&account, final.Account, previous.TakerPays, previous.TakerGets)
		case node.ModifiedNode != nil && node.ModifiedNode.LedgerEntryType == OFFER:
			// No change?
			if node.ModifiedNode.PreviousFields == nil {
				reason = "no change"
				break
			}
			// Partially consumed offer
			var (
				previous = node.ModifiedNode.PreviousFields.(*Offer)
				current  = node.ModifiedNode.FinalFields.(*Offer)
			)
			if previous.TakerPays == nil {
				reason = "Modified Offer PreviousFields missing TakerPays"
				break
			}
			if current.TakerPays == nil {
				reason = "Modified Offer FinalFields missing TakerPays"
				break
			}
			paid, err := previous.TakerPays.Subtract(current.TakerPays)
			if err != nil {
				return nil, err
			}
			if previous.TakerGets == nil {
				reason = "Modified Offer PreviousFields missing TakerGets"
				break
			}
			if current.TakerGets == nil {
				reason = "Modified Offer FinalFields missing TakerGets"
				break
			}
			got, err := previous.TakerGets.Subtract(current.TakerGets)
			if err != nil {
				return nil, err
			}
			trades.Add(&account, current.Account, paid, got)
		}
		if reason != "" {
			fmt.Println(txm.LedgerSequence, txm.GetHash().String(), reason)
		}
	}
	sort.Sort(trades)
	return trades, nil
}

func (txm *TransactionWithMetaData) Balances() (BalanceSlice, error) {
	if txm.GetTransactionType() != OFFER_CREATE && txm.GetTransactionType() != PAYMENT {
		return nil, nil
	}
	var (
		balances BalanceSlice
		account  = txm.Transaction.GetBase().Account
	)
	for _, node := range txm.MetaData.AffectedNodes {
		switch {
		case node.CreatedNode != nil:
			switch node.CreatedNode.LedgerEntryType {
			case ACCOUNT_ROOT:
				created := node.CreatedNode.NewFields.(*AccountRoot)
				balances.Add(created.Account, &zeroNative, created.Balance, &zeroCurrency)
			case RIPPLE_STATE:
				// New trust line
				state := node.CreatedNode.NewFields.(*RippleState)
				balances.Add(&account, &zeroNonNative, state.Balance.Value, &state.Balance.Currency)
			}
		case node.DeletedNode != nil:
			switch node.DeletedNode.LedgerEntryType {
			case RIPPLE_STATE:
				//?
			case ACCOUNT_ROOT:
				return nil, fmt.Errorf("Deleted AccountRoot!")
			}
		case node.ModifiedNode != nil:
			if node.ModifiedNode.PreviousFields == nil {
				// No change
				continue
			}
			switch node.ModifiedNode.LedgerEntryType {
			case ACCOUNT_ROOT:
				// Changed ICC Balance
				var (
					previous = node.ModifiedNode.PreviousFields.(*AccountRoot)
					current  = node.ModifiedNode.FinalFields.(*AccountRoot)
				)
				if previous.Balance == nil {
					// ownercount change
					continue
				}
				change, err := NewAmount(int64(current.Balance.num - previous.Balance.num))
				if err != nil {
					return nil, err
				}
				// Add fee and see if change is non-zero
				if current.Account.Equals(account) {
					change.Value, err = change.Value.Add(txm.GetBase().Fee)
					if err != nil {
						return nil, err
					}
				}
				if change.num != 0 {
					balances.Add(current.Account, current.Balance, change.Value, &zeroCurrency)
				}
			case RIPPLE_STATE:
				// Changed non-native balance
				var (
					previous = node.ModifiedNode.PreviousFields.(*RippleState)
					current  = node.ModifiedNode.FinalFields.(*RippleState)
				)
				if previous.Balance == nil {
					//flag change
					continue
				}
				change, err := current.Balance.Subtract(previous.Balance)
				if err != nil {
					return nil, err
				}
				balances.Add(&current.HighLimit.Issuer, current.Balance.Value, change.Value, &current.Balance.Currency)
				balances.Add(&current.LowLimit.Issuer, current.Balance.Value.Negate(), change.Value.Negate(), &current.Balance.Currency)
			}
		}
	}
	sort.Sort(balances)
	return balances, nil
}
