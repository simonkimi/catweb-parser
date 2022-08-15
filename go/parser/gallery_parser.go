package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/utils"
	"google.golang.org/protobuf/proto"
)

func GalleryParser(rpc *protobuf.RpcRequest) ([]byte, error) {
	parser := &protobuf.GalleryParser{}
	err := proto.Unmarshal(rpc.ParserData, parser)
	if err != nil {
		return nil, err
	}

	dom, root, global, err := InitParser(rpc, parser.ExtraSelector)

	// 基础信息
	model := &protobuf.GalleryRpcModel{
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
		Badges: utils.Map(dom.Nodes(parser.BadgeSelector, root), func(e any) *protobuf.TagRpcModel {
			return &protobuf.TagRpcModel{
				Text:     dom.String(parser.BadgeText, e),
				Category: dom.String(parser.BadgeCategory, e),
			}
		}),
		Comments: utils.Map(dom.Nodes(parser.CommentSelector, root), func(e any) *protobuf.GalleryRpcModel_Comment {
			return &protobuf.GalleryRpcModel_Comment{
				Username: dom.String(parser.Comment.Username, e),
				Content:  dom.String(parser.Comment.Content, e),
				Time:     dom.String(parser.Comment.Time, e),
				Score:    dom.String(parser.Comment.Score, e),
				Avatar:   ImageParser(dom, parser.Comment.Avatar, e),
			}
		}),
		Items: utils.Map(dom.Nodes(parser.ThumbnailSelector, root), func(e any) *protobuf.GalleryRpcModel_Item {
			return &protobuf.GalleryRpcModel_Item{
				PreviewImg: ImageParser(dom, parser.Thumbnail, e),
				Target:     dom.String(parser.Target, e),
			}
		}),
		CoverImg:  ImageParser(dom, parser.CoverImg, root),
		NextPage:  dom.String(parser.NextPage, root),
		GlobalEnv: global,
		LocalEnv:  dom.Env,

		EnableSuccess: parser.SuccessSelector.Selector != "",
		EnableFail:    parser.FailedSelector.Selector != "",

		IsSuccess:     dom.String(parser.SuccessSelector, root) != "",
		FailedMessage: dom.String(parser.FailedSelector, root),
	}

	marshal, err := proto.Marshal(model)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
