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

func ErrInvalidParam(id string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "Invalid " + id + " parameter",
	}
}

func SuccessCreated() ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Created successfully",
	}
}

func SuccessUpdated() ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Updated successfully",
	}
}

func SuccessDeleted() ServerResponse {
	return ServerResponse{
		Status:  "success",
		Message: "Deleted successfully",
	}
}

func InvalidData(err error) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: "Invalid data, " + err.Error(),
	}
}

func CustomError(message string) ServerResponse {
	return ServerResponse{
		Status:  "error",
		Message: message,
	}
}
