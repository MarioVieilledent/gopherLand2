package input

type KeyPressed struct {
	Nickname string `json:"nickname"` // Nickname of player doing movements
	Up       bool   `json:"up"`
	Right    bool   `json:"right"`
	Down     bool   `json:"down"`
	Left     bool   `json:"left"`
}
