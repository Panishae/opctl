package data

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeHandler.go --fake-name FakeHandler ./ Handler

import (
	"net/http"

	"github.com/opctl/opctl/sdks/go/internal/urlpath"
	"github.com/opctl/opctl/sdks/go/node/api/handler/data/ref"
	"github.com/opctl/opctl/sdks/go/node/core"
)

type Handler interface {
	Handle(
		httpResp http.ResponseWriter,
		httpReq *http.Request,
	)
}

// NewHandler returns an initialized Handler instance
func NewHandler(
	core core.Core,
) Handler {
	return _handler{
		refHandler: ref.NewHandler(core),
	}
}

type _handler struct {
	refHandler ref.Handler
}

func (hdlr _handler) Handle(
	httpResp http.ResponseWriter,
	httpReq *http.Request,
) {
	pathSegment, err := urlpath.NextSegment(httpReq.URL)
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusBadRequest)
		return
	}

	switch pathSegment {
	case "":
		http.NotFoundHandler().ServeHTTP(httpResp, httpReq)
		return
	default:
		hdlr.refHandler.Handle(
			pathSegment,
			httpResp,
			httpReq,
		)
	}
}
