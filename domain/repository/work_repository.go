package repository

import (
	"works/domain/model"
	"github.com/jinzhu/gorm"
)

type IWorker interface {
	InitTable() error
	CreateWorker(worker *model.Workers) (int64,error)
	UpdateWorker(worker *model.Workers) (int64,error)
	DeleteWorkerByID(int64) error
	FindWorkerByID(int64)(model.Workers,error)
	FindWorkerByNums(int64)(model.Workers,error)
	FindWorkersByName(string)([]model.Workers,error)
	FindAll()([]model.Workers,error)

}
func NewWorkerRepository(db *gorm.DB)IWorker{
	return &WorkersRepository{mysqlDB: db}
}
type WorkersRepository struct{
	mysqlDB *gorm.DB
}
func (w *WorkersRepository) InitTable() error{
	if w.mysqlDB.HasTable(&model.Workers{}){
		return nil
	}
	return w.mysqlDB.CreateTable(&model.Workers{}).Error
}
func (w *WorkersRepository) CreateWorker(worker *model.Workers) (int64,error){
	return worker.ID,w.mysqlDB.Model(worker).Create(&worker).Error
}
func (w *WorkersRepository) UpdateWorker(worker *model.Workers) (int64,error){
	return worker.ID,w.mysqlDB.Model(worker).Update(&worker).Error
}
func (w *WorkersRepository) DeleteWorkerByID(id int64) error{
	return w.mysqlDB.Where("id = ?",id).Delete(&model.Workers{}).Error
}
func (w *WorkersRepository) FindWorkerByID(id int64) (worker model.Workers,err error){
	return worker,w.mysqlDB.Model(&model.Workers{}).Where("id  = ?",id).Find(&worker).Error
}
func (w *WorkersRepository)FindWorkerByNums(nums int64)(worker model.Workers,err error){
	return worker,w.mysqlDB.Model(&model.Workers{}).Where("nums = ?",nums).Find(&worker).Error
}
func (w *WorkersRepository) FindWorkersByName(name string) (workers []model.Workers,err error){
	return workers,w.mysqlDB.Model(&model.Workers{}).Where("name  = ?",name).Find(&workers).Error
}
func (w *WorkersRepository) FindAll() (workers []model.Workers,err error){
	return workers,w.mysqlDB.Model(&model.Workers{}).Find(&workers).Error
}
