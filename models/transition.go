package models

type Transition struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	HasScreen     bool   `json:"hasScreen"`
	IsGlobal      bool   `json:"isGlobal"`
	IsInitial     bool   `json:"isInitial"`
	IsAvailable   bool   `json:"isAvailable"`
	IsConditional bool   `json:"isConditional"`
	IsLooped      bool   `json:"isLooped"`
	To            struct {
		Self           string `json:"self"`
		Id             string `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		IconUrl        string `json:"iconUrl"`
		StatusCategory struct {
			Self      string `json:"self"`
			Id        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"to"`
}

type ListTransitionsResponse struct {
	Expand      string       `json:"expand"`
	Transitions []Transition `json:"transitions"`
}
