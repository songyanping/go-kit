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

type MyEvent struct {
	Name string
	Time time.Time
}

type MyEventString struct {
	Name string
	Time string
}

func TestTimeSortStructsByFieldTime(t *testing.T) {

	var events = []*MyEvent{}
	for i := 0; i < 5; i++ {
		events = append(events, &MyEvent{
			Name: fmt.Sprintf("Event%d", i+1),
			Time: time.Date(2022, 10, 1, 12, i+1, 0, 0, time.UTC),
		})
	}

	sortedEvents, err := utils.TimeSortStructsByFieldTime(events, "Time", false) // 按时间降序排列
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, event := range sortedEvents.([]*MyEvent) { // 类型断言转换回原始类型
		fmt.Println(event.Name, event.Time)
	}
}

func TestTimeSortStructsByFieldString(t *testing.T) {

	events := []*MyEventString{
		{Name: "Event1", Time: "2022-10-12 22:11:00"},
		{Name: "Event2", Time: "2022-09-11 20:10:00"},
		{Name: "Event3", Time: "2022-11-13 23:12:00"},
	}

	sortedEvents, err := utils.TimeSortStructsByFieldString(events, "Time", false) // 按时间降序排列
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, event := range sortedEvents.([]*MyEventString) { // 类型断言转换回原始类型
		fmt.Println(event.Name, event.Time)
	}
}

func TestTimeConvertTimestampString(t *testing.T) {

	stringTime, _ := utils.TimeConvertSecondsString(1721983237)
	fmt.Println(stringTime)
}

func TestTimeStrFormatTime(t *testing.T) {

	stringTime, _ := utils.TimeStrFormatTime("2024-07-30 10:30:00")
	stringTime2, _ := utils.TimeStrFormatTime("2024-07-30 10:10:00")
	beforeMinutes := stringTime2.Minute() - stringTime.Minute()
	fmt.Println(beforeMinutes)
}

func TestGetStartAndEndOfWeek(t *testing.T) {
	// 获取当前时间
	//now := time.Now()
	now := time.Date(2024, 8, 28, 2, 0, 0, 0, time.Local)

	// 计算7天前的时间
	sevenDaysAgo := now.AddDate(0, 0, -6)

	// 格式化日期
	startDateStr := sevenDaysAgo.Format("2006/01/02")
	endDateStr := now.Format("2006/01/02")

	// 拼接日期周期字符串
	dateRangeStr := fmt.Sprintf("%s-%s", startDateStr, endDateStr)
	fmt.Println(dateRangeStr)
}
