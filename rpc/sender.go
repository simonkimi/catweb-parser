package rpc

import (
	"github.com/golang/protobuf/proto"
	"github.com/simonkimi/catweb-parser/gen/protobuf"
)

func ParseRpcData(data []byte) error {
	rpcMessage := &protobuf.RpcRequest{}

	if err := proto.Unmarshal(data, rpcMessage); err != nil {
		return err
	}

	switch rpcMessage.Type {
	case protobuf.RpcType_RPC_TYPE_IMAGE_PARSER:
	case protobuf.RpcType_RPC_TYPE_LIST_VIEW_PARSER:
	case protobuf.RpcType_RPC_TYPE_GALLERY_PARSER:
	}

	return nil
}
