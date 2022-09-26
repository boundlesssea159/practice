package practice1

// 商品
type Sku struct {
	// id
	id int
	// 名称
	name string
	// 体积
	volume float32
	// 是否核心
	isCore bool
}

func (this *Sku) IsCore() bool {
	return this.isCore
}

func (this *Sku) GetId() int {
	return this.id
}

func (this *Sku) GetVolume() float32 {
	return this.volume
}
