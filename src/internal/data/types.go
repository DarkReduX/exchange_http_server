package data

type SymbolPrice struct {
	Uuid   int64   `json:"uuid"`
	Symbol string  `json:"symbol"`
	Bid    float32 `json:"bid"`
	Ask    float32 `json:"ask"`
}

type Position struct {
	UUID   int64   `json:"uuid"`
	Symbol string  `json:"symbol"`
	Open   float32 `json:"open"`
	Close  float32 `json:"close"`
	IsBay  bool    `json:"isbay"`
}

func (p Position) PNL(lastPrice SymbolPrice) float32 {
	if p.IsBay {
		return p.Open - lastPrice.Bid
	}
	return lastPrice.Ask - p.Open
}
