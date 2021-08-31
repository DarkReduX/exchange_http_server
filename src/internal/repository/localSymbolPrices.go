package repository

import (
	"main/src/internal/data"
)

var Repository = symbolPriceStore{
	Data:                 map[string]data.SymbolPrice{},
	IsUpdated:            false,
	OpenedPos:            map[string]float32{},
	SortedPositionsByPnl: map[float64]data.Position{},
	SortedPositionsKey:   []float64{},
}

type symbolPriceStore struct {
	GeneralPnl           float32
	Balance              float32
	UserToken            string
	Data                 map[string]data.SymbolPrice
	IsUpdated            bool
	OpenedPos            map[string]float32
	Positions            map[string]data.Position
	SortedPositionsByPnl map[float64]data.Position
	SortedPositionsKey   []float64
}
