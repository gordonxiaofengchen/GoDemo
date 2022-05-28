package biz

import (
	"demo/dal"
	"demo/model"
)

type staffBiz struct{
}

var (
	StaffBiz staffBiz
	defaultQuerySliceSize = 10
	)

func (staffBiz) InsertOne(staff *model.Staff) error{
	return dal.InsertOne(staff)
}

func (staffBiz) QueryAll(p *dal.Page) (bool, error){
	staffs := make([]*model.Staff, defaultQuerySliceSize)
	p.List = staffs
	return dal.QueryAllByPage(p)
}

func (staffBiz) QueryById(id string) (*model.Staff, bool, error){
	staff := &model.Staff{}
	ok, err := dal.QueryById(staff, id)
	return staff, ok, err
}

func (staffBiz) DeleteById(id string) (bool, error){
	return dal.DeleteById(id)
}

func (staffBiz) UpdateWhole(staff *model.Staff) (bool, error){
	return dal.UpdateWhole(staff)
}
