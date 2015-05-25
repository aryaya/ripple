package data

type TransactionResult int16

const (
	// 0: S Success (success)
	// Causes:
	// - Success.
	// Implications:
	// - Applied
	// - Forwarded
	tesSUCCESS TransactionResult = 0
)
const (
	// 100 .. 159 C Claim fee only (ripple transaction with no good paths, pay to non-existent account, no path)
	// Causes:
	// - Success, but does not achieve optimal result.
	// - Invalid transaction or no effect, but claim fee to use the sequence number.
	// Implications:
	// - Applied
	// - Forwarded
	// Only allowed as a return code of appliedTransaction when !tapRetry. Otherwise, treated as terRETRY.
	//
	// DO NOT CHANGE THESE NUMBERS: They appear in ledger meta data.
	tecCLAIM TransactionResult = iota + 100
	tecPATH_PARTIAL
	tecUNFUNDED_ADD
	tecUNFUNDED_OFFER
	tecUNFUNDED_PAYMENT
	tecFAILED_PROCESSING
)
const (
	tecDIR_FULL TransactionResult = iota + 121
	tecINSUF_RESERVE_LINE
	tecINSUF_RESERVE_OFFER
	tecNO_DST
	tecNO_DST_INSUF_XRP
	tecNO_LINE_INSUF_RESERVE
	tecNO_LINE_REDUNDANT
	tecPATH_DRY
	tecUNFUNDED
	tecMASTER_DISABLED
	tecNO_REGULAR_KEY
	tecOWNERS
	tecNO_ISSUER
	tecNO_AUTH
	tecNO_LINE
	tecINSUFF_FEE
	tecFROZEN
	tecNO_TARGET
	tecNO_PERMISSION
	tecNO_ENTRY
	tecINSUFFICIENT_RESERVE
	tecNEED_MASTER_KEY
	tecDST_TAG_NEEDED
	tecINTERNAL
)

const (
	// Transaction Errors
	// -399 .. -300: L Local error (transaction fee inadequate, exceeds local limit)
	// Only valid during non-consensus processing.
	// Implications:
	// - Not forwarded
	// - No fee check
	telLOCAL_ERROR TransactionResult = iota - 399
	telBAD_DOMAIN
	telBAD_PATH_COUNT
	telBAD_PUBLIC_KEY
	telFAILED_PROCESSING
	telINSUF_FEE_P
	telNO_DST_PARTIAL
)
const (
	// -299 .. -200: M Malformed (bad signature)
	// Causes:
	// - Transaction corrupt.
	// Implications:
	// - Not applied
	// - Not forwarded
	// - Reject
	// - Can not succeed in any imagined ledger.
	temMALFORMED TransactionResult = iota - 299
	temBAD_AMOUNT
	temBAD_AUTH_MASTER
	temBAD_CURRENCY
	temBAD_FEE
	temBAD_EXPIRATION
	temBAD_ISSUER
	temBAD_LIMIT
	temBAD_OFFER
	temBAD_PATH
	temBAD_PATH_LOOP
	temBAD_PUBLISH
	temBAD_TRANSFER_RATE
	temBAD_SEND_XRP_LIMIT
	temBAD_SEND_XRP_MAX
	temBAD_SEND_XRP_NO_DIRECT
	temBAD_SEND_XRP_PARTIAL
	temBAD_SEND_XRP_PATHS
	temBAD_SIGNATURE
	temBAD_SRC_ACCOUNT
	temBAD_SEQUENCE
	temDST_IS_SRC
	temDST_NEEDED
	temINVALID
	temINVALID_FLAG
	temREDUNDANT
	temREDUNDANT_SEND_MAX
	temRIPPLE_EMPTY
	temUNCERTAIN // An intermediate result used internally, should never be returned.
	temUNKNOWN
)
const (
	// -199 .. -100: F Failure (sequence number previously used)
	// Causes:
	// - Transaction cannot succeed because of ledger state.
	// - Unexpected ledger state.
	// - C++ exception.
	// Implications:
	// - Not applied
	// - Not forwarded
	// - Could succeed in an imagined ledger.
	tefFAILURE TransactionResult = iota - 199
	tefALREADY
	tefBAD_ADD_AUTH
	tefBAD_AUTH
	tefBAD_CLAIM_ID
	tefBAD_GEN_AUTH
	tefBAD_LEDGER
	tefCLAIMED
	tefCREATED
	tefDST_TAG_NEEDED
	tefEXCEPTION
	tefGEN_IN_USE
	tefINTERNAL
	tefNO_AUTH_REQUIRED // Can't set auth if auth is not required.
	tefPAST_SEQ
	tefWRONG_PRIOR
	tefMASTER_DISABLED
	tefMAX_LEDGER
)
const (
	// -99 .. -1: R Retry (sequence too high, no funds for txn fee, originating account non-existent)
	// Causes:
	// - Prior application of another, possibly non-existant, another transaction could allow this transaction to succeed.
	// Implications:
	// - Not applied
	// - Not forwarded
	// - Might succeed later
	// - Hold
	// - Makes hole in sequence which jams transactions.
	terRETRY       TransactionResult = iota - 99
	terFUNDS_SPENT                   // This is a free transaction, therefore don't burden network.
	terINSUF_FEE_B                   // Can't pay fee TransactionError     = -99 therefore don't burden network.
	terNO_ACCOUNT                    // Can't pay fee, therefore don't burden network.
	terNO_AUTH                       // Not authorized to hold IOUs.
	terNO_LINE                       // Internal flag.
	terOWNERS                        // Can't succeed with non-zero owner count.
	terPRE_SEQ                       // Can't pay fee, no point in forwarding, therefore don't burden network.
	terLAST                          // Process after all other transactions
	terNO_RIPPLE                     // Rippling not allowed
)

var resultNames = map[TransactionResult]struct {
	Token string
	Human string
}{
	tesSUCCESS:                {"tesSUCCESS", "The transaction was applied."},
	tecCLAIM:                  {"tecCLAIM", "Fee claimed. Sequence used. No action."},
	tecDIR_FULL:               {"tecDIR_FULL", "Can not add entry to full directory."},
	tecFAILED_PROCESSING:      {"tecFAILED_PROCESSING", "Failed to correctly process transaction."},
	tecINSUF_RESERVE_LINE:     {"tecINSUF_RESERVE_LINE", "Insufficient reserve to add trust line."},
	tecINSUF_RESERVE_OFFER:    {"tecINSUF_RESERVE_OFFER", "Insufficient reserve to create offer."},
	tecNO_DST:                 {"tecNO_DST", "Destination does not exist. Send ICC to create it."},
	tecNO_DST_INSUF_XRP:       {"tecNO_DST_INSUF_XRP", "Destination does not exist. Too little ICC sent to create it."},
	tecNO_LINE_INSUF_RESERVE:  {"tecNO_LINE_INSUF_RESERVE", "No such line. Too little reserve to create it."},
	tecNO_LINE_REDUNDANT:      {"tecNO_LINE_REDUNDANT", "Can't set non-existant line to default."},
	tecPATH_DRY:               {"tecPATH_DRY", "Path could not send partial amount."},
	tecPATH_PARTIAL:           {"tecPATH_PARTIAL", "Path could not send full amount."},
	tecMASTER_DISABLED:        {"tecMASTER_DISABLED", "Master key is disabled."},
	tecNO_REGULAR_KEY:         {"tecNO_REGULAR_KEY", "Regular key is not set."},
	tecUNFUNDED:               {"tecUNFUNDED", "One of _ADD, _OFFER, or _SEND. Deprecated."},
	tecUNFUNDED_ADD:           {"tecUNFUNDED_ADD", "Insufficient ICC balance for WalletAdd."},
	tecUNFUNDED_OFFER:         {"tecUNFUNDED_OFFER", "Insufficient balance to fund created offer."},
	tecUNFUNDED_PAYMENT:       {"tecUNFUNDED_PAYMENT", "Insufficient ICC balance to send."},
	tecOWNERS:                 {"tecOWNERS", "Non-zero owner count."},
	tecNO_ISSUER:              {"tecNO_ISSUER", "Issuer account does not exist."},
	tecNO_AUTH:                {"tecNO_AUTH", "Not authorized to hold asset."},
	tecNO_LINE:                {"tecNO_LINE", "No such line."},
	tecINSUFF_FEE:             {"tecINSUFF_FEE", "Insufficient balance to pay fee."},
	tecFROZEN:                 {"tecFROZEN", "Asset is frozen."},
	tecNO_TARGET:              {"tecNO_TARGET", "Target account does not exist."},
	tecNO_PERMISSION:          {"tecNO_PERMISSION", "No permission to perform requested operation."},
	tecNO_ENTRY:               {"tecNO_ENTRY", "No matching entry found."},
	tecINSUFFICIENT_RESERVE:   {"tecINSUFFICIENT_RESERVE", "Insufficient reserve to complete requested operation."},
	tecNEED_MASTER_KEY:        {"tecNEED_MASTER_KEY", "The operation requires the use of the Master Key."},
	tecDST_TAG_NEEDED:         {"tecDST_TAG_NEEDED", "A destination tag is required."},
	tecINTERNAL:               {"tecINTERNAL", "An internal error has occurred during processing."},
	tefFAILURE:                {"tefFAILURE", "Failed to apply."},
	tefALREADY:                {"tefALREADY", "The exact transaction was already in this ledger."},
	tefBAD_ADD_AUTH:           {"tefBAD_ADD_AUTH", "Not authorized to add account."},
	tefBAD_AUTH:               {"tefBAD_AUTH", "Transaction's public key is not authorized."},
	tefBAD_CLAIM_ID:           {"tefBAD_CLAIM_ID", "Malformed: Bad claim id."},
	tefBAD_GEN_AUTH:           {"tefBAD_GEN_AUTH", "Not authorized to claim generator."},
	tefBAD_LEDGER:             {"tefBAD_LEDGER", "Ledger in unexpected state."},
	tefCLAIMED:                {"tefCLAIMED", "Can not claim a previously claimed account."},
	tefCREATED:                {"tefCREATED", "Can't add an already created account."},
	tefDST_TAG_NEEDED:         {"tefDST_TAG_NEEDED", "Destination tag required."},
	tefEXCEPTION:              {"tefEXCEPTION", "Unexpected program state."},
	tefGEN_IN_USE:             {"tefGEN_IN_USE", "Generator already in use."},
	tefINTERNAL:               {"tefINTERNAL", "Internal error."},
	tefNO_AUTH_REQUIRED:       {"tefNO_AUTH_REQUIRED", "Auth is not required."},
	tefPAST_SEQ:               {"tefPAST_SEQ", "This sequence number has already past."},
	tefWRONG_PRIOR:            {"tefWRONG_PRIOR", "This previous transaction does not match."},
	tefMASTER_DISABLED:        {"tefMASTER_DISABLED", "Master key is disabled."},
	tefMAX_LEDGER:             {"tefMAX_LEDGER", "Ledger sequence too high."},
	telLOCAL_ERROR:            {"telLOCAL_ERROR", "Local failure."},
	telBAD_DOMAIN:             {"telBAD_DOMAIN", "Domain too long."},
	telBAD_PATH_COUNT:         {"telBAD_PATH_COUNT", "Malformed: Too many paths."},
	telBAD_PUBLIC_KEY:         {"telBAD_PUBLIC_KEY", "Public key too long."},
	telFAILED_PROCESSING:      {"telFAILED_PROCESSING", "Failed to correctly process transaction."},
	telINSUF_FEE_P:            {"telINSUF_FEE_P", "Fee insufficient."},
	telNO_DST_PARTIAL:         {"telNO_DST_PARTIAL", "Partial payment to create account not allowed."},
	temMALFORMED:              {"temMALFORMED", "Malformed transaction."},
	temBAD_AMOUNT:             {"temBAD_AMOUNT", "Can only send positive amounts."},
	temBAD_AUTH_MASTER:        {"temBAD_AUTH_MASTER", "Auth for unclaimed account needs correct master key."},
	temBAD_CURRENCY:           {"temBAD_CURRENCY", "Malformed: Bad currency."},
	temBAD_FEE:                {"temBAD_FEE", "Invalid fee, negative or not ICC."},
	temBAD_EXPIRATION:         {"temBAD_EXPIRATION", "Malformed: Bad expiration."},
	temBAD_ISSUER:             {"temBAD_ISSUER", "Malformed: Bad issuer."},
	temBAD_LIMIT:              {"temBAD_LIMIT", "Limits must be non-negative."},
	temBAD_OFFER:              {"temBAD_OFFER", "Malformed: Bad offer."},
	temBAD_PATH:               {"temBAD_PATH", "Malformed: Bad path."},
	temBAD_PATH_LOOP:          {"temBAD_PATH_LOOP", "Malformed: Loop in path."},
	temBAD_PUBLISH:            {"temBAD_PUBLISH", "Malformed: Bad publish."},
	temBAD_SIGNATURE:          {"temBAD_SIGNATURE", "Malformed: Bad signature."},
	temBAD_SRC_ACCOUNT:        {"temBAD_SRC_ACCOUNT", "Malformed: Bad source account."},
	temBAD_TRANSFER_RATE:      {"temBAD_TRANSFER_RATE", "Malformed: Transfer rate must be >= 1.0"},
	temBAD_SEQUENCE:           {"temBAD_SEQUENCE", "Malformed: Sequence is not in the past."},
	temBAD_SEND_XRP_LIMIT:     {"temBAD_SEND_XRP_LIMIT", "Malformed: Limit quality is not allowed for ICC to ICC."},
	temBAD_SEND_XRP_MAX:       {"temBAD_SEND_XRP_MAX", "Malformed: Send max is not allowed for ICC to ICC."},
	temBAD_SEND_XRP_NO_DIRECT: {"temBAD_SEND_XRP_NO_DIRECT", "Malformed: No Ripple direct is not allowed for ICC to ICC."},
	temBAD_SEND_XRP_PARTIAL:   {"temBAD_SEND_XRP_PARTIAL", "Malformed: Partial payment is not allowed for ICC to ICC."},
	temBAD_SEND_XRP_PATHS:     {"temBAD_SEND_XRP_PATHS", "Malformed: Paths are not allowed for ICC to ICC."},
	temDST_IS_SRC:             {"temDST_IS_SRC", "Destination may not be source."},
	temDST_NEEDED:             {"temDST_NEEDED", "Destination not specified."},
	temINVALID:                {"temINVALID", "The transaction is ill-formed."},
	temINVALID_FLAG:           {"temINVALID_FLAG", "The transaction has an invalid flag."},
	temREDUNDANT:              {"temREDUNDANT", "Sends same currency to self."},
	temREDUNDANT_SEND_MAX:     {"temREDUNDANT_SEND_MAX", "Send max is redundant."},
	temRIPPLE_EMPTY:           {"temRIPPLE_EMPTY", "PathSet with no paths."},
	temUNCERTAIN:              {"temUNCERTAIN", "In process of determining result. Never returned."},
	temUNKNOWN:                {"temUNKNOWN", "The transactions requires logic not implemented yet."},
	terRETRY:                  {"terRETRY", "Retry transaction."},
	terFUNDS_SPENT:            {"terFUNDS_SPENT", "Can't set password, password set funds already spent."},
	terINSUF_FEE_B:            {"terINSUF_FEE_B", "Account balance can't pay fee."},
	terLAST:                   {"terLAST", "Process last."},
	terNO_RIPPLE:              {"terNO_RIPPLE", "Path does not permit rippling."},
	terNO_ACCOUNT:             {"terNO_ACCOUNT", "The source account does not exist."},
	terNO_AUTH:                {"terNO_AUTH", "Not authorized to hold IOUs."},
	terNO_LINE:                {"terNO_LINE", "No such line."},
	terPRE_SEQ:                {"terPRE_SEQ", "Missing/inapplicable prior transaction."},
	terOWNERS:                 {"terOWNERS", "Non-zero owner count."},
}

var reverseResults map[string]TransactionResult

func init() {
	reverseResults = make(map[string]TransactionResult)
	for result, name := range resultNames {
		reverseResults[name.Token] = result
	}
}

func (r TransactionResult) String() string {
	return resultNames[r].Token
}

func (r TransactionResult) Human() string {
	return resultNames[r].Human
}

func (r TransactionResult) Success() bool {
	return r == tesSUCCESS || r == tecCLAIM
}

func (r TransactionResult) Symbol() string {
	switch r {
	case tesSUCCESS, tecCLAIM:
		return "✓"
	case tecPATH_PARTIAL, tecPATH_DRY:
		return "½"
	case tecUNFUNDED, tecUNFUNDED_ADD, tecUNFUNDED_OFFER, tecUNFUNDED_PAYMENT:
		return "$"
	default:
		return "✗"
	}
}
