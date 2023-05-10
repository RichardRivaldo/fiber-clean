package models

type Query struct {
	Sort       interface{} `json:"sort"`
	Filter     interface{} `json:"filter"`
	Projection interface{} `json:"projection"`
}
