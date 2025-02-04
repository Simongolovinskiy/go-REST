package model

type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"` // 'deposit', 'transfer_in', 'transfer_out'
	CreatedAt string  `json:"created_at"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
