package category

type CategoryInput struct {
	Name string `json:"name" form:"name" validate:"required"`
}
