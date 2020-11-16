package common

type success struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data interface{}) *success {
	return &success{Data: data}
}
