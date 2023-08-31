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

func newTaskTemplateStage(db *gorm.DB, opts ...gen.DOOption) taskTemplateStage {
	_taskTemplateStage := taskTemplateStage{}

	_taskTemplateStage.taskTemplateStageDo.UseDB(db, opts...)
	_taskTemplateStage.taskTemplateStageDo.UseModel(&model.TaskTemplateStage{})

	tableName := _taskTemplateStage.taskTemplateStageDo.TableName()
	_taskTemplateStage.ALL = field.NewAsterisk(tableName)
	_taskTemplateStage.TaskTemplateID = field.NewInt32(tableName, "task_template_id")
	_taskTemplateStage.TaskStageID = field.NewInt32(tableName, "task_stage_id")

	_taskTemplateStage.fillFieldMap()

	return _taskTemplateStage
}

type taskTemplateStage struct {
	taskTemplateStageDo taskTemplateStageDo

	ALL            field.Asterisk
	TaskTemplateID field.Int32
	TaskStageID    field.Int32

	fieldMap map[string]field.Expr
}

func (t taskTemplateStage) Table(newTableName string) *taskTemplateStage {
	t.taskTemplateStageDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskTemplateStage) As(alias string) *taskTemplateStage {
	t.taskTemplateStageDo.DO = *(t.taskTemplateStageDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskTemplateStage) updateTableName(table string) *taskTemplateStage {
	t.ALL = field.NewAsterisk(table)
	t.TaskTemplateID = field.NewInt32(table, "task_template_id")
	t.TaskStageID = field.NewInt32(table, "task_stage_id")

	t.fillFieldMap()

	return t
}

func (t *taskTemplateStage) WithContext(ctx context.Context) *taskTemplateStageDo {
	return t.taskTemplateStageDo.WithContext(ctx)
}

func (t taskTemplateStage) TableName() string { return t.taskTemplateStageDo.TableName() }

func (t taskTemplateStage) Alias() string { return t.taskTemplateStageDo.Alias() }

func (t taskTemplateStage) Columns(cols ...field.Expr) gen.Columns {
	return t.taskTemplateStageDo.Columns(cols...)
}

func (t *taskTemplateStage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskTemplateStage) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["task_template_id"] = t.TaskTemplateID
	t.fieldMap["task_stage_id"] = t.TaskStageID
}

func (t taskTemplateStage) clone(db *gorm.DB) taskTemplateStage {
	t.taskTemplateStageDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t taskTemplateStage) replaceDB(db *gorm.DB) taskTemplateStage {
	t.taskTemplateStageDo.ReplaceDB(db)
	return t
}

type taskTemplateStageDo struct{ gen.DO }

func (t taskTemplateStageDo) Debug() *taskTemplateStageDo {
	return t.withDO(t.DO.Debug())
}

func (t taskTemplateStageDo) WithContext(ctx context.Context) *taskTemplateStageDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskTemplateStageDo) ReadDB() *taskTemplateStageDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskTemplateStageDo) WriteDB() *taskTemplateStageDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskTemplateStageDo) Session(config *gorm.Session) *taskTemplateStageDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskTemplateStageDo) Clauses(conds ...clause.Expression) *taskTemplateStageDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskTemplateStageDo) Returning(value interface{}, columns ...string) *taskTemplateStageDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskTemplateStageDo) Not(conds ...gen.Condition) *taskTemplateStageDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskTemplateStageDo) Or(conds ...gen.Condition) *taskTemplateStageDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskTemplateStageDo) Select(conds ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskTemplateStageDo) Where(conds ...gen.Condition) *taskTemplateStageDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskTemplateStageDo) Order(conds ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskTemplateStageDo) Distinct(cols ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskTemplateStageDo) Omit(cols ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskTemplateStageDo) Join(table schema.Tabler, on ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskTemplateStageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskTemplateStageDo) RightJoin(table schema.Tabler, on ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskTemplateStageDo) Group(cols ...field.Expr) *taskTemplateStageDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskTemplateStageDo) Having(conds ...gen.Condition) *taskTemplateStageDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskTemplateStageDo) Limit(limit int) *taskTemplateStageDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskTemplateStageDo) Offset(offset int) *taskTemplateStageDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskTemplateStageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *taskTemplateStageDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskTemplateStageDo) Unscoped() *taskTemplateStageDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskTemplateStageDo) Create(values ...*model.TaskTemplateStage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskTemplateStageDo) CreateInBatches(values []*model.TaskTemplateStage, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskTemplateStageDo) Save(values ...*model.TaskTemplateStage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskTemplateStageDo) First() (*model.TaskTemplateStage, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplateStage), nil
	}
}

func (t taskTemplateStageDo) Take() (*model.TaskTemplateStage, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplateStage), nil
	}
}

func (t taskTemplateStageDo) Last() (*model.TaskTemplateStage, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplateStage), nil
	}
}

func (t taskTemplateStageDo) Find() ([]*model.TaskTemplateStage, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskTemplateStage), err
}

func (t taskTemplateStageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskTemplateStage, err error) {
	buf := make([]*model.TaskTemplateStage, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskTemplateStageDo) FindInBatches(result *[]*model.TaskTemplateStage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskTemplateStageDo) Attrs(attrs ...field.AssignExpr) *taskTemplateStageDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskTemplateStageDo) Assign(attrs ...field.AssignExpr) *taskTemplateStageDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskTemplateStageDo) Joins(fields ...field.RelationField) *taskTemplateStageDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskTemplateStageDo) Preload(fields ...field.RelationField) *taskTemplateStageDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskTemplateStageDo) FirstOrInit() (*model.TaskTemplateStage, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplateStage), nil
	}
}

func (t taskTemplateStageDo) FirstOrCreate() (*model.TaskTemplateStage, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskTemplateStage), nil
	}
}

func (t taskTemplateStageDo) FindByPage(offset int, limit int) (result []*model.TaskTemplateStage, count int64, err error) {
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

func (t taskTemplateStageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskTemplateStageDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskTemplateStageDo) Delete(models ...*model.TaskTemplateStage) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskTemplateStageDo) withDO(do gen.Dao) *taskTemplateStageDo {
	t.DO = *do.(*gen.DO)
	return t
}
