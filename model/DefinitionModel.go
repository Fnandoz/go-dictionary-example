package model

type DefinitionModel struct {
	WordType   string `json:"type"`
	Definition string
	Example    string
	ImageUrl  string `json:"image_url"`
	Emoji      string
}
