package model

type Transaksi struct {
	Id       string `json:"id"`
	Status string `json:"status"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	Maded string `json:"maded"`
	Payed string `json:"payed"`
}
