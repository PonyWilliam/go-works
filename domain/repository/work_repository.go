package repository

import (
	"crypto/md5"
	"fmt"
	"github.com/PonyWilliam/go-works/domain/model"
	"github.com/jinzhu/gorm"
	"io"
)

type IWorker interface {
	InitTable() error
	CreateWorker(worker *model.Workers) (int64,error)
	UpdateWorker(worker *model.Workers) (int64,error)
	DeleteWorkerByID(int64) error
	FindWorkerByID(int64)(model.Workers,error)
	FindWorkerByNums(int64)(model.Workers,error)
	FindWorkersByName(string)([]model.Workers,error)
	FindWorkerByUserName(string)(model.Workers,error)
	FindAll()([]model.Workers,error)
	CheckSum(string,string)bool
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
	fmt.Println(worker.Username)
	worker.Password = EncodeMD5(worker.Password)
	fmt.Println(worker.Username)
	return worker.ID,w.mysqlDB.Model(worker).Create(&worker).Error
}
func (w *WorkersRepository) UpdateWorker(worker *model.Workers) (int64,error){
	worker.Password = EncodeMD5(worker.Password)
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
func (w *WorkersRepository) FindWorkerByUserName(username string)(worker model.Workers,err error){
	return worker,w.mysqlDB.Model(&model.Workers{}).Where("username = ?",username).Error
}
func (w *WorkersRepository)CheckSum(username string,pwd string)bool{

	temp := EncodeMD5(pwd)
	worker := model.Workers{}
	fmt.Println("username:" + username)
	w.mysqlDB.Where("username = ?",username).Find(&worker)
	fmt.Println("password:" + temp)
	fmt.Println("sql:" + worker.Password)
	if worker.Password == temp{
		return true
	}else{
		return false
	}
}
func EncodeMD5(pwd string)string{
	h := md5.New()
	_, _ = io.WriteString(h,pwd)
	sum := fmt.Sprintf("%x",h.Sum([]byte("123")))//独立签名
	return sum
}