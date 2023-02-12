package common

// PriceLevel is a common structure for bids and asks in the
// order book.
type PriceLevel struct {
	Price    Float64
	Quantity Float64
}

// Parse parses this PriceLevel's Price and Quantity and
// returns them both.  It also returns an error if either
// fails to parse.
func (p *PriceLevel) Parse() (float64, float64, error) {
	// price, err := strconv.ParseFloat(p.Price, 64)
	// if err != nil {
	// 	return 0, 0, err
	// }
	// quantity, err := strconv.ParseFloat(p.Quantity, 64)
	// if err != nil {
	// 	return price, 0, err
	// }
	// return price, quantity, nil
	return float64(p.Price), float64(p.Quantity), nil
}
