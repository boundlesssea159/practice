package practice1

import (
	"math"
)

// 冷冻柜
type Refrigerator struct {
	// id
	id int
	// 剩余库容
	residueVolume float32
	// 库容上限
	totalVolume float32
	// 核心品的约束比例
	coreRatio float32
	// 非核心品的约束比例
	nonCoreRatio float32
	// 剩余商品
	residualSkus map[int]SkuAmount
}

func (this *Refrigerator) Limit(amount float32, isCore bool) float32 {
	if isCore {
		return this.coreRatio * amount
	}
	return this.nonCoreRatio * amount
}

func (this *Refrigerator) computeCoreRatio(planVolumes float32) {
	if this.residueVolume <= 0 {
		this.coreRatio = 0
		return
	}
	this.coreRatio = this.computeRatio(planVolumes)
	return
}

func (this *Refrigerator) computeNonCoreRatio(planVolumes float32) {
	if this.residueVolume <= 0 {
		this.nonCoreRatio = 0
		return
	}
	this.nonCoreRatio = this.computeRatio(planVolumes)
	return
}

func (this *Refrigerator) computeRatio(planVolumes float32) float32 {
	if this.residueVolume >= planVolumes {
		return 1.0
	}
	return float32(math.Max(float64(this.residueVolume/planVolumes), 1.0))
}

//// 计算剩余库容体积
//func (this *Refrigerator) computeResidualVolume() {
//	var skuVolumes float32
//	for id, sku := range this.residualSkus {
//		amount, _ := this.residualAmount[id]
//		skuVolumes += float32(amount) * sku.Volume
//	}
//	this.residueVolume = this.totalVolume - skuVolumes
//}

// 添加商品
func (this *Refrigerator) AddSku(sku SkuAmount) {
	if sku.GetSkuId() == 0 {
		return
	}
	if residualSku, exist := this.residualSkus[sku.GetSkuId()]; !exist {
		this.residualSkus[sku.GetSkuId()] = sku
	} else {
		residualSku.AddAmount(sku.amount)
	}
	planVolumes := float32(sku.amount) * sku.GetSkuVolume()
	// 优先保证核心品的熔断
	this.computeCoreRatio(planVolumes)
	this.residueVolume -= planVolumes
	// 再保证非核心品的熔断
	this.computeNonCoreRatio(planVolumes)
}

func (this *Refrigerator) GetId() int {
	return this.id
}
