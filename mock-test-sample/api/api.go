package api

type Fetcher interface {
	FetchDataFromHoge() string
}

type Client struct {
	HC Fetcher
}

type HogeClient struct {
}

func (c *Client) Fetch() string {
	data := c.HC.FetchDataFromHoge()
	res := "data: " + data
	return res
}

func (hc *HogeClient) FetchDataFromHoge() string {
	// ほんとはここで外部API通信等
	return "hogehoge"
}
