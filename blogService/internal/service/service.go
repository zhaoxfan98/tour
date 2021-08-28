package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/zhaoxfan98/blog/global"
	"github.com/zhaoxfan98/blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//新增数据库连接实例的上下文信息注册
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
