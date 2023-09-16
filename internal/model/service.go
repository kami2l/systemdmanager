package model

type Service struct {
	Unit string `json:"unit"`

	//  Reflects whether the unit definition was properly loaded.
	Load string `json:"load"`

	// The high-level unit activation state, i.e. generalization of SUB.
	Active string `json:"active"`

	// The low-level unit activation state, values depend on unit type.
	Sub         string `json:"sub"`
	Description string `json:"description"`
}
