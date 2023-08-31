package time_template

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util"
	"command_parser_schedule/util/logFile"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
	"testing"
)

func setUpHandler() (l logFile.LogFile, modelConfig initial.GinApp) {
	l = logFile.NewLogFile("test", "handler.log")
	dbs := dbs.NewDbs(l, true)
	modelConfig = initial.NewGinApp(l, dbs)
	Inject(modelConfig)
	return
}

func TestGetList(t *testing.T) {
	l, app := setUpHandler()
	l.Info().Println("test handler get list")
	w := util.PerformRequest(app.GetRouter(), "GET", "/time_template/", nil)
	require.Equal(t, 200, w.Code)
	var response []TimeTemplate
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.Nil(t, err)
	util.P("============================")
	util.P(string(response[0].TimeData.EndTime))
}
