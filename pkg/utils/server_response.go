package utils

type ServerResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrParseJson() ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "Cannot parse JSON",
	}
}

func ErrParseID() ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "Cannot parse ID",
	}
}

func DuplicateData(field string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: field + " already exists",
	}
}

func ServerError(err error) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: err.Error(),
	}
}

func NotFound(field string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "No " + field + " found",
	}
}

func IDNotFound(field string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: field + " not found",
	}
}

func ErrParseParam(id string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "Invalid " + id + " param",
	}
}

func SuccessCreated(data interface{}) ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Created successfully",
		Data:    data,
	}
}

func SuccessUpdated(data interface{}) ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Updated successfully",
		Data:    data,
	}
}

func SuccessDeleted(data interface{}) ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Deleted successfully",
		Data:    data,
	}
}
