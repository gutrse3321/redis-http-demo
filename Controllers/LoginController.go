package Controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"login-demo/Cache"
	VMlogin "login-demo/ViewModel"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Login(ctx iris.Context) {
	userName := ctx.Request().FormValue("username")
	passWord := ctx.Request().FormValue("password")

	result, err := requestJavaServerService(userName, passWord)
	if err != nil {
		ctx.JSON(SendJson(iris.StatusBadGateway, err))
	} else {
		ctx.JSON(result)
	}
}

/**
 * 请求java服务端接口，返回请求成功或错误的值
 * @method requestJavaServerService
 * @param username string
 * @param password string
 * @return interface{}
 * @return error
 */
func requestJavaServerService(username, password string) (interface{}, error) {
	client := &http.Client{}
	var err error
	var urlAddr = "http://tomonori.cc"
	userInfo := url.Values{}
	userInfo.Add("userName", username)
	userInfo.Add("password", password)

	request, err := http.NewRequest(
		"POST",
		urlAddr,
		strings.NewReader(userInfo.Encode()),
	)
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	request.Header.Set("cms-token", "null")
	request.Header.Set("cms-channel", "0")
	if err != nil {
		return nil, err
	}

	parseObj, err := parseResponseBody(client, request)
	if err != nil {
		return nil, err
	}

	return parseObj, nil
}


type tempBody struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}
/**
 * 处理请求成功的body
 * @method parseResponseBody
 * @param client *http.Client
 * @param request *http.Request
 * @return obj interface{}
 * @return err error
 */
func parseResponseBody(client *http.Client, request *http.Request) (obj interface{}, err error) {
	var resultObj VMlogin.Vresult
	var errorObj VMlogin.Verror
	var tempObj tempBody
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tempObj)
	if err != nil {
		return nil, err
	}
	if tempObj.Code == 200 {
		_ = json.Unmarshal(body, &resultObj)
		obj = resultObj
		if err = cacheManagerService(resultObj); err != nil {
			return nil, err
		}
	} else {
		_ = json.Unmarshal(body, &errorObj)
		obj = errorObj
	}
	return obj, nil
}

/**
 * 设置处理缓存的工厂
 * @method cacheManagerService
 * @param resultObj VMlogin.Vresult
 * @return error
 */
func cacheManagerService(resultObj VMlogin.Vresult) error {
	var oldResultObj VMlogin.Vresult
	var redisKey string
	var redisValue string
	var err error

	// 组合查询用户信息的声明和赋值
	redisKey = FormatRedisString(REDIS_KEY_USER, resultObj.Data.UserId)
	redisJsonValue, _ := json.Marshal(resultObj.Data)
	redisValue = string(redisJsonValue)

	// 查询是否存在此用户
	infoVal, _ := Cache.Instance().Get(redisKey)

	// json反序列化查询到的json字符串，将存储在缓存中的值赋值到一个变量上
	_ = json.Unmarshal([]byte(infoVal), &oldResultObj.Data)
	// 如果返回空字符串，则是新用户，创建新的缓存
	// 否则根据缓存的token组合去将旧的缓存删除，覆盖新的用户信息和创建新的token
	if infoVal == "" {
		// 设置新的用户信息缓存
		err = Cache.Instance().Set(redisKey, redisValue, 744 * time.Hour)
		if err != nil {
			return err
		}

		// 设置新的token信息缓存
		redisKey = FormatRedisString(REDIS_KEY_TOKEN, resultObj.Data.Token)
		redisValue = strconv.Itoa(resultObj.Data.UserId)
		err = Cache.Instance().Set(redisKey, redisValue, 744 * time.Hour)
		if err != nil {
			return err
		}
	} else {
		// 声明并赋值旧的token信息
		oldToken := oldResultObj.Data.Token
		// 删除旧的token
		err = Cache.Instance().Del(FormatRedisString(REDIS_KEY_TOKEN, oldToken))
		if err != nil {
			return err
		}

		// 设置覆盖新的用户信息
		err = Cache.Instance().Set(redisKey, redisValue, 744 * time.Hour)
		if err != nil {
			return err
		}

		// 设置新的token
		redisKey = FormatRedisString(REDIS_KEY_TOKEN, resultObj.Data.Token)
		redisValue = strconv.Itoa(resultObj.Data.UserId)
		err = Cache.Instance().Set(redisKey, redisValue, 744 * time.Hour)
		if err != nil {
			return err
		}
	}
	return nil
}