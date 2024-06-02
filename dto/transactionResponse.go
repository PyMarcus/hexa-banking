package dto

type TransactionResponse struct {
	TransactionId    string  `json:"transaction_id"`
	AccountId        string  `json:"account_id"`
	TransactionValue float64 `json:"transaction_value"`
	TransactionType  string  `json:"transaction_type"`
	TransactionDate  string  `json:"transaction_date"`
}