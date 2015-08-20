package decimal

import (
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	StartFloat = 10000.55
	Factor     = 0.333
)

func MultiplyBigFloat() (calculated *big.Float) {
	bigFloatStartFloat := big.NewFloat(StartFloat)
	bigFloatFactor := big.NewFloat(Factor)

	calculated = bigFloatStartFloat.Mul(bigFloatStartFloat, bigFloatFactor)

	return
}

func MultiplyBigFloatWithoutNew(bigFloatStartFloat, bigFloatFactor *big.Float) (calculated *big.Float) {
	calculated = bigFloatStartFloat.Mul(bigFloatStartFloat, bigFloatFactor)

	return
}

func MultiplyDecimal() (calculated decimal.Decimal) {
	decimalStartFloat := decimal.NewFromFloat(StartFloat)
	decimalFactor := decimal.NewFromFloat(Factor)

	calculated = decimalStartFloat.Mul(decimalFactor)

	return
}

func MultiplyDecimalWithoutNew(decimalStartFloat, decimalFactor decimal.Decimal) (calculated decimal.Decimal) {
	calculated = decimalStartFloat.Mul(decimalFactor)

	return
}

func MultiplyFloat64() (calculated float64) {
	calculated = StartFloat * Factor

	return
}

func MultiplyFloat32() (calculated float32) {
	calculated = StartFloat * Factor

	return
}

func AddBigFloat() (calculated *big.Float) {
	bigFloatStartFloat := big.NewFloat(StartFloat)
	bigFloatFactor := big.NewFloat(Factor)

	calculated = bigFloatStartFloat.Add(bigFloatStartFloat, bigFloatFactor)

	return
}

func AddBigFloatWithoutNew(bigFloatStartFloat, bigFloatFactor *big.Float) (calculated *big.Float) {
	calculated = bigFloatStartFloat.Add(bigFloatStartFloat, bigFloatFactor)

	return
}

func AddDecimal() (calculated decimal.Decimal) {
	decimalStartFloat := decimal.NewFromFloat(StartFloat)
	decimalFactor := decimal.NewFromFloat(Factor)

	calculated = decimalStartFloat.Add(decimalFactor)

	return
}

func AddDecimalWithoutNew(decimalStartFloat, decimalFactor decimal.Decimal) (calculated decimal.Decimal) {
	calculated = decimalStartFloat.Add(decimalFactor)

	return
}

func AddFloat64() (calculated float64) {
	calculated = StartFloat + Factor

	return
}

func AddFloat32() (calculated float32) {
	calculated = StartFloat + Factor

	return
}
