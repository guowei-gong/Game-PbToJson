package pbToJson

import (
	"encoding/json"
	"github.com/jhump/protoreflect/desc/protoparse"
	"log"
	"os"
	"strings"
)

func GetProtoMessageToJSON(dir string) {
	// 加载
	if len(dir) == 0 {
		dir = "."
	}

	filesInfos, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("filepath glob err: %v", err)
		return
	}

	var files []string
	for _, fileInfo := range filesInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".proto") {
			files = append(files, fileInfo.Name())
		}
	}

	p := protoparse.Parser{}
	descriptors, err := p.ParseFiles(files...)
	if err != nil {
		log.Fatalf("parse files err: %v", err)
		return
	}

	var routes []string
	// 解析
	for _, descriptor := range descriptors {
		for _, t := range descriptor.GetMessageTypes() {

			message := t.GetName()

			if strings.HasSuffix(message, "Req") {
				routes = append(routes, t.GetName())
			}

			if strings.HasSuffix(message, "Res") {
				routes = append(routes, t.GetName())
			}

		}
	}

	// 写入 JSON 文件
	jsonData, err := json.MarshalIndent(routes, "", " ")
	if err != nil {
		log.Fatalf("encoding to JSON err: %v", err)
	}

	err = os.WriteFile("gen_file.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("writing to file err: %v", err)
	}

}
