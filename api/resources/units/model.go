package units

type Unit struct {
	ID           int    `json:"id"`
	SingularName string `json:"singular_name"`
	CreatedAt    string `json:"created_at"`
}

type UnitDTO struct {
	SingularName string `json:"singular_name"`
}
