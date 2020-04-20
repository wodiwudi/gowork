package three

import (
	"fmt"
	"testing"
	"time"
)

type MemberCoinsList struct {
	Amount   string `json:"amount"`
	CreateAt string `json:"create_at"`
	Desc     string `json:"desc"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
}

func TestDateShow(t *testing.T) {
	testData := []MemberCoinsList{
		{
			Amount:   "24",
			CreateAt: "2020-03-11 10:17:52",
			Desc:     "幸运金币",
			Name:     "幸运金币",
			Remark:   "",
		}, {
			Amount:   "28",
			CreateAt: "2020-03-11 10:13:14",
			Desc:     "幸运金币",
			Name:     "幸运金币",
			Remark:   "",
		}, {
			Amount:   "38",
			CreateAt: "2020-03-11 08:37:35",
			Desc:     "每日签到奖励",
			Name:     "每日签到奖励",
			Remark:   "",
		}, {
			Amount:   "88",
			CreateAt: "2020-03-10 08:39:25",
			Desc:     "每日签到奖励",
			Name:     "每日签到奖励",
			Remark:   "",
		},
		{
			Amount:   "27",
			CreateAt: "2020-03-09 16:12:41",
			Desc:     "幸运金币",
			Name:     "幸运金币",
			Remark:   "",
		},
	}
	data := CoinSearchByTime(testData, "2020-03-10", "2020-03-11")
	fmt.Println(data)

}
func CoinSearchByTime(results []MemberCoinsList, startTime string, endTime string) []MemberCoinsList {
	startTimeFormat, _ := time.Parse("2006-01-02", startTime)
	startTimeStr := startTimeFormat.Unix()
	endTimeFormat, _ := time.Parse("2006-01-02", endTime)
	endTimeStr := endTimeFormat.Unix() + 86400
	var data []MemberCoinsList
	//时间筛选
	for _, result := range results {
		searchTime, _ := time.Parse("2006-01-02 15:04:05", result.CreateAt)
		if searchTime.Unix() > startTimeStr && searchTime.Unix() < endTimeStr {
			data = append(data, result)
		}
	}
	return data
}
