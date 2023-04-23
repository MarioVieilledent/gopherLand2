package input

type KeyPressed struct {
	Up    bool `json:"up"`
	Right bool `json:"right"`
	Down  bool `json:"down"`
	Left  bool `json:"left"`
}
