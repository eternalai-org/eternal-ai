package models

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/shopspring/decimal"
)

const (
	ETH_ZERO_ADDRESS    = "0x0000000000000000000000000000000000000000"
	BURN_ADDRESS        = "0x000000000000000000000000000000000000dEaD"
	GENERTAL_NETWORK_ID = 0
	ETHEREUM_NETWORK_ID = 1

	DURATION_1D  = 24 * 3600
	DURATION_30D = 30 * 24 * 3600
	DURATION_90D = 90 * 24 * 3600

	LOCAL_CHAIN_ID            = uint64(1337)
	ETERNAL_AI_CHAIN_ID       = uint64(43338)
	BTC_CHAIN_ID              = uint64(0)
	ETHEREUM_CHAIN_ID         = uint64(1)
	FANS_CHAIN_ID             = uint64(45761)
	BASE_CHAIN_ID             = uint64(8453)
	SHARDAI_CHAIN_ID          = uint64(222671)
	SOLANA_CHAIN_ID           = uint64(1111)
	SOLANA_CHAIN_ID_OLD       = uint64(101)
	HERMES_CHAIN_ID           = uint64(45762)
	ARBITRUM_CHAIN_ID         = uint64(42161)
	ZKSYNC_CHAIN_ID           = uint64(324)
	POLYGON_CHAIN_ID          = uint64(137)
	BSC_CHAIN_ID              = uint64(56)
	SEPOLIA_CHAIN_ID          = uint64(11155111)
	DAGI_CHAIN_ID             = uint64(222672)
	APE_CHAIN_ID              = uint64(33139)
	AVALANCHE_C_CHAIN_ID      = uint64(43114)
	ABSTRACT_TESTNET_CHAIN_ID = uint64(11124)
	BITTENSOR_CHAIN_ID        = uint64(964)
	DUCK_CHAIN_ID             = uint64(5545)
	TRON_CHAIN_ID             = uint64(728126428)
	MODE_CHAIN_ID             = uint64(34443)
	ZETA_CHAIN_ID             = uint64(7000)
	STORY_CHAIN_ID            = uint64(1514)
	HYPE_CHAIN_ID             = uint64(999)
	MONAD_TESTNET_CHAIN_ID    = uint64(10143)
	MEGAETH_TESTNET_CHAIN_ID  = uint64(1338)
	BASE_SEPOLIA_CHAIN_ID     = uint64(84532)
	CELO_CHAIN_ID             = uint64(42220)

	LOWER_PRICE_USD = 2800
	UPPER_PRICE_USD = 1000000000
	TOKEN_SUPPLY    = 1000000000
)

var CHAIN_NAME_MAP = map[uint64]string{
	BASE_CHAIN_ID:             "BASE",
	SHARDAI_CHAIN_ID:          "BITCOIN",
	FANS_CHAIN_ID:             "FANS",
	ETHEREUM_CHAIN_ID:         "ETHEREUM",
	HERMES_CHAIN_ID:           "SYMBIOSIS",
	ARBITRUM_CHAIN_ID:         "ARBITRUM",
	SOLANA_CHAIN_ID:           "SOLANA",
	ZKSYNC_CHAIN_ID:           "ZKSYNC",
	POLYGON_CHAIN_ID:          "POLYGON",
	BSC_CHAIN_ID:              "BSC",
	SEPOLIA_CHAIN_ID:          "SEPOLIA",
	APE_CHAIN_ID:              "APECHAIN",
	AVALANCHE_C_CHAIN_ID:      "AVALANCHE C-CHAIN",
	ABSTRACT_TESTNET_CHAIN_ID: "ABSTRACT TESTNET",
	BITTENSOR_CHAIN_ID:        "BITTENSOR",
	DUCK_CHAIN_ID:             "DUCK CHAIN",
	TRON_CHAIN_ID:             "TRON",
	MODE_CHAIN_ID:             "MODE",
	LOCAL_CHAIN_ID:            "LOCAL",
	ZETA_CHAIN_ID:             "ZETA",
	STORY_CHAIN_ID:            "STORY",
	HYPE_CHAIN_ID:             "HYPER EVM",
	MONAD_TESTNET_CHAIN_ID:    "MONAD TESTNET",
	MEGAETH_TESTNET_CHAIN_ID:  "MEGAETH TESTNET",
	BASE_SEPOLIA_CHAIN_ID:     "BASE SEPOLIA",
	CELO_CHAIN_ID:             "CELO",
}

var MAP_TOOLSET_NAME = map[string]string{
	"":                             "Default",
	"default":                      "Default",
	"reply_mentions":               "Reply",
	"reply_mentions_farcaster":     "Reply",
	"shadow_reply":                 "Reply",
	"reply_non_mentions_farcaster": "Engage",
	"follow":                       "Follow",
	"post":                         "Post",
	"post_farcaster":               "Post",
	"lucky_moneys":                 "Airdrop",
	"trade_analytics_twitter":      "Trading Analytics",
	"trade_analytics_mentions":     "Trading Analytics",
}

var MAP_CHAIN_ID_TO_LLM_MODEL = map[uint64]map[string]string{
	DAGI_CHAIN_ID: { // AGI chain
		"neuralmagic/Meta-Llama-3.1-405B-Instruct-quantized.w4a16": "200001",
	},
	BASE_CHAIN_ID: { // Base chain
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8":                  "700002",
		"neuralmagic/Meta-Llama-3.1-405B-Instruct-quantized.w4a16": "700001",
	},
	HERMES_CHAIN_ID: { // uncensored chain
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "500001",
		"unsloth/Llama-3.3-70B-Instruct-bnb-4bit": "500006",
		"PrimeIntellect/INTELLECT-1-Instruct":     "500007",
	},
	SHARDAI_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "600001",
	},
	ARBITRUM_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "110001",
	},
	ZKSYNC_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "120001",
	},
	POLYGON_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "130001",
	},
	SOLANA_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "990001",
	},
	APE_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "990001",
	},
	AVALANCHE_C_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "990001",
	},
	ABSTRACT_TESTNET_CHAIN_ID: {
		"NousResearch/Hermes-3-Llama-3.1-70B-FP8": "990001",
	},
}

func ONE_ETHER() *big.Int {
	return big.NewInt(1e18)
}

func ParseBool(v string) bool {
	ok, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return ok
}

func ConvertFloatToReserveAmount(amt float64) uint64 {
	if amt < 0 {
		panic("invalid amount")
	}
	return decimal.NewFromFloat(amt).Shift(2).Truncate(0).BigInt().Uint64()
}

func ConvertPriceAmount(amt float64) float64 {
	if amt < 0 {
		panic("invalid amount")
	}
	amt, _ = decimal.NewFromFloat(amt).Round(8).Float64()
	return amt
}

func ValidateFiatAmount(amt float64) error {
	if amt < 0 {
		return errors.New("invalid amount")
	}
	if amt != ConvertFiatAmount(amt) {
		return errors.New("fiat amount is invalid")
	}
	return nil
}

func ConvertFiatAmount(amt float64) float64 {
	if amt < 0 {
		panic("invalid amount")
	}
	amt, _ = decimal.NewFromFloat(amt).Round(2).Float64()
	return amt
}

func ValidateNftCurrencyAmount(amt float64, decimals uint) error {
	if amt < 0 {
		return errors.New("amount is small than 0")
	}
	newAmt, _ := decimal.NewFromFloat(amt).Round(int32(decimals)).Float64()
	if newAmt != amt {
		return errors.New("amount is invalid decimals")
	}
	return nil
}

func ConvertNumberFloat(amt float64, decimals uint) float64 {
	if amt < 0 {
		panic(errors.New("amount is small than 0"))
	}
	newAmt, _ := decimal.NewFromFloat(amt).Round(int32(decimals)).Float64()
	return newAmt
}

func ConvertWeiToBigFloat(amt *big.Int, decimals uint) *big.Float {
	if amt == nil {
		return big.NewFloat(0.0)
	}
	// if amt.Cmp(big.NewInt(0)) < 0 {
	// 	panic(errors.New("amount is small than 0"))
	// }
	amtFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	decimalFloat := new(big.Float).SetPrec(1024).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	retFloat := new(big.Float).Quo(amtFloat, decimalFloat)
	return retFloat
}

func ConvertWeiStringToBigFloat(amtString string, decimals uint) *big.Float {
	if amtString == "s" {
		return big.NewFloat(0.0)
	}
	amt, ok := new(big.Int).SetString(amtString, 10)
	if !ok {
		panic("not Ok")
	}
	if amt.Cmp(big.NewInt(0)) < 0 {
		panic(errors.New("amount is small than 0"))
	}
	amtFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	decimalFloat := new(big.Float).SetPrec(1024).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	retFloat := new(big.Float).Quo(amtFloat, decimalFloat)
	return retFloat
}

func ConvertStringToBigFloat(amtString string) *big.Float {
	if amtString == "s" {
		return big.NewFloat(0.0)
	}
	amt, ok := new(big.Float).SetPrec(1024).SetString(amtString)
	if !ok {
		panic("not Ok")
	}
	return amt
}

func ConvertWeiToBigFloatNegative(amt *big.Int, decimals uint) *big.Float {
	// if amt.Cmp(big.NewInt(0)) < 0 {
	// 	panic(errors.New("amount is small than 0"))
	// }
	if amt == nil {
		return big.NewFloat(0.0)
	}
	amtFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	decimalFloat := new(big.Float).SetPrec(1024).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	retFloat := new(big.Float).Quo(amtFloat, decimalFloat)
	return retFloat
}

func ConvertBigFloatToWei(amt *big.Float, decimals uint) *big.Int {
	if amt.Cmp(big.NewFloat(0)) < 0 {
		panic(errors.New("amount is small than 0"))
	}
	newAmt, err := decimal.NewFromString(amt.Text('f', 64))
	if err != nil {
		panic(err)
	}
	newAmt = newAmt.Shift(int32(decimals)).Truncate(0)
	return newAmt.BigInt()
}

func ConvertBigFloatToWeiFull(amt *big.Float, decimals uint) *big.Int {
	newAmt, err := decimal.NewFromString(amt.Text('f', 64))
	if err != nil {
		panic(err)
	}
	newAmt = newAmt.Shift(int32(decimals)).Truncate(0)
	return newAmt.BigInt()
}

func ConvertCryptoCurrencyAmount(amt float64) float64 {
	if amt < 0 {
		panic("invalid amount")
	}
	amt, _ = decimal.NewFromFloat(amt).Round(8).Float64()
	return amt
}

func ConvertReserveAmountToFloat(reserveAmt uint64) float64 {
	amt, _ := decimal.New(int64(reserveAmt), -2).Float64()
	return amt
}

func ConvertFloatToCollateralAmount(amt float64) uint64 {
	if amt < 0 {
		panic("amount invalid")
	}
	return decimal.NewFromFloat(amt).Shift(8).Truncate(0).BigInt().Uint64()
}

func ParseString2FloatAmountArr(s, sep string) []float64 {
	rets := []float64{}
	s = strings.TrimSpace(s)
	if s != "" {
		ss := strings.Split(s, sep)
		if len(ss) > 0 {
			for _, n := range ss {
				dm, err := decimal.NewFromString(n)
				if err != nil {
					panic(err)
				}
				val, _ := dm.Truncate(2).Float64()
				rets = append(rets, val)
			}
		}
	}
	return rets
}

func MulFloats(val1 float64, vals ...float64) float64 {
	val := decimal.NewFromFloat(val1)
	for _, v := range vals {
		val = val.Mul(decimal.NewFromFloat(v))
	}
	num, _ := val.Float64()
	return num
}

func DivFloats(val1 float64, vals ...float64) float64 {
	val := decimal.NewFromFloat(val1)
	for _, v := range vals {
		val = val.Div(decimal.NewFromFloat(v))
	}
	num, _ := val.Float64()
	return num
}

func AddFloats(val1 float64, vals ...float64) float64 {
	val := decimal.NewFromFloat(val1)
	for _, v := range vals {
		val = val.Add(decimal.NewFromFloat(v))
	}
	num, _ := val.Float64()
	return num
}

func SubFloats(val1 float64, vals ...float64) float64 {
	val := decimal.NewFromFloat(val1)
	for _, v := range vals {
		val = val.Sub(decimal.NewFromFloat(v))
	}
	num, _ := val.Float64()
	return num
}

func ConvertStringToFloat(s string) (float64, error) {
	num, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}
	amount, _ := num.Float64()
	return amount, nil
}

func ConvertString2BigInt(s string) (*big.Int, error) {
	n, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("%s is not number", s)
	}
	return n, nil
}

func ConvertStringNumber2BigInt(s string) *big.Int {
	n, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		panic(fmt.Errorf("%s is not number", s))
	}
	return n
}

func ToEtherPriceFiatAmount(num big.Float, pr float64) float64 {
	rs, _ := big.NewFloat(0).Mul(&num, big.NewFloat(pr)).Float64()
	return ConvertFiatAmount(rs)
}

func ToEtherAmount(e *big.Int) big.Float {
	if e == nil {
		return big.Float{}
	}
	return *decimal.NewFromBigInt(e, -18).BigFloat()
}

func ToEtherWeiAmount(num big.Float) big.Int {
	dn, err := decimal.NewFromString(big.NewFloat(0).Mul(&num, big.NewFloat(1e18)).String())
	if err != nil {
		panic(err)
	}
	return *dn.BigInt()
}

func ToBigInt(s string) big.Int {
	if s == "" {
		return big.Int{}
	}
	n, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		panic(errors.New("numer is invalid"))
	}
	return *n
}

func Number2BigInt(s string, decimals int) *big.Int {
	dn, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	dn = dn.Shift(int32(decimals)).Round(0)
	return dn.BigInt()
}

func MulBigFloats(val1 *big.Float, vals ...*big.Float) *big.Float {
	val := val1
	for _, v := range vals {
		val = new(big.Float).Mul(val, v)
	}
	return val
}

func MulBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		val = new(big.Int).Mul(val, v)
	}
	return val
}

func MinBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		if val.Cmp(v) > 0 {
			val = v
		}
	}
	return val
}

func MaxBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		if val.Cmp(v) < 0 {
			val = v
		}
	}
	return val
}

func MaxBigFloats(val1 *big.Float, vals ...*big.Float) *big.Float {
	val := val1
	for _, v := range vals {
		if val.Cmp(v) < 0 {
			val = v
		}
	}
	return val
}

func AddBigFloats(val1 *big.Float, vals ...*big.Float) *big.Float {
	val := val1
	for _, v := range vals {
		val = new(big.Float).Add(val, v)
	}
	return val
}

func AddBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		val = new(big.Int).Add(val, v)
	}
	return val
}

func SubBigFloats(val1 *big.Float, vals ...*big.Float) *big.Float {
	val := val1
	for _, v := range vals {
		val = new(big.Float).Sub(val, v)
	}
	return val
}

func SubBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		val = new(big.Int).Sub(val, v)
	}
	return val
}

func QuoBigFloats(val1 *big.Float, vals ...*big.Float) *big.Float {
	val := val1
	for _, v := range vals {
		if v.Cmp(big.NewFloat(0)) == 0 {
			panic(errors.New("divide zero"))
		}
		val = new(big.Float).Quo(val, v)
	}
	return val
}

func QuoBigInts(val1 *big.Int, vals ...*big.Int) *big.Int {
	val := val1
	for _, v := range vals {
		if v.Cmp(big.NewInt(0)) == 0 {
			panic(errors.New("divide zero"))
		}
		val = new(big.Int).Quo(val, v)
	}
	return val
}

func EqualBigFloats(val1 *big.Float, val2 *big.Float) bool {
	return val1.Text('f', 64) == val2.Text('f', 64)
}

func NegativeBigFloat(val *big.Float) *big.Float {
	return SubBigFloats(big.NewFloat(0), val)
}

func FormatFloatNumber(f string, amt float64) string {
	return FormatStringNumber(fmt.Sprintf(f, amt))
}

func FormatStringNumber(amt string) string {
	if strings.Contains(amt, ".") {
		amt = strings.TrimRight(amt, "0")
		amt = strings.TrimRight(amt, ".")
	}
	return amt
}

func FormatBigFloatNumber(amt *big.Float) string {
	return FormatStringNumber(amt.Text('f', 64))
}

func FormatEmailTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func ConvertSqrtPriceX96ToPrice(amt *big.Int, decimals uint) *big.Float {
	// if amt.Cmp(big.NewInt(0)) < 0 {
	// 	panic(errors.New("amount is small than 0"))
	// }
	pow := new(big.Int).Exp(amt, big.NewInt(2), nil)
	x96 := new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil)

	powFloat := new(big.Float).SetPrec(1024).SetInt(pow)
	x96Float := new(big.Float).SetPrec(1024).SetInt(x96)
	retFloat := new(big.Float).Quo(powFloat, x96Float)

	return retFloat

}

func ConvertX96ToNumber(amt *big.Int, decimals uint) *big.Float {
	// if amt.Cmp(big.NewInt(0)) < 0 {
	// 	panic(errors.New("amount is small than 0"))
	// }
	// pow := new(big.Int).Exp(amt, big.NewInt(2), nil)
	x96 := new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)

	powFloat := new(big.Float).SetPrec(1024).SetInt(amt)
	x96Float := new(big.Float).SetPrec(1024).SetInt(x96)
	retFloat := new(big.Float).Quo(powFloat, x96Float)

	return retFloat

}

func EtherToWeiWithDecimals(eth *big.Float, decimals int) *big.Int {
	return ConvertBigFloatToWei(eth, uint(decimals))
}

func EtherToWei(eth *big.Float) *big.Int {
	return ConvertBigFloatToWei(eth, 18)
}

func RoundHalfUp(num *big.Float, decimals int) *big.Float {
	scale := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	scaled := new(big.Float).Mul(num, scale)

	roundedInt, _ := scaled.Int(nil) // Extract the *big.Int value
	roundedFloat := new(big.Float).SetInt(roundedInt)

	quotient := new(big.Float).Quo(roundedFloat, scale)

	return quotient
}

func AbsBigFloat(val *big.Float) *big.Float {
	return new(big.Float).Abs(val)
}

func ConvertSqrtPriceX96ToPriceEx(amt *big.Int, decimals uint, zeroForOne bool) *big.Float {
	pow := new(big.Int).Exp(amt, big.NewInt(2), nil)
	x96 := new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil)
	powFloat := new(big.Float).SetPrec(1024).SetInt(pow)
	x96Float := new(big.Float).SetPrec(1024).SetInt(x96)
	retFloat := new(big.Float).Quo(powFloat, x96Float)
	if zeroForOne {
		retFloat = QuoBigFloats(
			big.NewFloat(1),
			retFloat,
		)
	}
	return retFloat
}

func GetAlphaDBName() string {
	dbName := "dev_nbc_perp"
	if configs.GetConfig().Env == "mainnet" {
		dbName = "prod_nbc_perp"
	}
	return dbName
}

func IsNativeToken(tokenAddress string) bool {
	return strings.EqualFold("0x0000000000000000000000000000000000000000", tokenAddress)
}

func PriceToTick(price float64, spacing int64) int64 {
	const base = 1.0001
	tick := int64(math.Log(price) / math.Log(base))
	return (tick / spacing) * spacing
}

func RandFloatInRage(min, max float64) float64 {
	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*(max-min)
}

func RandSeed() uint64 {
	min := int(1)
	max := int(10e6)
	return uint64(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min+1) + min)
}

func GetSqrtPriceX96ForLiquidityAndAmount0(sqrtPriceX96AMin *big.Int, sqrtPriceX96BMax *big.Int, liquidity *big.Int, amountMax *big.Int) *big.Int {
	if sqrtPriceX96AMin.Cmp(sqrtPriceX96BMax) >= 0 {
		panic("wrong sqrtPriceX96A sqrtPriceX96b")
	}
	amountCheckMax := helpers.GetAmount0ForLiquidity(sqrtPriceX96AMin, sqrtPriceX96BMax, liquidity)
	if amountCheckMax.Cmp(amountMax) < 0 {
		return sqrtPriceX96AMin
	}
	sqrtPriceX96A := sqrtPriceX96AMin
	sqrtPriceX96B := sqrtPriceX96BMax
	var sqrtPriceX96 *big.Int
	for i := 0; i < 2560; i++ {
		sqrtPriceX96Tmp := QuoBigInts(AddBigInts(sqrtPriceX96A, sqrtPriceX96B), big.NewInt(2))
		amount := helpers.GetAmount0ForLiquidity(sqrtPriceX96Tmp, sqrtPriceX96BMax, liquidity)
		if amount.Cmp(amountMax) <= 0 {
			if SubBigInts(amountMax, amount).Cmp(big.NewInt(10000)) <= 0 {
				sqrtPriceX96 = sqrtPriceX96Tmp
				break
			}
			sqrtPriceX96B = sqrtPriceX96Tmp
		} else {
			sqrtPriceX96A = sqrtPriceX96Tmp
		}
	}
	if sqrtPriceX96 == nil {
		panic("sqrtPriceX96 not found")
	}
	return sqrtPriceX96
}

func GetSqrtPriceX96ForLiquidityAndAmount1(sqrtPriceX96AMin *big.Int, sqrtPriceX96BMax *big.Int, liquidity *big.Int, amountMax *big.Int) *big.Int {
	if sqrtPriceX96AMin.Cmp(sqrtPriceX96BMax) >= 0 {
		panic("wrong sqrtPriceX96A sqrtPriceX96b")
	}
	amountCheckMax := helpers.GetAmount1ForLiquidity(sqrtPriceX96AMin, sqrtPriceX96BMax, liquidity)
	if amountCheckMax.Cmp(amountMax) < 0 {
		return sqrtPriceX96BMax
	}
	sqrtPriceX96A := sqrtPriceX96AMin
	sqrtPriceX96B := sqrtPriceX96BMax
	var sqrtPriceX96 *big.Int
	for i := 0; i < 2560; i++ {
		sqrtPriceX96Tmp := QuoBigInts(AddBigInts(sqrtPriceX96A, sqrtPriceX96B), big.NewInt(2))
		amount := helpers.GetAmount1ForLiquidity(sqrtPriceX96AMin, sqrtPriceX96Tmp, liquidity)
		if amount.Cmp(amountMax) <= 0 {
			if SubBigInts(amountMax, amount).Cmp(big.NewInt(10000)) <= 0 {
				sqrtPriceX96 = sqrtPriceX96Tmp
				break
			}
			sqrtPriceX96A = sqrtPriceX96Tmp
		} else {
			sqrtPriceX96B = sqrtPriceX96Tmp
		}
	}
	if sqrtPriceX96 == nil {
		panic("sqrtPriceX96 not found")
	}
	return sqrtPriceX96
}

func GetChainName(chainID uint64) string {
	return CHAIN_NAME_MAP[chainID]
}

func GetChainID(chainName string) uint64 {
	switch strings.ToLower(chainName) {
	case "base":
		{
			return BASE_CHAIN_ID
		}
	case "ethereum":
		{
			return ETHEREUM_NETWORK_ID
		}
	case "solana":
		{
			return SOLANA_CHAIN_ID
		}
	case "arbitrum":
		{
			return ARBITRUM_CHAIN_ID
		}
	case "bsc", "bnbchain", "binancechain", "binance", "bnb":
		{
			return BSC_CHAIN_ID
		}
	case "polygon":
		{
			return POLYGON_CHAIN_ID
		}
	case "avax", "avalanche":
		{
			return AVALANCHE_C_CHAIN_ID
		}
	case "apechain", "ape_chain":
		{
			return APE_CHAIN_ID
		}
	case "zksync_era", "zksync":
		{
			return ZKSYNC_CHAIN_ID
		}
	case "abstract testnet", "abstract_testnet":
		{
			return ABSTRACT_TESTNET_CHAIN_ID
		}
	case "bittensor":
		{
			return BITTENSOR_CHAIN_ID
		}
	case "bitcoin":
		{
			return HERMES_CHAIN_ID
		}
	case "duck_chain", "duckchain":
		{
			return DUCK_CHAIN_ID
		}
	default:
		return GENERTAL_NETWORK_ID
	}
}

func GetTradeUrl(tokenNetworkID uint64, tokenAddress, dexID string) string {
	if tokenNetworkID == BASE_CHAIN_ID {
		return fmt.Sprintf(`https://app.uniswap.org/explore/tokens/base/%s`, tokenAddress)
	} else if tokenNetworkID == SOLANA_CHAIN_ID {
		if dexID == "raydium" {
			return fmt.Sprintf("https://raydium.io/swap/?inputMint=sol&outputMint=%s", tokenAddress)
		} else {
			return fmt.Sprintf("https://pump.fun/coin/%s", tokenAddress)
		}
	} else if tokenNetworkID == ARBITRUM_CHAIN_ID {
		return fmt.Sprintf(`https://app.camelot.exchange/?token1=0xDB8C67e6CA293F43C75e106c70b97033cC2909E3&token2=%s`, tokenAddress)
	} else if tokenNetworkID == BSC_CHAIN_ID {
		return fmt.Sprintf(`https://pancakeswap.finance/?outputCurrency=%s&inputCurrency=0x4B6bF1d365ea1A8d916Da37FaFd4ae8C86d061D7`, tokenAddress)
	}
	return ""
}

func GetDexUrl(tokenNetworkID uint64, tokenAddress string) string {
	if tokenAddress != "" {
		if tokenNetworkID == BASE_CHAIN_ID {
			return fmt.Sprintf(`https://dexscreener.com/base/%s`, tokenAddress)
		} else if tokenNetworkID == ARBITRUM_CHAIN_ID {
			return fmt.Sprintf(`https://dexscreener.com/arbitrum/%s`, tokenAddress)
		} else if tokenNetworkID == SOLANA_CHAIN_ID {
			return fmt.Sprintf(`https://dexscreener.com/solana/%s`, tokenAddress)
		} else if tokenNetworkID == BSC_CHAIN_ID {
			return fmt.Sprintf(`https://dexscreener.com/bsc/%s`, tokenAddress)
		} else if tokenNetworkID == APE_CHAIN_ID {
			return fmt.Sprintf(`https://dexscreener.com/apechain/%s`, tokenAddress)
		}
	}
	return ""
}

func GetImageUrl(imageUrl string) string {
	if strings.HasPrefix(imageUrl, "ipfs://") {
		imageUrl = strings.Replace(imageUrl, "ipfs://", "https://gateway.lighthouse.storage/ipfs/", -1)
	}
	return imageUrl
}

var (
	ETERNALAI_LIQUIDITY_SUPPORTED_NETWORKS = []map[string]any{
		{
			"network": "ETH",
			"name":    "Ethereum",
			"icon":    "https://eternalai.org/icons/coins/ic-eth.svg",
			"id":      ETHEREUM_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "EtherScan",
					"url":    "https://etherscan.io/",
					"icon":   "https://eternalai.org/icons/coins/ic-ether-scan.png",
					"okLink": "https://www.okx.com/web3/explorer/eth/",
				},
			},
		},
		{
			"network": "BASE",
			"name":    "Base",
			"icon":    "https://eternalai.org/icons/blockchains/ic_base.svg",
			"id":      BASE_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "BaseScan",
					"url":    "https://basescan.org/",
					"icon":   "https://eternalai.org/icons/coins/ic-base-scan.png",
					"okLink": "https://www.okx.com/web3/explorer/base/",
				},
			},
		},
		{
			"network": "ARBITRUM",
			"name":    "Arbitrum",
			"icon":    "https://eternalai.org/icons/blockchains/ic_arbitrum.svg",
			"id":      ARBITRUM_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "ArbitrumScan",
					"url":    "https://arbiscan.io/",
					"icon":   "https://eternalai.org/icons/blockchains/ic_arbitrum.svg",
					"okLink": "https://www.okx.com/web3/explorer/arbitrum/",
				},
			},
		},
		{
			"network": "SOL",
			"name":    "Solana",
			"icon":    "https://eternalai.org/icons/blockchains/ic_solana.svg",
			"id":      SOLANA_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "SolScan",
					"url":    "https://solscan.io/",
					"icon":   "https://eternalai.org/icons/coins/ic-sol-scan.svg",
					"okLink": "https://www.okx.com/web3/explorer/sol/",
				},
			},
		},
		{
			"network": "BSC",
			"name":    "BNB Smart Chain",
			"icon":    "https://eternalai.org/icons/coins/ic-bsc.png",
			"id":      BSC_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "BscScan",
					"url":    "https://bscscan.com/",
					"icon":   "https://eternalai.org/icons/coins/ic-bsc-scan.svg",
					"okLink": "https://www.okx.com/web3/explorer/bsc/",
				},
			},
		},
		{
			"network": "APE",
			"name":    "ApeChain",
			"icon":    "https://eternalai.org/icons/coins/ic-ape.svg",
			"id":      APE_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "ApeScan",
					"url":    "https://apescan.io/",
					"icon":   "https://eternalai.org/icons/coins/ic-ape.svg",
					"okLink": "https://www.okx.com/web3/explorer/apechain/",
				},
			},
		},
		{
			"network": "AVAX",
			"name":    "Avalanche C-Chain",
			"icon":    "https://eternalai.org/icons/blockchains/ic_avax.svg",
			"id":      AVALANCHE_C_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "AVAXScan",
					"url":    "https://subnets.avax.network/c-chain",
					"icon":   "https://eternalai.org/icons/blockchains/ic_avax.svg",
					"okLink": "https://www.okx.com/web3/explorer/avax/",
				},
			},
		},
		{
			"network": "Celo",
			"name":    "Celo",
			"icon":    "https://eternalai.org/icons/blockchains/ic-celo.svg",
			"id":      CELO_CHAIN_ID,
			"explorer": []map[string]any{
				{
					"name":   "CeloScan",
					"url":    "https://celoscan.io/",
					"icon":   "https://eternalai.org/icons/blockchains/ic-celo.svg",
					"okLink": "",
				},
			},
		},
	}
)
