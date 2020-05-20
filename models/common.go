package models

import "gomall-center/pkg/web"

//PagerArgs common pager args
type PagerArgs struct {
	PageSize    int
	PageNum     int
	OrderField  string
	OrderMethod string
	KeyWord     string
	Filter      map[string]interface{}
}

//Bind Pager bind
func (pager *PagerArgs) Bind(ctx *web.Context) error {
	if err := ctx.ShouldBindQuery(pager); err != nil {
		return err
	}
	if pager.PageSize == 0 {
		pager.PageSize = 10
	}
	if pager.PageNum == 0 {
		pager.PageNum = 1
	}
	return nil
}

//PagerData ...
type PagerData struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}
