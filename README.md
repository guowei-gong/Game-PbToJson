## GamePbToJson

将 proto 文件中, 所有以 `Req` 和 `Res` 结尾的 Message 名称提取到 Json 文件中。

提取到 Json 是为了避免文件路径的问题, 任何使用该 Json 的场景只需要定位 Json 文件路径即可。

### Example

输入 proto 文件:
```proto
syntax = "proto3";

package route;

option go_package = "daycare_go/internal/pb/v1/route";

message HelloWorld {
}

message HelloWorldRequest {
}

message HelloWorldResponse {
}
```

输出 json 文件:
```json
[
  "HelloWorldRequest",
  "HelloWorldResponse"
]
```
