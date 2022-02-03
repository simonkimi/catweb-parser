package parser

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
	"golang.org/x/net/html"
	"google.golang.org/protobuf/proto"
	"strings"
)

func ListParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.ListViewParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(rpc.Data, "{") || strings.HasPrefix(rpc.Data, "[") {
		// TODO Json
		return nil, errors.New("TODO: Json")
	} else {
		dom := &selector.DomSelector{Env: rpc.Env}
		model := &protobuf.ListRpcModel{}
		globalEnv := make(map[string]string)

		root, err := htmlquery.Parse(strings.NewReader(rpc.Data))
		if err != nil {
			return nil, err
		}

		// 开始解析
		local, global := dom.GetEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_NONE, root)
		for k, v := range global {
			globalEnv[k] = v
			dom.Env[k] = v
		}
		for k, v := range local {
			dom.Env[k] = v
		}

		model.Items = DomMap(dom.Nodes(parser.ItemSelector, root), func(e *html.Node) *protobuf.ListRpcModel_Item {
			return &protobuf.ListRpcModel_Item{
				//IdCode:     dom.String(parser.IdCode, e),
				//Title:      dom.String(parser.Title, e),
				//Subtitle:   dom.String(parser.Subtitle, e),
				//ImgCount:   dom.Int(parser.ImgCount, e),
				//Paper:      dom.String(parser.Paper, e),
				Star: dom.Double(parser.Star, e),
				//UploadTime: dom.String(parser.UploadTime, e),
				//Tag: &protobuf.ListRpcModel_Tag{
				//	Text:  dom.String(parser.Tag, e),
				//	Color: dom.Color(parser.Tag, e),
				//},
				//Badges: DomMap(dom.Nodes(parser.BadgeSelector, e), func(e *html.Node) *protobuf.ListRpcModel_Tag {
				//	return &protobuf.ListRpcModel_Tag{
				//		Text:  dom.String(parser.BadgeText, e),
				//		Color: dom.Color(parser.BadgeColor, e),
				//	}
				//}),
				//PreviewImg: ImageParser(dom, parser.PreviewImg, e),
				//NextPage:   dom.String(parser.NextPage, e),
				Env: dom.LocalEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_LIST_ITEM, e),
			}
		})

		model.NextPage = dom.String(parser.NextPage, root)
		model.GlobalEnv = globalEnv
		model.LocalEnv = dom.Env

		marshal, err := proto.Marshal(model)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}
