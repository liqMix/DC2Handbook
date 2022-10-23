package types

type Photo struct {
	ID string
	Name string
	ImageURL string
	Chapter uint8
	Missable bool
	MissableNote string
	Location string
	IsScoop bool
	Memo string
}