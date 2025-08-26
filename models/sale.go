package models

type Sale struct {
    ID           string  `json:"id" bson:"_id,omitempty"`
    ItemID       string  `json:"item_id"`
    QuantitySold int     `json:"quantity_sold"`
    TotalAmount  float64 `json:"total_amount"`
}