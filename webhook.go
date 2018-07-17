package main

import "time"

type WebhookBody struct {
	Type string
	Data WebhookData
}

type WebhookData struct {
	AccountID   string `json:"account_id"`
	Amount      int    // Pennies
	Created     time.Time
	Currency    string
	Description string
	ID          string
	Category    string
	IsLoad      bool `json:"is_load"`
	Settled     time.Time
	Merchant    Merchant
}

type Merchant struct {
	Address  Address
	Created  time.Time
	GroupID  string `json:"group_id"`
	ID       string
	Logo     string
	Emoji    string
	Name     string
	Category string
}

type Address struct {
	Address   string
	City      string
	Country   string
	Latitude  float64
	Longitude float64
	Postcode  string
	Region    string
}
