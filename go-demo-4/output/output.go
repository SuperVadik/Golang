package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errValue, ok := value.(error)
	if ok {
		color.Red(errValue.Error())
		return
	}
	color.Red("Не известный тип: %T", value)
}

func PrintErrorWithTypeSwitch(value any) {
	switch v := value.(type) {
	case string:
		color.Red(v)
	case int:
		color.Red("Код ошибки: %d", v)
	case error:
		color.Red(v.Error())
	default:
		color.Red("Не известный тип: %T", v)
	}
}

func PromtData[T any](promt []T) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
