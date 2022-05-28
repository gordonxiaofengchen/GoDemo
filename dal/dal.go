package dal

import (
	"fmt"
)

type Page struct{
	Total int64 `json:"total"`
	PageSize int `json:"pageSize"`
	PageNum int `json:"pageNum"`
	List interface{} `json:"list"`
}

var (
	emptyObj interface{}
)

func idCond(id string) string{
	return fmt.Sprintf("ID = %s", id)
}

func InsertOne(obj interface{}) error {
	tx := OrmDb.Create(obj)
	return tx.Error
}

func QueryAll(obj interface{}) (bool, error){
	tx := OrmDb.Find(obj)
	return tx.RowsAffected > 0, tx.Error
}

func QueryAllByPage(p *Page) (bool, error){
	if p == nil || p.PageNum == 0 || p.PageSize == 0 {
		return false, nil
	}else{
		var total int64
		offset := p.PageSize * (p.PageNum - 1)
		if err := OrmDb.Count(&total).Error; err != nil{
			p.Total = 0
			return false, err
		}
		tx := OrmDb.Order("id").Limit(p.PageSize).Offset(offset).Find(p.List)
		p.Total = total
		return tx.RowsAffected > 0, tx.Error
	}
}

func QueryById(obj interface{}, id string) (bool, error){
	idCond := idCond(id)
	tx := OrmDb.Find(obj, idCond)
	return tx.RowsAffected > 0, tx.Error
}

func DeleteById(id string) (bool, error){
	idCond := idCond(id)
	tx := OrmDb.Delete(emptyObj, idCond)
	return tx.RowsAffected > 0, tx.Error
}

func UpdateWhole(obj interface{}) (bool, error){
	tx := OrmDb.Updates(obj)
	return tx.RowsAffected > 0, tx.Error
}

/*func UpdateNonZero(obj interface{}) (bool, error) {
	OrmDb.update
}*/