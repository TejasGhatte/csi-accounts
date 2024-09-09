package schemas
type Did struct {
	Scheme           string `json:"scheme" validate:"required"`
	Method           string `json:"method" validate:"required"`
	MethodSpecificID string `json:"methodSpecificId" validate:"required"`
}

