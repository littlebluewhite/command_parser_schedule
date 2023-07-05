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

func newHTTPSCommand(db *gorm.DB, opts ...gen.DOOption) hTTPSCommand {
	_hTTPSCommand := hTTPSCommand{}

	_hTTPSCommand.hTTPSCommandDo.UseDB(db, opts...)
	_hTTPSCommand.hTTPSCommandDo.UseModel(&model.HTTPSCommand{})

	tableName := _hTTPSCommand.hTTPSCommandDo.TableName()
	_hTTPSCommand.ALL = field.NewAsterisk(tableName)
	_hTTPSCommand.ID = field.NewInt32(tableName, "id")
	_hTTPSCommand.CommandID = field.NewInt32(tableName, "command_id")
	_hTTPSCommand.Method = field.NewString(tableName, "method")
	_hTTPSCommand.URL = field.NewString(tableName, "url")
	_hTTPSCommand.AuthorizationType = field.NewString(tableName, "authorization_type")
	_hTTPSCommand.Params = field.NewBytes(tableName, "params")
	_hTTPSCommand.Header = field.NewBytes(tableName, "header")
	_hTTPSCommand.BodyType = field.NewString(tableName, "body_type")
	_hTTPSCommand.Body = field.NewBytes(tableName, "body")

	_hTTPSCommand.fillFieldMap()

	return _hTTPSCommand
}

type hTTPSCommand struct {
	hTTPSCommandDo hTTPSCommandDo

	ALL               field.Asterisk
	ID                field.Int32
	CommandID         field.Int32
	Method            field.String
	URL               field.String
	AuthorizationType field.String
	Params            field.Bytes
	Header            field.Bytes
	BodyType          field.String
	Body              field.Bytes

	fieldMap map[string]field.Expr
}

func (h hTTPSCommand) Table(newTableName string) *hTTPSCommand {
	h.hTTPSCommandDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hTTPSCommand) As(alias string) *hTTPSCommand {
	h.hTTPSCommandDo.DO = *(h.hTTPSCommandDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hTTPSCommand) updateTableName(table string) *hTTPSCommand {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt32(table, "id")
	h.CommandID = field.NewInt32(table, "command_id")
	h.Method = field.NewString(table, "method")
	h.URL = field.NewString(table, "url")
	h.AuthorizationType = field.NewString(table, "authorization_type")
	h.Params = field.NewBytes(table, "params")
	h.Header = field.NewBytes(table, "header")
	h.BodyType = field.NewString(table, "body_type")
	h.Body = field.NewBytes(table, "body")

	h.fillFieldMap()

	return h
}

func (h *hTTPSCommand) WithContext(ctx context.Context) *hTTPSCommandDo {
	return h.hTTPSCommandDo.WithContext(ctx)
}

func (h hTTPSCommand) TableName() string { return h.hTTPSCommandDo.TableName() }

func (h hTTPSCommand) Alias() string { return h.hTTPSCommandDo.Alias() }

func (h *hTTPSCommand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hTTPSCommand) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 9)
	h.fieldMap["id"] = h.ID
	h.fieldMap["command_id"] = h.CommandID
	h.fieldMap["method"] = h.Method
	h.fieldMap["url"] = h.URL
	h.fieldMap["authorization_type"] = h.AuthorizationType
	h.fieldMap["params"] = h.Params
	h.fieldMap["header"] = h.Header
	h.fieldMap["body_type"] = h.BodyType
	h.fieldMap["body"] = h.Body
}

func (h hTTPSCommand) clone(db *gorm.DB) hTTPSCommand {
	h.hTTPSCommandDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hTTPSCommand) replaceDB(db *gorm.DB) hTTPSCommand {
	h.hTTPSCommandDo.ReplaceDB(db)
	return h
}

type hTTPSCommandDo struct{ gen.DO }

func (h hTTPSCommandDo) Debug() *hTTPSCommandDo {
	return h.withDO(h.DO.Debug())
}

func (h hTTPSCommandDo) WithContext(ctx context.Context) *hTTPSCommandDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hTTPSCommandDo) ReadDB() *hTTPSCommandDo {
	return h.Clauses(dbresolver.Read)
}

func (h hTTPSCommandDo) WriteDB() *hTTPSCommandDo {
	return h.Clauses(dbresolver.Write)
}

func (h hTTPSCommandDo) Session(config *gorm.Session) *hTTPSCommandDo {
	return h.withDO(h.DO.Session(config))
}

func (h hTTPSCommandDo) Clauses(conds ...clause.Expression) *hTTPSCommandDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hTTPSCommandDo) Returning(value interface{}, columns ...string) *hTTPSCommandDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hTTPSCommandDo) Not(conds ...gen.Condition) *hTTPSCommandDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hTTPSCommandDo) Or(conds ...gen.Condition) *hTTPSCommandDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hTTPSCommandDo) Select(conds ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hTTPSCommandDo) Where(conds ...gen.Condition) *hTTPSCommandDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hTTPSCommandDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *hTTPSCommandDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h hTTPSCommandDo) Order(conds ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hTTPSCommandDo) Distinct(cols ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hTTPSCommandDo) Omit(cols ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hTTPSCommandDo) Join(table schema.Tabler, on ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hTTPSCommandDo) LeftJoin(table schema.Tabler, on ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hTTPSCommandDo) RightJoin(table schema.Tabler, on ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hTTPSCommandDo) Group(cols ...field.Expr) *hTTPSCommandDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hTTPSCommandDo) Having(conds ...gen.Condition) *hTTPSCommandDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hTTPSCommandDo) Limit(limit int) *hTTPSCommandDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hTTPSCommandDo) Offset(offset int) *hTTPSCommandDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hTTPSCommandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *hTTPSCommandDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hTTPSCommandDo) Unscoped() *hTTPSCommandDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hTTPSCommandDo) Create(values ...*model.HTTPSCommand) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hTTPSCommandDo) CreateInBatches(values []*model.HTTPSCommand, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hTTPSCommandDo) Save(values ...*model.HTTPSCommand) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hTTPSCommandDo) First() (*model.HTTPSCommand, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HTTPSCommand), nil
	}
}

func (h hTTPSCommandDo) Take() (*model.HTTPSCommand, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HTTPSCommand), nil
	}
}

func (h hTTPSCommandDo) Last() (*model.HTTPSCommand, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HTTPSCommand), nil
	}
}

func (h hTTPSCommandDo) Find() ([]*model.HTTPSCommand, error) {
	result, err := h.DO.Find()
	return result.([]*model.HTTPSCommand), err
}

func (h hTTPSCommandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HTTPSCommand, err error) {
	buf := make([]*model.HTTPSCommand, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hTTPSCommandDo) FindInBatches(result *[]*model.HTTPSCommand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hTTPSCommandDo) Attrs(attrs ...field.AssignExpr) *hTTPSCommandDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hTTPSCommandDo) Assign(attrs ...field.AssignExpr) *hTTPSCommandDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hTTPSCommandDo) Joins(fields ...field.RelationField) *hTTPSCommandDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hTTPSCommandDo) Preload(fields ...field.RelationField) *hTTPSCommandDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hTTPSCommandDo) FirstOrInit() (*model.HTTPSCommand, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HTTPSCommand), nil
	}
}

func (h hTTPSCommandDo) FirstOrCreate() (*model.HTTPSCommand, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HTTPSCommand), nil
	}
}

func (h hTTPSCommandDo) FindByPage(offset int, limit int) (result []*model.HTTPSCommand, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hTTPSCommandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hTTPSCommandDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hTTPSCommandDo) Delete(models ...*model.HTTPSCommand) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hTTPSCommandDo) withDO(do gen.Dao) *hTTPSCommandDo {
	h.DO = *do.(*gen.DO)
	return h
}
