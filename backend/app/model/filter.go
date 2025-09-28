package model

type TableFilter struct {
	Search string
	Page   int
	Limit  int
	Status string
	Role   string
}
