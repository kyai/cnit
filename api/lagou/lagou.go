package lagou

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kyai/gurl"
)

type Response struct {
	Success bool `json:"success"`
	Content struct {
		PositionResult struct {
			Result []struct {
				FirstType       string `json:"firstType"`
				SecondType      string `json:"secondType"`
				ThirdType       string `json:"thirdType"`
				CompanyFullName string `json:"companyFullName"`
				Salary          string `json:"salary"`
			} `json:"result"`
		} `json:"positionResult"`
	} `json:"content"`
}

func Get(city, keyword string, page int) (res *Response, err error) {
	URL := "https://www.lagou.com/jobs/positionAjax.json"

	param := map[string]string{
		"city":                city,
		"needAddtionalResult": "false",
	}

	data := map[string]string{
		"first": "false",
		"pn":    strconv.Itoa(page),
		"kd":    keyword,
	}

	web := fmt.Sprintf("https://www.lagou.com/jobs/list_%s?city=%s", keyword, city)

	header := map[string]string{
		"Cookie":  cookieStr(web),
		"Referer": web,
	}

	body, err := gurl.New("POST", URL).Param(param).Data(data, gurl.FORM).Header(header).Do()
	if err != nil {
		return
	}

	res = &Response{}
	err = json.Unmarshal(body, res)

	return
}

func cookieStr(url string) (str string) {
	resp, _ := http.Get(url)
	for _, cookie := range resp.Cookies() {
		str += cookie.String() + ";"
	}
	return
}
