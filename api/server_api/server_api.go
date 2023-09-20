// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package server_api

import (
	"context"
	
	"gov2panel/api/server_api/v1"
)

type IServerApiV1 interface {
	Config(ctx context.Context, req *v1.ConfigReq) (res *v1.ConfigRes, err error)
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
}


