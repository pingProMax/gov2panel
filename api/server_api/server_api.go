// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package server_api

import (
	"context"
	
	"gov2panel/api/server_api/v1"
)

type IServerApiV1 interface {
	Api(ctx context.Context, req *v1.ApiReq) (res *v1.ApiRes, err error)
}


