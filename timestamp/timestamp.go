package timestamp

import (
	"fmt"
	"strconv"
)

func divmod(divisor int64, dividend int64) (int64, int64) {
	return divisor / dividend, divisor % dividend
}

type DivStr struct {
	Dividend int64
	Str      string
}

var list = [6]DivStr{DivStr{1000, ":"}, DivStr{60, ":"},
	DivStr{60, ":"}, DivStr{24, " "}, DivStr{30, "-"}, DivStr{12, "-"}}

func MsecDateFormat(msec string) string {
	// 毫秒时间戳转换为 YY-MM-dd h:m:s:ms
	result := ""
	millisecond, err := strconv.ParseInt(msec, 10, 64)
	if err != nil {
		return result
	}
	for _, divStr := range list {
		tmp, div := divmod(millisecond, divStr.Dividend)
		millisecond = tmp
		result = divStr.Str + strconv.FormatInt(div, 10) + result
	}
	result = strconv.FormatInt(millisecond, 10) + result
	return result
}

func main() {
	fmt.Println(MsecDateFormat("150359602")) // 0-0-1 17:45:59:602
}
