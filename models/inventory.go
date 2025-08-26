package models

type Inventory struct {
    ID       string  `json:"id" bson:"_id,omitempty"`
    ItemName string  `json:"item_name"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}