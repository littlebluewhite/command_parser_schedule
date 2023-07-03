// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		HeaderTemplate: newHeaderTemplate(db, opts...),
		TimeDatum:      newTimeDatum(db, opts...),
		TimeTemplate:   newTimeTemplate(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	HeaderTemplate headerTemplate
	TimeDatum      timeDatum
	TimeTemplate   timeTemplate
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		HeaderTemplate: q.HeaderTemplate.clone(db),
		TimeDatum:      q.TimeDatum.clone(db),
		TimeTemplate:   q.TimeTemplate.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		HeaderTemplate: q.HeaderTemplate.replaceDB(db),
		TimeDatum:      q.TimeDatum.replaceDB(db),
		TimeTemplate:   q.TimeTemplate.replaceDB(db),
	}
}

type queryCtx struct {
	HeaderTemplate *headerTemplateDo
	TimeDatum      *timeDatumDo
	TimeTemplate   *timeTemplateDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		HeaderTemplate: q.HeaderTemplate.WithContext(ctx),
		TimeDatum:      q.TimeDatum.WithContext(ctx),
		TimeTemplate:   q.TimeTemplate.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
