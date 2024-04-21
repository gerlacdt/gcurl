package http

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c *Client) Get(params Params) (response Result, err error) {
	return doRequest(params)
}

func (c *Client) Delete(params Params) (response Result, err error) {
	return doRequest(params)
}

func (c *Client) Post(params Params) (result Result, err error) {
	return doRequest(params)
}

func (c *Client) Put(params Params) (result Result, err error) {
	return doRequest(params)
}
