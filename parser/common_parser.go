package parser

import (
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/selector"
	"golang.org/x/net/html"
)

func ImageParser(dom *selector.DomSelector, parser *protobuf.ImageSelector, node *html.Node) *protobuf.ImageRpcModel {
	if parser == nil {
		return nil
	}
	return &protobuf.ImageRpcModel{
		Url:    dom.String(parser.ImgUrl, node),
		Width:  dom.Double(parser.ImgWidth, node),
		Height: dom.Double(parser.ImgHeight, node),
		ImgX:   dom.Double(parser.ImgX, node),
		ImgY:   dom.Double(parser.ImgY, node),
		Target: dom.String(parser.Target, node),
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
