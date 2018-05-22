package znotify

import (
	"encoding/hex"
	"net/url"
	"time"
	"io/ioutil"
	"strings"
	"net/http"
	"fmt"
	"crypto/md5"
	"strconv"
)

/**
 * account: 用户中心->语音通知->产品总览->APIID
 * password: 用户中心->语音通知->产品总览->APIKEY
 * mobile: 手机号码
 * content: 短信内容
 */
func SendVoice(appId,appKey,mobile,content string){
	md5Str :=func(s string) string {
		h := md5.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	}
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	v.Set("account", appId)
	v.Set("password", md5Str(appId+appKey+mobile+content+_now))
	v.Set("mobile", mobile)
	v.Set("content", content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://api.vm.ihuyi.com/webservice/voice.php?method=Submit&format=json", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构
	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}

