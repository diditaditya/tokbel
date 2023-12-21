package category

type CategoryRequest struct {
	Type string `json:"type" validate:"required"`
}
