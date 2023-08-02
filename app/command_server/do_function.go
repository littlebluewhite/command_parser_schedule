package command_server

import (
	"bytes"
	"command_parser_schedule/util"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
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
	asserts := make([]assertResult, 0, len(m.MConditions))
	for _, condition := range m.MConditions {
		ar := stringAnalyze(result.respData, condition.SearchRule)
		assert := assertValue(ar, condition)
	}
}

func stringAnalyze(data []byte, rule string) (result analyzeResult) {
	r := strings.Split(rule, ".")
	// "root.person.[all]array.name
	var f []any
	var arrayFlag bool
	var d any
	e := json.Unmarshal(data, &d)
	if e != nil {
		return
	}
	f = append(f, d)
	for _, word := range r[1:] {
		var handleFunc func(word string, find []any) ([]any, bool)
		if strings.Index(word, "array") == -1 {
			handleFunc = handleKey
		} else {
			handleFunc = handleArray
		}
		f, arrayFlag = handleFunc(word, f)
	}
	if arrayFlag {
		result.arrayResult = f
	} else {
		result.valueResult = f[0]
	}
	if len(f) > 0 {
		result.getSuccess = true
	}
	return
}

func assertValue(ar analyzeResult, condition mCondition) (a assertResult) {
	a.order = condition.Order
	a.preLogicType = condition.PreLogicType
	if ar.getSuccess == false {
		return
	}
	if ar.valueResult != nil && util.Contains([]string{condition.CalculateType}, valueCalculate) {
		a.assertSuccess = assertSingle(ar.valueResult, condition.Value, condition.CalculateType)
	} else if ar.arrayResult != nil && util.Contains([]string{condition.CalculateType}, sliceCalculate) {
		a.assertSuccess = assertArray(ar.arrayResult, condition.Value, condition.CalculateType)
	}
	return
}

func assertSingle(result any, cv, c string) (r bool) {
	switch result.(type) {
	case string:
		r = assertString(result.(string), cv, c)
	case float64:
		r = assertNumber(result.(float64), cv, c)
	}
	return
}

func assertString(v string, cv, c string) (r bool) {
	switch c {
	case "=":
		if v == cv {
			r = true
		}
	case "!=":
		if v != cv {
			r = true
		}
	default:
		vNum, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return
		}
		cNum, err := strconv.ParseFloat(cv, 64)
		if err != nil {
			return
		}
		switch c {
		case "<":
			if vNum < cNum {
				r = true
			}
		case "<=":
			if vNum <= cNum {
				r = true
			}
		case ">":
			if vNum > cNum {
				r = true
			}
		case ">=":
			if vNum >= cNum {
				r = true
			}
		}
	}
	return
}

func assertNumber(v float64, cv, c string) (r bool) {
	cNum, err := strconv.ParseFloat(cv, 64)
	if err != nil {
		return
	}
	switch c {
	case "=":
		if v == cNum {
			r = true
		}
	case "!=":
		if v != cNum {
			r = true
		}
	case "<":
		if v < cNum {
			r = true
		}
	case "<=":
		if v <= cNum {
			r = true
		}
	case ">":
		if v > cNum {
			r = true
		}
	case ">=":
		if v >= cNum {
			r = true
		}
	}
	return
}

func assertArray(result []any, cv, calculateType string) (r bool) {
	for _, v := range result {
		switch v.(type) {
		case string:
			r = assertArraySingleString(v.(string), cv, calculateType)
		case float64:
			r = assertArraySingleFloat64(v.(float64), cv, calculateType)
		}
	}
}

func assertArraySingleString(v string, cv, calculateType string) (r bool) {

}

func assertArraySingleFloat64(v float64, cv, c string) (r bool) {

}

func handleArray(word string, find []any) (result []any, flag bool) {
	re, _ := regexp.Compile(`\[([0-9]*)]`)
	indexes := re.FindStringSubmatchIndex(word)
	index := word[indexes[2]:indexes[3]]
	if index == "" {
		result = handleArrayAll(find)
		flag = true
	} else {
		result = handleArrayIndex(index, find)
	}
	return
}

func handleArrayAll(find []any) (result []any) {
	for _, item := range find {
		s, ok := item.([]any)
		if !ok {
			continue
		}
		for _, v := range s {
			result = append(result, v)
		}
	}
	return
}

func handleArrayIndex(index string, find []any) (result []any) {
	for _, item := range find {
		num, err := strconv.ParseInt(index, 10, 64)
		if err != nil {
			continue
		}
		s, ok := item.([]any)
		if !ok {
			continue
		}
		if num < 0 || int(num) >= len(s) {
			continue
		}
		result = append(result, s[num])
	}
	return
}

func handleKey(word string, find []any) (result []any, flag bool) {
	for _, item := range find {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		item, ok = m[word]
		if !ok {
			continue
		}
		result = append(result, item)
	}
	return
}
