package decimal

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

/************/
/* Multiply */
/************/

func BenchmarkMultiplyBigFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MultiplyBigFloat()
	}
}

func BenchmarkMultiplyBigFloatWithoutNew(b *testing.B) {

	bigFloatStartFloat := big.NewFloat(StartFloat)
	bigFloatFactor := big.NewFloat(Factor)

	for n := 0; n < b.N; n++ {
		MultiplyBigFloatWithoutNew(bigFloatStartFloat, bigFloatFactor)
	}
}

func BenchmarkMultiplyDecimal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MultiplyDecimal()
	}
}

func BenchmarkMultiplyDecimalWithoutNew(b *testing.B) {

	decimalStartFloat := decimal.NewFromFloat(StartFloat)
	decimalFactor := decimal.NewFromFloat(Factor)

	for n := 0; n < b.N; n++ {
		MultiplyDecimalWithoutNew(decimalStartFloat, decimalFactor)
	}
}

func BenchmarkMultiplyFloat32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MultiplyFloat64()
	}
}

func BenchmarkMultiplyFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddFloat32()
	}
}

/*******/
/* Add */
/*******/

func BenchmarkAddBigFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddBigFloat()
	}
}

func BenchmarkAddBigFloatWithoutNew(b *testing.B) {

	bigFloatStartFloat := big.NewFloat(StartFloat)
	bigFloatFactor := big.NewFloat(Factor)

	for n := 0; n < b.N; n++ {
		AddBigFloatWithoutNew(bigFloatStartFloat, bigFloatFactor)
	}
}

func BenchmarkAddDecimal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddDecimal()
	}
}

func BenchmarkAddDecimalWithoutNew(b *testing.B) {

	decimalStartFloat := decimal.NewFromFloat(StartFloat)
	decimalFactor := decimal.NewFromFloat(Factor)

	for n := 0; n < b.N; n++ {
		AddDecimalWithoutNew(decimalStartFloat, decimalFactor)
	}
}

func BenchmarkAddFloat32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddFloat64()
	}
}

func BenchmarkAddFloat64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddFloat32()
	}
}
