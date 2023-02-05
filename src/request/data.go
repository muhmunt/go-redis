package request

type (
	DataRequest struct {
		Key string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}
)
