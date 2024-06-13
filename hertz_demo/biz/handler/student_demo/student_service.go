// Code generated by hertz generator.

package student_demo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/han-y-s/hertz_demo/caller"
)

// 泛化调用
// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req student_demo.Student
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	req, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		hlog.CtxInfof(ctx, "get req error: %v", err)
		return
	}
	customReq, err := generic.FromHTTPRequest(req)

	if err != nil {
		hlog.CtxInfof(ctx, "get custom request error: %v", err)
		return
	}
	resp, err := caller.GeneralCli.GenericCall(ctx, "", customReq)
	if err != nil {
		c.JSON(200, utils.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(consts.StatusOK, resp)
	}
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	//var req student_demo.QueryReq
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	req, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		hlog.CtxInfof(ctx, "get req error: %v", err)
		return
	}
	customReq, err := generic.FromHTTPRequest(req)

	if err != nil {
		hlog.CtxInfof(ctx, "get custom request error: %v", err)
		return
	}
	resp, err := caller.GeneralCli.GenericCall(ctx, "", customReq)
	if err != nil {
		c.JSON(200, utils.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(consts.StatusOK, resp)
	}
}

//普通的rpc调用
//// Register .
//// @router /add-student-info [POST]
//func Register(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req student_demo.Student
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//	hlog.CtxInfof(ctx, string(c.Request.RequestURI()))
//	rpcReq := &student_demo2.Student{
//		Id:   req.ID,
//		Name: req.Name,
//		College: &student_demo2.College{
//			Name:    req.College.Name,
//			Address: req.College.Address,
//		},
//		Email: req.Email,
//	}
//	resp, err := caller.Rpccli.Register(ctx, rpcReq)
//	if err != nil {
//		c.JSON(200, utils.H{
//			"error": err.Error(),
//		})
//	} else {
//		c.JSON(consts.StatusOK, resp)
//	}
//}
//// Query .
//// @router /query [GET]
//func Query(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req student_demo.QueryReq
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		c.String(consts.StatusBadRequest, err.Error())
//		return
//	}
//	rpcReq := &student_demo2.QueryReq{
//		Id: req.ID,
//	}
//	resp, err := caller.Rpccli.Query(ctx, rpcReq)
//	if err != nil {
//		c.JSON(200, utils.H{
//			"error": err.Error(),
//		})
//	} else {
//		c.JSON(consts.StatusOK, resp)
//	}
//}
