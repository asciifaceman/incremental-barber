package data

import (
	"math/big"

	"fyne.io/fyne/v2"
	"github.com/asciifaceman/incremental-barber/pkg/bundles"
)

type Tool struct {
	Name            string
	Owned           bool
	Icon            fyne.Resource
	Cost            *big.Float
	QualityModifier float32
	ComfortModifier float32
}

func NewToolPackage() *ToolPackage {
	return &ToolPackage{
		Scissors: &Tool{
			Name:            "Scissors",
			Owned:           true,
			Icon:            bundles.ToolsScissorsPng,
			QualityModifier: 1.0,
			ComfortModifier: 1.0,
		},
		Cape: &Tool{
			Name:            "Cape",
			Owned:           false,
			Icon:            nil,
			Cost:            big.NewFloat(20),
			QualityModifier: 1.0,
			ComfortModifier: 1.1,
		},
		Comb: &Tool{
			Name:            "Comb",
			Owned:           false,
			Icon:            nil,
			Cost:            big.NewFloat(5),
			QualityModifier: 1.1,
			ComfortModifier: 1.0,
		},
	}
}

// ToolPackage represents the set of tools a barber might have
type ToolPackage struct {
	Scissors *Tool
	Cape     *Tool
	Comb     *Tool
}
