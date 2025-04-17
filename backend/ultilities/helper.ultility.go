package ultilities

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type BindType uint8

var (
	BindTypeJson  BindType = 0
	BindTypeUri   BindType = 1
	BindTypeQuery BindType = 2
	BindTypeForm  BindType = 4
)

func BindRequest(ctxt *gin.Context, bind BindType, reqt any) error {
	switch bind {
	case BindTypeUri:
		return ctxt.ShouldBindUri(reqt)
	case BindTypeQuery:
		return ctxt.ShouldBindQuery(reqt)
	case BindTypeForm:
		return ctxt.ShouldBind(reqt)
	default:
		return ctxt.ShouldBindJSON(reqt)
	}
}

func ParseObjectFromJson(jsonData []byte, object any) error {
	return json.Unmarshal(jsonData, object)
}
