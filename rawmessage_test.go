package pbtypes_test

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/nxpkg/pbtypes"
)

var _ json.Marshaler = (*pbtypes.RawMessage)(nil)
var _ json.Unmarshaler = (*pbtypes.RawMessage)(nil)

var _ proto.Marshaler = (*pbtypes.RawMessage)(nil)
var _ proto.Unmarshaler = (*pbtypes.RawMessage)(nil)
