package models

type Cart struct {
    Items []Product  `json:"items"`
    Total float32 `json:"total"`
}
