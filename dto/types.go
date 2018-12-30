package dto

type LogRequest struct {
	URL      string `json:"url"`
	Referrer string `json:"referrer"`
}

type LogResponse struct{}
