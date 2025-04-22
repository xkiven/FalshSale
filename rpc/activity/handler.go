package main

import (
	activity_service "FlashSale/kitex_gen/FlashSale/activity_service"
	"FlashSale/logic"
	"FlashSale/svc"
	"context"
)

// ActivityServiceImpl implements the last service interface defined in the IDL.
type ActivityServiceImpl struct {
	sc *svc.ServiceContext
}

// NewActivityServiceImpl 创建 ActivityServiceImpl 实例
func NewActivityServiceImpl(sc *svc.ServiceContext) *ActivityServiceImpl {
	return &ActivityServiceImpl{
		sc: sc,
	}
}

// CreateActivity implements the ActivityServiceImpl interface.
func (s *ActivityServiceImpl) CreateActivity(ctx context.Context, req *activity_service.CreateActivityRequest) (resp *activity_service.CreateActivityResponse, err error) {
	// TODO: Your code here...
	resp, err = logic.CreateActivity(ctx, s.sc.MySQLClient, s.sc.RedisClient, req)
	return resp, err
}
