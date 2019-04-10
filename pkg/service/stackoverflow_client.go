package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/justina777/website-stackflow/pkg/schema"

	resty "gopkg.in/resty.v1"
)

type StackOverflowClient struct {
}

const (
	API_HOST          = "https://api.stackexchange.com/"
	API_VERSION       = "2.2"
	FUNCTION_QUESTION = "questions"
	PAGE_SIZE         = 10
	MAXPAGE           = 100
)

func (s *StackOverflowClient) Fetch(sort string, fromDate time.Time, page int) *schema.Questions {

	url := fmt.Sprintf("%s%s/%s", API_HOST, API_VERSION, FUNCTION_QUESTION)
	if page > MAXPAGE {
		page = MAXPAGE
	}
	resp, err := resty.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		SetQueryParams(map[string]string{
			"order":    "desc",
			"sort":     sort,
			"page":     strconv.Itoa(page),
			"site":     "stackoverflow",
			"tagged":   "android",
			"pagesize": strconv.Itoa(PAGE_SIZE),
			"fromdate": fromDate.Format("2006-01-02"),
			"filter":   "!9Z(-wno.B",
		}).
		Get(url)

	if err != nil {
		fmt.Println("error is ", err)
	}

	if resp.StatusCode() != 200 {
		fmt.Println("error is ", string(resp.Body()))
	}
	result := &schema.Questions{}
	err = json.Unmarshal(resp.Body(), result)
	fmt.Println("backoff time ", result.Backoff)
	return result
}

func (s *StackOverflowClient) GetQuestion(url string) string {

	resp, err := resty.R().
		SetHeader("accept-encoding", "gzip, deflate, br").
		SetHeader("accept-language", "en-US,en;q=0.9").
		SetHeader("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36").
		Get(url)

	if err != nil {
		fmt.Println("error is ", err)
	}

	if resp.StatusCode() != 200 {
		fmt.Println("error is ", string(resp.Body()))
	}

	return string(resp.Body())
}
