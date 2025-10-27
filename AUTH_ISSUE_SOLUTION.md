# JWT 401认证问题解决方案

## 问题根源

经过深入分析，401错误的真正原因是：

1. **后端使用内存存储token**：服务端将生成的token保存在内存中（`tokenStore`）
2. **容器重启导致token丢失**：当Docker容器重启后，内存中的token全部丢失
3. **前端token无效**：虽然前端localStorage中还有token，但后端已经不认识了

## 临时解决方案

### 每次部署后重新登录
```bash
# 1. 清除浏览器缓存
Ctrl + Shift + Delete

# 2. 重新访问登录页
http://localhost:8080/login.html

# 3. 使用admin/admin123登录
```

### 测试认证流程
```bash
test-auth-flow.bat
```

## 永久解决方案

### 方案1：使用真正的JWT（推荐）

修改 `service/auth_service.go` 使用JWT库：

```go
import "github.com/dgrijalva/jwt-go"

func generateToken(userID, username string) string {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    
    tokenString, _ := token.SignedString([]byte(jwtSecret))
    return tokenString
}

func ValidateToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtSecret), nil
    })
    
    if err != nil {
        return "", err
    }
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["user_id"].(string), nil
    }
    
    return "", fmt.Errorf("invalid token")
}
```

### 方案2：使用Redis存储token

1. 在docker-compose中添加Redis：
```yaml
redis:
  image: redis:alpine
  container_name: at-mmbs-redis
  ports:
    - "6379:6379"
  networks:
    - mmbs-network
```

2. 修改token存储逻辑使用Redis

### 方案3：使用数据库存储token

创建token表并将token持久化到数据库。

## 快速验证

1. **检查容器是否重启过**：
   ```bash
   docker logs at-mmbs-app | grep "starting"
   ```

2. **验证当前token状态**：
   ```bash
   # 在浏览器控制台执行
   localStorage.getItem('authToken')
   ```

3. **手动测试API**：
   ```bash
   # 替换YOUR_TOKEN为实际token
   curl -H "Authorization: Bearer YOUR_TOKEN" http://localhost:8080/api/users/profile
   ```

## 注意事项

- 目前的实现是一个简化版本，不适合生产环境
- 每次容器重启都需要重新登录
- 建议实施真正的JWT方案以彻底解决此问题