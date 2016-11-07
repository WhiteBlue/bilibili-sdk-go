package service

import (
	"encoding/json"
	"github.com/whiteblue/bilibili-go/client/utils"
)

type BaseParam struct {
	Appkey string
	Secret string
}

type BaseService struct {
	Params BaseParam
	Client utils.HttpClient
}

//mdzz..... Message和Error为什么要分两个
type apiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ApiError struct {
	Msg string
}

func (a *ApiError) Error() string {
	return a.Msg
}

func (b *BaseService) doRequest(url string, params map[string]string) ([]byte, error) {
	params["appkey"] = b.Params.Appkey
	//generate bilibili sign code
	query, sign := utils.EncodeSign(params, b.Params.Secret)
	reqUrl := url + "?" + query + "&sign=" + sign
	retByte, err := b.Client.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	var badRet apiResponse
	err = json.Unmarshal(retByte, &badRet)
	if err != nil {
		return nil, &ApiError{Msg: "api encode error"}
	}

	if badRet.Code == 0 {
		//api return success
		return retByte, nil
	}

	//api return error
	return nil, &ApiError{Msg: badRet.Message + badRet.Error}
}
