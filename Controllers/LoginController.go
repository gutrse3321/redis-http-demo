package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Login(ctx iris.Context) {
	userName := ctx.Request().FormValue("username")
	passWord := ctx.Request().FormValue("password")

	userInfo := url.Values{}
	userInfo.Add("userName", userName)
	userInfo.Add("password", passWord)

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		"https://planet.aimymusic.com/cms/passport/login",
		//"https://dev.aimymusic.com/aimyplay/cms/passport/login",
		strings.NewReader(userInfo.Encode()),
	)
	request.Header.Set("cms-token", "null")
	request.Header.Set("cms-channel", "0")

	if err != nil {
		fmt.Errorf("%e\n", err)
	}

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	stdout := os.Stdout
	_, err = io.Copy(stdout, resp.Body)

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var resultObj interface{}
		err := json.Unmarshal(body, &resultObj)
		if err == nil {
			ctx.JSON(resultObj)
		}
	}

}