package fhttp

import (
	"fmt"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/pkg/logger"
	"probabilisticTimeSeriesModeling/pkg/utils"

	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type Client struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewClient(cfg *config.Config, logger logger.Logger) *Client {
	return &Client{
		cfg:    cfg,
		logger: logger,
	}
}

var timeout = time.Second * 5

func (h *Client) Request(
	method string,
	url string,
	body []byte,
	queryParams map[string]string,
	headers map[string]string,
) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}
	if err = client.DoTimeout(req, res, timeout); err != nil {
		h.logger.Warnf("Error( %s ) url: %s %s ", err.Error(), url, utils.GetStructJSON(queryParams))
		err = errors.New("timeout")
		return
	}
	req.SetConnectionClose()
	statusCode = res.StatusCode()
	//h.logger.Debugf("%s |%s%s%s| %s\nstatusCode: %d", method, url, string(body), utils.GetStructJSON(queryParams), utils.GetStructJSON(headers), statusCode)

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	return
}
