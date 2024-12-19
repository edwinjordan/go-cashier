package helpers

import (
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
)

/*
	to get data type from any struct

	fstring values:
	snake
	kebab
	camel
*/
func GetStructDataType(s interface{}, fString string) map[string]interface{} {
	a := spew.Sdump(&s)
	dataType := make(map[string]interface{})
	for _, v := range strings.Split(a, "\n") {
		for _, vv := range strings.Split(v, ",") {
			val := strings.Split(vv, ": ")
			if len(val) > 1 {
				res := strings.Replace(strings.Split(val[1], " ")[0], "(", "", -1)
				res = strings.Replace(res, ")", "", -1)
				index := ""
				switch fString {
				case "snake":
					index = strcase.ToSnake(strings.Trim(val[0], " "))
				case "kebab":
					index = strcase.ToKebab(strings.Trim(val[0], " "))
				case "camel":
					index = strcase.ToCamel(strings.Trim(val[0], " "))
				default:
					index = strcase.ToSnake(strings.Trim(val[0], " "))
				}

				dataType[index] = res
			}
		}
	}

	return dataType
}

func CheckIndex(data []string, index int) interface{} {
	if len(data) > index {
		return data[index]
	} else {
		return ""
	}
}

func ConvertToInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		res = 0
	}
	return res
}

func ConvertToFloat(val string) float64 {
	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		res = 0
	}
	return res
}
