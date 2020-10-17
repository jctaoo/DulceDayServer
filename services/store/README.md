# store - 商店(i.e. 内购的后台支持)模块
---
## 运行模式:
非该程序的模块定期从 AppStoreConnect 获取商品信息并写入持久化数据库
该程序以持久化数据库中的数据为准

## 内购项目:
- 开启宽限期
### D币:
  1. 600D币:
    - 参考名称: 6rmb的D币-600D
    - Identifier: com.jctaoo.DulceDay.600dcoin
    - 显示名称: 600枚D币
    - 描述: 600枚用于app内礼物购买等消费的D币
    - 价格: ¥6.00
  2. 1800D币:
    - 参考名称: 18rmb的D币-1800D
    - Identifier: com.jctaoo.DulceDay.1800dcoin
    - 显示名称: 1800枚D币
    - 描述: 1800枚用于app内礼物购买等消费的D币
    - 价格: ¥18.00
   3. 6800D币:
     - 参考名称: 68rmb的D币-6800D
     - Identifier: com.jctaoo.DulceDay.6800dcoin
     - 显示名称: 6800枚D币
     - 描述: 6800枚用于app内礼物购买等消费的D币
     - 价格: ¥68.00
  4. 15800D币:
    - 参考名称: 158rmb的D币-15800D
    - Identifier: com.jctaoo.DulceDay.15800dcoin
    - 显示名称: 15800枚D币
    - 描述: 15800枚用于app内礼物购买等消费的D币
    - 价格: ¥158.00
  5. 23300D币:
    - 参考名称: 233rmb的D币-23300D
    - Identifier: com.jctaoo.DulceDay.23300dcoin
    - 显示名称:  23300枚D币
    - 描述:  23300枚用于app内礼物购买等消费的D币
    - 价格: ¥233.00
  6. 38800D币:
    - 参考名称: 388rmb的D币-38800D
    - Identifier: com.jctaoo.DulceDay.38800dcoin
    - 显示名称:  38800枚D币
    - 描述:  38800枚用于app内礼物购买等消费的D币
    - 价格: ¥388.00
### 会员:
  1. 月度会员：
    - 参考名称: 1月度会员
    - Identifier: com.jctaoo.DulceDay.onemonthvip
    - 显示名称:  月度会员
    - 价格: ¥25.00
  2. 季度会员：
    - 参考名称: 1季度会员
    - Identifier: com.jctaoo.DulceDay.onequartervip
    - 显示名称: 季度会员
    - 价格: ¥68.00
  3. 年度会员：
    - 参考名称: 1年度会员
    - Identifier: com.jctaoo.DulceDay.oneyearvip
    - 显示名称: 年度会员
    - 价格: ¥168.00
### 会员(自动续期):
  - 订阅组: 自动续期会员
  1. 包月会员：
    - 参考名称: 包月会员
    - Identifier: com.jctaoo.DulceDay.c.monthvip
    - 显示名称:  连续包月
    - 价格: ¥15.00
  2. 季度会员：
    - 参考名称: 包季会员
    - Identifier: com.jctaoo.DulceDay.c.quartervip
    - 显示名称: 连续包季
    - 价格: ¥45.00
  3. 年度会员：
    - 参考名称: 包年会员
    - Identifier: com.jctaoo.DulceDay.c.yearvip
    - 显示名称: 连续包年
    - 价格: ¥148.00

## 子模块
- PurchaseProvider: 管理 Moment
- Repository(StoreRepository): 管理该模块的持久化数据（因为Store在此处有歧义，所以换成 Repository）
