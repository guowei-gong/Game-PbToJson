package pbToJson

import "testing"

func TestGetRoutes(t *testing.T) {
	GetProtoMessageToJSON("pb 文件夹的绝对路径, 例如: /Users/guowei/go/src/你的项目/pb/v1")
}
