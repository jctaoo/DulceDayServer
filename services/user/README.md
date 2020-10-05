# user - 鉴权与授权模块
---
## 运行模式
- 该系统的 Token 字符串采用 JWT 生成与验证。
- 注册：往 User 表中写入相应记录
- 登录/鉴权：根据要登陆的用户 (包括设备名，IP)
  1. 检验 UserId 字符串是否在 BlackList 中
  2. 检查密码是否正确
  3. 查看是否有相应的 Token 记录，如果有，就更新 Token 并返回新的 Token 字符串和 IP，否则写入新的 Token 记录
- 授权：
  1. 检验 Token 字符串是否在 RevokeTokens 中
  2. 检验 Token 字符串是否在 InActiveTokens 中
  3. JWT 检查 Token 字符串
  4. 检验 UserId 字符串是否在 BlackList 中
- 将 User 列入黑名单：
  1. 将 UserID 存入 BlackList
- 撤回 Token:
  > 出于安全问题，用户可能会撤回某些 Token 以强制其他客户端下线
  1. 将 Token 字符串存入 RevokeTokens 中
  2. 删除相应 Token 记录 (软删)

## 持久化数据库表
- User: 用于存放用于用户登录的信息，账户密码，权限，VIP等(区别与用户信息)
- Token: 存放用户的鉴权信息

## 缓存数据库列表
- RevokeTokens: 存放撤销的 Token 字符串（因为单纯 JWT 无法撤回，即使删掉持久化数据库中的记录也没用）
- InActiveTokens: 存放未激活的 Token 字符串 (可用于强制邮箱激活或二维码登录等功能)
- BlackList: 存放被该模块列入黑名单的 UserId
- IPBlackList: 存放被列入黑名单的 IP

## 子模块
- Service(UserService): 管理 User
- TokenGranter: 管理 Token
- TokenAdaptor: Token 的无状态相关，如生成 Token 字符串，目前采用 JWT
- EncryptionAdaptor
- TokenStore
- UserStore
