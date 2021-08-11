package iutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLatestDaysBaseNow(t *testing.T) {
	fn := func(days []time.Time) {
		for i := range days {
			t.Log(days[i].Format(Y_M_D))
		}
	}

	t.Log("-7 days asc")
	fn(GetLatestDaysBaseNow(-7))

	t.Log("-7 days desc")
	fn(GetLatestDaysBaseNow(-7, true))

	t.Log(" 0 days")
	fn(GetLatestDaysBaseNow(0))
	fn(GetLatestDaysBaseNow(0, true))

	t.Log("+7 days asc")
	fn(GetLatestDaysBaseNow(+7))

	t.Log("+7 days desc")
	fn(GetLatestDaysBaseNow(+7, true))
}

func TestWeekInfo(t *testing.T) {
	// 上周
	t.Log(WeekInfo(time.Now().AddDate(0, 0, -7)))
	// 本周
	t.Log(WeekInfo())
}

func TestYmdStr(t *testing.T) {
	t.Log(YmdStr(2020, 11, 26))
	t.Log(YmdStr(2020, 11, 26, "/"))
}

func TestLastMonthDay(t *testing.T) {
	loc, _ := time.LoadLocation("Local")

	now := time.Date(2019, 3, 1, 0, 0, 0, 0, loc)
	target := LastMonthDay(2, 29, now)
	expected := time.Date(2016, 2, 29, 0, 0, 0, 0, loc)
	assert.Equal(t, expected.Unix(), target.Unix())

	now1 := time.Date(2021, 3, 1, 0, 0, 0, 0, loc)
	target1 := LastMonthDay(2, 28, now1)
	expected1 := time.Date(2021, 2, 28, 0, 0, 0, 0, loc)
	assert.Equal(t, expected1.Unix(), target1.Unix())

	now2 := time.Date(2021, 3, 1, 0, 0, 0, 0, loc)
	target2 := LastMonthDay(4, 28, now2)
	expected2 := time.Date(2020, 4, 28, 0, 0, 0, 0, loc)
	assert.Equal(t, expected2.Unix(), target2.Unix())
}
