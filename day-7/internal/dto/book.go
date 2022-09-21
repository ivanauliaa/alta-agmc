package dto

type (
	CreateBookRequest struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   int    `json:"year" validate:"required"`
	}

	UpdateBookRequest struct {
		Title  *string `json:"title"`
		Author *string `json:"author"`
		Year   *int    `json:"year"`
	}
)
