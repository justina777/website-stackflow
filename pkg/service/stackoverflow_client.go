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
)

func (s *StackOverflowClient) Fetch(sort string, fromDate time.Time) *schema.Questions {

	url := fmt.Sprintf("%s%s/%s", API_HOST, API_VERSION, FUNCTION_QUESTION)
	resp, err := resty.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		SetQueryParams(map[string]string{
			"order":    "desc",
			"sort":     sort,
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

	return result
}
