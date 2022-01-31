package parser

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
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
		model := &protobuf.DetailRpcModel{}
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

		// 基础信息
		model.Title = dom.String(parser.Title, root)
		model.Subtitle = dom.String(parser.Subtitle, root)
		model.UploadTime = dom.String(parser.UploadTime, root)
		model.Language = dom.String(parser.Language, root)
		model.Description = dom.String(parser.Description, root)

		model.ImageCount = dom.Int(parser.ImgCount, root)
		model.CountPrePage = dom.Int(parser.CountPrePage, root)
		model.Star = dom.Double(parser.Star, root)

		// Tag
		model.Tag = &protobuf.DetailRpcModel_Tag{
			Text:  dom.String(parser.Tag, root),
			Color: dom.Color(parser.TagColor, root),
		}

		// 徽章
		var badgeList []*protobuf.DetailRpcModel_Badge
		for _, e := range dom.Nodes(parser.BadgeSelector, root) {
			badgeList = append(badgeList, &protobuf.DetailRpcModel_Badge{
				Text:     dom.String(parser.BadgeText, e),
				Category: dom.String(parser.BadgeCategory, e),
			})
		}
		model.Badges = badgeList

		// 评论
		var commentList []*protobuf.DetailRpcModel_Comment
		for _, e := range dom.Nodes(parser.CommentSelector, root) {
			commentList = append(commentList, &protobuf.DetailRpcModel_Comment{
				Username: dom.String(parser.Comment.Username, e),
				Content:  dom.String(parser.Comment.Content, e),
				Time:     dom.String(parser.Comment.Time, e),
				Score:    dom.String(parser.Comment.Score, e),
				Avatar:   ImageParser(dom, parser.Comment.Avatar, e),
			})
		}
		model.Comments = commentList

		// 缩略图
		var imageList []*protobuf.ImageRpcModel
		for _, e := range dom.Nodes(parser.ThumbnailSelector, root) {
			imageList = append(imageList, ImageParser(dom, parser.Thumbnail, e))
		}
		model.PreviewImg = imageList

		model.GlobalEnv = globalEnv
		model.LocalEnv = dom.Env

		marshal, err := proto.Marshal(model)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}
