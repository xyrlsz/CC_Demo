## 初始化命令

```bash
go mod init openccapi
mkdir rpc/opencc
cd rpc/opencc
kitex -module openccapi -service opencc.convertor -I ../../  ../../idl/opencc_service.proto
```


 