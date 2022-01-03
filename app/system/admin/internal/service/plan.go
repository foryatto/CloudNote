package service

import (
	"CloudNote/app/dao"
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
)

var Plan = planService{}

type planService struct{}

func (p *planService) Add(ownerId string, param *define.PlanAddReq) error {
	if param.Content == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	if param.Title == "" {
		endIndex := len(param.Content) % 10
		if endIndex == 0 {
			endIndex = 10
		}
		param.Title = param.Content[:endIndex] + "..."
	}
	var saveData *model.Plan
	err := gconv.Scan(param, &saveData)
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	saveData.OwnerId = ownerId
	saveData.PlanId = guuid.New().String()
	_, err = dao.Plan.Ctx(context.TODO()).Insert(saveData)
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	return nil
}

func (p *planService) QueryList(ownerId string, param *define.PlanQueryReq) (interface{}, error) {
	sql := dao.Plan.Ctx(context.TODO()).Where(g.Map{
		dao.Plan.Columns.OwnerId:   ownerId,
		dao.Plan.Columns.Completed: param.Completed,
	})
	if param.Page >= 1 && param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	if param.PlanId != "" {
		sql = sql.Where(dao.Plan.Columns.PlanId, param.PlanId)
	}
	if param.Title != "" {
		sql = sql.WhereLike(dao.Plan.Columns.Title, "%"+param.Title+"%")
	}
	if param.Content != "" {
		sql = sql.WhereLike(dao.Plan.Columns.Content, "%"+param.Content+"%")
	}

	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}

	var result []*define.PlanQueryResp
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

func (p *planService) UpdateTitle(ownerId string, param *define.PlanUpdateTitleReq) error {
	if param.Title == "" || param.PlanId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Plan.Ctx(context.TODO()).Where(g.Map{
		dao.Plan.Columns.OwnerId: ownerId,
		dao.Plan.Columns.PlanId:  param.PlanId,
	}).Update(g.Map{
		dao.Plan.Columns.Title: param.Title,
	})
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	row, err := update.RowsAffected()
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	if row == 0 {
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}

func (p *planService) UpdateContent(ownerId string, param *define.PlanUpdateContentReq) error {
	if param.Content == "" || param.PlanId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Plan.Ctx(context.TODO()).Where(g.Map{
		dao.Plan.Columns.OwnerId: ownerId,
		dao.Plan.Columns.PlanId:  param.PlanId,
	}).Update(g.Map{
		dao.Plan.Columns.Content: param.Content,
	})
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	row, err := update.RowsAffected()
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	if row == 0 {
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}

func (p *planService) UpdateStatus(ownerId string, param *define.PlanUpdateStatusReq) error {
	if param.PlanId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Plan.Ctx(context.TODO()).Where(g.Map{
		dao.Plan.Columns.OwnerId: ownerId,
		dao.Plan.Columns.PlanId:  param.PlanId,
	}).Update(g.Map{
		dao.Plan.Columns.Completed: param.Completed,
	})
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	row, err := update.RowsAffected()
	if err != nil {
		g.Log().Line().Warning(err)
		return err
	}
	if row == 0 {
		return shared.NewError(model.ERR_DATABASE_UPDATE)
	}
	return nil
}

func (p *planService) Delete(ownerId string, param *define.PlanDeleteReq) error {
	if len(param.PlanIds) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	for _, id := range param.PlanIds {
		_, err := dao.Plan.Ctx(context.TODO()).Where(g.Map{
			dao.Plan.Columns.OwnerId: ownerId,
			dao.Plan.Columns.PlanId:  id,
		}).Delete()
		if err != nil {
			g.Log().Line().Warning(err)
			return err
		}
	}
	return nil
}
