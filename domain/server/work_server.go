package server

import (
	"github.com/PonyWilliam/go-works/domain/model"
	"github.com/PonyWilliam/go-works/domain/repository"
)

type IWorkerServices interface {
	CreateWorker(worker *model.Workers) (int64,error)
	UpdateWorker(worker *model.Workers) (int64,error)
	DeleteWorkerByID(int64) error
	FindWorkerByID(int64)(model.Workers,error)
	FindWorkerByNums(int64)(model.Workers,error)
	FindWorkersByName(string)([]model.Workers,error)
	FindAll()([]model.Workers,error)
	CheckSum(string,string) bool
}
func NewWorkerServices(worker repository.IWorker)IWorkerServices{
	return &WorkServices{worker}
}

type WorkServices struct{
	worker repository.IWorker
}
func(w *WorkServices) CreateWorker(worker *model.Workers) (int64,error){
	return w.worker.CreateWorker(worker)
}
func(w *WorkServices) UpdateWorker(worker *model.Workers) (int64,error){
	return w.worker.UpdateWorker(worker)
}
func(w *WorkServices) DeleteWorkerByID(id int64) error{
	return w.worker.DeleteWorkerByID(id)
}
func(w *WorkServices) FindWorkerByID(id int64)(model.Workers,error){
	return w.worker.FindWorkerByID(id)
}
func(w *WorkServices) FindWorkerByNums(nums int64)(model.Workers,error){
	return w.worker.FindWorkerByNums(nums)
}
func(w *WorkServices) FindWorkersByName(name string)([]model.Workers,error){
	return w.worker.FindWorkersByName(name)
}
func(w *WorkServices) FindAll()([]model.Workers,error){
	return w.worker.FindAll()
}
func(w *WorkServices)CheckSum(username string,password string)bool{
	return w.worker.CheckSum(username,password)
}
