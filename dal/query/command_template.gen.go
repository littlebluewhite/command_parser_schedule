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

func newCommandTemplate(db *gorm.DB, opts ...gen.DOOption) commandTemplate {
	_commandTemplate := commandTemplate{}

	_commandTemplate.commandTemplateDo.UseDB(db, opts...)
	_commandTemplate.commandTemplateDo.UseModel(&model.CommandTemplate{})

	tableName := _commandTemplate.commandTemplateDo.TableName()
	_commandTemplate.ALL = field.NewAsterisk(tableName)
	_commandTemplate.ID = field.NewInt32(tableName, "id")
	_commandTemplate.Name = field.NewString(tableName, "name")
	_commandTemplate.Protocol = field.NewString(tableName, "protocol")
	_commandTemplate.Description = field.NewString(tableName, "description")
	_commandTemplate.Host = field.NewString(tableName, "host")
	_commandTemplate.Port = field.NewString(tableName, "port")
	_commandTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_commandTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_commandTemplate.Http = commandTemplateHasOneHttp{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Http", "model.HTTPSCommand"),
	}

	_commandTemplate.Mqtt = commandTemplateHasOneMqtt{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Mqtt", "model.MqttCommand"),
	}

	_commandTemplate.Websocket = commandTemplateHasOneWebsocket{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Websocket", "model.WebsocketCommand"),
	}

	_commandTemplate.Redis = commandTemplateHasOneRedis{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Redis", "model.RedisCommand"),
	}

	_commandTemplate.Monitor = commandTemplateHasOneMonitor{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Monitor", "model.Monitor"),
		MConditions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Monitor.MConditions", "model.MCondition"),
		},
	}

	_commandTemplate.fillFieldMap()

	return _commandTemplate
}

type commandTemplate struct {
	commandTemplateDo commandTemplateDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	Protocol    field.String
	Description field.String
	Host        field.String
	Port        field.String
	UpdatedAt   field.Time
	CreatedAt   field.Time
	Http        commandTemplateHasOneHttp

	Mqtt commandTemplateHasOneMqtt

	Websocket commandTemplateHasOneWebsocket

	Redis commandTemplateHasOneRedis

	Monitor commandTemplateHasOneMonitor

	fieldMap map[string]field.Expr
}

func (c commandTemplate) Table(newTableName string) *commandTemplate {
	c.commandTemplateDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commandTemplate) As(alias string) *commandTemplate {
	c.commandTemplateDo.DO = *(c.commandTemplateDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commandTemplate) updateTableName(table string) *commandTemplate {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Name = field.NewString(table, "name")
	c.Protocol = field.NewString(table, "protocol")
	c.Description = field.NewString(table, "description")
	c.Host = field.NewString(table, "host")
	c.Port = field.NewString(table, "port")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.CreatedAt = field.NewTime(table, "created_at")

	c.fillFieldMap()

	return c
}

func (c *commandTemplate) WithContext(ctx context.Context) *commandTemplateDo {
	return c.commandTemplateDo.WithContext(ctx)
}

func (c commandTemplate) TableName() string { return c.commandTemplateDo.TableName() }

func (c commandTemplate) Alias() string { return c.commandTemplateDo.Alias() }

func (c *commandTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commandTemplate) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 13)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["protocol"] = c.Protocol
	c.fieldMap["description"] = c.Description
	c.fieldMap["host"] = c.Host
	c.fieldMap["port"] = c.Port
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["created_at"] = c.CreatedAt

}

func (c commandTemplate) clone(db *gorm.DB) commandTemplate {
	c.commandTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commandTemplate) replaceDB(db *gorm.DB) commandTemplate {
	c.commandTemplateDo.ReplaceDB(db)
	return c
}

type commandTemplateHasOneHttp struct {
	db *gorm.DB

	field.RelationField
}

func (a commandTemplateHasOneHttp) Where(conds ...field.Expr) *commandTemplateHasOneHttp {
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

func (a commandTemplateHasOneHttp) WithContext(ctx context.Context) *commandTemplateHasOneHttp {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a commandTemplateHasOneHttp) Session(session *gorm.Session) *commandTemplateHasOneHttp {
	a.db = a.db.Session(session)
	return &a
}

func (a commandTemplateHasOneHttp) Model(m *model.CommandTemplate) *commandTemplateHasOneHttpTx {
	return &commandTemplateHasOneHttpTx{a.db.Model(m).Association(a.Name())}
}

type commandTemplateHasOneHttpTx struct{ tx *gorm.Association }

func (a commandTemplateHasOneHttpTx) Find() (result *model.HTTPSCommand, err error) {
	return result, a.tx.Find(&result)
}

func (a commandTemplateHasOneHttpTx) Append(values ...*model.HTTPSCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a commandTemplateHasOneHttpTx) Replace(values ...*model.HTTPSCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a commandTemplateHasOneHttpTx) Delete(values ...*model.HTTPSCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a commandTemplateHasOneHttpTx) Clear() error {
	return a.tx.Clear()
}

func (a commandTemplateHasOneHttpTx) Count() int64 {
	return a.tx.Count()
}

type commandTemplateHasOneMqtt struct {
	db *gorm.DB

	field.RelationField
}

func (a commandTemplateHasOneMqtt) Where(conds ...field.Expr) *commandTemplateHasOneMqtt {
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

func (a commandTemplateHasOneMqtt) WithContext(ctx context.Context) *commandTemplateHasOneMqtt {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a commandTemplateHasOneMqtt) Session(session *gorm.Session) *commandTemplateHasOneMqtt {
	a.db = a.db.Session(session)
	return &a
}

func (a commandTemplateHasOneMqtt) Model(m *model.CommandTemplate) *commandTemplateHasOneMqttTx {
	return &commandTemplateHasOneMqttTx{a.db.Model(m).Association(a.Name())}
}

type commandTemplateHasOneMqttTx struct{ tx *gorm.Association }

func (a commandTemplateHasOneMqttTx) Find() (result *model.MqttCommand, err error) {
	return result, a.tx.Find(&result)
}

func (a commandTemplateHasOneMqttTx) Append(values ...*model.MqttCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a commandTemplateHasOneMqttTx) Replace(values ...*model.MqttCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a commandTemplateHasOneMqttTx) Delete(values ...*model.MqttCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a commandTemplateHasOneMqttTx) Clear() error {
	return a.tx.Clear()
}

func (a commandTemplateHasOneMqttTx) Count() int64 {
	return a.tx.Count()
}

type commandTemplateHasOneWebsocket struct {
	db *gorm.DB

	field.RelationField
}

func (a commandTemplateHasOneWebsocket) Where(conds ...field.Expr) *commandTemplateHasOneWebsocket {
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

func (a commandTemplateHasOneWebsocket) WithContext(ctx context.Context) *commandTemplateHasOneWebsocket {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a commandTemplateHasOneWebsocket) Session(session *gorm.Session) *commandTemplateHasOneWebsocket {
	a.db = a.db.Session(session)
	return &a
}

func (a commandTemplateHasOneWebsocket) Model(m *model.CommandTemplate) *commandTemplateHasOneWebsocketTx {
	return &commandTemplateHasOneWebsocketTx{a.db.Model(m).Association(a.Name())}
}

type commandTemplateHasOneWebsocketTx struct{ tx *gorm.Association }

func (a commandTemplateHasOneWebsocketTx) Find() (result *model.WebsocketCommand, err error) {
	return result, a.tx.Find(&result)
}

func (a commandTemplateHasOneWebsocketTx) Append(values ...*model.WebsocketCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a commandTemplateHasOneWebsocketTx) Replace(values ...*model.WebsocketCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a commandTemplateHasOneWebsocketTx) Delete(values ...*model.WebsocketCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a commandTemplateHasOneWebsocketTx) Clear() error {
	return a.tx.Clear()
}

func (a commandTemplateHasOneWebsocketTx) Count() int64 {
	return a.tx.Count()
}

type commandTemplateHasOneRedis struct {
	db *gorm.DB

	field.RelationField
}

func (a commandTemplateHasOneRedis) Where(conds ...field.Expr) *commandTemplateHasOneRedis {
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

func (a commandTemplateHasOneRedis) WithContext(ctx context.Context) *commandTemplateHasOneRedis {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a commandTemplateHasOneRedis) Session(session *gorm.Session) *commandTemplateHasOneRedis {
	a.db = a.db.Session(session)
	return &a
}

func (a commandTemplateHasOneRedis) Model(m *model.CommandTemplate) *commandTemplateHasOneRedisTx {
	return &commandTemplateHasOneRedisTx{a.db.Model(m).Association(a.Name())}
}

type commandTemplateHasOneRedisTx struct{ tx *gorm.Association }

func (a commandTemplateHasOneRedisTx) Find() (result *model.RedisCommand, err error) {
	return result, a.tx.Find(&result)
}

func (a commandTemplateHasOneRedisTx) Append(values ...*model.RedisCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a commandTemplateHasOneRedisTx) Replace(values ...*model.RedisCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a commandTemplateHasOneRedisTx) Delete(values ...*model.RedisCommand) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a commandTemplateHasOneRedisTx) Clear() error {
	return a.tx.Clear()
}

func (a commandTemplateHasOneRedisTx) Count() int64 {
	return a.tx.Count()
}

type commandTemplateHasOneMonitor struct {
	db *gorm.DB

	field.RelationField

	MConditions struct {
		field.RelationField
	}
}

func (a commandTemplateHasOneMonitor) Where(conds ...field.Expr) *commandTemplateHasOneMonitor {
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

func (a commandTemplateHasOneMonitor) WithContext(ctx context.Context) *commandTemplateHasOneMonitor {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a commandTemplateHasOneMonitor) Session(session *gorm.Session) *commandTemplateHasOneMonitor {
	a.db = a.db.Session(session)
	return &a
}

func (a commandTemplateHasOneMonitor) Model(m *model.CommandTemplate) *commandTemplateHasOneMonitorTx {
	return &commandTemplateHasOneMonitorTx{a.db.Model(m).Association(a.Name())}
}

type commandTemplateHasOneMonitorTx struct{ tx *gorm.Association }

func (a commandTemplateHasOneMonitorTx) Find() (result *model.Monitor, err error) {
	return result, a.tx.Find(&result)
}

func (a commandTemplateHasOneMonitorTx) Append(values ...*model.Monitor) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a commandTemplateHasOneMonitorTx) Replace(values ...*model.Monitor) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a commandTemplateHasOneMonitorTx) Delete(values ...*model.Monitor) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a commandTemplateHasOneMonitorTx) Clear() error {
	return a.tx.Clear()
}

func (a commandTemplateHasOneMonitorTx) Count() int64 {
	return a.tx.Count()
}

type commandTemplateDo struct{ gen.DO }

func (c commandTemplateDo) Debug() *commandTemplateDo {
	return c.withDO(c.DO.Debug())
}

func (c commandTemplateDo) WithContext(ctx context.Context) *commandTemplateDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commandTemplateDo) ReadDB() *commandTemplateDo {
	return c.Clauses(dbresolver.Read)
}

func (c commandTemplateDo) WriteDB() *commandTemplateDo {
	return c.Clauses(dbresolver.Write)
}

func (c commandTemplateDo) Session(config *gorm.Session) *commandTemplateDo {
	return c.withDO(c.DO.Session(config))
}

func (c commandTemplateDo) Clauses(conds ...clause.Expression) *commandTemplateDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commandTemplateDo) Returning(value interface{}, columns ...string) *commandTemplateDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commandTemplateDo) Not(conds ...gen.Condition) *commandTemplateDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commandTemplateDo) Or(conds ...gen.Condition) *commandTemplateDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commandTemplateDo) Select(conds ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commandTemplateDo) Where(conds ...gen.Condition) *commandTemplateDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commandTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *commandTemplateDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commandTemplateDo) Order(conds ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commandTemplateDo) Distinct(cols ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commandTemplateDo) Omit(cols ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commandTemplateDo) Join(table schema.Tabler, on ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commandTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commandTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commandTemplateDo) Group(cols ...field.Expr) *commandTemplateDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commandTemplateDo) Having(conds ...gen.Condition) *commandTemplateDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commandTemplateDo) Limit(limit int) *commandTemplateDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commandTemplateDo) Offset(offset int) *commandTemplateDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commandTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commandTemplateDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commandTemplateDo) Unscoped() *commandTemplateDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commandTemplateDo) Create(values ...*model.CommandTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commandTemplateDo) CreateInBatches(values []*model.CommandTemplate, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commandTemplateDo) Save(values ...*model.CommandTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commandTemplateDo) First() (*model.CommandTemplate, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandTemplate), nil
	}
}

func (c commandTemplateDo) Take() (*model.CommandTemplate, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandTemplate), nil
	}
}

func (c commandTemplateDo) Last() (*model.CommandTemplate, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandTemplate), nil
	}
}

func (c commandTemplateDo) Find() ([]*model.CommandTemplate, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommandTemplate), err
}

func (c commandTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommandTemplate, err error) {
	buf := make([]*model.CommandTemplate, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commandTemplateDo) FindInBatches(result *[]*model.CommandTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commandTemplateDo) Attrs(attrs ...field.AssignExpr) *commandTemplateDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commandTemplateDo) Assign(attrs ...field.AssignExpr) *commandTemplateDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commandTemplateDo) Joins(fields ...field.RelationField) *commandTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commandTemplateDo) Preload(fields ...field.RelationField) *commandTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commandTemplateDo) FirstOrInit() (*model.CommandTemplate, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandTemplate), nil
	}
}

func (c commandTemplateDo) FirstOrCreate() (*model.CommandTemplate, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommandTemplate), nil
	}
}

func (c commandTemplateDo) FindByPage(offset int, limit int) (result []*model.CommandTemplate, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commandTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commandTemplateDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commandTemplateDo) Delete(models ...*model.CommandTemplate) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commandTemplateDo) withDO(do gen.Dao) *commandTemplateDo {
	c.DO = *do.(*gen.DO)
	return c
}
