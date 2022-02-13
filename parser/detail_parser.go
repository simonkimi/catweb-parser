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

func DetailParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.DetailParser{}
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
		model := &protobuf.DetailRpcModel{
			Title:        dom.String(parser.Title, root),
			Subtitle:     dom.String(parser.Subtitle, root),
			UploadTime:   dom.String(parser.UploadTime, root),
			Language:     dom.String(parser.Language, root),
			Description:  dom.String(parser.Description, root),
			ImageCount:   dom.Int(parser.ImgCount, root),
			CountPrePage: dom.Int(parser.CountPrePage, root),
			Star:         dom.Double(parser.Star, root),
			Tag: &protobuf.TagRpcModel{
				Text:  dom.String(parser.Tag, root),
				Color: dom.Color(parser.TagColor, root),
			},
			Badges: utils.Map(dom.Nodes(parser.BadgeSelector, root), func(e *html.Node) *protobuf.DetailRpcModel_Badge {
				return &protobuf.DetailRpcModel_Badge{
					Text:     dom.String(parser.BadgeText, e),
					Category: dom.String(parser.BadgeCategory, e),
				}
			}),
			Comments: utils.Map(dom.Nodes(parser.CommentSelector, root), func(e *html.Node) *protobuf.DetailRpcModel_Comment {
				return &protobuf.DetailRpcModel_Comment{
					Username: dom.String(parser.Comment.Username, e),
					Content:  dom.String(parser.Comment.Content, e),
					Time:     dom.String(parser.Comment.Time, e),
					Score:    dom.String(parser.Comment.Score, e),
					Avatar:   ImageParser(dom, parser.Comment.Avatar, e),
				}
			}),
			PreviewImg: utils.Map(dom.Nodes(parser.ThumbnailSelector, root), func(e *html.Node) *protobuf.ImageRpcModel {
				return ImageParser(dom, parser.Thumbnail, e)
			}),
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
