package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"login-demo/Cache"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type data struct {
	CreateTime int64 `json:"createTime"`
	UserId int `json:"userId"`
	Token string `json:"token"`
	RealName string `json:realName`
}
type result struct {
	Code int `json:"code"`
	Data data `json:"data"`
}
func Login(ctx iris.Context) {
	userName := ctx.Request().FormValue("username")
	passWord := ctx.Request().FormValue("password")

	userInfo := url.Values{}
	userInfo.Add("userName", userName)
	userInfo.Add("password", passWord)

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
<<<<<<< HEAD
		"http://tomonori.cc",
=======
		"https://tomonori.cc",
>>>>>>> master
		strings.NewReader(userInfo.Encode()),
	)
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	request.Header.Set("cms-token", "null")
	request.Header.Set("cms-channel", "0")
	if err != nil {
		ctx.JSON(err)
	}

	resp, _ := client.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var resultObj result
		var redisKey string
		var redisValue string
		var err error

		err = json.Unmarshal(body, &resultObj)
		if err == nil {
			ctx.JSON(resultObj)
		}

<<<<<<< HEAD
		redisKey = fmt.Sprintf("%s:%d", "AIMY_BIGDATA_USER", resultObj.Data.UserId)
		redisJsonValue, _ := json.Marshal(resultObj.Data)
		redisValue = string(redisJsonValue)

		res, err := Cache.Instance().Get(redisKey)
		if err != nil {
			ctx.JSON(err)
		}
		_ = json.Unmarshal([]byte(res), &resultObj.Data)

		if res == "" {
			_, err = Cache.Instance().Set(redisKey, redisValue, 0)
			if err != nil {
				ctx.JSON(err)
			}

			redisKey = fmt.Sprintf("%s:%s", "AIMY_BIGDATA_TOKEN", resultObj.Data.Token)
			redisValue = strconv.Itoa(resultObj.Data.UserId)
			_, err = Cache.Instance().Set(redisKey, redisValue, 0)
			if err != nil {
				ctx.JSON(err)
			}
		} else {
			oldToken := resultObj.Data.Token
			fmt.Println(oldToken)
			_, err = Cache.Instance().Set(redisKey, redisValue, 0)
			if err != nil {
				ctx.JSON(err)
			}
		}
	}
}
=======
}
>>>>>>> master
