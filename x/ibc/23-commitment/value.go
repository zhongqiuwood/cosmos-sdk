package commitment

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/store/mapping"
)

type Base struct {
	cdc    *codec.Codec
	prefix []byte
}

func NewBase(cdc *codec.Codec) Base {
	return Base{
		cdc: cdc,
	}
}

func (base Base) Store(ctx Context) Store {
	return ctx.Store()
}

type Value struct {
	base Base
	key  []byte
}

func NewValue(base Base, key []byte) Value {
	return Value{base, key}
}

func (v Value) Is(ctx Context, value interface{}) bool {
	return v.base.Store(ctx).Prove(v.key, v.base.cdc.MustMarshalBinaryBare(value))
}

func (v Value) IsRaw(ctx Context, value []byte) bool {
	return v.base.Store(ctx).Prove(v.key, value)
}

type Enum struct {
	Value
}

func NewEnum(v Value) Enum {
	return Enum{v}
}

func (v Enum) Is(ctx Context, value byte) bool {
	return v.Value.IsRaw(ctx, []byte{value})
}

type Integer struct {
	Value

	enc mapping.IntEncoding
}

func NewInteger(v Value, enc mapping.IntEncoding) Integer {
	return Integer{v, enc}
}

func (v Integer) Is(ctx Context, value uint64) bool {
	return v.Value.IsRaw(ctx, mapping.EncodeInt(value, v.enc))
}
