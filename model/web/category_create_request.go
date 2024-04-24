package web

type CategoryCreateRequest struct {
	Category string `validate:"required,min=1,max=100" json:"name"`
}
