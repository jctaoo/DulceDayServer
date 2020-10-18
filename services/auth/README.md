# auth - 鉴权与授权模块
---
## 运行模式
- 该系统的 Token 字符串采用 JWT 生成与验证。
- 注册：往 AuthUser 表中写入相应记录
- 登录/鉴权：根据要登陆的用户 (包括设备名，IP)
  1. 检验 AuthUserId 字符串是否在 BlackList 中
  2. 检查密码是否正确
  3. 查看是否有相应的 Token 记录，如果有，就将原 Token 放入 Revoke 列表, 更新 Token 并返回新的 Token 字符串和 IP，否则写入新的 Token 记录
- 授权：
  1. 检验 Token 字符串是否在 RevokeTokens 中
  2. 检验 Token 字符串是否在 InActiveTokens 中
  3. JWT 检查 Token 字符串
  4. 检验 AuthUserId 字符串是否在 BlackList 中
- 将 AuthUser 列入黑名单：
  1. 将 AuthUserID 存入 BlackList
- 撤回 Token:
  > 出于安全问题，用户可能会撤回某些 Token 以强制其他客户端下线
  1. 将 Token 字符串存入 RevokeTokens 中
  2. 删除相应 Token 记录 (软删)
- 敏感信息 Token (人机认证同理):
  > 当某个请求需要此类 Token 则不需要一般授权的 Token
  1. 生成 TokenAuth 存入持久化数据库，其中 TokenStr 中放入是否属于敏感验证的标识
  2. 将 TokenStr 记录进 InActiveTokens 中
  3. 生成验证码，验证方式(邮箱等)为键，验证码为值，当客户端传递的验证方式和验证码正确时，返回 1 生成的 TokenStr
  4. 验证成功后从 InActiveTokens 中移除，若不成功 Revoke 相应 Token
  5. 删除验证码键值
  > 疑问：TokenAuth 如何与验证码关联，在缓存数据库中的值中存 ‘Code:TokenStr’，也可以在 1 的时候返回 TokenStr，这样在 4 的时候就不要从 InActiveTokens 中移除
  > 验证时校验 IP 和 deviceName
  > 注意常规检查黑名单等列表

## 持久化数据库表
- AuthUser: 用于存放用于用户登录的信息，账户密码，权限，VIP等(区别与用户信息)
- Token: 存放用户的鉴权信息

## 缓存数据库列表
- RevokeTokens: 存放撤销的 Token 字符串（因为单纯 JWT 无法撤回，即使删掉持久化数据库中的记录也没用）
- InActiveTokens: 存放未激活的 Token 字符串 (可用于强制邮箱激活或二维码登录等功能)
- BlackList: 存放被该模块列入黑名单的 AuthUserId
- IPBlackList: 存放被列入黑名单的 IP
- VerificationCode: 验证码键值对，其中值同时存放验证码和对应的 TokenStr

## 子模块
- Service(AuthUser): 管理 AuthUser
- TokenGranter: 管理 Token
- TokenAdaptor: Token 的无状态相关，如生成 Token 字符串，目前采用 JWT
- EncryptionAdaptor
- TokenStore
- AuthUserStore
