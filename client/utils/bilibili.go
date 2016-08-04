package utils

func EncodeSign(params map[string]string, secret string) (string, string) {
	queryString := httpBuildQuery(params)
	return queryString, Md5(queryString + secret)
}
