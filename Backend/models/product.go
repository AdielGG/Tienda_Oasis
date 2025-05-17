package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Quantity    float32 `json:"quantity"`
	Description string  `json:"description"`
	Image       []byte  `json:"image"`
}
