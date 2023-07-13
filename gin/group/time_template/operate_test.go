package time_template

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/util/logFile"
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/datatypes"
	"testing"
	"time"
)

func setUpOperate() (o Operate, l logFile.LogFile) {
	l = logFile.NewLogFile("test", "operate.log")
	dbs := dbs.NewDbs(l, true)
	o = NewOperate(dbs)
	return
}

func TestQuery(t *testing.T) {
	o, l := setUpOperate()
	t.Run("test find", func(t *testing.T) {
		l.Info().Println("tset operate time template list")
		var i int32 = 20
		testTimeTemplates := []*model.TimeTemplate{
			{Name: "test1",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "test2",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "test3",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "test4",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
		}
		timeTemplates, err := o.Create(testTimeTemplates)
		require.Nil(t, err)
		tIds := make([]int32, 0, 4)
		for _, item := range timeTemplates {
			tIds = append(tIds, item.ID)
		}
		timeTemplates2, err := o.Find(tIds)
		require.Nil(t, err)
		require.Equal(t, len(timeTemplates), 4)
		require.Equal(t, timeTemplates2[0].Name, "test1")
		require.Equal(t, timeTemplates2[1].Name, "test2")
		require.Equal(t, timeTemplates2[2].Name, "test3")
		require.Equal(t, timeTemplates2[3].Name, "test4")
	})
	t.Run("test List", func(t *testing.T) {
		l.Info().Println("test time templates list")
		timeTemplates, err := o.List()
		require.Nil(t, err)
		require.Equal(t, int32(4), timeTemplates[0].ID)
	})
}

func TestCreate(t *testing.T) {
	o, l := setUpOperate()
	t.Run("create success", func(t *testing.T) {
		l.Info().Println("test operate time template create")
		var i int32 = 300
		testTimeTemplate := []*model.TimeTemplate{
			{Name: "test6", TimeData: model.TimeDatum{
				RepeatType:      nil,
				StartDate:       time.Date(2023, 6, 16, 0, 0, 0, 0, time.Local),
				StartTime:       []byte("12:15:12"),
				EndTime:         []byte("13:21:13"),
				IntervalSeconds: &i,
				ConditionType:   nil,
				TCondition:      []byte("[1, 7, 3, 4]"),
			}},
		}
		result, err := o.Create(testTimeTemplate)
		fmt.Println(result)
		require.Nil(t, err)
		require.Equal(t, result[0].Name, "test6")
	})
	t.Run("create fail", func(t *testing.T) {

		l.Info().Println("test operate time template create")
		var i int32 = 300
		testTimeTemplate := []*model.TimeTemplate{
			{Name: "test6", TimeData: model.TimeDatum{
				RepeatType:      nil,
				StartDate:       time.Date(2023, 6, 19, 0, 0, 0, 0, time.Local),
				StartTime:       []byte("08:12:12"),
				EndTime:         []byte("13:09:13"),
				IntervalSeconds: &i,
				ConditionType:   nil,
				TCondition:      []byte("[1, 8, 3, 4]"),
			}},
		}
		result, err := o.Create(testTimeTemplate)
		fmt.Println(result)
		require.Nil(t, result)
		require.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	o, l := setUpOperate()
	t.Run("update", func(t *testing.T) {
		var s string = "monthly_day"
		name := "test1"
		startTime := datatypes.NewTime(8, 12, 12, 0)
		testTimeTemplate := []*TimeTemplateUpdate{
			{Name: &name, ID: 1,
				TimeData: &TimeDatumUpdate{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       &startTime,
					EndTime:         datatypes.NewTime(16, 55, 16, 0),
					IntervalSeconds: nil,
					ConditionType:   &s,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
		}
		l.Info().Println("test operate time template update")
		err := o.Update(testTimeTemplate)
		require.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	o, l := setUpOperate()
	t.Run("delete", func(t *testing.T) {
		var i int32 = 20
		testTimeTemplates := []*model.TimeTemplate{
			{Name: "apple",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "dog",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "banana",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
			{Name: "cherry",
				TimeData: model.TimeDatum{
					RepeatType:      nil,
					StartDate:       time.Date(2023, 6, 18, 0, 0, 0, 0, time.Local),
					StartTime:       []byte("05:12:12"),
					EndTime:         []byte("16:09:16"),
					IntervalSeconds: &i,
					ConditionType:   nil,
					TCondition:      []byte("[5, 1, 7]"),
				},
			},
		}
		l.Info().Println("test operate time template delete")
		timeTemplates, err := o.Create(testTimeTemplates)
		require.Nil(t, err)
		ids := make([]int32, 0, len(timeTemplates))
		for _, tt := range timeTemplates {
			ids = append(ids, tt.ID)
		}
		err = o.Delete(ids)
		require.Nil(t, err)
	})
}
