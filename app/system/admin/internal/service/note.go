package service

import (
	"CloudNote/app/dao"
	"CloudNote/app/model"
	"CloudNote/app/shared"
	"CloudNote/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/guuid"
	"time"
)

var Note = noteService{}

type noteService struct{}

// Add 新增笔记
func (n *noteService) Add(ownerId string, param *define.NoteAddReq) error {
	if param.Content == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	if param.Title == "" {
		var endIndex int
		if len(param.Content) < 10 {
			endIndex = len(param.Content)
		} else {
			endIndex = 10
		}
		param.Title = param.Content[:endIndex] + "..."
	}
	saveDate := model.Note{
		NoteId:     guuid.New().String(),
		Title:      param.Title,
		Content:    param.Content,
		CategoryId: param.CategoryId,
		Shared:     param.Shared,
		OwnerId:    ownerId,
	}
	_, err := dao.Note.Ctx(context.TODO()).OmitEmpty().Insert(saveDate)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLimit 更新笔记权限(是否公开)
func (n *noteService) UpdateLimit(ownerId string, param *define.NoteUpdateLimitReq) error {
	if param.NoteId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.NoteId:  param.NoteId,
		dao.Note.Columns.OwnerId: ownerId,
	}).Update(g.Map{
		dao.Note.Columns.Shared: param.Shared,
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

// UpdateTitle 更新笔记标题
func (n *noteService) UpdateTitle(ownerId string, param *define.NoteUpdateTitleReq) error {
	if param.NoteId == "" || param.Title == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.NoteId:  param.NoteId,
		dao.Note.Columns.OwnerId: ownerId,
	}).Update(g.Map{
		dao.Note.Columns.Title: param.Title,
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

// UpdateContent 更新笔记内容
func (n *noteService) UpdateContent(ownerId string, param *define.NoteUpdateContentReq) error {
	if param.NoteId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.NoteId:  param.NoteId,
		dao.Note.Columns.OwnerId: ownerId,
	}).Update(g.Map{
		dao.Note.Columns.Content: param.Content,
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

// UpdateCategory 更新笔记的分类
func (n *noteService) UpdateCategory(ownerId string, param *define.NoteUpdateCategoryReq) error {
	if param.NoteId == "" || param.CategoryId == "" {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.NoteId:  param.NoteId,
		dao.Note.Columns.OwnerId: ownerId,
	}).Update(g.Map{
		dao.Note.Columns.CategoryId: param.CategoryId,
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

// Delete 删除笔记(移至回收站)
func (n *noteService) Delete(ownerId string, param *define.NoteDeleteReq) error {
	if len(param.NoteIds) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	for _, noteId := range param.NoteIds {
		update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
			dao.Note.Columns.NoteId:  noteId,
			dao.Note.Columns.OwnerId: ownerId,
		}).Update(g.Map{
			dao.Note.Columns.Trash:       true,
			dao.Note.Columns.DeletedTime: gtime.New(time.Now()).String(),
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
	}
	return nil
}

// RecoverFromTrash 从回收站中恢复
func (n *noteService) RecoverFromTrash(ownerId string, param *define.NoteRecoverReq) error {
	if len(param.NoteIds) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	for _, noteId := range param.NoteIds {
		update, err := dao.Note.Ctx(context.TODO()).Where(g.Map{
			dao.Note.Columns.NoteId:  noteId,
			dao.Note.Columns.OwnerId: ownerId,
		}).Update(g.Map{
			dao.Note.Columns.Trash: false,
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
	}
	return nil
}

// DeleteFromTrash 从回收站中移除
func (n *noteService) DeleteFromTrash(ownerId string, param *define.NoteDeleteFromTrashReq) error {
	if len(param.NoteIds) == 0 {
		return shared.NewError(model.ERR_INVALID_PARAM)
	}
	for _, noteId := range param.NoteIds {
		_, err := dao.Note.Ctx(context.TODO()).Delete(g.Map{
			dao.Note.Columns.NoteId:  noteId,
			dao.Note.Columns.OwnerId: ownerId,
			dao.Note.Columns.Trash:   true,
		})
		if err != nil {
			g.Log().Line().Warning(err)
			return err
		}
	}
	return nil
}

// BaseQuery 查询笔记列表
func (n *noteService) BaseQuery(ownerId string, param *define.NoteBaseQueryReq) (interface{}, error) {
	if len(ownerId) == 0 {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.OwnerId: ownerId,
	})
	if param.Trash == true {
		sql = sql.Where(dao.Note.Columns.Trash, param.Trash)
	} else {
		sql = sql.Where(dao.Note.Columns.Trash, false)
	}

	if len(param.CategoryId) > 0 {
		sql = sql.Where(dao.Note.Columns.CategoryId, param.CategoryId)
	}

	//sql = sql.Where(dao.Note.Columns.Shared, param.Shared)

	if param.Page >= 1 && param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	if param.NoteId != "" {
		sql = sql.Where(dao.Note.Columns.NoteId, param.NoteId)
	}
	if param.Title != "" {
		sql = sql.WhereLike(dao.Note.Columns.Title, "%"+param.Title+"%")
	}
	if param.Content != "" {
		sql = sql.WhereLike(dao.Note.Columns.Content, "%"+param.Content+"%")
	}

	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	// todo 新增功能：返回分类
	var result []*define.NoteBaseQueryResp
	var noteData []*model.Note
	err = sql.Fields(dao.Note.Columns.Title,
		dao.Note.Columns.NoteId,
		dao.Note.Columns.UpdatedAt,
		dao.Note.Columns.CategoryId,
	).Scan(&noteData)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}

	err = gconv.Scan(noteData, &result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	for i := 0; i < len(noteData); i++ {
		if len(noteData[i].CategoryId) == 0 {
			continue
		}
		categoryTemp, err := Category.QueryById(noteData[i].CategoryId)
		if err != nil {
			g.Log().Line().Warning(err)
			return nil, err
		}
		if categoryTemp != nil {
			result[i].Category = *categoryTemp
		}
	}

	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

// DetailQuery 查询笔记详细内容
func (n *noteService) DetailQuery(ownerId string, param *define.NoteDetailQueryReq) (interface{}, error) {
	if param.NoteId == "" || len(ownerId) == 0 {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.OwnerId: ownerId,
		dao.Note.Columns.NoteId:  param.NoteId,
	})

	var result *model.Note
	var noteDetail *define.Note
	err := sql.Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	gconv.Scan(result, &noteDetail)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	if result.CategoryId != "" {
		categoryTemp, err := Category.QueryById(result.CategoryId)
		if err != nil {
			g.Log().Line().Warning(err)
			return nil, err
		}
		noteDetail.Category = *categoryTemp
	}
	return noteDetail, nil
}

// BaseQueryByUserId 查询用户公开的笔记
func (n *noteService) BaseQueryByUserId(param *define.NotePublicBaseQueryReq) (interface{}, error) {
	if len(param.UserId) == 0 {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.OwnerId: param.UserId,
		dao.Note.Columns.Shared:  true,
		dao.Note.Columns.Trash:   false,
	})

	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}

	var result []*define.NoteBaseQueryResp
	err = sql.Fields(dao.Note.Columns.Title,
		dao.Note.Columns.NoteId,
	).Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

// DetailQueryByUserId 查询用户公开笔记的详细内容
func (n *noteService) DetailQueryByUserId(param *define.NotePublicDetailQueryReq) (interface{}, error) {
	if param.UserId == "" || len(param.NoteId) == 0 {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Note.Ctx(context.TODO()).Where(g.Map{
		dao.Note.Columns.OwnerId: param.UserId,
		dao.Note.Columns.NoteId:  param.NoteId,
		dao.Note.Columns.Shared:  true,
		dao.Note.Columns.Trash:   false,
	})

	var result *model.Note
	err := sql.Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	return result, nil
}
