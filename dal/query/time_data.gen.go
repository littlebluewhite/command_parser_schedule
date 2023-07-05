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

func newTimeDatum(db *gorm.DB, opts ...gen.DOOption) timeDatum {
	_timeDatum := timeDatum{}

	_timeDatum.timeDatumDo.UseDB(db, opts...)
	_timeDatum.timeDatumDo.UseModel(&model.TimeDatum{})

	tableName := _timeDatum.timeDatumDo.TableName()
	_timeDatum.ALL = field.NewAsterisk(tableName)
	_timeDatum.ID = field.NewInt32(tableName, "id")
	_timeDatum.RepeatType = field.NewString(tableName, "repeat_type")
	_timeDatum.StartDate = field.NewTime(tableName, "start_date")
	_timeDatum.EndDate = field.NewTime(tableName, "end_date")
	_timeDatum.StartTime = field.NewBytes(tableName, "start_time")
	_timeDatum.EndTime = field.NewBytes(tableName, "end_time")
	_timeDatum.IntervalSeconds = field.NewInt32(tableName, "interval_seconds")
	_timeDatum.ConditionType = field.NewString(tableName, "condition_type")
	_timeDatum.TCondition = field.NewBytes(tableName, "t_condition")

	_timeDatum.fillFieldMap()

	return _timeDatum
}

type timeDatum struct {
	timeDatumDo timeDatumDo

	ALL             field.Asterisk
	ID              field.Int32
	RepeatType      field.String
	StartDate       field.Time
	EndDate         field.Time
	StartTime       field.Bytes
	EndTime         field.Bytes
	IntervalSeconds field.Int32
	ConditionType   field.String
	TCondition      field.Bytes

	fieldMap map[string]field.Expr
}

func (t timeDatum) Table(newTableName string) *timeDatum {
	t.timeDatumDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeDatum) As(alias string) *timeDatum {
	t.timeDatumDo.DO = *(t.timeDatumDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeDatum) updateTableName(table string) *timeDatum {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.RepeatType = field.NewString(table, "repeat_type")
	t.StartDate = field.NewTime(table, "start_date")
	t.EndDate = field.NewTime(table, "end_date")
	t.StartTime = field.NewBytes(table, "start_time")
	t.EndTime = field.NewBytes(table, "end_time")
	t.IntervalSeconds = field.NewInt32(table, "interval_seconds")
	t.ConditionType = field.NewString(table, "condition_type")
	t.TCondition = field.NewBytes(table, "t_condition")

	t.fillFieldMap()

	return t
}

func (t *timeDatum) WithContext(ctx context.Context) *timeDatumDo {
	return t.timeDatumDo.WithContext(ctx)
}

func (t timeDatum) TableName() string { return t.timeDatumDo.TableName() }

func (t timeDatum) Alias() string { return t.timeDatumDo.Alias() }

func (t *timeDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeDatum) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.ID
	t.fieldMap["repeat_type"] = t.RepeatType
	t.fieldMap["start_date"] = t.StartDate
	t.fieldMap["end_date"] = t.EndDate
	t.fieldMap["start_time"] = t.StartTime
	t.fieldMap["end_time"] = t.EndTime
	t.fieldMap["interval_seconds"] = t.IntervalSeconds
	t.fieldMap["condition_type"] = t.ConditionType
	t.fieldMap["t_condition"] = t.TCondition
}

func (t timeDatum) clone(db *gorm.DB) timeDatum {
	t.timeDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeDatum) replaceDB(db *gorm.DB) timeDatum {
	t.timeDatumDo.ReplaceDB(db)
	return t
}

type timeDatumDo struct{ gen.DO }

func (t timeDatumDo) Debug() *timeDatumDo {
	return t.withDO(t.DO.Debug())
}

func (t timeDatumDo) WithContext(ctx context.Context) *timeDatumDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeDatumDo) ReadDB() *timeDatumDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeDatumDo) WriteDB() *timeDatumDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeDatumDo) Session(config *gorm.Session) *timeDatumDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeDatumDo) Clauses(conds ...clause.Expression) *timeDatumDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeDatumDo) Returning(value interface{}, columns ...string) *timeDatumDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeDatumDo) Not(conds ...gen.Condition) *timeDatumDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeDatumDo) Or(conds ...gen.Condition) *timeDatumDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeDatumDo) Select(conds ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeDatumDo) Where(conds ...gen.Condition) *timeDatumDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeDatumDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *timeDatumDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timeDatumDo) Order(conds ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeDatumDo) Distinct(cols ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeDatumDo) Omit(cols ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeDatumDo) Join(table schema.Tabler, on ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeDatumDo) Group(cols ...field.Expr) *timeDatumDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeDatumDo) Having(conds ...gen.Condition) *timeDatumDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeDatumDo) Limit(limit int) *timeDatumDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeDatumDo) Offset(offset int) *timeDatumDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timeDatumDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeDatumDo) Unscoped() *timeDatumDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeDatumDo) Create(values ...*model.TimeDatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeDatumDo) CreateInBatches(values []*model.TimeDatum, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeDatumDo) Save(values ...*model.TimeDatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeDatumDo) First() (*model.TimeDatum, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeDatum), nil
	}
}

func (t timeDatumDo) Take() (*model.TimeDatum, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeDatum), nil
	}
}

func (t timeDatumDo) Last() (*model.TimeDatum, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeDatum), nil
	}
}

func (t timeDatumDo) Find() ([]*model.TimeDatum, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimeDatum), err
}

func (t timeDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimeDatum, err error) {
	buf := make([]*model.TimeDatum, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeDatumDo) FindInBatches(result *[]*model.TimeDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeDatumDo) Attrs(attrs ...field.AssignExpr) *timeDatumDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeDatumDo) Assign(attrs ...field.AssignExpr) *timeDatumDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeDatumDo) Joins(fields ...field.RelationField) *timeDatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeDatumDo) Preload(fields ...field.RelationField) *timeDatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeDatumDo) FirstOrInit() (*model.TimeDatum, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeDatum), nil
	}
}

func (t timeDatumDo) FirstOrCreate() (*model.TimeDatum, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeDatum), nil
	}
}

func (t timeDatumDo) FindByPage(offset int, limit int) (result []*model.TimeDatum, count int64, err error) {
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

func (t timeDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeDatumDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeDatumDo) Delete(models ...*model.TimeDatum) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeDatumDo) withDO(do gen.Dao) *timeDatumDo {
	t.DO = *do.(*gen.DO)
	return t
}
