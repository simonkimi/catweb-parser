package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"google.golang.org/protobuf/proto"
)

func AutoCompleteParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.AutoCompleteParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	dom, root, global, err := InitParser(rpc, parser.ExtraSelector)

	// 基础信息
	model := &protobuf.AutoCompleteRpcModel{
		Items: utils.Map(dom.Nodes(parser.ItemSelector, root), func(e any) *protobuf.AutoCompleteRpcModel_Item {
			return &protobuf.AutoCompleteRpcModel_Item{
				Title:    dom.String(parser.ItemTitle, e),
				Subtitle: dom.String(parser.ItemSubtitle, e),
				Complete: dom.String(parser.ItemComplete, e),
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
