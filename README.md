# 鉴权责任链

CLI → Chain → 多个 Auth Handler；任一步失败应返回错误。

## 测试

```bash
set GOTOOLCHAIN=local
go test ./... -count=1
```
