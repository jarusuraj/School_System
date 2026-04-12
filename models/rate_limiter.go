package models
type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}
type RateLimitConfig struct {
	Limit  float64
	Status string
	Body   string
}