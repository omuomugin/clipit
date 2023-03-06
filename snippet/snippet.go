package snippet

type Snippet struct {
	Cmd       string     `json:"cmd"`
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Key     string   `json:"key"`
	Choices []string `json:"choices"`
}
