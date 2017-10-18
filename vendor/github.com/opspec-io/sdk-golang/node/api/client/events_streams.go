package client

import (
	"encoding/json"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/api"
	"path"
	"strings"
	"time"
)

func (c client) GetEventStream(
	req *model.GetEventStreamReq,
) (chan model.Event, error) {

	reqUrl := c.baseUrl
	reqUrl.Scheme = "ws"
	reqUrl.Path = path.Join(reqUrl.Path, api.URLEvents_Stream)

	if nil != req.Filter {
		queryValues := reqUrl.Query()
		if nil != req.Filter.Since {
			queryValues.Add("since", req.Filter.Since.Format(time.RFC3339))
		}
		if nil != req.Filter.Roots {
			queryValues.Add("roots", strings.Join(req.Filter.Roots, ","))
		}
		reqUrl.RawQuery = queryValues.Encode()
	}

	wsConn, _, err := c.wsDialer.Dial(
		reqUrl.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	eventStream := make(chan model.Event, 1000)
	go func() {
		// ensure web socket closed on exit
		defer wsConn.Close()

		// ensure channel closed on exit
		defer close(eventStream)

		for {

			_, bytes, err := wsConn.ReadMessage()
			if nil != err {
				return
			}

			var event model.Event
			err = json.Unmarshal(bytes, &event)
			if nil != err {
				return
			}
			eventStream <- event

		}
	}()

	return eventStream, err

}
