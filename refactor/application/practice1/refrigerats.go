package practice1

// 冷冻柜聚合（对外提供统一抽象）
type Refrigerates struct {
	// 冷冻库容柜集合
	refrigerators map[int]Refrigerator
	// 核心品的约束比例
	coreRatio float32
	// 非核心品的约束比例
	nonCoreRatio float32
	// 总库容
	totalVolumes float32
	// 剩余可用库容
	residualVolumes float32
}

func (this *Refrigerates) AddRefrigerator(refrigerator Refrigerator) {
	if this.refrigerators == nil {
		this.refrigerators = make(map[int]Refrigerator)
	}
	this.refrigerators[refrigerator.GetId()] = refrigerator
	this.totalVolumes += refrigerator.totalVolume
	this.residualVolumes += refrigerator.residueVolume
}

func (this *Refrigerates) GetToTalVolume() float32 {
	return this.totalVolumes
}

func (this *Refrigerates) GetResidueVolume() float32 {
	return this.residualVolumes
}

func (this *Refrigerates) ReduceRefrigerator(refrigerator Refrigerator) {
	delete(this.refrigerators, refrigerator.GetId())
	this.totalVolumes -= refrigerator.totalVolume
	this.residualVolumes -= refrigerator.residueVolume
}
