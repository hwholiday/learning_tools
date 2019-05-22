package huawei

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
	"net/url"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(10*time.Minute, 15*time.Minute)

func NewHuaweiPush(tokenUrl, appId, appSecret, pushUrl string, auto bool) *HuaweiPush {
	return &HuaweiPush{tokenUrl, appId, appSecret, pushUrl, auto}
}

func (p *HuaweiPush) GetToken() (*ResToken, error) {
	var (
		req          http.Request
		bodyStr      string
		err          error
		res          *http.Response
		resToken     *ResToken
		resTokenByte []byte
	)
	req.ParseForm()
	req.Form.Add("grant_type", "client_credentials")
	req.Form.Add("client_secret", p.AppSecret)
	req.Form.Add("client_id", p.AppId)
	bodyStr = strings.TrimSpace(req.Form.Encode())
	if res, err = http.Post(p.TokenUrl, "application/x-www-form-urlencoded", strings.NewReader(bodyStr)); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if resTokenByte, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resTokenByte, &resToken); err != nil {
		fmt.Println("GetToken: ", string(resTokenByte))
		return nil, err
	}
	if p.AutoCacheAccessToken {
		c.Set("AccessToken", resToken.AccessToken, time.Duration(resToken.ExpiresIn-100)*time.Second)
	}
	return resToken, nil
}

func (p *HuaweiPush) Push(in *ReqPush) (*ResPush, error) {
	var (
		req                  http.Request
		bodyStr, AccessToken string
		err                  error
		res                  *http.Response
		resPush              *ResPush
		resPushByte          []byte
		urlPath              *url.URL
	)
	req.ParseForm()
	if p.AutoCacheAccessToken == true {
		info, has := c.Get("AccessToken")
		if !has {
			data, err := p.GetToken()
			if err != nil {
				return nil, err
			}
			AccessToken = data.AccessToken
		} else {
			AccessToken = fmt.Sprint(info)
		}
	} else {
		AccessToken = in.AccessToken
	}
	req.Form.Add("access_token", AccessToken)
	if in.NspSvc == "" {
		req.Form.Add("nsp_svc", "openpush.message.api.send")
	} else {
		req.Form.Add("nsp_svc", in.NspSvc)
	}
	if in.NspTs == "" {
		req.Form.Add("nsp_ts", fmt.Sprint(time.Now().Unix()))
	} else {
		req.Form.Add("nsp_ts", in.NspTs)
	}
	req.Form.Add("device_token_list", fmt.Sprintf("[\"%s\"]", strings.Join(in.DeviceTokenList, "\",\"")))
	if in.ExpireTime != "" {
		req.Form.Add("expire_time", in.ExpireTime)
	}
	req.Form.Add("payload", in.Payload)
	bodyStr = strings.TrimSpace(req.Form.Encode())
	if urlPath, err = url.Parse(fmt.Sprintf("%s?nsp_ctx=%v", p.PushUrl, fmt.Sprintf("{\"ver\":%s, \"appId\":%s}", in.Ver, p.AppId))); err != nil {
		return nil, err
	}
	if res, err = http.Post(fmt.Sprintf("%s?%s", p.PushUrl, urlPath.Query().Encode()), "application/x-www-form-urlencoded", strings.NewReader(bodyStr)); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if resPushByte, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resPushByte, &resPush); err != nil {
		fmt.Println("Push :", string(resPushByte))
		return nil, err
	}
	return resPush, nil
}
