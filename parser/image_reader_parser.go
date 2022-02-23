package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"google.golang.org/protobuf/proto"
)

func ImageReaderParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.ImageReaderParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	dom, root, global, err := InitParser(rpc, parser.ExtraSelector)

	// 基础信息
	model := &protobuf.ImageReaderRpcModel{
		Image:          ImageParser(dom, parser.Image, root),
		LargerImageUrl: dom.String(parser.LargerImage, root),
		RawImageUrl:    dom.String(parser.RawImage, root),
		UploadTime:     dom.String(parser.UploadTime, root),
		Source:         dom.String(parser.Source, root),
		Rating:         dom.String(parser.Rating, root),
		Score:          dom.String(parser.Score, root),
		Badges: utils.Map(dom.Nodes(parser.BadgeSelector, root), func(e any) *protobuf.TagRpcModel {
			return &protobuf.TagRpcModel{
				Text:     dom.String(parser.BadgeText, e),
				Category: dom.String(parser.BadgeCategory, e),
			}
		}),
		LocalEnv:  dom.Env,
		GlobalEnv: global,
	}

	marshal, err := proto.Marshal(model)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
