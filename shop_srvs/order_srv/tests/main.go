package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/order_srv/proto"
)

var orderClient proto.OrderClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	orderClient = proto.NewOrderClient(conn)
}

func TestCreateCartItem(userId, nums, goodsId int32) {
	rsp, err := orderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		Nums:    nums,
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}

func TestCartItemList(userId int32) {
	rsp, err := orderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, cartItem := range rsp.Data {
		fmt.Println(cartItem.Id)
	}
}

func TestUpdateCartItem(id int32) {
	_, err := orderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		Id:      id,
		Checked: true,
	})
	if err != nil {
		panic(err)
	}
}

func TestGetOrderDetail(orderId int32) {
	rsp, err := orderClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id: orderId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.OrderInfo.OrderSn)

	for _, good := range rsp.Goods {
		fmt.Println(good.GoodsName)
	}

}

func TestOrderList() {
	rsp, err := orderClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId: 12,
	})
	if err != nil {
		panic(err)
	}

	for _, order := range rsp.Data {
		fmt.Println(order.OrderSn)
	}
}

func TestCreateOrder(userId int) {
	_, err := orderClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(userId),
		Address: "湖南",
		Name:    "bigYYJ",
		Mobile:  "18787878787",
		Post:    "搜索树",
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()

	//for i := 4; i < 15; i++ {
	//	TestCreateCartItem(int32(i), 2, 422+int32(i))
	//}
	//TestCartItemList(1)
	//TestUpdateCartItem(2)
	//TestCreateOrder(13)
	//TestGetOrderDetail(11)
	TestOrderList()
	conn.Close()
}
