// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2KnowledgeDao is the data access object for table v2_knowledge.
type V2KnowledgeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns V2KnowledgeColumns // columns contains all the column names of Table for convenient usage.
}

// V2KnowledgeColumns defines and stores column names for table v2_knowledge.
type V2KnowledgeColumns struct {
	Id        string //
	Category  string // 分類名
	Title     string // 標題
	Body      string // 內容
	OrderId   string // 排序
	Show      string // 顯示
	CreatedAt string //
	UpdatedAt string //
}

// v2KnowledgeColumns holds the columns for table v2_knowledge.
var v2KnowledgeColumns = V2KnowledgeColumns{
	Id:        "id",
	Category:  "category",
	Title:     "title",
	Body:      "body",
	OrderId:   "order_id",
	Show:      "show",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewV2KnowledgeDao creates and returns a new DAO object for table data access.
func NewV2KnowledgeDao() *V2KnowledgeDao {
	return &V2KnowledgeDao{
		group:   "default",
		table:   "v2_knowledge",
		columns: v2KnowledgeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2KnowledgeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2KnowledgeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2KnowledgeDao) Columns() V2KnowledgeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2KnowledgeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2KnowledgeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2KnowledgeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
