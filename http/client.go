package http

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c *Client) Get(params paramsInternal) (response Result, err error) {
	return doRequest(params)
}

func (c *Client) Delete(params paramsInternal) (response Result, err error) {
	return doRequest(params)
}

func (c *Client) Post(params paramsInternal) (result Result, err error) {
	return doRequest(params)
}

func (c *Client) Put(params paramsInternal) (result Result, err error) {
	return doRequest(params)
}
