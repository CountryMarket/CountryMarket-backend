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
    "user_id": 3, // 要修改的用户的 ID
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

# 商户 API

每个 product 有一个自增 ID

## *增加一个商品 POST /shop/addProduct

(需要是商户账号，即permission有2)

请求：

```json
{
    "price": 233, // 价格
    "title": "过期猪小排100g", // 标题
    "description": "过期的猪小排。", // 描述
    "pictureNumber": 2 // 可选，详情页幻灯片图片页数
}
```

响应：无

## *更新一个商品 PUT /shop/updateProduct

(需要是商户账号，即permission有2，且更新的商品是申请用户的)

请求：

```json
{
    "price": 233, // 价格
    "title": "过期猪小排100g", // 标题
    "description": "过期的猪小排。", // 描述
    "pictureNumber": 2, // 可选，详情页幻灯片图片页数
    "id": 2 // 要改的商品 id
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
    "ownerUserId": 2 // 谁的 product，为 user id
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
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2 // 谁的 product，为 user id
        },
        {
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "pictureNumber": 2, // 可选，详情页幻灯片图片页数
            "ownerUserId": 2 // 谁的 product，为 user id
        }
    ]
}
```

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
            "ownerUserId": 2 // 谁的 product，为 user id
        },
        {
            "id": 4, // 商品 id
            "count": 4, // 加入的数量
            "price": 233, // 价格
            "title": "过期猪小排100g", // 标题
            "description": "过期的猪小排。", // 描述
            "ownerUserId": 2 // 谁的 product，为 user id
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

## *结算购物车中的部分商品 POST /cart/buyProduct

给出商品 id，结算并生成订单

请求：

```json
{
    "productIds": [1, 3, 5, 7, 8] // 需要结算的 商品 id
}
```

响应：

```json
{
    "idOrder": 3 // 生成的订单 id
}
```

# 商品 API

## 查询 Tab List GET /shop/tabList

## 查询一个 Tab 的商品 GET /shop/tabProducts

## **修改一个 Tab POST /shop/tabModify

## **增加一个 Tab POST /shop/tabAdd

## **隐藏 / 显示一个 Tab POST /shop/tabXor

## **修改首页推荐商品列表 POST /shop/modifyHomeTab

## **修改首页推荐图片个数 POST /shop/modifyHomePicture

## 查询首页推荐图片个数 GET /shop/homePicture

## 查询首页推荐商品列表 GET /shop/homeTab


# 订单 API