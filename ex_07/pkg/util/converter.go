package util

import "github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"

var (
	weightConversions = map[constant.WeightUnit]map[constant.WeightUnit]float64{
		constant.WeightUnitKg: {
			constant.WeightUnitGram: 1000,
			constant.WeightUnitKg:   1,
			constant.WeightUnitLBS:  2.20462,
		},
		constant.WeightUnitGram: {
			constant.WeightUnitGram: 1,
			constant.WeightUnitKg:   0.001,
			constant.WeightUnitLBS:  0.0022,
		},
		constant.WeightUnitLBS: {
			constant.WeightUnitGram: 453.59237,
			constant.WeightUnitKg:   0.45359,
			constant.WeightUnitLBS:  1,
		},
	}
)

func ConvertWeight(weight float64, from, to constant.WeightUnit) float64 {
	return weightConversions[from][to] * weight
}
