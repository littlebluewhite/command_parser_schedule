package command_server

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *commandServer) requestProtocol(ctx context.Context, com command) (result doResult) {
	for {
		select {
		case <-ctx.Done():
			result.status = Failure
			result.message = "command not match monitor timeout"
			return
		default:
			switch com.Template.Protocol {
			case https.String():
				result = c.doHttp(ctx, com)
			case websocket.String():
			case mqtt.String():
			case redisTopic.String():
			default:
			}
			if com.Template.Monitor == nil {
				result.status = Success
				return
			} else {
				monitorData(result, *com.Template.Monitor)
			}
		}
	}
}

func (c *commandServer) doHttp(ctx context.Context, com command) (result doResult) {
	// TODO: add variable function
	var body io.Reader
	h := com.Template.Http
	var contentType string
	if h.Body != nil {
		switch *h.BodyType {
		case "json":
			body = bytes.NewBuffer(*h.Body)
			contentType = "application/json"
		case "form_data":
			//TODO form data body
			contentType = "multipart/form-data"
		case "x_www_form_urlencoded":
			//TODO x_www_form_urlencoded body
			contentType = "application/x-www-form-urlencoded"
		default:
		}
	}
	header := make([]httpHeader, 0, 20)
	req, e := http.NewRequestWithContext(ctx, h.Method, h.URL, body)
	if e != nil {
		result.status = Failure
		result.message = "http request timeout"
		return
	}
	if h.Header != nil {
		if e := json.Unmarshal(h.Header, &header); e != nil {
			c.l.Error().Printf("id: %d header unmarshal failed", com.CommandId)
		}
	}
	for _, item := range header {
		if item.IsActive {
			req.Header.Set(item.Key, item.Value)
		}
	}
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	var resp *http.Response
	if resp1, e := client.Do(req); e != nil {
		result.respData = []byte{}
		c.l.Error().Printf("id: %d request failed", com.CommandId)
	} else {
		resp = resp1
	}
	result.statusCode = resp.StatusCode
	if respBody1, e := io.ReadAll(resp.Body); e != nil {
		result.respData = []byte{}
		c.l.Error().Printf("id: %d request body failed", com.CommandId)
		return
	} else {
		result.respData = respBody1
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			c.l.Error().Println("Response body closed failed")
		}
	}()
	c.l.Info().Printf("id: %d request status: %v\nrequest result: %s, co\n", com.CommandId)
	return
}

func monitorData(result doResult, m monitor) doResult {
	if result.statusCode != int(m.StatusCode) {
		result.message = "status code error"
		return result
	}
	for _, condition := range m.MConditions {
		v := stringAnalyze(result.respData, condition.SearchRule)
	}
}

func stringAnalyze(data []byte, rule string) (result analyzeResult) {
	r := strings.Split(rule, ".")
	// "root.person.[all]array.name
	var f any
	e := json.Unmarshal(data, &f)
	if e != nil {
		return
	}
	for _, word := range r[1:] {
		if strings.Index(word, "array") == -1 {
			m, ok := f.(map[string]interface{})
			if !ok {
				return
			}
			f, ok = m[word]
			if !ok {
				return
			}
		} else {

		}
	}
	return
}
