package files

//go:generate counterfeiter -o ./fakeInterpreter.go --fake-name FakeInterpreter ./ Interpreter

import (
	"fmt"
	"github.com/golang-interfaces/ios"
	"github.com/golang-utils/filecopier"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/interpreter/expression"
	"path/filepath"
	"strings"
)

type Interpreter interface {
	Interpret(
		opHandle model.DataHandle,
		scope map[string]*model.Value,
		scgContainerCallFiles map[string]interface{},
		scratchDirPath string,
	) (map[string]string, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter(
	rootFSPath string,
) Interpreter {
	return _interpreter{
		expression: expression.New(),
		fileCopier: filecopier.New(),
		os:         ios.New(),
		rootFSPath: rootFSPath,
	}
}

type _interpreter struct {
	expression expression.Expression
	fileCopier filecopier.FileCopier
	os         ios.IOS
	rootFSPath string
}

func (itp _interpreter) Interpret(
	opHandle model.DataHandle,
	scope map[string]*model.Value,
	scgContainerCallFiles map[string]interface{},
	scratchDirPath string,
) (map[string]string, error) {
	dcgContainerCallFiles := map[string]string{}
fileLoop:
	for scgContainerFilePath, fileExpression := range scgContainerCallFiles {

		if nil == fileExpression {
			// bound implicitly
			fileExpression = fmt.Sprintf("$(%v)", scgContainerFilePath)
		}

		fileValue, err := itp.expression.EvalToFile(
			scope,
			fileExpression,
			opHandle,
			scratchDirPath,
		)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				scgContainerFilePath,
				fileExpression,
				err,
			)
		}

		if !strings.HasPrefix(*fileValue.File, itp.rootFSPath) {
			// bound to non rootFS file
			dcgContainerCallFiles[scgContainerFilePath] = *fileValue.File
			continue fileLoop
		}
		dcgContainerCallFiles[scgContainerFilePath] = filepath.Join(scratchDirPath, scgContainerFilePath)

		// create file
		if err := itp.os.MkdirAll(
			filepath.Dir(dcgContainerCallFiles[scgContainerFilePath]),
			0777,
		); nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				scgContainerFilePath,
				fileExpression,
				err,
			)
		}

		err = itp.fileCopier.OS(
			*fileValue.File,
			dcgContainerCallFiles[scgContainerFilePath],
		)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				scgContainerFilePath,
				fileExpression,
				err,
			)
		}

	}
	return dcgContainerCallFiles, nil
}
