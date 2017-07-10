package files

import (
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"io"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func (f _Files) Interpret(
	pkgPath string,
	scope map[string]*model.Value,
	scgContainerCallFiles map[string]string,
	scratchDirPath string,
) (map[string]string, error) {
	dcgContainerCallFiles := map[string]string{}
fileLoop:
	for scgContainerFilePath, scgContainerFileBind := range scgContainerCallFiles {

		if "" == scgContainerFileBind {
			// bound implicitly
			scgContainerFileBind = scgContainerFilePath
		}

		isBoundToPkgContent := strings.HasPrefix(scgContainerFileBind, "/")
		value, isBoundToScope := scope[scgContainerFileBind]

		var contentSrc io.Reader
		var err error
		switch {
		case isBoundToPkgContent:
			// bound to pkg file
			dcgContainerCallFiles[scgContainerFilePath] = filepath.Join(scratchDirPath, scgContainerFilePath)
			err = f.fileCopier.OS(
				filepath.Join(pkgPath, scgContainerFileBind),
				dcgContainerCallFiles[scgContainerFilePath],
			)
			if nil != err {
				return nil, fmt.Errorf(
					"Unable to bind file '%v' to pkg content '%v'; error was: %v",
					scgContainerFilePath,
					scgContainerFileBind,
					err.Error(),
				)
			}
			continue fileLoop
		case isBoundToScope:
			switch {
			case nil == value:
				return nil, fmt.Errorf(
					"Unable to bind file '%v' to '%v'; '%v' null",
					scgContainerFilePath,
					scgContainerFileBind,
					scgContainerFileBind,
				)
			case nil != value.Number:
				contentSrc = strings.NewReader(strconv.FormatFloat(*value.Number, 'f', -1, 64))
			case nil != value.String:
				contentSrc = strings.NewReader(*value.String)
			case nil != value.File:
				if strings.HasPrefix(*value.File, f.rootFSPath) {
					// bound to rootFS file
					dcgContainerCallFiles[scgContainerFilePath] = filepath.Join(scratchDirPath, scgContainerFilePath)
					err = f.fileCopier.OS(
						filepath.Join(*value.File),
						dcgContainerCallFiles[scgContainerFilePath],
					)
					if nil != err {
						return nil, fmt.Errorf(
							"Unable to bind file '%v' to '%v'; error was: %v",
							scgContainerFilePath,
							scgContainerFileBind,
							err.Error(),
						)
					}
					if nil != err {
						return nil, err
					}
				} else {
					// bound to non rootFS file
					dcgContainerCallFiles[scgContainerFilePath] = *value.File
				}
				continue fileLoop
			default:
				return nil, fmt.Errorf(
					"Unable to bind file '%v' to '%v'; '%v' not a file, number, or string",
					scgContainerFilePath,
					scgContainerFileBind,
					scgContainerFileBind,
				)
			}
		default:
			// unbound
			contentSrc = strings.NewReader("")
		}
		dcgContainerCallFiles[scgContainerFilePath] = filepath.Join(scratchDirPath, scgContainerFilePath)

		// create file
		if err := f.os.MkdirAll(
			path.Dir(dcgContainerCallFiles[scgContainerFilePath]),
			0700,
		); nil != err {
			return nil, err
		}
		outputFile, err := f.os.Create(dcgContainerCallFiles[scgContainerFilePath])
		if nil != err {
			return nil, err
		}

		// copy content to file
		_, err = f.io.Copy(outputFile, contentSrc)
		outputFile.Close()
		if nil != err {
			return nil, err
		}

	}
	return dcgContainerCallFiles, nil
}
