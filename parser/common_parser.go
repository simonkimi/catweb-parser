package parser

import (
	"github.com/antchfx/htmlquery"
	"github.com/ohler55/ojg/oj"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
	"strings"
)

func ImageParser(dom *selector.DomSelector, parser *protobuf.ImageSelector, node any) *protobuf.ImageRpcModel {
	if parser == nil {
		return nil
	}
	return &protobuf.ImageRpcModel{
		Url:    dom.String(parser.ImgUrl, node),
		Width:  dom.Double(parser.ImgWidth, node),
		Height: dom.Double(parser.ImgHeight, node),
		ImgX:   dom.Double(parser.ImgX, node),
		ImgY:   dom.Double(parser.ImgY, node),
	}
}

func MergeEnv(global map[string]string, local map[string]string, newGlobal map[string]string, newLocal map[string]string) {
	for k, v := range newGlobal {
		global[k] = v
		local[k] = v
	}
	for k, v := range newLocal {
		local[k] = v
	}
}

func InitParser(rpc *protobuf.RpcRequest, extraSelector []*protobuf.ExtraSelector) (dom *selector.DomSelector, root any, global map[string]string, err error) {
	globalEnv := make(map[string]string)
	dom = &selector.DomSelector{Env: rpc.Env}
	if strings.HasPrefix(rpc.Data, "{") || strings.HasPrefix(rpc.Data, "[") {
		root, err = oj.ParseString(rpc.Data)
	} else {
		root, err = htmlquery.Parse(strings.NewReader(rpc.Data))
	}
	if err != nil {
		return
	}
	// 开始解析
	local, global := dom.GetEnv(extraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_NONE, root)
	MergeEnv(globalEnv, dom.Env, global, local)
	return
}
