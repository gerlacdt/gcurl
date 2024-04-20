package http

func Delete(params Params) (response Result, err error) {
	return request(params)
}
