package bracketed

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeCoerceToArrayOrObjecter.go --fake-name fakeCoerceToArrayOrObjecter ./ coerceToArrayOrObjecter

import (
	"github.com/opctl/opctl/sdks/go/data/coerce"
	"github.com/opctl/opctl/sdks/go/model"
)

// coerceToArrayOrObjecter coerces data ao an object or an array value
type coerceToArrayOrObjecter interface {
	CoerceToArrayOrObject(
		data *model.Value,
	) (*model.Value, error)
}

func newCoerceToArrayOrObjecter() coerceToArrayOrObjecter {
	return _coerceToArrayOrObjecter{
		coerce: coerce.New(),
	}
}

type _coerceToArrayOrObjecter struct {
	coerce coerce.Coerce
}

func (cao _coerceToArrayOrObjecter) CoerceToArrayOrObject(
	data *model.Value,
) (*model.Value, error) {

	if dataAsArray, err := cao.coerce.ToArray(data); nil == err {
		// array coercible type
		return dataAsArray, nil
	}

	return cao.coerce.ToObject(data)
}
