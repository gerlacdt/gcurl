package http

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c *Client) Get(params Params) (response Result, err error) {
	return request(params)
}

func (c *Client) Delete(params Params) (response Result, err error) {
	return request(params)
}

func (c *Client) Post(params ParamsWithBody) (result Result, err error) {
	return requestWithBody(params)
}

func (c *Client) Put(params ParamsWithBody) (result Result, err error) {
	return requestWithBody(params)
}
