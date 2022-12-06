package httpclient

import "github.com/parnurzeal/gorequest"

const (
	getCommand  = "GET"
	postCommand = "POST"
)

var agent *gorequest.SuperAgent

func init() {
	agent = gorequest.New()
}

func Send(action, url string, data string) (string, error) {
	var (
		res string
		err error
	)

	switch action {
	case getCommand:
		res, err = sendGet(url)
	case postCommand:
		res, err = sendPost(url, data)
	default:
		return "", ErrAction
	}

	if err != nil {
		return "", err
	}
	return res, nil
}

func sendGet(url string) (string, error) {
	_, body, errs := agent.Get(url).End()
	if len(errs) > 0 {
		return "", ErrRequest
	}
	return body, nil
}

func sendPost(url string, data string) (string, error) {
	_, body, errs := agent.Post(url).Send(data).End()
	if len(errs) > 0 {
		return "", ErrRequest
	}
	return body, nil
}
