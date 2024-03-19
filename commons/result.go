package commons

type Result struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func ResultSuccess(data interface{}) Result {
	return Result{data, "success", 200}
}

func ResultByCodeAndMessage(code int, message string) Result {
	return Result{"", message, code}
}

func ResultByAll(code int, message string, data interface{}) Result {
	return Result{data, message, code}
}
