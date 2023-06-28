package main

import (
	"command_parser_schedule/app/sql"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		/* Mode: gen.WithoutContext,*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
	})

	db, err := sql.NewDB("mySQL", "gen_sql.log", "db")
	if err != nil {
		panic(err)
	}

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary, or it will panic
	g.UseDB(db)
	timeData := g.GenerateModel("time_data",
		gen.FieldType("m_condition", "json.RawMessage"),
		gen.FieldType("start_time", "[]byte"),
		gen.FieldType("end_time", "[]byte"))
	timeTemplate := g.GenerateModel("time_template", gen.FieldRelate(field.BelongsTo, "TimeData", timeData,
		&field.RelateConfig{
			GORMTag: map[string]string{"foreignKey": "time_data_id"},
		}))
	g.ApplyBasic(timeData, timeTemplate)

	// execute the action of code generation
	g.Execute()
}
