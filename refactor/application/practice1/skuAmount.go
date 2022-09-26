package practice1

type SkuAmount struct {
	sku    Sku
	amount int
}

func (this *SkuAmount) GetSkuId() int {
	return this.sku.Id
}

func (this *SkuAmount) GetSkuVolume() float32 {
	return this.sku.Volume
}

func (this *SkuAmount) IsCore() bool {
	return this.sku.IsCore
}

func (this *SkuAmount) AddAmount(amount int) {
	this.amount += amount
}
