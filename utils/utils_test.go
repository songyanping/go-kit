package utils_test

import (
	"fmt"
	"github.com/songyanping/go-kit/utils"
	"testing"
	"time"
)

func TestTimeFormatNow(t *testing.T) {

	//layout参数可以是"time.DateTime"或"time.RFC3339"
	//timeString, err := utils.TimeFormatNow(time.DateTime)
	timeString, err := utils.TimeFormatNow(time.RFC3339)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeString)

	timeString, err = utils.TimeStrFormatCST("2024-07-18T10:44:22+08:00")
	fmt.Println(timeString)
	tt, _ := utils.TimeStrFormatTime("2024-07-18 10:44:22")
	fmt.Println(tt.Local())
}
