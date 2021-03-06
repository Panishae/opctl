package core

import (
	"context"
	"github.com/opctl/opctl/sdks/go/model"
)

func (this _core) StartOp(
	ctx context.Context,
	req model.StartOpReq,
) (string, error) {
	opHandle, err := this.data.Resolve(
		ctx,
		req.Op.Ref,
		this.data.NewFSProvider(),
		this.data.NewGitProvider(this.dataCachePath, req.Op.PullCreds),
	)
	if nil != err {
		return "", err
	}

	opID, err := this.uniqueStringFactory.Construct()
	if nil != err {
		// end run immediately on any error
		return "", err
	}

	// construct scgOpCall
	scgOpCall := &model.SCGOpCall{
		Ref:     opHandle.Ref(),
		Inputs:  map[string]interface{}{},
		Outputs: map[string]string{},
	}

	// pull Creds
	if nil != req.Op.PullCreds {
		scgOpCall.PullCreds = &model.SCGPullCreds{
			Username: req.Op.PullCreds.Username,
			Password: req.Op.PullCreds.Password,
		}
	}

	for name := range req.Args {
		// implicitly bind
		scgOpCall.Inputs[name] = ""
	}

	opFile, err := this.opFileGetter.Get(
		ctx,
		opHandle,
	)
	if nil != err {
		return "", err
	}
	for name := range opFile.Outputs {
		// implicitly bind
		scgOpCall.Outputs[name] = ""
	}

	go func() {
		this.caller.Call(
			// call in background context
			context.Background(),
			opID,
			req.Args,
			&model.SCG{
				Op: scgOpCall,
			},
			opHandle,
			nil,
			opID,
		)
	}()

	return opID, nil

}
