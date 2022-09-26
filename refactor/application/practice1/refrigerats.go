package practice1

import (
	"math"
)

// 冷冻柜聚合（对外提供统一抽象）
type Refrigerates struct {
	// 核心品的约束比例
	coreRatio float32
	// 非核心品的约束比例
	nonCoreRatio float32
	// 总库容
	totalVolumes float32
	// 剩余可用库容
	residualVolumes float32
	// 剩余商品
	residualSkus map[int]SkuAmount
}

func NewRefrigerates() *Refrigerates {
	return &Refrigerates{
		residualSkus: make(map[int]SkuAmount),
	}
}

func (this *Refrigerates) IncrementTotalVolumes(totalVolumes float32) {
	this.totalVolumes += totalVolumes
}

func (this *Refrigerates) DecrementTotalVolumes(totalVolumes float32) {
	this.totalVolumes -= totalVolumes
}

func (this *Refrigerates) IncrementResidualVolumes(residualVolumes float32) {
	this.residualVolumes += residualVolumes
}

func (this *Refrigerates) DecrementResidualVolumes(residualVolumes float32) {
	this.residualVolumes -= residualVolumes
}

func (this *Refrigerates) GetToTalVolume() float32 {
	return this.totalVolumes
}

func (this *Refrigerates) GetResidueVolume() float32 {
	return this.residualVolumes
}

// 添加商品
func (this *Refrigerates) AddSku(sku SkuAmount) {
	if sku.GetSkuId() == 0 {
		return
	}
	if residualSku, exist := this.residualSkus[sku.GetSkuId()]; !exist {
		this.residualSkus[sku.GetSkuId()] = sku
	} else {
		residualSku.AddAmount(sku.amount)
		this.residualSkus[sku.GetSkuId()] = residualSku
	}
	return
}

// 商品熔断
func (this *Refrigerates) SkuLimit() {
	// 计算熔断系数
	this.computeSkuRatio()
	// 遍历商品进行熔断
	this.reComputeSkuAmount()
}

func (this *Refrigerates) computeSkuRatio() {
	this.computeCoreSkuRatio()
	this.computeNonCoreSkuRatio()
}

func (this *Refrigerates) computeCoreSkuRatio() {
	var volumes float32
	for _, skuAmount := range this.residualSkus {
		if skuAmount.IsCore() {
			volumes += skuAmount.ComputeVolume()
		}
	}
	if this.residualVolumes >= volumes {
		this.coreRatio = 1.0
		return
	}
	this.coreRatio = float32(math.Min(float64(this.residualVolumes/volumes), float64(1)))
}

func (this *Refrigerates) computeNonCoreSkuRatio() {
	var coreSkuVolumes, nonCoreSkuVolumes float32
	for _, skuAmount := range this.residualSkus {
		if skuAmount.IsCore() {
			coreSkuVolumes += skuAmount.ComputeVolume()
			continue
		}
		nonCoreSkuVolumes += skuAmount.ComputeVolume()
	}
	residualVolumes := this.residualVolumes
	residualVolumes -= coreSkuVolumes
	if residualVolumes >= nonCoreSkuVolumes {
		this.nonCoreRatio = 1.0
		return
	}
	this.nonCoreRatio = float32(math.Min(float64(residualVolumes/nonCoreSkuVolumes), float64(1)))
}

func (this *Refrigerates) GetSkuAmountById(id int) SkuAmount {
	if v, exist := this.residualSkus[id]; exist {
		return v
	}
	return SkuAmount{}
}

func (this *Refrigerates) reComputeSkuAmount() {
	for id, _ := range this.residualSkus {
		sku, _ := this.residualSkus[id]
		if sku.IsCore() {
			sku.UpdateAmount(int(math.Ceil(float64(float32(sku.GetAmount()) * this.coreRatio))))
		} else {
			sku.UpdateAmount(int(math.Ceil(float64(float32(sku.GetAmount()) * this.nonCoreRatio))))
		}
		this.residualSkus[id] = sku
	}
}
