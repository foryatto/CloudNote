package service

import (
	"CloudNote/app/dao"
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/guuid"
)

var Category = categoryService{}

type categoryService struct{}

func (c *categoryService) Add(ownerId string, param *define.CategoryAddReq) error {
	if len(param.Title) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	saveDate := model.Category{
		CategoryId:  guuid.New().String(),
		Title:       param.Title,
		Description: param.Description,
		OwnerId:     ownerId,
	}
	_, err := dao.Category.Ctx(context.TODO()).Insert(saveDate)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryService) Delete(ownerId string, param *define.CategoryDeleteReq) error {
	if len(param.CategoryIds) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	for _, id := range param.CategoryIds {
		_, err := dao.Category.Ctx(context.TODO()).Where(g.Map{
			dao.Category.Columns.OwnerId:    ownerId,
			dao.Category.Columns.CategoryId: id,
		}).Delete()
		if err != nil {
			g.Log().Line().Warning(err)
			return err
		}
	}
	return nil
}

func (c *categoryService) QueryList(ownerId string, param *define.CategoryQueryReq) (interface{}, error) {
	sql := dao.Category.Ctx(context.TODO()).Where(g.Map{
		dao.Category.Columns.OwnerId: ownerId,
	})

	if param.Page >= 1 && param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	if param.CategoryId != "" {
		sql = sql.Where(dao.Category.Columns.CategoryId, param.CategoryId)
	}
	if param.Title != "" {
		sql = sql.WhereLike(dao.Category.Columns.Title, "%"+param.Title+"%")
	}
	if param.Description != "" {
		sql = sql.WhereLike(dao.Category.Columns.Description, "%"+param.Description+"%")
	}

	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}

	var result []*define.CategoryQueryResp
	err = sql.Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

func (c *categoryService) QueryById(id string) (*define.Category, error) {
	sql := dao.Category.Ctx(context.TODO()).Where(g.Map{
		dao.Category.Columns.CategoryId: id,
	})

	var result *define.Category
	err := sql.Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	return result, nil
}

func (c *categoryService) UpdateTitle(ownerId string, param *define.CategoryUpdateTitleReq) error {
	if len(param.Title) == 0 || len(param.CategoryId) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Category.Ctx(context.TODO()).Where(g.Map{
		dao.Category.Columns.CategoryId: param.CategoryId,
		dao.Category.Columns.OwnerId:    ownerId,
	}).Update(g.Map{
		dao.Category.Columns.Title: param.Title,
	})
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	updateResult, err := update.RowsAffected()
	if err != nil || updateResult == 0 {
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}

func (c *categoryService) UpdateDescription(ownerId string, param *define.CategoryUpdateDescriptionReq) error {
	if len(param.Description) == 0 || len(param.CategoryId) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Category.Ctx(context.TODO()).Where(g.Map{
		dao.Category.Columns.CategoryId: param.CategoryId,
		dao.Category.Columns.OwnerId:    ownerId,
	}).Update(g.Map{
		dao.Category.Columns.Description: param.Description,
	})
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	updateResult, err := update.RowsAffected()
	if err != nil || updateResult == 0 {
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}
