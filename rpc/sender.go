package rpc

import (
	"errors"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/parser"
	"google.golang.org/protobuf/proto"
)

func ParseRpcData(data []byte) ([]byte, error) {
	rpcMessage := &protobuf.RpcRequest{}
	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		return nil, err
	}

	if rpcMessage.Env == nil {
		rpcMessage.Env = make(map[string]string)
	}

	switch rpcMessage.Type {
	case protobuf.RpcType_RPC_TYPE_LIST_VIEW_PARSER:
		return parser.ListParser(rpcMessage)
	case protobuf.RpcType_RPC_TYPE_IMAGE_PARSER:
		return parser.ImageReaderParser(rpcMessage)
	case protobuf.RpcType_RPC_TYPE_GALLERY_PARSER:
		return parser.GalleryParser(rpcMessage)
	}

	return nil, errors.New("TODO: Unsupported: " + rpcMessage.Type.String())
}

func ExecSelector(input []byte) []byte {
	var errmsg string
	data, err := ParseRpcData(input)
	if err != nil {
		errmsg = err.Error()
	} else {
		errmsg = ""
	}
	model := &protobuf.RpcResponse{
		Data:  data,
		Error: errmsg,
	}
	marshal, err := proto.Marshal(model)
	if err != nil {
		panic(err)
	}
	return marshal
}
