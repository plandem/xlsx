package rule

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for TimePeriodType
const (
	_ primitives.TimePeriodType = iota
	TimePeriodToday
	TimePeriodYesterday
	TimePeriodTomorrow
	TimePeriodLast7Days
	TimePeriodThisMonth
	TimePeriodLastMonth
	TimePeriodNextMonth
	TimePeriodThisWeek
	TimePeriodLastWeek
	TimePeriodNextWeek
)

func init() {
	primitives.FromTimePeriodType = map[primitives.TimePeriodType]string{
		TimePeriodToday:     "today",
		TimePeriodYesterday: "yesterday",
		TimePeriodTomorrow:  "tomorrow",
		TimePeriodLast7Days: "last7Days",
		TimePeriodThisMonth: "thisMonth",
		TimePeriodLastMonth: "lastMonth",
		TimePeriodNextMonth: "nextMonth",
		TimePeriodThisWeek:  "thisWeek",
		TimePeriodLastWeek:  "lastWeek",
		TimePeriodNextWeek:  "nextWeek",
	}

	primitives.ToTimePeriodType = make(map[string]primitives.TimePeriodType, len(primitives.FromTimePeriodType))
	for k, v := range primitives.FromTimePeriodType {
		primitives.ToTimePeriodType[v] = k
	}
}
