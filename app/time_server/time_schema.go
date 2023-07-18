package time_server

var allWeekDay = [...]int{0, 1, 2, 3, 4, 5, 6}
var allMonthDay = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

type RepeatType int

const (
	daily RepeatType = iota
	weekly
	monthly
)

func (r RepeatType) String() string {
	return [...]string{"daily", "weekly", "monthly"}[r]
}

type ConditionType int

const (
	monthDay ConditionType = iota
	weeklyDay
	weeklyFirst
	weeklySecond
	weeklyThird
	weeklyFourth
)

func (c ConditionType) String() string {
	return [...]string{
		"monthly_day", "weekly_day", "weekly_first",
		"weekly_second", "weekly_third", "weekly_fourth"}[c]
}
