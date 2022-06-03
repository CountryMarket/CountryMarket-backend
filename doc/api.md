# 后端 API 文档

## 数据交换格式

### 身份验证

在 Header 中加入 `Authorization` 字段进行验证，将获取的 JWT 令牌作为 Bearer Token 加入该字段的值，例如：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U
```

以下接口中，标题带有 `*` 标记的为需要身份验证的接口。

### 响应格式

响应使用 JSON 格式，例如：

```json
{
    "success": true,
    "info": '',
    "data": {
        // ...
    }
}
```

### URL 前缀

文档中所有接口 URL 都包含前缀 `/api/v1`。

# 用户 API

以下开头打 * 需要头部提供 token

每个 user 有一个自增 ID

## 小程序 code 发送 GET /user/code

本方法发送小程序的 code 到后端换取 token。

请求：

```json
{
    "code": "code", // 小程序 GET 到的 code
    "nickName": "FlyInTheSky", // 用户昵称
    "avatarUrl": "http://xxx/"  // 头像地址
}
```

响应

```json
{
    "token": "token" // token
}
```

## *检验 token 合法性 GET /user/validate

本方法检测当前 token 的合法性

请求：无

响应：若 success 为 false，且 info 字段为 `"invalid token" `或 `"auth error"`，则失效

## *获取用户信息 GET /user/profile

本方法获取用户信息

请求：无

响应：

```json
{
    "nickName": "FlyInTheSky",   // 名字
    "avatarUrl": "http://xxx/",  // 头像地址
    "phoneNumber": "13866654321", // 电话号码，无则返回空字符串
    "permission": 3 // int，包含1表示普通用户，2表示商户，4表示root，使用位运算
}
```

## **修改用户权限 POST /user/modifyPermission

(仅 Root)

请求：

```json
{
    "userId": 3, // 要修改的用户的 ID
    "permission": 1 // 修改为的权限
}
```

响应：无

# 地址 API

## *新增收货信息 POST /address/addAddress

请求：

```json
{
	"name": "", // 姓名
    "address": "", // 地址
    "phoneNumber": "" // 电话号码
}
```

响应：无

## *修改收货信息 POST /address/modifyAddress

请求：

```json
{
    "addressId" : 3, // 要修改的 收货信息 Id
	"name": "", // 姓名
    "address": "", // 地址
    "phoneNumber": "" // 电话号码
}
```

响应：无

## *删除收货信息 POST /address/deleteAddress

请求：

```json
{
    "addressId": 3 // 要删除的 收货信息 Id
}
```

响应：无

## *查询用户收货信息列表 GET /address/address

请求：无

响应：

```json
{
    "address": [
        {
            "addressId" : 2, // 要修改的 收货信息 Id
            "name": "", // 姓名
            "address": "", // 地址
            "phoneNumber": "" // 电话号码
        },
        {
            "addressId" : 4, // 要修改的 收货信息 Id
            "name": "", // 姓名
            "address": "", // 地址
            "phoneNumber": "" // 电话号码
        }
    ]
}
```

## *修改用户默认收货信息 POST /address/modifyDefault

请求：

```json
{
    "address_id": 4 // 需要改到的地址 Id
}
```

响应：无

## *获取用户默认收货信息 GET /address/default

请求：无

响应：

```json
{
    "address_id": 4 // 默认地址 Id
}
```

# 商户 API

每个 product 有一个自增 ID

## *增加一个商品 POST /shop/addProduct

(需要是商户账号，即permission有2)

请求：

```json
{
    "price": 233, // 价格
    "title": "没过期猪小排100g", // 标题
    "description": "过期的猪小排。", // 描述
    "pictureNumber": 2, // 可选，详情页幻灯片图片页数
    "stock": 32, // 库存，不提供默认为 0
    "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    "detailPictureNumber": 3 // 可选，详情信息处的图片页数
}
```

响应：

```json
{
    "id": 4 // 新增商品的 id
}
```

## *更新一个商品 POST /shop/updateProduct

(需要是商户账号，即permission有2，且更新的商品是申请用户的)

请求：

```json
{
    "price": 233, // 价格
    "title": "过期猪小排100g", // 标题
    "description": "过期的猪小排。", // 描述
    "pictureNumber": 2, // 可选，详情页幻灯片图片页数
    "stock": 32, // 库存，不提供默认为 0
    "id": 2, // 要改的商品 id
    "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    "detailPictureNumber": 3 // 可选，详情信息处的图片页数
}
```

响应：无

## 查询一个商品信息 GET /shop/product

请求：

```json
{
    "id": 2 // 商品 (product) 对应 id
}
```

响应：

```json
{
    "id": 2, // 商品 id
    "price": 233, // 价格
    "title": "过期猪小排100g", // 标题
    "description": "过期的猪小排。", // 描述
    "pictureNumber": 2, // 可选，详情页幻灯片图片页数
    "ownerUserId": 2, // 谁的 product，为 user id
    "stock": 32, // 库存，不提供默认为 0
    "isDrop": false, // 是否下架
    "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    "detailPictureNumber": 3 // 可选，详情信息处的图片页数
}
```

## *查询一个 Owner 的所有商品 GET /shop/ownerProducts

请求：

```json
{
	"from": 0, // 开始 index
	"length": 10 // 取多少个
} // [from, from + length) 的商品将被返回
```

响应：

```json
{
    "Products": [
        {
            "id": 2,
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false, // 是否下架
            "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    		"detailPictureNumber": 3 // 可选，详情信息处的图片页数
        },
        {
            "id": 4,
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false, // 是否下架
            "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    		"detailPictureNumber": 3 // 可选，详情信息处的图片页数
        }
    ]
}
```

## *下架一个商品 POST /shop/dropProduct

请求：

```json
{
    "id": 4 // 下架的商品
}
```

响应：无

## *上架一个商品 POST /shop/putProduct

请求：

```json
{
    "id": 4 // 上架的商品
}
```

响应：无

# 购物车 API

## *查询用户购物车商品 GET /cart/userProducts

用于购物车页面

请求：

```json
{
	"from": 0, // 开始 index
	"length": 10 // 取多少个
} // [from, from + length) 的商品将被返回
```

响应：

```json
{
    "Products": [
        {
            "id": 3, // 商品 id
            "count": 3, // 加入的数量
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        },
        {
            "id": 4, // 商品 id
            "count": 4, // 加入的数量
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        }
    ]
}
```

## *查询某商品是否在用户购物车中 GET /cart/inCart

用于商品页面，查找一个商品是否在用户购物车中，有则显示加了多少个

请求：

```json
{
    "productId": 2 // 商品 (product) 对应 id
}
```

响应：

```json
{
	"inCart": true, // 是否在购物车中
    "count": 3 // 加入的数量
}
```

## *用商品ID查询购物车中的商品 POST /cart/getCart

请求：

```json
{
    "product_ids": [1, 3, 4, 6]
}
```

响应：

```json
{
    "products": [
         {
            "id": 4, // 商品 id
            "count": 4, // 加入的数量
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        },
         {
            "id": 2, // 商品 id
            "count": 4, // 加入的数量
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        }
        // ...
    ]
}
```

## *增加购物车中的商品数量 POST /cart/addProduct

给出商品 id，增加其商品数量

请求：

```json
{
    "productId": 3 // 商品 id
}
```

响应：

```json
{
    "count": 3 // 当前此商品加入的数量
}
```

## *减少购物车中的商品数量 POST /cart/reduceProduct

给定购物车商品 id，减少其商品数量

请求：

```json
{
    "productId": 3 // 商品 id
}
```

响应：

```json
{
    "count": 2 // 当前此商品加入的数量
}
```

## *修改购物车中的商品数量 POST /cart/modifyProduct

给出商品 id，修改其商品数量

请求：

```json
{
    "productId": 3, // 商品 id
    "modifyCount": 4 // 修改到多少
}
```

响应：

```json
{
    "count": 4 // 当前此商品加入的数量
}
```

# 商品 API

## 查询 Tab List GET /product/tabList

请求：无

响应：

```json
{
    "tabs": [
        {
            "tabId": 1, // Tab Id
            "name": "xxx" // Tab 名称
        },
        {
            "tabId": 2, // Tab Id
            "name": "xxx" // Tab 名称            
        },
        {
            "tabId": 4, // Tab Id
            "name": "xxx" // Tab 名称            
        },
    ]
}
```

## 查询一个 Tab 的商品 GET /product/tabProducts

请求：

```json
{
    "tabId": 3 // 要查询的 tab 的 tab id
}
```

响应：

```json
{
    "products": [
        {
            "id": 2, // 商品 id
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        },
        {
            "id": 3, // 商品 id
            "price": 333, // 价格
            "title": "过期牛小排100g", // 标题
            "description": "过期的牛小排。", // 描述
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        }
    ]
}
```

## *修改一个 Tab POST /product/tabModify

请求：

```json
{
    "tabId": 4, // 要修改的 tabId
    "name": "hh", // 修改后的名字
    "products": [1, 2, 4, 8] // 修改后 Tab 包含的商品
}
```

响应：无

## **增加一个 Tab POST /product/tabAdd

请求：

```json
{
    "name": "hh", // 修改后的名字
    "products": [1, 3, 4, 6] // Tab 包含的商品
}
```

响应：无

## **删除一个 Tab POST /product/tabDelete

请求：

```json
{
    "tabId": 4 // 要删除的 tabId
}
```

响应：无

## **修改首页推荐商品列表 POST /product/modifyHomeTab

## **修改首页推荐图片个数 POST /product/modifyHomePicture

## 查询首页推荐图片个数 GET /product/homePicture

## 查询首页推荐商品列表 GET /product/homeTab

请求：无

响应：

```json
{
    products: [
        {
            "id": 2, // 商品 id
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        },
        {
            "id": 3, // 商品 id
            "price": 333, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 4, // 可选，详情页幻灯片图片页数
            "ownerUserId": 4, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false // 是否下架
        }
    ]
}
```

# 订单 API

## *结算购物车中的部分商品 POST /order/generateOrder

给出商品 id，结算并生成多个订单

请求：

```json
{
    "product_ids": [1, 3, 5, 7, 8], // 需要结算的 商品 id
    "transportation_price": 5.5, // 运费
    "name": "hhh", // 收件人名字
    "phone_number": "18089876534", // 收件人电话
    "address": "广东省深圳市南山区", // 收件人地址
    "message": "快点发货" // 用户留言
}
```

响应：无

## *查询用户所有订单 GET /order/userOrder

请求：

```json
{
	"from": 0, // 开始 index
	"length": 10, // 取多少个
    "status": 1 // 查询什么状态的
} // [from, from + length) 的订单将被返回
```

响应：

```json
{
    "orders": [
        {
            "order_id": 1, // order id
            "owner_user_id": 3, // 哪个用户定的订单
            "owner_shop_user_id": 4, // 哪个商家接的订单
            "now_status": 2, // 现在的状态，待支付1、待收货2、待评价3、已完成4
            "total_price": 433.2, // 最后的总价
            "transportation_price": 30.1, // 包括的运费
            "discount_price": 10.3, // 减少的打折费
            "product_and_count": [ // 商品和数量
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 3 // 数量
                },
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 6
                }
            ],
            "person_name": "name", // 收件人名字
            "person_phone_number": "13877766543", // 收件人电话
            "person_address": "address", // 收件人地址
            "order_time": 1531292871, // 下单时间
            "pay_time": 1531292871, // 支付时间，Unix 时间戳
            "verify_time": 1531292871, // 确定收货时间，Unix 时间戳
            "tracking_number": [ // 快递号
                "75844515448966",
                "75234515448968",
                "75584515448966"
            ],
            "message": "快点发货", // 留言
            "user_phone_number": "13876545678", // 下单人手机号
            "shop_message": "商家退单" // 商户消息
        },
        {
            "order_id": 2, // order id
            "owner_user_id": 3, // 哪个用户定的订单
            "owner_shop_user_id": 4, // 哪个商家接的订单
            "now_status": 2, // 现在的状态，待支付1、待收货2、待评价3、已完成4
            "total_price": 433.2, // 最后的总价
            "transportation_price": 30.1, // 包括的运费
            "discount_price": 10.3, // 减少的打折费
            "product_and_count": [ // 商品和数量
                {
                    "product_id": 4, // 商品 id
                    "count": 3 // 数量
                },
                {
                    "product_id": 4,
                    "count": 6
                }
            ],
            "person_name": "name", // 收件人名字
            "person_phone_number": "13877766543", // 收件人电话
            "person_address": "address", // 收件人地址
            "order_time": 1531292871, // 下单时间
            "pay_time": 1531292871, // 支付时间，Unix 时间戳
            "verify_time": 1531292871, // 确定收货时间，Unix 时间戳
            "tracking_number": [ // 快递号
                "75844515448966",
                "75234515448968",
                "75584515448966"
            ],
            "message": "快点发货", // 留言
            "user_phone_number": "13876545678", // 下单人手机号
            "shop_message": "商家退单" // 商户消息
        }
    ]
}
```

## *根据订单 ID 查询一个订单信息 GET /order/orderInfo

请求：

```json
{
    order_id: 4 // 查询的订单 id
}
```

响应：

```json
{
    "order_id": 1, // order id
	"owner_user_id": 3, // 哪个用户定的订单
	"owner_shop_user_id": 4, // 哪个商家接的订单
	"now_status": 2, // 现在的状态，待支付1、待收货2、待评价3、已完成4
	"total_price": 433.2, // 最后的总价
	"transportation_price": 30.1, // 包括的运费
	"discount_price": 10.3, // 减少的打折费
	"product_and_count": [ // 商品和数量
        {
            "products": {
                // 此处与商品的相同  
            },
            "count": 3 // 数量
        },
        {
            "products": {
                // 此处与商品的相同  
            },
            "count": 6
        }
    ],
	"person_name": "name", // 收件人名字
	"person_phone_number": "13877766543", // 收件人电话
	"person_address": "address", // 收件人地址
    "order_time": 1531292871, // 下单时间
	"pay_time": 1531292871, // 支付时间，Unix 时间戳
	"verify_time": 1531292871, // 确定收货时间，Unix 时间戳
	"tracking_number": [ // 快递号
        "75844515448966",
        "75234515448968",
        "75584515448966"
    ],
	"message": "快点发货", // 留言
    "user_phone_number": "13876545678", // 下单人手机号
    "shop_message": "商家退单" // 商户消息
}
```

## *查询某个商户的所有订单 GET /order/shopOrder

shop 端调用

请求：

```json
{
	"from": 0, // 开始 index
	"length": 10, // 取多少个
    "status": 1 // 查询什么状态的
} // [from, from + length) 的订单将被返回
```

响应：

```json
{
    "orders": [
        {
            "order_id": 1, // order id
            "owner_user_id": 3, // 哪个用户定的订单
            "owner_shop_user_id": 4, // 哪个商家接的订单
            "now_status": 2, // 现在的状态，待支付1、待收货2、待评价3、已完成4
            "total_price": 433.2, // 最后的总价
            "transportation_price": 30.1, // 包括的运费
            "discount_price": 10.3, // 减少的打折费
            "product_and_count": [ // 商品和数量
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 3 // 数量
                },
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 6
                }
            ],
            "person_name": "name", // 收件人名字
            "person_phone_number", "13877766543", // 收件人电话
            "person_address": "address", // 收件人地址
            "order_time": 1531292871, // 下单时间
            "pay_time": 1531292871, // 支付时间，Unix 时间戳
            "verify_time": 1531292871, // 确定收货时间，Unix 时间戳
            "tracking_number": [ // 快递号
                "75844515448966",
                "75234515448968",
                "75584515448966"
            ],
            "message": "快点发货", // 留言
            "user_phone_number": "13876545678", // 下单人手机号
            "shop_message": "商家退单" // 商户消息
        },
        {
            "order_id": 3, // order id
            "owner_user_id": 3, // 哪个用户定的订单
            "owner_shop_user_id": 4, // 哪个商家接的订单
            "now_status": 2, // 现在的状态，待支付1、待收货2、待评价3、已完成4
            "total_price": 433.2, // 最后的总价
            "transportation_price": 30.1, // 包括的运费
            "discount_price": 10.3, // 减少的打折费
            "product_and_count": [ // 商品和数量
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 3 // 数量
                },
                {
                    "products": {
                      // 此处与商品的相同  
                    },
                    "count": 6
                }
            ],
            "person_name": "name", // 收件人名字
            "person_phone_number", "13877766543", // 收件人电话
            "person_address": "address", // 收件人地址
            "order_time": 1531292871, // 下单时间
            "pay_time": 1531292871, // 支付时间，Unix 时间戳
            "verify_time": 1531292871, // 确定收货时间，Unix 时间戳
            "tracking_number": [ // 快递号
                "75844515448966",
                "75234515448968",
                "75584515448966"
            ],
            "message": "快点发货", // 留言
            "user_phone_number": "13876545678", // 下单人手机号
            "shop_message": "商家退单" // 商户消息
        }
    ]
}
```

## *删除订单 POST /order/deleteOrder

请求：

```json
{
    "order_id": 4 // 要删除的订单 id
}
```

响应：无

## *更新订单状态 POST /order/changeStatus

shop 端调用

请求：

```json
{
    "order_id": 4, // 要更新的订单 id
    "status": 2, // 新的 status
    "pay_time": 1531292871, // 可选，按需调用修改，Unix 时间戳
    "verify_time": 1531292871, // 可选，按需调用修改，Unix 时间戳
    "shop_message": "商户退单" // 可选，按需调用修改，商户消息，不修改
    // 可选项不修改就不要这个字段
}
```

响应：无

## *增加订单物流号 POST /order/addTrackingNumber

shop 端调用

请求：

```json
{
    "order_id": 4, // 增加物流编号的订单 id
    "tracking_number": [
        "75844515448966",
        "75234515448968",
        "75584515448966"
    ] // 物流编号
}
```

响应：无

# 评价 API

## *发表评价 POST /comment/add

请求：

```json
{
    "comments": [
        {
            "product_id": 3, // 商品 ID
    		"comment": "很香很甜" // 评价内容
        },
        {
            "product_id": 5, // 商品 ID
    		"comment": "很好" // 评价内容
        }
    ]
}
```

响应：无

## 查看商品评价 GET /comment/product

请求：

```json
{
	"product_id": 4 // 要查询评价的商品 Id
}
```

响应：

```json
{
	"comments": [
        {
            "user_id": 4, // 评价用户
            "comment": "哈哈" // 评价内容
        },
        {
            "user_id": 3, // 评价用户
            "comment": "你好" // 评价内容
        }
    ]
}
```

## *删除评价 POST /comment/delete

请求：

```json
{
	"comment_id": 4 // 要删除的评价的 Id
}
```

响应：无

# 搜索 API

## 搜索商品标题 POST /search

请求：

```json
{
	"key": "鸡" // 要查询的关键字
}
```

响应：

```json
{
    "Products": [
        {
            "id": 2,
            "price": 233, // 价格
            "title": "鸡小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false, // 是否下架
            "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    		"detailPictureNumber": 3 // 可选，详情信息处的图片页数
        },
        {
            "id": 4,
            "price": 233, // 价格
            "title": "鸡小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2, // 谁的 product，为 user id
            "stock": 32, // 库存
            "isDrop": false, // 是否下架
            "detail": "新鲜的天的猪小排，哈哈哈哈哈哈", // 详情信息，描述
    		"detailPictureNumber": 3 // 可选，详情信息处的图片页数
        }
    ]
}
```

 
