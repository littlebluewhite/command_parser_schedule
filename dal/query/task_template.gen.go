// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"command_parser_schedule/dal/model"
)

func newTaskTemplate(db *gorm.DB, opts ...gen.DOOption) taskTemplate {
	_taskTemplate := taskTemplate{}

	_taskTemplate.taskTemplateDo.UseDB(db, opts...)
	_taskTemplate.taskTemplateDo.UseModel(&model.TaskTemplate{})

	tableName := _taskTemplate.taskTemplateDo.TableName()
	_taskTemplate.ALL = field.NewAsterisk(tableName)
	_taskTemplate.ID = field.NewInt32(tableName, "id")
	_taskTemplate.Name = field.NewString(tableName, "name")
	_taskTemplate.Variable = field.NewBytes(tableName, "variable")
	_taskTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_taskTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_taskTemplate.Tags = field.NewBytes(tableName, "tags")
	_taskTemplate.Stages = taskTemplateManyToManyStages{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Stages", "model.TaskStage"),
		CommandTemplate: struct {
			field.RelationField
			Http struct {
				field.RelationField
			}
			Mqtt struct {
				field.RelationField
			}
			Websocket struct {
				field.RelationField
			}
			Redis struct {
				field.RelationField
			}
			Monitor struct {
				field.RelationField
				MConditions struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Stages.CommandTemplate", "model.CommandTemplate"),
			Http: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Stages.CommandTemplate.Http", "model.HTTPSCommand"),
			},
			Mqtt: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Stages.CommandTemplate.Mqtt", "model.MqttCommand"),
			},
			Websocket: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Stages.CommandTemplate.Websocket", "model.WebsocketCommand"),
			},
			Redis: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Stages.CommandTemplate.Redis", "model.RedisCommand"),
			},
			Monitor: struct {
				field.RelationField
				MConditions struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Stages.CommandTemplate.Monitor", "model.Monitor"),
				MConditions: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Stages.CommandTemplate.Monitor.MConditions", "model.MCondition"),
				},
			},
		},
	}

	_taskTemplate.fillFieldMap()

	return _taskTemplate
}

type taskTemplate struct {
	taskTemplateDo taskTemplateDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String
	Variable  field.Bytes
	UpdatedAt field.Time
	CreatedAt field.Time
	Tags      field.Bytes
	Stages    taskTemplateManyToManyStages

	fieldMap map[string]field.Expr
}

func (t taskTemplate) Table(newTableName string) *taskTemplate {
	t.taskTemplateDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskTemplate) As(alias string) *taskTemplate {
	t.taskTemplateDo.DO = *(t.taskTemplateDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskTemplate) updateTableName(table string) *taskTemplate {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Name = field.NewString(table, "name")
	t.Variable = field.NewBytes(table, "variable")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.Tags = field.NewBytes(table, "tags")

	t.fillFieldMap()

	return t
}

func (t *taskTemplate) WithContext(ctx context.Context) *taskTemplateDo {
	return t.taskTemplateDo.WithContext(ctx)
}

func (t taskTemplate) TableName() string { return t.taskTemplateDo.TableName() }

func (t taskTemplate) Alias() string { return t.taskTemplateDo.Alias() }

func (t taskTemplate) Columns(cols ...field.Expr) gen.Columns {
	return t.taskTemplateDo.Columns(cols...)
}

func (t *taskTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskTemplate) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
	t.fieldMap["variable"] = t.Variable
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["tags"] = t.Tags

}

func (t taskTemplate) clone(db *gorm.DB) taskTemplate {
	t.taskTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t taskTemplate) replaceDB(db *gorm.DB) taskTemplate {
	t.taskTemplateDo.ReplaceDB(db)
	return t
}

type taskTemplateManyToManyStages struct {
	db *gorm.DB

	field.RelationField

	CommandTemplate struct {
		field.RelationField
		Http struct {
			field.RelationField
		}
		Mqtt struct {
			field.RelationField
		}
		Websocket struct {
			field.RelationField
		}
		Redis struct {
			field.RelationField
		}
		Monitor struct {
			field.RelationField
			MConditions struct {
				field.RelationField
			}
		}
	}
}

func (a taskTemplateManyToManyStages) Where(conds ...field.Expr) *taskTemplateManyToManyStages {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a taskTemplateManyToManyStages) WithContext(ctx context.Context) *taskTemplateManyToManyStages {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a taskTemplateManyToManyStages) Session(session *gorm.Session) *taskTemplateManyToManyStages {
	a.db = a.db.Session(session)
	return &a
}

func (a taskTemplateManyToManyStages) Model(m *model.TaskTemplate) *taskTemplateManyToManyStagesTx {
	return &taskTemplateManyToManyStagesTx{a.db.Model(m).Association(a.Name())}
}

type taskTemplateManyToManyStagesTx struct{ tx *gorm.Association }

func (a taskTemplateManyToManyStagesTx) Find() (result []*model.TaskStage, err error) {
	return result, a.tx.Find(&result)
}

func (a taskTemplateManyToManyStagesTx) Append(values ...*model.TaskStage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a taskTemplateManyToManyStagesTx) Replace(values ...*model.TaskStage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a taskTemplateManyToManyStagesTx) Delete(values ...*model.TaskStage) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a taskTemplateManyToManyStagesTx) Clear() error {
	return a.tx.Clear()
}

func (a taskTemplateManyToManyStagesTx) Count() int64 {
	return a.tx.Count()
}

type taskTemplateDo struct{ gen.DO }

func (t taskTemplateDo) Debug() *taskTemplateDo {
	return t.withDO(t.DO.Debug())
}

func (t taskTemplateDo) WithContext(ctx context.Context) *taskTemplateDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskTemplateDo) ReadDB() *taskTemplateDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskTemplateDo) WriteDB() *taskTemplateDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskTemplateDo) Session(config *gorm.Session) *taskTemplateDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskTemplateDo) Clauses(conds ...clause.Expression) *taskTemplateDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskTemplateDo) Returning(value interface{}, columns ...string) *taskTemplateDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskTemplateDo) Not(conds ...gen.Condition) *taskTemplateDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskTemplateDo) Or(conds ...gen.Condition) *taskTemplateDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskTemplateDo) Select(conds ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskTemplateDo) Where(conds ...gen.Condition) *taskTemplateDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskTemplateDo) Order(conds ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskTemplateDo) Distinct(cols ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskTemplateDo) Omit(cols ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskTemplateDo) Join(table schema.Tabler, on ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskTemplateDo) Group(cols ...field.Expr) *taskTemplateDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskTemplateDo) Having(conds ...gen.Condition) *taskTemplateDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskTemplateDo) Limit(limit int) *taskTemplateDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskTemplateDo) Offset(offset int) *taskTemplateDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *taskTemplateDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskTemplateDo) Unscoped() *taskTemplateDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskTemplateDo) Create(values ...*model.TaskTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskTemplateDo) CreateInBatches(values []*model.TaskTemplate, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskTemplateDo) Save(values ...*model.TaskTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskTemplateDo) First() (*model.TaskTemplate, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplate), nil
	}
}

func (t taskTemplateDo) Take() (*model.TaskTemplate, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplate), nil
	}
}

func (t taskTemplateDo) Last() (*model.TaskTemplate, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplate), nil
	}
}

func (t taskTemplateDo) Find() ([]*model.TaskTemplate, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskTemplate), err
}

func (t taskTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskTemplate, err error) {
	buf := make([]*model.TaskTemplate, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskTemplateDo) FindInBatches(result *[]*model.TaskTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskTemplateDo) Attrs(attrs ...field.AssignExpr) *taskTemplateDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskTemplateDo) Assign(attrs ...field.AssignExpr) *taskTemplateDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskTemplateDo) Joins(fields ...field.RelationField) *taskTemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskTemplateDo) Preload(fields ...field.RelationField) *taskTemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskTemplateDo) FirstOrInit() (*model.TaskTemplate, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplate), nil
	}
}

func (t taskTemplateDo) FirstOrCreate() (*model.TaskTemplate, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplate), nil
	}
}

func (t taskTemplateDo) FindByPage(offset int, limit int) (result []*model.TaskTemplate, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t taskTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskTemplateDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskTemplateDo) Delete(models ...*model.TaskTemplate) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskTemplateDo) withDO(do gen.Dao) *taskTemplateDo {
	t.DO = *do.(*gen.DO)
	return t
}
