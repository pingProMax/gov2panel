package user

import (
	"context"
	"fmt"
	"gov2panel/internal/dao"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
	"gov2panel/internal/service"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
)

//userMap  k是用户id v结构体是model.UserTraffic
//这个处理是 用户流量上报问题的，避免每一分钟节点上报流量都更新数据库

func init() {
	ctx := gctx.New()

	//每一个小时执行一次  持久化所有 userMap 数据
	gcron.Add(ctx, "0 0 */1 * * *", func(ctx context.Context) {
		service.User().MSaveAllRam()
	}, "SAVE_USER_MAP_CRON")

}

// 内存map数据
var userMap = gmap.NewHashMap(true)

// 启动 把有效用户 存入到内存
func (s *sUser) MSaveToRam() (err error) {

	userArr := make([]*entity.V2User, 0)
	err = s.Cornerstone.GetDB().
		WhereGT(dao.V2User.Columns().GroupId, 0).
		Where(fmt.Sprintf("%s > %s + %s", dao.V2User.Columns().TransferEnable, dao.V2User.Columns().U, dao.V2User.Columns().D)).
		WhereGT(dao.V2User.Columns().ExpiredAt, time.Now()).
		Where(dao.V2User.Columns().Banned, -1).
		WhereNotIn(dao.V2User.Columns().Uuid, "").
		Scan(&userArr)
	if err != nil {
		return
	}

	for _, user := range userArr {
		userMap.Set(user.Id, model.UserToUserTraffic(user))
	}

	fmt.Println(userArr)

	return nil
}

// 更新用户 流量使用情况2 直接更新缓存（原来有一个直接更新数据库UpUserUAndDBy）
func (s *sUser) MUpUserUAndBy(data []*model.UserTraffic) (err error) {

	for _, u := range data {
		if !userMap.GetVar(u.UID).IsNil() {
			//存在缓存

			var userTraffic model.UserTraffic
			err = userMap.GetVar(u.UID).Struct(&userTraffic)
			if err == nil {
				userTraffic.Upload = userTraffic.Upload + u.Upload
				userTraffic.Download = userTraffic.Download + u.Download

				//流量判断、到期时间判断、用户权限组、用户状态
				if (userTraffic.Upload+userTraffic.Download) >= userTraffic.TransferEnable || userTraffic.ExpiredAt.Unix() <= time.Now().Unix() || userTraffic.GroupId <= 0 || userTraffic.Banned == 1 {
					userMap.Remove(userTraffic.UID)                   //map中删除
					s.UpUserDUTBy([]*model.UserTraffic{&userTraffic}) //保存数据库
				} else {
					userMap.Set(userTraffic.UID, userTraffic)
				}

			}
		} else {
			newU, err := s.GetUserById(u.UID)
			if err == nil {
				if newU.Id != 0 {
					userMap.Set(newU.Id, model.UserToUserTraffic(newU))
				}
			}
		}

	}

	//用户流量7天使用缓存
	err = s.UpUserDay7Flow(data)
	if err != nil {
		return
	}
	return
}

// 所有数据持久化
func (s *sUser) MSaveAllRam() (err error) {
	data := make([]*model.UserTraffic, 0)
	for _, v := range userMap.Keys() {
		var user model.UserTraffic
		err = userMap.GetVar(v).Struct(&user)
		if err == nil {
			data = append(data, &user)
		}
	}

	//保存到数据库
	err = s.UpUserDUTBy(data)
	if err != nil {
		return err
	}

	return
}

// 更新/添加 缓存
func (s *sUser) MUpUserMap(data *model.UserTraffic) {
	userMap.Set(data.UID, data)

}

// 先把原有缓存更新到数据库,再查询查询数据库更新到缓存
func (s *sUser) MUpDbAndUserMap(uid int) (err error) {
	var userTraffic model.UserTraffic
	err = userMap.GetVar(uid).Struct(&userTraffic)
	if err != nil {
		return
	}

	err = s.UpUserDUTBy([]*model.UserTraffic{&userTraffic})
	if err != nil {
		return
	}

	user, err := s.GetUserById(uid)
	if err != nil {
		return err
	}

	s.MUpUserMap(model.UserToUserTraffic(user))
	return
}

// 删除 缓存
func (s *sUser) MDelUserMap(id int) {
	userMap.Remove(id)
}

// 权限组获取用户
func (s *sUser) MGetUserByGroupId(GroupId int) (d []*model.UserTraffic) {
	d = make([]*model.UserTraffic, 0)
	for _, v := range userMap.Keys() {
		var user model.UserTraffic
		err := userMap.GetVar(v).Struct(&user)
		if err == nil {
			if user.GroupId == GroupId {
				d = append(d, &user)
			}
		}
	}

	return
}
