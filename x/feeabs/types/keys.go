package types

const (
	// Module name store the name of the module
	ModuleName = "feeabs"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the feeabs module
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_feeabs"

	// Contract: Coin denoms cannot contain this character
	KeySeparator = "|"
)

var (
	OsmosisTwapExchangeRate          = []byte{0x01} // Key for the exchange rate of osmosis (to native token)
	KeyChannelID                     = []byte{0x02} // Key for IBC channel to osmosis
	KeyHostChainChainConfigByFeeAbs  = []byte{0x03} // Key for IBC channel to osmosis
	KeyHostChainChainConfigByOsmosis = []byte{0x04} // Key for IBC channel to osmosis
	KeyPrefixEpoch                   = []byte{0x05} // KeyPrefixEpoch defines prefix key for storing epochs.
	KeyTokenDenomPair                = []byte{0x06} // Key store token denom pair on feeabs and osmosis
)

func GetKeyHostZoneConfigByFeeabsIBCDenom(feeabsIbcDenom string) []byte {
	return append(KeyHostChainChainConfigByFeeAbs, []byte(feeabsIbcDenom)...)
}

func GetKeyHostZoneConfigByOsmosisIBCDenom(osmosisIbcDenom string) []byte {
	return append(KeyHostChainChainConfigByOsmosis, []byte(osmosisIbcDenom)...)
}

func GetKeyTwapExchangeRate(ibcDenom string) []byte {
	return append(OsmosisTwapExchangeRate, []byte(ibcDenom)...)
}
