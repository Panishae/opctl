package manifest

import (
	"bytes"
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"gopkg.in/yaml.v2"
	"io"
)

func (this _Manifest) Unmarshal(
	manifestReader io.Reader,
) (*model.PkgManifest, error) {

	var err error

	// 1) ensure valid
	errs := this.Validate(manifestReader)
	if len(errs) > 0 {
		messageBuffer := bytes.NewBufferString(
			fmt.Sprint(`
-
  Error(s):`))
		for _, validationError := range errs {
			messageBuffer.WriteString(fmt.Sprintf(`
    - %v`, validationError.Error()))
		}
		err = fmt.Errorf(
			`%v
-`, messageBuffer.String())
	}
	if nil != err {
		return nil, err
	}

	// 2) build
	packageManifestBytes, err := this.ioUtil.ReadAll(
		manifestReader,
	)
	if nil != err {
		return nil, err
	}

	pkgManifest := model.PkgManifest{}
	return &pkgManifest, yaml.Unmarshal(packageManifestBytes, &pkgManifest)

}
