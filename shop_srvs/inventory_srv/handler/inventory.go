package handler

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"shop_srvs/inventory_srv/global"
	"shop_srvs/inventory_srv/model"
	"shop_srvs/inventory_srv/proto"

	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存， 如果我要更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).FirstOrCreate(&inv, model.Inventory{
		Goods:  req.GoodsId,
		Stocks: req.Num,
	})

	if inv.Stocks != req.Num {
		inv.Stocks = req.Num
		global.DB.Save(&inv)
	}

	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}
func (*InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//无论如何一定要先提交事物，然后再释放锁
	pool := goredis.NewPool(global.RedisClient)
	rs := redsync.New(pool)

	// 创建事务
	tx := global.DB.Begin()

	// 先获取所有商品的锁
	mutexes := make([]*redsync.Mutex, 0)
	for _, goodInfo := range req.GoodsInfo {
		mutex := rs.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId))
		if err := mutex.Lock(); err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "获取分布式锁失败")
		}
		mutexes = append(mutexes, mutex)
	}

	// 处理扣减
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := tx.Where("goods = ?", goodInfo.GoodsId).First(&inv); result.Error != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "库存不存在")
		}

		if inv.Stocks < goodInfo.Num {
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}

		// 原子更新
		if err := tx.Model(&model.Inventory{}).
			Where("goods = ?", goodInfo.GoodsId).
			Update("stocks", gorm.Expr("stocks - ?", goodInfo.Num)).Error; err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "更新库存失败")
		}
	}

	// 提交事务后再释放锁
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "事务提交失败")
	}

	for _, mutex := range mutexes {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			zap.S().Infof("释放锁失败: %v", err)
		}
	}

	return &emptypb.Empty{}, nil
}

func (*InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还： 1：订单超时归还 2. 订单创建失败，归还之前扣减的库存 3. 手动归还
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		if err := tx.Model(&model.Inventory{}).
			Where("goods = ?", goodInfo.GoodsId).
			Update("stocks", gorm.Expr("stocks + ?", goodInfo.Num)).
			Error; err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "更新库存失败: %v", err)
		}
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}
