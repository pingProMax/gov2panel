package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gview"

	"gov2panel/internal/controller/admin"
	"gov2panel/internal/controller/public"
	"gov2panel/internal/controller/server_api"
	user_c "gov2panel/internal/controller/user"
	"gov2panel/internal/middleware"
	"gov2panel/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			adminiPath, err := g.Cfg().Get(ctx, "admini_path")
			if err != nil {
				panic(err.Error())
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(func(r *ghttp.Request) { //设置参数
					r.Response.CORSDefault() //跨域处理
					d, err := service.Setting().GetSettingAllMap()
					if err != nil {
						panic(err.Error())
					}
					d["admin_path"] = adminiPath
					r.Assigns(gview.Params{
						"setting":     d,
						"admini_path": adminiPath,
					})

					r.SetCtxVar("setting", d)

					r.Middleware.Next()
				})

				group.Middleware(ghttp.MiddlewareHandlerResponse) //处理处理程序响应对象及其错误的默认中间件
				group.Bind(
					public.NewV1(),
				)

				group.Group("/api/server", func(group *ghttp.RouterGroup) {
					group.Middleware(func(r *ghttp.Request) { //设置参数
						// 中间件处理逻辑
						d, err := service.Setting().GetSettingAllMap()
						if err != nil {
							panic(err.Error())
						}

						if r.Get("token").String() != d["api_key"].String() {
							r.Response.WriteExit(`{"message": "token is error"}`)
						}

						r.Middleware.Next()
					}) //token
					group.Bind(
						server_api.NewV1(),
					)

				})

				group.Group("/"+adminiPath.String(), func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.AdminJWTAuth) //权限处理
					group.Bind(
						admin.NewV1(),
					)

				})

				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.UserJWTAuth) //权限处理
					group.Bind(
						user_c.NewV1(),
					)

				})

			})

			//每天6点执行  更新过期用户的权限组和流量
			service.User().ClearExpiredUserGroupIdAndUDTransferEnable()
			_, err = gcron.Add(ctx, "0 0 6 * * *", func(ctx context.Context) {
				service.User().ClearExpiredUserGroupIdAndUDTransferEnable()
			}, "CEUP_CRON")
			if err != nil {
				panic(err)
			}

			// 启动 把有效用户 存入到内存
			err = service.User().MSaveToRam()
			if err != nil {
				panic(err)
			}

			//启动时 从文件加载用户 7天流量使用数据 到内存
			err = service.User().LoadUserDay7FlowFromFile(gctx.GetInitCtx(), "./resource/user_day7_flow.json")
			if err != nil {
				panic(err)
			}

			s.Run()
			return nil
		},
	}
)
