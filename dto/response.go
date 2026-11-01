package dto

type ResponseSuccess[T any] struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    T      `json:"data,omitempty"`
}

type ResponseError struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Error   interface{} `json:"error,omitempty"`
}

func CreateResponseError(code int, message string) ResponseError {
    return ResponseError{
        Code:    code,
        Message: message,
        Error:   nil,
    }
}

func CreateResponseErrorData(code int, message string, errDetails interface{}) ResponseError {
    return ResponseError{
        Code:    code,
        Message: message,
        Error:   errDetails,
    }
}

func CreateResponseSuccess(message string) ResponseSuccess[any] {
    return ResponseSuccess[any]{
        Code:    200,
        Message: message,
        Data:    nil,
    }
}

func CreateResponseSuccessWithData[T any](message string, data T) ResponseSuccess[T] {
    return ResponseSuccess[T]{
        Code:    200,
        Message: message,
        Data:    data,
    }
}