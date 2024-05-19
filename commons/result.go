package commons

type Result struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

const (
	SUCCESS_CODE        int    = 200
	DEFAULT_ERROR_CODE  int    = 500
	DEFAULT_SUCCESS_MSG string = "success"
	DEFAULT_ERROR_MSG   string = "error"
	DEFAULT_EMPTY_DATA  string = ""
)

func ResultSuccess(data interface{}) Result {
	return Result{data, DEFAULT_SUCCESS_MSG, SUCCESS_CODE}
}

func ResultSuccessByMessage(message string, data interface{}) Result {
	return Result{data, message, SUCCESS_CODE}
}

func ResultByCodeAndMessage(code int, message string) Result {
	return Result{DEFAULT_EMPTY_DATA, message, code}
}

func ResultByAll(code int, message string, data interface{}) Result {
	return Result{data, message, code}
}

func ResultError(message string) Result {
	return Result{DEFAULT_EMPTY_DATA, message, DEFAULT_ERROR_CODE}
}
