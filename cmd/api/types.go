package main

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}
