package docker

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/opctl/opctl/util/pubsub"
	"github.com/opspec-io/sdk-golang/model"
	"golang.org/x/net/context"
	"io"
)

func (this _containerProvider) pullImage(
	dcgContainerImage *model.DCGContainerCallImage,
	containerId string,
	rootOpId string,
	eventPublisher pubsub.EventPublisher,
) error {

	imagePullOptions := types.ImagePullOptions{}
	if nil != dcgContainerImage.PullCreds &&
		"" != dcgContainerImage.PullCreds.Username &&
		"" != dcgContainerImage.PullCreds.Password {
		var err error
		imagePullOptions.RegistryAuth, err = constructRegistryAuth(
			dcgContainerImage.PullCreds.Username,
			dcgContainerImage.PullCreds.Password,
		)
		if nil != err {
			return err
		}
	}

	imagePullResp, err := this.dockerClient.ImagePull(
		context.Background(),
		dcgContainerImage.Ref,
		imagePullOptions,
	)
	defer imagePullResp.Close()
	if nil != err {
		return err
	}

	stdOutWriter := NewStdOutWriteCloser(eventPublisher, containerId, rootOpId)
	defer stdOutWriter.Close()

	dec := json.NewDecoder(imagePullResp)
	for {
		var jm jsonmessage.JSONMessage
		if err = dec.Decode(&jm); nil != err {
			if err == io.EOF {
				err = nil
			}
			return err
		}
		jm.Display(stdOutWriter, nil)
	}
}
