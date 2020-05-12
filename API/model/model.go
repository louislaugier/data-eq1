package model

// Response type stores a response's meta and data
type Response struct {
	StatusCode int16  `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`
	} `json:"meta"`
	Data []interface{} `json:"data"`
}
