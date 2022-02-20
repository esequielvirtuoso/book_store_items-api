package items

// TODO: refactor to separate into sellers, items, inventory, and customers

// Item defines the items characteristics
type Item struct {
	ID                string      `json:"item"`
	SellerID          int64       `json:"seller_id"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	UnitPrice         float32     `json:"unit_price"`
	AvailableQuantity int64       `json:"available_quantity"`
	SoldQuantity      int64       `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description defines the items descriptions structure
type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

// Pictures defines the items pictures information
type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}
