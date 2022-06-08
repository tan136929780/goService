package json

import "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func GetJson() jsoniter.API {
	return json
}
