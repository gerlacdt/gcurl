package http

func Get(params Params) (response Result, err error) {
	return request(params)
}
