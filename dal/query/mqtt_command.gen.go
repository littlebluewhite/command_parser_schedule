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

func newMqttCommand(db *gorm.DB, opts ...gen.DOOption) mqttCommand {
	_mqttCommand := mqttCommand{}

	_mqttCommand.mqttCommandDo.UseDB(db, opts...)
	_mqttCommand.mqttCommandDo.UseModel(&model.MqttCommand{})

	tableName := _mqttCommand.mqttCommandDo.TableName()
	_mqttCommand.ALL = field.NewAsterisk(tableName)
	_mqttCommand.ID = field.NewInt32(tableName, "id")
	_mqttCommand.CommandID = field.NewInt32(tableName, "command_id")
	_mqttCommand.Topic = field.NewString(tableName, "topic")
	_mqttCommand.Header = field.NewString(tableName, "header")
	_mqttCommand.Message = field.NewString(tableName, "message")
	_mqttCommand.Type = field.NewString(tableName, "type")

	_mqttCommand.fillFieldMap()

	return _mqttCommand
}

type mqttCommand struct {
	mqttCommandDo mqttCommandDo

	ALL       field.Asterisk
	ID        field.Int32
	CommandID field.Int32
	Topic     field.String
	Header    field.String
	Message   field.String
	Type      field.String

	fieldMap map[string]field.Expr
}

func (m mqttCommand) Table(newTableName string) *mqttCommand {
	m.mqttCommandDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mqttCommand) As(alias string) *mqttCommand {
	m.mqttCommandDo.DO = *(m.mqttCommandDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mqttCommand) updateTableName(table string) *mqttCommand {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.CommandID = field.NewInt32(table, "command_id")
	m.Topic = field.NewString(table, "topic")
	m.Header = field.NewString(table, "header")
	m.Message = field.NewString(table, "message")
	m.Type = field.NewString(table, "type")

	m.fillFieldMap()

	return m
}

func (m *mqttCommand) WithContext(ctx context.Context) *mqttCommandDo {
	return m.mqttCommandDo.WithContext(ctx)
}

func (m mqttCommand) TableName() string { return m.mqttCommandDo.TableName() }

func (m mqttCommand) Alias() string { return m.mqttCommandDo.Alias() }

func (m *mqttCommand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mqttCommand) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["command_id"] = m.CommandID
	m.fieldMap["topic"] = m.Topic
	m.fieldMap["header"] = m.Header
	m.fieldMap["message"] = m.Message
	m.fieldMap["type"] = m.Type
}

func (m mqttCommand) clone(db *gorm.DB) mqttCommand {
	m.mqttCommandDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mqttCommand) replaceDB(db *gorm.DB) mqttCommand {
	m.mqttCommandDo.ReplaceDB(db)
	return m
}

type mqttCommandDo struct{ gen.DO }

func (m mqttCommandDo) Debug() *mqttCommandDo {
	return m.withDO(m.DO.Debug())
}

func (m mqttCommandDo) WithContext(ctx context.Context) *mqttCommandDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mqttCommandDo) ReadDB() *mqttCommandDo {
	return m.Clauses(dbresolver.Read)
}

func (m mqttCommandDo) WriteDB() *mqttCommandDo {
	return m.Clauses(dbresolver.Write)
}

func (m mqttCommandDo) Session(config *gorm.Session) *mqttCommandDo {
	return m.withDO(m.DO.Session(config))
}

func (m mqttCommandDo) Clauses(conds ...clause.Expression) *mqttCommandDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mqttCommandDo) Returning(value interface{}, columns ...string) *mqttCommandDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mqttCommandDo) Not(conds ...gen.Condition) *mqttCommandDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mqttCommandDo) Or(conds ...gen.Condition) *mqttCommandDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mqttCommandDo) Select(conds ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mqttCommandDo) Where(conds ...gen.Condition) *mqttCommandDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mqttCommandDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *mqttCommandDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m mqttCommandDo) Order(conds ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mqttCommandDo) Distinct(cols ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mqttCommandDo) Omit(cols ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mqttCommandDo) Join(table schema.Tabler, on ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mqttCommandDo) LeftJoin(table schema.Tabler, on ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mqttCommandDo) RightJoin(table schema.Tabler, on ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mqttCommandDo) Group(cols ...field.Expr) *mqttCommandDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mqttCommandDo) Having(conds ...gen.Condition) *mqttCommandDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mqttCommandDo) Limit(limit int) *mqttCommandDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mqttCommandDo) Offset(offset int) *mqttCommandDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mqttCommandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *mqttCommandDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mqttCommandDo) Unscoped() *mqttCommandDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mqttCommandDo) Create(values ...*model.MqttCommand) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mqttCommandDo) CreateInBatches(values []*model.MqttCommand, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mqttCommandDo) Save(values ...*model.MqttCommand) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mqttCommandDo) First() (*model.MqttCommand, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MqttCommand), nil
	}
}

func (m mqttCommandDo) Take() (*model.MqttCommand, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MqttCommand), nil
	}
}

func (m mqttCommandDo) Last() (*model.MqttCommand, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MqttCommand), nil
	}
}

func (m mqttCommandDo) Find() ([]*model.MqttCommand, error) {
	result, err := m.DO.Find()
	return result.([]*model.MqttCommand), err
}

func (m mqttCommandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MqttCommand, err error) {
	buf := make([]*model.MqttCommand, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mqttCommandDo) FindInBatches(result *[]*model.MqttCommand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mqttCommandDo) Attrs(attrs ...field.AssignExpr) *mqttCommandDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mqttCommandDo) Assign(attrs ...field.AssignExpr) *mqttCommandDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mqttCommandDo) Joins(fields ...field.RelationField) *mqttCommandDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mqttCommandDo) Preload(fields ...field.RelationField) *mqttCommandDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mqttCommandDo) FirstOrInit() (*model.MqttCommand, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MqttCommand), nil
	}
}

func (m mqttCommandDo) FirstOrCreate() (*model.MqttCommand, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MqttCommand), nil
	}
}

func (m mqttCommandDo) FindByPage(offset int, limit int) (result []*model.MqttCommand, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mqttCommandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mqttCommandDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mqttCommandDo) Delete(models ...*model.MqttCommand) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mqttCommandDo) withDO(do gen.Dao) *mqttCommandDo {
	m.DO = *do.(*gen.DO)
	return m
}
