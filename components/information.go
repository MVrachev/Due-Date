package components

// Information is used to send data through the websockets
type Information struct {
	Priority    string `json:"priority"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	Day         string `json:"day"`
}
