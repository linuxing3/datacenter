# jwt

## 如何进行jwt校验

jwt网络托证token的签发流程


![jwt](/assets/jwt.drawio.svg)

+ 创建声明claims的哈希映射，截止日期，用户id等信息

+ 创建HS256签发的基本token

+ 将声明claims哈希映射添加到token

+ 将秘钥secretKey转化为字节

+ 使用secretKey字节码，将token转化为签名字符串

```go
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))