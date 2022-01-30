package rpc

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
	"github.com/simonkimi/catweb-parser/parser"
)

func ParseRpcData(data []byte) ([]byte, error) {
	rpcMessage := &protobuf.RpcRequest{}
	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		return nil, err
	}

	switch rpcMessage.Type {
	case protobuf.RpcType_RPC_TYPE_LIST_VIEW_PARSER:
		return parser.ListParser(rpcMessage)
	case protobuf.RpcType_RPC_TYPE_IMAGE_PARSER:
		return nil, errors.New("TODO: RpcType_RPC_TYPE_IMAGE_PARSER")
	case protobuf.RpcType_RPC_TYPE_GALLERY_PARSER:
		return nil, errors.New("TODO: RpcType_RPC_TYPE_GALLERY_PARSER")
	}

	return nil, errors.New("TODO: 你怎么到的这里")
}
