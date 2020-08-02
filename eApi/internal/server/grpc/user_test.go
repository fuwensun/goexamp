package grpc

import (
	"context"
	"testing"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/eApi/internal/pkg/err"
	"github.com/aivuca/goms/eApi/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}
	// monkey
	uid := m.GetUid()
	Patch(m.GetUid, func() int64 {
		return uid
	})
	Convey("TestCreateUser should succ", t, func() {
		//mock
		user := &m.User{
			Uid:  uid,
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctx, usert)

		//断言
		So(uidt.Uid, ShouldEqual, uid)
		So(err, ShouldEqual, nil)
	})

	Convey("TestCreateUser should failed", t, func() {
		//mock
		user := &m.User{
			Uid:  uid,
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			CreateUser(gomock.Any(), user).
			Return(e.ErrInternalError)

		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		uidt, err := srv.CreateUser(ctx, usert)
		//断言
		So(uidt.Uid, ShouldEqual, 0) //todo
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestReadUser should succ", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, nil)

		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		usert, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, nil)
		So(usert.Uid, ShouldEqual, user.Uid)
		So(usert.Name, ShouldEqual, user.Name)
		So(usert.Sex, ShouldEqual, user.Sex)
	})

	Convey("TestReadUser should failed", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}

		svcm.EXPECT().
			ReadUser(gomock.Any(), user.Uid).
			Return(user, e.ErrInternalError)

		//构建 req
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.ReadUser(ctx, uidt)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestUpdateUser should succ", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(nil)

		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestUpdateUser should failed", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			UpdateUser(gomock.Any(), user).
			Return(e.ErrInternalError)

		//构建 req
		usert := &api.UserT{
			Uid:  user.Uid,
			Name: user.Name,
			Sex:  user.Sex,
		}
		//发起 req
		_, err := srv.UpdateUser(ctx, usert)
		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	srv := Server{svc: svcm}

	Convey("TestDeleteUser should succ", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(nil)

		//构建 req
		var ctx = context.Background()
		usert := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, usert)
		//断言
		So(err, ShouldEqual, nil)
	})

	Convey("TestDeleteUser should failed", t, func() {
		//mock
		user := &m.User{
			Uid:  m.GetUid(),
			Name: m.GetName(),
			Sex:  m.GetSex(),
		}
		svcm.EXPECT().
			DeleteUser(gomock.Any(), user.Uid).
			Return(e.ErrInternalError)

		//构建 req
		var ctx = context.Background()
		uidt := &api.UidT{
			Uid: user.Uid,
		}
		//发起 req
		_, err := srv.DeleteUser(ctx, uidt)

		//断言
		So(err, ShouldEqual, e.ErrInternalError)
	})
}
