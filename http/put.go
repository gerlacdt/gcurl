package http

func Put(params ParamsWithBody) (result Result, err error) {
	return requestWithBody(params)
}
