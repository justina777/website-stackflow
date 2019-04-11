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

	if resp.StatusCode() == 301 {
		fmt.Println("error is ", resp.StatusCode(), " ", resp.Header(), " ", string(resp.Body()))
	} else if resp.StatusCode() != 200 {
		fmt.Println("error is ", resp.StatusCode(), " ,err is ", err, ", ", string(resp.Body()))
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
		SetHeader("cookie", "prov=5e56334e-146b-e335-0b16-d34f251d0000; _ga=GA1.2.1277608371.1544000152; __qca=P0-2078927393-1544000152440; __gads=ID=c2e064d30acd89e3:T=1544000153:S=ALNI_MYUoMVZ-DsUItNFOro-dKbCHlxm2A; notice-ctt=4%3B1544768195148; _gid=GA1.2.1761761962.1554869132; se-consent=%7b%22s%22%3a1%2c%22d%22%3a%222019-04-11T04%3a11%3a08.0872901Z%22%7d; _gat=1; hero-dismissed=1554955874709!stk_a").
		SetHeader("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36").
		Get(url)

	if resp.StatusCode() == 301 || resp.StatusCode() == 302 {
		url = resp.Header().Get("Location")
		fmt.Println("301 error is ", resp.StatusCode(), " ", resp.Header().Get("Location"), " ", string(resp.Body()))
		resp, err = resty.R().
			SetHeader("accept-encoding", "gzip, deflate, br").
			SetHeader("accept-language", "en-US,en;q=0.9").
			SetHeader("cookie", "prov=5e56334e-146b-e335-0b16-d34f251d0000; _ga=GA1.2.1277608371.1544000152; __qca=P0-2078927393-1544000152440; __gads=ID=c2e064d30acd89e3:T=1544000153:S=ALNI_MYUoMVZ-DsUItNFOro-dKbCHlxm2A; notice-ctt=4%3B1544768195148; _gid=GA1.2.1761761962.1554869132; se-consent=%7b%22s%22%3a1%2c%22d%22%3a%222019-04-11T04%3a11%3a08.0872901Z%22%7d; _gat=1; hero-dismissed=1554955874709!stk_a").
			SetHeader("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36").
			Get(url)
	} else if resp.StatusCode() != 200 {
		fmt.Println("!= 200 error is ", resp.StatusCode(), " ,err is ", err, ", ", string(resp.Body()))
	}

	return string(resp.Body())
}
