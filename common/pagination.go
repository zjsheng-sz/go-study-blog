package common

import "gorm.io/gorm"

type Pagination struct {
	Page     int `form:"page" json:"page"`         // 当前页码
	PageSize int `form:"pageSize" json:"pageSize"` // 每页数量
}

type PaginationResult struct {
	List      interface{} `json:"list"`      // 数据列表
	Total     int64       `json:"total"`     // 总记录数
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页数量
	TotalPage int         `json:"totalPage"` // 总页数
}

// Paginate GORM 分页作用域
func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 设置默认值
		if p.Page <= 0 {
			p.Page = 1
		}

		// 限制每页最大数量
		if p.PageSize > 100 {
			p.PageSize = 100
		} else if p.PageSize <= 0 {
			p.PageSize = 10
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
