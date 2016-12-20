package client

import (
	"encoding/json"
)

type BaseParam struct {
	Appkey string
	Secret string
}

type BaseService struct {
	Params BaseParam
	Client HttpClient
}

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
	params["actionKey"] = "appkey"
	params["access_key"] = b.Params.Secret
	//generate bilibili sign code
	query, sign := EncodeSign(params, b.Params.Secret)
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
