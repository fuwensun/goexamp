package grpc

import (
	"context"

	"github.com/aivuca/goms/eLog/api"
	. "github.com/aivuca/goms/eLog/internal/model"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

var empty = &api.Empty{}

func handValidateError(err error) error {
	for _, ev := range err.(validator.ValidationErrors) {
		log.Info().
			Msgf("%v err => %v", ev.StructField(), ev.Value())
		return UserErrMap[ev.Namespace()]
	}
	return nil
}

// CreateUser create user.
func (srv *Server) CreateUser(c context.Context, u *api.UserT) (*api.UidT, error) {
	svc := srv.svc
	res := &api.UidT{}
	// 记录参数
	log.Info().Msgf("start to create user, arg: {%v}", u)

	user := &User{}
	user.Uid = GetUid()
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		// 记录异常
		log.Info().
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return res, handValidateError(err)
	}
	// 记录中间结果
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	if err := svc.CreateUser(c, user); err != nil {
		// 记录异常
		log.Info().
			Int64("user_id", user.Uid).
			Msgf("fail to create user, data: %v, error: %v", *user, err)
		return res, ErrInternalError
	}
	res.Uid = user.Uid
	// 记录返回结果
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create user, user = %v", *user)
	return res, nil
}

// ReadUser read user.
func (srv *Server) ReadUser(c context.Context, uid *api.UidT) (*api.UserT, error) {
	svc := srv.svc
	res := &api.UserT{}

	log.Info().Msgf("start to read user, arg: {%v}", uid)

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Info().
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return res, handValidateError(err)
	}
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	u, err := svc.ReadUser(c, user.Uid)
	if err != nil {
		log.Info().
			Int64("user_id", res.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return res, ErrInternalError
	}

	res.Uid = u.Uid
	res.Name = u.Name
	res.Sex = u.Sex
	log.Info().
		Int64("user_id", res.Uid).
		Msgf("succ to read user, user = %v", *user)
	return res, nil
}

// UpdateUser update user.
func (srv *Server) UpdateUser(c context.Context, u *api.UserT) (*api.Empty, error) {
	svc := srv.svc

	log.Info().Msgf("start to update user, arg: {%v}", u)

	user := &User{}
	user.Uid = u.Uid
	user.Name = u.Name
	user.Sex = u.Sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Info().
			Msgf("fail to validate data, data: %v, error: %v", *user, err)
		return empty, handValidateError(err)
	}
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, user = %v", *user)

	err := svc.UpdateUser(c, user)
	if err != nil {
		log.Info().
			Int64("user_id", user.Uid).
			Msgf("fail to update user, data: %v, error: %v", *user, err)
		return empty, ErrInternalError
	}
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to update user, user = %v", *user)
	return empty, nil
}

// DeleteUser delete user.
func (srv *Server) DeleteUser(c context.Context, uid *api.UidT) (*api.Empty, error) {
	svc := srv.svc

	log.Info().Msgf("start to delete user, arg: {%v}", uid)

	user := &User{}
	user.Uid = uid.Uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		log.Info().
			Msgf("fail to validate data, data: %v, error: %v", user.Uid, err)
		return empty, handValidateError(err)
	}
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to create data, uid = %v", user.Uid)

	err := svc.DeleteUser(c, user.Uid)
	if err != nil {
		log.Info().
			Int64("user_id", user.Uid).
			Msgf("fail to read user, data: %v, error: %v", user.Uid, err)
		return empty, ErrInternalError
	}
	log.Info().
		Int64("user_id", user.Uid).
		Msgf("succ to read user, user = %v", *user)
	return empty, nil
}
