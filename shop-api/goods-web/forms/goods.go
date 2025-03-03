package forms

type GoodsForm struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=100"`
	GoodsSn     string   `form:"goods_sn" json:"goods_sn" binding:"required,min=2,lt=20"`
	Stocks      int32    `form:"stocks" json:"stocks" binding:"required,min=1"`
	CategoryId  int32    `form:"category" json:"category" binding:"required"`
	MarketPrice float32  `form:"market_price" json:"market_price" binding:"required,min=0"`
	ShopPrice   float32  `form:"shop_price" json:"shop_price" binding:"required,min=0"`
	GoodsBrief  string   `form:"goods_brief" json:"goods_brief" binding:"required,min=3"`
	Images      []string `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string `form:"desc_images" json:"desc_images" binding:"required,min=1"`
	ShipFree    *bool    `form:"ship_free" json:"ship_free" binding:"required"`
	FrontImage  string   `form:"front_image" json:"front_image" binding:"required,url"`
	Brand       int32    `form:"brand" json:"brand" binding:"required"`
}

//
//message CreateGoodsInfo {
//int32 id = 1;
//string name = 2;
//string goodsSn = 3;
//int32 stocks = 7; //库存，
//float marketPrice = 8;
//float shopPrice = 9;
//string goodsBrief = 10;
//string goodsDesc = 11;
//bool shipFree = 12;
//repeated string images = 13;
//repeated string descImages = 14;
//string goodsFrontImage = 15;
//bool isNew = 16;
//bool isHot = 17;
//bool onSale = 18;
//int32 categoryId = 19;
//int32 brandId = 20;
//}
