package response

type DataResponse struct {
	Meta Meta `json:"meta"`
	Value interface{} `json:"data"`
}