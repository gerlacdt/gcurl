package http

func Post(params ParamsWithBody) (result Result, err error) {
	return requestWithBody(params)
}
