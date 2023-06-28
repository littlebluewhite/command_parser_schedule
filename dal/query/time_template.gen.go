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

func newTimeTemplate(db *gorm.DB, opts ...gen.DOOption) timeTemplate {
	_timeTemplate := timeTemplate{}

	_timeTemplate.timeTemplateDo.UseDB(db, opts...)
	_timeTemplate.timeTemplateDo.UseModel(&model.TimeTemplate{})

	tableName := _timeTemplate.timeTemplateDo.TableName()
	_timeTemplate.ALL = field.NewAsterisk(tableName)
	_timeTemplate.ID = field.NewInt32(tableName, "id")
	_timeTemplate.Name = field.NewString(tableName, "name")
	_timeTemplate.TimeDataID = field.NewInt32(tableName, "time_data_id")
	_timeTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_timeTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_timeTemplate.TimeData = timeTemplateBelongsToTimeData{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("TimeData", "model.TimeDatum"),
	}

	_timeTemplate.fillFieldMap()

	return _timeTemplate
}

type timeTemplate struct {
	timeTemplateDo timeTemplateDo

	ALL        field.Asterisk
	ID         field.Int32
	Name       field.String
	TimeDataID field.Int32
	UpdatedAt  field.Time
	CreatedAt  field.Time
	TimeData   timeTemplateBelongsToTimeData

	fieldMap map[string]field.Expr
}

func (t timeTemplate) Table(newTableName string) *timeTemplate {
	t.timeTemplateDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeTemplate) As(alias string) *timeTemplate {
	t.timeTemplateDo.DO = *(t.timeTemplateDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeTemplate) updateTableName(table string) *timeTemplate {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Name = field.NewString(table, "name")
	t.TimeDataID = field.NewInt32(table, "time_data_id")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *timeTemplate) WithContext(ctx context.Context) *timeTemplateDo {
	return t.timeTemplateDo.WithContext(ctx)
}

func (t timeTemplate) TableName() string { return t.timeTemplateDo.TableName() }

func (t timeTemplate) Alias() string { return t.timeTemplateDo.Alias() }

func (t *timeTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeTemplate) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
	t.fieldMap["time_data_id"] = t.TimeDataID
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["created_at"] = t.CreatedAt

}

func (t timeTemplate) clone(db *gorm.DB) timeTemplate {
	t.timeTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeTemplate) replaceDB(db *gorm.DB) timeTemplate {
	t.timeTemplateDo.ReplaceDB(db)
	return t
}

type timeTemplateBelongsToTimeData struct {
	db *gorm.DB

	field.RelationField
}

func (a timeTemplateBelongsToTimeData) Where(conds ...field.Expr) *timeTemplateBelongsToTimeData {
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

func (a timeTemplateBelongsToTimeData) WithContext(ctx context.Context) *timeTemplateBelongsToTimeData {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a timeTemplateBelongsToTimeData) Session(session *gorm.Session) *timeTemplateBelongsToTimeData {
	a.db = a.db.Session(session)
	return &a
}

func (a timeTemplateBelongsToTimeData) Model(m *model.TimeTemplate) *timeTemplateBelongsToTimeDataTx {
	return &timeTemplateBelongsToTimeDataTx{a.db.Model(m).Association(a.Name())}
}

type timeTemplateBelongsToTimeDataTx struct{ tx *gorm.Association }

func (a timeTemplateBelongsToTimeDataTx) Find() (result *model.TimeDatum, err error) {
	return result, a.tx.Find(&result)
}

func (a timeTemplateBelongsToTimeDataTx) Append(values ...*model.TimeDatum) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a timeTemplateBelongsToTimeDataTx) Replace(values ...*model.TimeDatum) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a timeTemplateBelongsToTimeDataTx) Delete(values ...*model.TimeDatum) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a timeTemplateBelongsToTimeDataTx) Clear() error {
	return a.tx.Clear()
}

func (a timeTemplateBelongsToTimeDataTx) Count() int64 {
	return a.tx.Count()
}

type timeTemplateDo struct{ gen.DO }

func (t timeTemplateDo) Debug() *timeTemplateDo {
	return t.withDO(t.DO.Debug())
}

func (t timeTemplateDo) WithContext(ctx context.Context) *timeTemplateDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeTemplateDo) ReadDB() *timeTemplateDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeTemplateDo) WriteDB() *timeTemplateDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeTemplateDo) Session(config *gorm.Session) *timeTemplateDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeTemplateDo) Clauses(conds ...clause.Expression) *timeTemplateDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeTemplateDo) Returning(value interface{}, columns ...string) *timeTemplateDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeTemplateDo) Not(conds ...gen.Condition) *timeTemplateDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeTemplateDo) Or(conds ...gen.Condition) *timeTemplateDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeTemplateDo) Select(conds ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeTemplateDo) Where(conds ...gen.Condition) *timeTemplateDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *timeTemplateDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timeTemplateDo) Order(conds ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeTemplateDo) Distinct(cols ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeTemplateDo) Omit(cols ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeTemplateDo) Join(table schema.Tabler, on ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeTemplateDo) Group(cols ...field.Expr) *timeTemplateDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeTemplateDo) Having(conds ...gen.Condition) *timeTemplateDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeTemplateDo) Limit(limit int) *timeTemplateDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeTemplateDo) Offset(offset int) *timeTemplateDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timeTemplateDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeTemplateDo) Unscoped() *timeTemplateDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeTemplateDo) Create(values ...*model.TimeTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeTemplateDo) CreateInBatches(values []*model.TimeTemplate, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeTemplateDo) Save(values ...*model.TimeTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeTemplateDo) First() (*model.TimeTemplate, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTemplate), nil
	}
}

func (t timeTemplateDo) Take() (*model.TimeTemplate, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTemplate), nil
	}
}

func (t timeTemplateDo) Last() (*model.TimeTemplate, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTemplate), nil
	}
}

func (t timeTemplateDo) Find() ([]*model.TimeTemplate, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimeTemplate), err
}

func (t timeTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimeTemplate, err error) {
	buf := make([]*model.TimeTemplate, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeTemplateDo) FindInBatches(result *[]*model.TimeTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeTemplateDo) Attrs(attrs ...field.AssignExpr) *timeTemplateDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeTemplateDo) Assign(attrs ...field.AssignExpr) *timeTemplateDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeTemplateDo) Joins(fields ...field.RelationField) *timeTemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeTemplateDo) Preload(fields ...field.RelationField) *timeTemplateDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeTemplateDo) FirstOrInit() (*model.TimeTemplate, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTemplate), nil
	}
}

func (t timeTemplateDo) FirstOrCreate() (*model.TimeTemplate, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTemplate), nil
	}
}

func (t timeTemplateDo) FindByPage(offset int, limit int) (result []*model.TimeTemplate, count int64, err error) {
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

func (t timeTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeTemplateDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeTemplateDo) Delete(models ...*model.TimeTemplate) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeTemplateDo) withDO(do gen.Dao) *timeTemplateDo {
	t.DO = *do.(*gen.DO)
	return t
}
