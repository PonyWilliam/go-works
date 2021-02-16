package handler

import (
	"context"
	"fmt"
	common "github.com/PonyWilliam/go-common"
	"github.com/PonyWilliam/go-works/domain/model"
	work "github.com/PonyWilliam/go-works/domain/server"
	works "github.com/PonyWilliam/go-works/proto"
)

type Works struct{
	WorkService work.IWorkerServices
}
// Call is a single request handler called via client.Call or the generated client code
func(w *Works)CreateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Works)UpdateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Works)DeleteWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Workers)error{
	err := w.WorkService.DeleteWorkerByID(req.Id)
	if err!=nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	return nil
}
func(w *Works)FindWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Worker_Show)error{
	worker,err := w.WorkService.FindWorkerByID(req.Id)
	fmt.Println(worker)
	workers := &works.Response_Workers_Info{}
	if err!=nil{
		return err
	}
	err = common.SwapTo(worker, workers)
	if err != nil{
		return err
	}
	res.Worker = workers
	return nil
}
func(w *Works)FindWorkerByName(ctx context.Context,req *works.Request_Workers_Name,res *works.Response_Workers_Show)error{
	workers,err := w.WorkService.FindWorkersByName(req.Name)
	if err != nil{
		return err
	}
	for _,v := range workers{
		worker := &works.Response_Workers_Info{}
		err = common.SwapTo(v,worker)
		res.Workers = append(res.Workers,worker)
	}
	if err != nil{
		return err
	}
	return nil
}
func(w *Works)FindAll(ctx context.Context,req *works.Request_Null,res *works.Response_Workers_Show)error{
	workers,err := w.WorkService.FindAll()
	if err != nil{
		return err
	}
	for _,v := range workers{
		worker := &works.Response_Workers_Info{}
		err = common.SwapTo(v,worker)
		res.Workers = append(res.Workers,worker)
	}
	if err != nil{
		return err
	}
	return nil
}
func(w *Works)FindWorkerByNums(ctx context.Context,req *works.Request_Workers_Nums,res *works.Response_Worker_Show) error{
	worker,err := w.WorkService.FindWorkerByNums(req.Nums)
	fmt.Println(worker)
	workers := &works.Response_Workers_Info{}
	if err!=nil{
		return err
	}
	err = common.SwapTo(worker, workers)
	if err != nil{
		return err
	}
	res.Worker = workers
	return nil
}
func(w *Works)CheckSum(ctx context.Context,req *works.LoginRequest,rsp *works.LoginResponse) error{
	username := req.User
	password := req.Password
	if w.WorkService.CheckSum(username,password){
		rsp.Code = true
		rsp.Msg = "success"
		return nil
	}
	rsp.Code = false
	rsp.Msg = "password is not true"
	return nil
}