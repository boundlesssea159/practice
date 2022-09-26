package practice1

import (
	"math"
	"testing"
)

func TestRefrigerates_LimitSku(t *testing.T) {
	type arg struct {
		residualVolumes       float32
		skuAmounts            []SkuAmount
		expectCoreSkuRatio    float32
		expectNonCoreSkuRatio float32
	}
	args := []arg{
		{
			residualVolumes: 70,
			skuAmounts: []SkuAmount{
				{sku: Sku{id: 1, name: "核心品商品1", volume: 2, isCore: true}, amount: 8},
				{sku: Sku{id: 2, name: "核心品商品2", volume: 10, isCore: true}, amount: 4},
				{sku: Sku{id: 3, name: "非核心品商品3", volume: 5, isCore: false}, amount: 4},
				{sku: Sku{id: 4, name: "非核心品商品4", volume: 8, isCore: false}, amount: 4},
			},
			expectCoreSkuRatio:    1,
			expectNonCoreSkuRatio: 0.2672,
		},
	}

	for index, arg := range args {
		refrigerates := NewRefrigerates()
		refrigerates.IncrementResidualVolumes(arg.residualVolumes)
		for _, sku := range arg.skuAmounts {
			refrigerates.AddSku(sku)
		}
		refrigerates.SkuLimit()
		for _, sku := range arg.skuAmounts {
			newSku := refrigerates.GetSkuAmountById(sku.GetSkuId())
			if newSku.GetSkuId() == 0 {
				t.Fatalf("index %+v test fail.sku is nil", index)
			}
			if sku.IsCore() {
				if newSku.GetAmount() != int(math.Ceil(float64(float32(sku.GetAmount())*arg.expectCoreSkuRatio))) {
					t.Fatalf("index %+v test fail.core sku amount comupte wrong.skuId %+v expect %+v actual %+v",
						index, sku.GetSkuId(), int(math.Ceil(float64(float32(sku.GetAmount())*arg.expectCoreSkuRatio))), newSku.GetAmount())
				}
				continue
			}
			if newSku.GetAmount() != int(math.Ceil(float64(float32(sku.GetAmount())*arg.expectNonCoreSkuRatio))) {
				t.Fatalf("index %+v test fail.non core sku amount comupte wrong.skuId %+v expect %+v actual %+v",
					index, sku.GetSkuId(), int(math.Ceil(float64(float32(sku.GetAmount())*arg.expectNonCoreSkuRatio))), newSku.GetAmount())
			}
		}
	}
}
