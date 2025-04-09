package main

import (
	activity_service "FlashSale/kitex_gen/FlashSale/activity_service"
	"context"
)

// ActivityServiceImpl implements the last service interface defined in the IDL.
type ActivityServiceImpl struct{}

// CreateActivity implements the ActivityServiceImpl interface.
func (s *ActivityServiceImpl) CreateActivity(ctx context.Context, req *activity_service.CreateActivityRequest) (resp *activity_service.CreateActivityResponse, err error) {
	// TODO: Your code here...
	return
}
