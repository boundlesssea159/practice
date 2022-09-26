package practice1

type SkuAmount struct {
	sku    Sku
	amount int
}

func (this *SkuAmount) GetSkuId() int {
	return this.sku.GetId()
}

func (this *SkuAmount) GetSkuVolume() float32 {
	return this.sku.GetVolume()
}

func (this *SkuAmount) IsCore() bool {
	return this.sku.IsCore()
}

func (this *SkuAmount) AddAmount(amount int) {
	this.amount += amount
}

func (this *SkuAmount) GetAmount() int {
	return this.amount
}

func (this *SkuAmount) UpdateAmount(amount int) {
	this.amount = amount
}

func (this *SkuAmount) ComputeVolume() float32 {
	return float32(this.amount) * this.sku.GetVolume()
}
