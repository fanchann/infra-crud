package types

type (
	BookWithId struct {
		Id   int `json:"id"`
		Book Book
	}

	Book struct {
		Title string `json:"title"`
	}

	ApiResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
