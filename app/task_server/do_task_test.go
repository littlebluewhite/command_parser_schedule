package task_server

import (
	"command_parser_schedule/entry/e_task_template"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStages(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		ts := []e_task_template.TaskStage{
			{StageNumber: 1, Name: "w"},
			{StageNumber: 2, Name: "e"},
			{StageNumber: 6, Name: "r"},
			{StageNumber: 1, Name: "t"},
			{StageNumber: 3, Name: "y"},
		}
		sns, m := getStages(ts)
		fmt.Println(sns)
		fmt.Println(m)
		require.Contains(t, sns, int32(1))
		require.Contains(t, sns, int32(2))
		require.Contains(t, sns, int32(3))
		require.Contains(t, sns, int32(6))
	})
}
