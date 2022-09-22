package dto

type ByIdRequest struct {
	ID uint `param:"id" validate:"required"`
}

type ByUnixIDRequest struct {
	ID int64 `param:"id" validate:"required"`
}
