package mock

import (
	"time"

	"github.com/rs/zerolog/log"

	"ddd/pkg/helper/helpertype"
)

// 提供 fake 標準庫的時間型別
func StandardTime(fakeTime string) time.Time {
	return TimeNowFunc(fakeTime)()
}

// 提供 fake 自定義的時間型別
func Time(fakeTime string) helpertype.Time {
	return helpertype.Time{TimeNowFunc(fakeTime)()}
}

// 提供 fake 時間 Now 函數
func TimeNowFunc(fakeTime string) func() time.Time {
	return func() time.Time {
		t, err := helpertype.TimeTool{}.Parse(fakeTime)
		if err != nil {
			log.Fatal().Msgf("New fakeTimeNow function failed: %v\n%+[1]v", err)
		}
		return t
	}
}
