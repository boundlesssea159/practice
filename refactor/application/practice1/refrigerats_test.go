package practice1

import (
	"testing"
)

func TestRefrigerates_AddRefrigerator(t *testing.T) {
	type arg struct {
		totalVolume         float32
		residueVolume       float32
		refrigerators       []Refrigerator
		expectTotalVolume   float32
		expectResidueVolume float32
	}

	AddArgs := []arg{
		{
			totalVolume:   0,
			residueVolume: 0,
			refrigerators: []Refrigerator{
				{totalVolume: 100, residueVolume: 100},
				{totalVolume: 80, residueVolume: 70},
				{totalVolume: 10.5, residueVolume: 0.5},
			},
			expectTotalVolume:   190.5,
			expectResidueVolume: 170.5,
		},
	}

	for _, arg := range AddArgs {
		refrigerates := Refrigerates{totalVolumes: arg.totalVolume, residualVolumes: arg.residueVolume}
		for _, refrigerator := range arg.refrigerators {
			refrigerates.AddRefrigerator(refrigerator)
		}
		if refrigerates.GetToTalVolume() != arg.expectTotalVolume {
			t.Fatal("totalVolume wrong")
		}
		if refrigerates.GetResidueVolume() != arg.expectResidueVolume {
			t.Fatal("totalVolume wrong")
		}
	}

}

func TestRefrigerates_ReduceRefrigerator(t *testing.T) {
	type arg struct {
		totalVolume         float32
		residueVolume       float32
		refrigerators       []Refrigerator
		expectTotalVolume   float32
		expectResidueVolume float32
	}

	ReduceArgs := []arg{
		{
			totalVolume:   9.5,
			residueVolume: 29.5,
			refrigerators: []Refrigerator{
				{totalVolume: 9.5, residueVolume: 29.5},
			},
			expectTotalVolume:   0,
			expectResidueVolume: 0,
		},

		{
			totalVolume:   0,
			residueVolume: 0,
			refrigerators: []Refrigerator{
				{totalVolume: 9.5, residueVolume: 29.5},
			},
			expectTotalVolume:   -9.5,
			expectResidueVolume: -29.5,
		},
	}

	for _, arg := range ReduceArgs {
		refrigerates := Refrigerates{totalVolumes: arg.totalVolume, residualVolumes: arg.residueVolume}
		for _, refrigerator := range arg.refrigerators {
			refrigerates.ReduceRefrigerator(refrigerator)
		}
		if refrigerates.GetToTalVolume() != arg.expectTotalVolume {
			t.Fatal("totalVolume wrong")
		}
		if refrigerates.GetResidueVolume() != arg.expectResidueVolume {
			t.Fatal("totalVolume wrong")
		}
	}
}
