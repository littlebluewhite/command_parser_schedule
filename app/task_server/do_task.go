package task_server

import (
	"command_parser_schedule/entry/e_task"
	"command_parser_schedule/entry/e_task_template"
	"context"
	"sort"
)

func (t *taskServer) doTask(task e_task.Task) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stages := task.Template.Stages
	sns, stageMap := getStages(stages)
	for _, sn := range sns {
		s := stageMap[sn]
		doStages(s)
	}
}

// getStages return stage number array without duplicates and return the map (stage number as key stages as value)
func getStages(stages []e_task_template.TaskStage) (sns []int32, stageMap map[int32][]e_task_template.TaskStage) {
	snSet := make(map[int32]struct{})
	stageMap = make(map[int32][]e_task_template.TaskStage)
	for i := 0; i < len(stages); i++ {
		sn := stages[i].StageNumber
		var ts []e_task_template.TaskStage
		if _, ok := snSet[sn]; !ok {
			sns = append(sns, sn)
			snSet[sn] = struct{}{}
			ts = []e_task_template.TaskStage{stages[i]}
		} else {
			ts = stageMap[sn]
			ts = append(ts, stages[i])
		}
		stageMap[sn] = ts
	}
	sort.Slice(sns, func(i, j int) bool {
		return sns[i] < sns[j]
	})
	return
}

func doStages(s []e_task_template.TaskStage) {

}
