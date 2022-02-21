package parser

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
	"github.com/simonkimi/catweb-parser/utils"
	"golang.org/x/net/html"
	"google.golang.org/protobuf/proto"
	"strings"
)

func ImageReaderParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.ImageReaderParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(rpc.Data, "{") || strings.HasPrefix(rpc.Data, "[") {
		// TODO Json
		return nil, errors.New("TODO: Json")
	} else {
		dom := &selector.DomSelector{Env: rpc.Env}

		globalEnv := make(map[string]string)

		root, err := htmlquery.Parse(strings.NewReader(rpc.Data))
		if err != nil {
			return nil, err
		}

		// 开始解析
		local, global := dom.GetEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_NONE, root)
		MergeEnv(globalEnv, dom.Env, global, local)

		// 基础信息
		model := &protobuf.ImageReaderRpcModel{
			Image:          ImageParser(dom, parser.Image, root),
			LargerImageUrl: dom.String(parser.LargerImage, root),
			RawImageUrl:    dom.String(parser.RawImage, root),
			UploadTime:     dom.String(parser.UploadTime, root),
			Source:         dom.String(parser.Source, root),
			Rating:         dom.String(parser.Rating, root),
			Score:          dom.String(parser.Score, root),
			Badges: utils.Map(dom.Nodes(parser.BadgeSelector, root), func(e *html.Node) *protobuf.TagRpcModel {
				return &protobuf.TagRpcModel{
					Text:     dom.String(parser.BadgeText, e),
					Category: dom.String(parser.BadgeCategory, e),
				}
			}),
			LocalEnv:  dom.Env,
			GlobalEnv: globalEnv,
		}

		marshal, err := proto.Marshal(model)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}
