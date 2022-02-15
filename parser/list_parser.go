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

		globalEnv := make(map[string]string)

		root, err := htmlquery.Parse(strings.NewReader(rpc.Data))
		if err != nil {
			return nil, err
		}

		// 开始解析
		local, global := dom.GetEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_NONE, root)
		MergeEnv(globalEnv, dom.Env, global, local)

		model := &protobuf.ListRpcModel{
			Items: utils.Map(dom.Nodes(parser.ItemSelector, root), func(e *html.Node) *protobuf.ListRpcModel_Item {
				return &protobuf.ListRpcModel_Item{
					IdCode:     dom.String(parser.IdCode, e),
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
					Badges: utils.Map(dom.Nodes(parser.BadgeSelector, e), func(e *html.Node) *protobuf.TagRpcModel {
						return &protobuf.TagRpcModel{
							Text:  dom.String(parser.BadgeText, e),
							Color: dom.Color(parser.BadgeColor, e),
						}
					}),
					PreviewImg: ImageParser(dom, parser.PreviewImg, e),
					NextPage:   dom.String(parser.NextPage, e),
					Env:        dom.LocalEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_LIST_ITEM, e),
				}
			}),
			NextPage:  dom.String(parser.NextPage, root),
			GlobalEnv: globalEnv,
			LocalEnv:  dom.Env,
		}

		marshal, err := proto.Marshal(model)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}
