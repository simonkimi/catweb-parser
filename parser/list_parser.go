package parser

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
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
		var items []*protobuf.ListRpcModel_Item
		local, global := dom.GetEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_NONE, root)
		for k, v := range global {
			globalEnv[k] = v
			dom.Env[k] = v
		}
		for k, v := range local {
			dom.Env[k] = v
		}

		for _, e := range dom.Nodes(parser.ItemSelector, root) {
			var badges []*protobuf.ListRpcModel_Tag
			for _, e := range dom.Nodes(parser.BadgeSelector, e) {
				badges = append(badges, &protobuf.ListRpcModel_Tag{
					Text:  dom.String(parser.BadgeText, e),
					Color: dom.Color(parser.BadgeColor, e),
				})
			}

			items = append(items, &protobuf.ListRpcModel_Item{
				IdCode:     dom.String(parser.IdCode, e),
				Title:      dom.String(parser.Title, e),
				Subtitle:   dom.String(parser.Subtitle, e),
				ImgCount:   dom.Int(parser.ImgCount, e),
				Paper:      dom.String(parser.Paper, e),
				Star:       dom.Double(parser.Star, e),
				UploadTime: dom.String(parser.UploadTime, e),
				Tag: &protobuf.ListRpcModel_Tag{
					Text:  dom.String(parser.Tag, e),
					Color: dom.Color(parser.Tag, e),
				},
				Badges: badges,
				PreviewImg: &protobuf.ImageRpcModel{
					Url:    dom.String(parser.PreviewImg.ImgUrl, e),
					Width:  dom.Double(parser.PreviewImg.ImgWidth, e),
					Height: dom.Double(parser.PreviewImg.ImgHeight, e),
					ImgX:   dom.Double(parser.PreviewImg.ImgX, e),
					ImgY:   dom.Double(parser.PreviewImg.ImgY, e),
				},
				NextPage: dom.String(parser.NextPage, e),
				Env:      dom.LocalEnv(parser.ExtraSelector, protobuf.ExtraSelectorType_EXTRA_SELECTOR_TYPE_LIST_ITEM, e),
			})
		}
		model.Items = items
		model.GlobalEnv = globalEnv
		model.LocalEnv = dom.Env

		marshal, err := proto.Marshal(model)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}
