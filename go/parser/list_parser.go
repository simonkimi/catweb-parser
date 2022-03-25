package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"google.golang.org/protobuf/proto"
)

func ListParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.ListViewParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	dom, root, global, err := InitParser(rpc, parser.ExtraSelector)

	model := &protobuf.ListRpcModel{
		Items: utils.Map(dom.Nodes(parser.ItemSelector, root), func(e any) *protobuf.ListRpcModel_Item {
			item := &protobuf.ListRpcModel_Item{
				Title:      dom.String(parser.Title, e),
				Subtitle:   dom.String(parser.Subtitle, e),
				ImgCount:   dom.Int(parser.ImgCount, e),
				Paper:      dom.String(parser.Paper, e),
				Star:       dom.Double(parser.Star, e),
				UploadTime: dom.String(parser.UploadTime, e),
				Language:   dom.String(parser.Language, e),
				Tag: &protobuf.TagRpcModel{
					Text:  dom.String(parser.Tag, e),
					Color: dom.Color(parser.TagColor, e),
				},
				Badges: utils.Map(dom.Nodes(parser.BadgeSelector, e), func(e any) *protobuf.TagRpcModel {
					return &protobuf.TagRpcModel{
						Text:  dom.String(parser.BadgeText, e),
						Color: dom.Color(parser.BadgeColor, e),
					}
				}),
				PreviewImg: ImageParser(dom, parser.PreviewImg, e),
				Target:     dom.String(parser.Target, e),
				NextPage:   dom.String(parser.NextPage, e),
				Env:        dom.LocalEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_LIST_ITEM, e),
			}
			item.Env["idCode"] = dom.String(parser.IdCode, e)
			return item
		}),
		NextPage:  dom.String(parser.NextPage, root),
		GlobalEnv: global,
		LocalEnv:  dom.Env,
	}

	marshal, err := proto.Marshal(model)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
