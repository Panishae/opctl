// Code generated by go-bindata.
// sources:
// github.com/opctl/specs/opspec/opfile/jsonschema.json
// DO NOT EDIT!

package dotyml

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _githubComOpctlSpecsOpspecOpfileJsonschemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x7b\x73\xdb\x36\xf2\xff\xe7\x53\x60\xd4\xf4\x57\xfb\x97\x48\xb2\xd3\x26\x6d\xdd\xc9\x78\xdc\x3c\xee\x72\xd3\x3c\xa6\x79\xdc\x4c\x2d\x37\x03\x91\x2b\x0b\x35\x09\x30\x20\xe8\x47\xef\xf2\xdd\x6f\x00\x50\x7c\x09\x00\x29\x8a\xb4\xec\x44\xf9\x27\x16\xb1\x58\x2e\xf6\x85\xc5\x62\x01\xfe\xe7\x0e\x42\x83\xbb\xb1\x37\x87\x10\x0f\x0e\xd0\x60\x2e\x44\x74\x30\x1e\xff\x15\x33\x3a\xd4\x4f\x47\x8c\x9f\x8e\x7d\x8e\x67\x62\xb8\xf7\xe3\x58\x3f\xfb\x66\x70\x5f\xf5\x23\xfe\xa2\x4f\x7c\x30\x1e\xb3\x28\x8e\xc0\x1b\x11\x36\xde\x1b\xed\x8f\x7e\x1c\xb3\x68\x74\x15\x06\xa3\x14\x8d\x44\xa9\xbb\x09\x22\x02\x90\x1d\x5f\x47\xe8\x39\x09\x40\x3f\xf5\x21\xf6\x38\x89\x04\x61\x54\xb6\x3d\x85\x19\xa1\x10\x23\x4c\x11\x8b\x34\x44\xc4\x59\x04\x5c\x10\x88\x07\x07\x48\x12\x8e\xd0\x80\xe2\x10\xb2\x5f\xcb\x58\x5e\xe1\x10\x10\x9b\x21\x31\x87\x05\x1a\x05\x26\xae\x22\x45\x41\x2c\x38\xa1\xa7\x03\xf5\xf8\xb3\x6e\xad\xa0\xb0\x61\x7e\x9a\xff\x34\xbd\xe0\x2e\x87\x99\x04\xfb\x66\x9c\x53\x3d\x26\x34\x4a\x44\x3c\x8e\xb0\x10\xc0\xe9\x9b\xbc\xe1\xdb\x87\xbf\x0e\x3f\x8e\xf0\xf0\xef\xa3\xe1\x1f\x7b\xc3\x9f\xbf\x7d\xf8\xf4\xdb\x07\xbf\x16\x7b\xfa\x84\x97\x7e\x16\x68\x29\xd1\xae\xdf\x50\x24\x1b\xfb\x3e\x91\x70\x38\x78\x53\x64\xdf\x0c\x07\x31\xdc\xb7\x8c\xed\x0d\xe6\x38\x04\x01\x5c\x8e\x2c\x17\x80\x02\x5d\x22\xbe\xf0\x32\x84\x06\xc7\xc5\x71\x9c\xdc\x2b\x35\x22\x34\x60\x14\x5e\x4b\xbe\x1c\x17\x1e\xa2\x12\x88\x02\xe3\xf0\x29\x21\x1c\xfc\x25\x48\x3d\x24\xce\xf1\xd5\xa0\xf2\xfc\xa4\xf4\xfb\xf3\xfd\x75\x5e\x30\x65\x2c\x00\x4c\xfb\x7c\x85\x4f\x78\x9f\xe8\x67\xd2\xac\x7a\xc4\x4f\x93\x70\x0a\xbd\x8e\x80\x4d\xff\x02\x4f\xf4\xf9\x86\x98\x79\x67\xd0\xef\x1b\x0a\xde\xc5\xfa\x86\xc2\xaf\x93\xe2\xdb\x4c\xee\x2e\x6b\xd3\x26\x50\x7d\xdc\xd0\xdc\x33\xe0\xcc\x11\x2b\x74\xca\xea\x07\x4b\x40\x15\xdf\x70\x24\x41\x51\x64\xf7\x10\x4d\xe8\x37\x20\x5e\x06\xb8\x7e\x2f\x5a\x92\x4b\x75\x3c\x8a\xe2\x19\x4e\x02\x61\xa3\x76\x31\xa9\x18\xfd\x93\x19\x23\x89\xdf\x82\xc7\xc1\x8a\xb2\xc2\xfc\x17\x7a\xa2\x51\x2f\x40\x24\x46\xb1\xee\xbc\x8c\xb8\x40\x8d\xc5\x99\x99\xe9\xf1\x18\x8d\x05\xc7\x84\x0a\xb3\xd0\x96\x94\xe6\x49\xa1\x83\x93\x8c\xd4\x9e\x8d\x30\x35\x9a\x82\xca\x8a\xfd\x42\x40\x68\x07\x5c\x66\xda\xbf\xde\xbe\x7e\x85\xde\xaa\x20\x04\x1d\x57\xd0\xa0\x33\xb8\xba\x60\xdc\x3f\xd9\x59\x04\x31\x82\xb1\x20\x1e\x11\x10\x33\x15\xf8\xcc\x45\x18\xa4\xd1\xcf\x05\x27\xa7\x73\x31\x2c\x84\x46\xc3\x73\x1c\x10\x1f\x4b\x7c\xc3\xbd\xfd\x6f\x62\xf0\xd4\x9f\x8f\x46\xfb\x7b\xbb\xc6\x91\xa2\x1e\x94\x5a\x33\xb6\xf8\xa4\x20\xc4\xe2\xe3\xc2\x9f\x46\x2f\x61\xa4\xd7\xa0\x24\x6a\x10\x64\x1d\x21\x90\x1e\x59\xff\xb3\x83\xf3\x98\x5e\x19\xa2\x8f\xe2\x3f\xdb\x70\x6e\x9f\xe0\x1c\xc2\xab\x1d\x68\x9d\x70\x6f\x2b\x3f\xca\xb3\x6d\xb3\x96\x93\x95\xcc\x22\xc4\x97\x6b\xb9\xa7\x45\xff\x0e\x8d\x63\x2f\x33\x8e\x87\x6e\xbf\xb4\xf0\xd4\x84\x0a\x38\x05\x6e\x07\x0c\x09\x25\x61\x12\x0e\x0e\xd0\xde\x6a\xcc\x21\x74\x3d\xe6\xa4\xfd\xfb\x62\xce\xfe\x26\x99\x93\x50\xf2\x29\x81\xb5\xf8\x53\x40\xd1\xd7\xbc\xf6\x7d\x03\x16\x59\x03\x0e\x64\x31\x33\x23\x4b\x9c\x91\xec\x72\x2c\x73\xa7\x06\x65\x35\x0e\x71\x45\xf9\xd9\x08\x3a\x8b\xaf\x53\x84\xcd\x22\xec\x5f\x35\xf0\x36\xc6\x6e\x10\x10\x3f\xd5\xd0\xe8\x1c\x07\x09\xb4\x8d\x83\x3b\xd5\x1d\xb9\xb0\xef\x4c\x6f\x7c\xc2\x9b\xe9\xcc\x53\xc2\xc1\x13\x8c\x5f\xdb\xca\x2c\x23\x31\xc4\xfc\xcc\x67\x17\xd4\xcc\xfc\x0a\x99\x2f\x53\x60\x44\x28\x3a\x3e\xdf\x1b\x3d\xf8\x09\x3d\x61\x61\xc8\xa8\x6c\x40\xf1\x15\x15\xf8\x52\xbb\xac\x83\xf1\x58\x25\x13\x3d\xd5\x2c\x5f\xa2\xdc\x96\xec\x32\xde\x45\x84\x7a\x41\xe2\x13\x7a\x8a\xfe\xf1\xfc\x25\x12\x78\x1a\x00\x82\x4b\x01\x34\x26\xcc\x42\x89\x29\xe3\xe7\x90\x23\x5a\x53\x17\x7f\x91\xbc\xc7\xd3\x98\x05\x89\x00\x14\x61\x31\x47\x9c\x31\x01\x3e\xc2\x02\xf9\x84\x23\x8f\x51\x81\x09\x95\x63\xd0\x79\x52\xc4\xf8\x7d\x84\x11\x87\x00\x0b\x72\x9e\xf6\x91\xd3\x0d\x8f\x38\xc8\x8e\x33\xce\x42\x74\x31\x07\x0e\x69\xb6\x51\xad\x00\x05\xe6\x02\xfc\xee\xc6\xdc\x6e\x45\xea\x67\xfa\xb7\xf6\xaa\xb4\x53\x6b\x54\x79\xb0\xce\xcc\x51\x62\x6b\x66\x8f\xcf\x49\x00\x5b\x07\xbe\x35\x1a\xf3\xa0\x53\xa3\x91\xea\x74\xc3\xec\x25\xcd\xeb\x76\x66\x31\x1a\x5f\x33\x9b\x79\xa5\x60\xb7\x56\xe3\x52\x03\x73\xde\xbd\x4b\xa5\xd4\x6f\xd8\x78\x72\x51\x93\x71\x7d\xd9\xc5\x20\x50\x99\xa1\x76\x39\x45\xd9\xb9\xa7\x45\xe9\x83\x07\x0d\x56\x5c\x3a\xe1\x6c\x05\xab\x4f\xe7\x74\x6d\x0d\x5a\x7a\x96\x54\x8e\x2d\x3f\x63\xce\xc0\xd8\xd6\xcb\x8b\x5c\x5e\x3b\x89\xc9\xce\x7d\x49\xac\xc9\x1a\xf9\x6b\x94\x18\x50\x95\x16\x69\x25\x30\xd9\xb7\x2f\x79\x35\xc9\x89\xad\x2f\xaf\x3a\x0f\x9e\xf2\x6e\x25\x8e\xce\x18\x0f\xb1\xcd\xbf\x2b\x08\xf3\x7e\x7b\xf1\x9f\x33\x11\x9c\xb9\xe3\xba\xb4\x17\x32\x08\xf0\x77\xbd\x37\x1a\x17\xa7\x95\x29\xc8\x29\xbd\x11\xb6\x4a\x1c\xe6\x84\x4d\x55\xcb\x3e\x4a\x05\xb5\x78\xad\x03\xca\x9c\xed\x45\x5d\xe6\x87\xd3\xec\x60\xdb\xf4\xb0\xec\xde\x97\x25\x34\x30\x04\x97\xfa\x3a\xf2\xbe\x6b\x0d\x5a\x77\xef\x69\xd0\x3f\xf4\x35\xe8\x24\x10\x24\x0a\xa0\xfd\x24\x95\x63\xe8\x2b\xe1\xdd\xd3\xd0\x29\x73\xfa\x24\xd7\x98\x29\x13\x7d\x29\xf7\xc3\xeb\xdb\x92\x6d\x31\x99\xda\x78\xb9\x70\xe0\xad\xb8\xa9\x3a\xf7\xc5\xcf\x26\x86\xf3\x45\x87\x39\x86\xa7\x37\x6e\x27\x23\x85\xe9\x6c\x3d\xaf\xf1\x35\x5b\xcf\xbf\x56\xb0\xdb\xf5\x7c\x83\x75\x6b\x8f\xeb\x79\xfd\x86\x8d\xaf\xe7\x35\x19\x1b\xa8\x16\xb2\x14\xaa\xba\x79\x67\x29\x19\xca\x71\xad\xe2\x54\xe7\x98\xfa\x1c\x2e\xe2\x06\x6e\xf5\xd1\xe8\xe1\xe8\x91\xc3\xaf\xae\x1b\xce\x37\xd9\xa1\x45\x6b\x14\x8e\xdc\xba\x8a\x90\x6e\xa2\xfa\x6d\xf6\xe8\xfa\xc4\xbe\xcd\x1e\x7d\x9d\x12\xf3\x21\x02\xea\x03\xf5\xd6\xf0\xe4\x45\x1c\x7d\xad\xa9\x5c\xf5\x87\x6b\xba\xef\xdb\xe7\x5d\xfb\xad\x3f\xac\xdb\xa9\x2b\x10\xd2\xf3\x04\xb0\x4d\x6d\xd6\x64\xdf\x06\x34\x09\x02\x77\x0e\xcf\x16\x0a\x2f\xfe\xd9\x52\x73\xab\xb9\x91\x10\x5f\xae\x1f\x11\x96\x90\xf4\xe5\x48\x9a\x4c\x04\x7d\x96\x6a\x76\xc0\xa6\x22\x92\xbe\xd8\xd4\x24\x0d\xd1\x1b\x9b\xbe\xfa\x4c\x57\x8b\x89\x7f\x9b\xe9\xba\xa9\x92\x59\xd5\x97\x36\x58\x88\xd7\x88\x2d\xea\xdd\x3d\xb8\x56\xd3\x0d\x92\x0d\xa8\x45\x32\xa1\x98\xf6\x90\xaf\xa8\x4b\x7a\x98\xf9\x94\x9f\xc1\xf5\x9a\xf5\xaf\x3f\xdc\xe2\x8e\xba\x9a\x0a\xd4\x42\x70\x1d\x78\x6d\x68\x9e\xb5\x75\xaa\x02\x8f\x5c\xbb\x1c\x05\xda\x9a\xc7\x92\xc8\x1d\xd8\x2e\xf0\xa5\x1a\xb0\x26\x5b\x14\x9a\x2f\x83\x21\x1a\xdf\xba\xfc\xb8\x8a\xba\x65\xc7\xc3\xd1\x83\x95\xf8\xe1\xde\x62\x4e\x61\xdd\x93\x41\x09\xb4\x19\x7b\x5d\x1b\xd3\xfa\x5f\xbd\x00\x2e\x38\x11\xf0\x9a\x06\xa6\x63\xc3\x06\x70\x97\x14\x32\x54\x1d\xa7\x3f\xf7\xf7\x46\x8e\xa8\xb7\x40\x5c\xe3\x44\x66\xca\x1c\x67\xbb\xab\xd5\xc9\xd6\xeb\x3d\x1e\xa7\x94\x6a\xc5\xf9\xfc\x86\x0d\xa1\xd5\xee\xdb\x0d\x1b\x43\xab\xb8\xea\x86\x8d\x41\x3b\x9d\x36\x63\xe8\x79\x75\xee\xbe\x5c\xa4\x02\xec\x0c\x2c\xab\x88\xfa\x8a\x2f\x7f\xdc\x48\x7c\x79\xcb\x32\x81\xab\x29\x41\xe1\x4e\x8d\x56\xb2\x5f\xf4\xef\x4b\xe4\xae\xe5\x75\xf7\xf5\x82\xee\x00\xe1\x16\x97\x26\xa4\xd7\xaf\x74\x56\x9a\xa0\xf1\x35\x2b\x4d\x78\xab\x60\xbf\xd2\xd2\x84\x76\x75\x04\x9a\xbd\x37\xec\xb8\x4a\x6a\x1c\xdd\xe9\x90\xc2\xd7\x50\x87\x14\xec\x57\xaa\x43\x0d\xcb\x5b\xfa\x3f\x43\xa5\xdf\xb0\xf1\xf2\x16\x4d\xc6\xf6\xb8\xca\xed\x2c\x38\x68\x15\x0f\x6f\x0b\x0e\x32\xb0\x2f\x52\x62\xdb\x3d\xdd\xd6\xe1\x67\xeb\xe3\x2a\xd5\x29\x36\x89\x80\xc7\x20\xe4\xd4\x5a\xe2\xaf\xc6\xd4\x0b\x87\x5d\x0b\xba\xae\x4e\xd3\xf8\x58\xc0\x50\x90\xd0\x7c\x63\x85\x8d\x1b\xe5\x7c\xf9\x02\x05\xd2\xbc\xe8\x96\x07\xa3\xef\xeb\xb2\xc4\x3d\x1c\xcf\xc9\xb9\xe2\x80\x73\x1c\xd0\x69\x59\x6c\x92\x4b\x45\x46\xb8\x7c\x48\x42\x7c\x0a\x43\xe9\xae\x56\x11\xce\x11\xd2\xdd\x91\xea\x8e\x38\xcc\x80\x03\xf5\x00\xe1\x18\xf9\xea\x2a\x5b\x1f\x4d\xaf\xd0\xf1\x29\x11\xf3\x64\x3a\xf2\x58\x38\xd6\x1d\xc6\x3e\x91\x2c\x9c\x26\x12\xd3\x38\xeb\x97\xcb\xb3\xa6\x87\xe0\x00\x8b\x86\xfd\xd1\xfe\xf7\x39\x8a\xeb\x17\x60\x95\x81\x9b\x91\x23\x84\x98\xd4\xd4\x9a\x38\x7d\xb7\xec\xde\x97\x55\x39\xc2\x23\xd4\x8f\x50\x34\x37\x36\x23\x89\x39\x8b\x85\xba\xa0\xb9\xb5\x30\x16\x18\xfa\x92\x47\xcd\x76\x43\x0f\xf2\xc8\x78\xb2\x19\x91\x90\xe8\xfc\x87\xf6\xe2\x90\xbd\xfb\x12\x85\xa3\xca\x02\xf5\x23\x0a\xc5\x8b\x8d\x89\xe1\xd1\x5a\x62\x78\xd4\x97\x18\x1c\xd9\x4d\xd4\x9b\x18\x1e\x6d\x48\x0c\x09\x27\xed\xa5\x90\x70\xd2\x97\x10\x1c\x55\x2b\xa8\x1f\x21\x48\x4e\x6c\x46\x06\x31\x84\xe7\x2b\x9e\x2c\x3f\x42\x31\x84\x98\x0a\xe2\xa1\x73\xe0\x31\x61\xb4\x1a\x66\x69\xa4\x52\x06\xf9\x75\x62\xd9\xa3\xf1\xb5\x73\x37\x1d\x63\x3b\x06\x5b\x5a\x56\x3e\x78\xfe\x1b\xd0\x53\x31\x5f\xa3\xe6\x55\x23\xe8\x69\x5d\xdb\xa4\x50\x6b\x85\x3a\xce\xfd\xd5\xb8\x43\xe8\x9a\xdc\x59\x20\xe8\x89\x3b\x4d\xb6\x19\x57\xa9\x72\xb5\xc2\xe4\x89\xdd\x6d\x25\x6c\xdf\x19\xa9\x6d\x25\xec\x4d\x95\x4c\xcb\xba\x85\x35\xab\x15\x7a\x92\xda\x4f\x0d\x84\x56\x33\xd5\xe5\xa9\xbb\x01\x87\x53\xb8\xbc\x4d\x77\xfb\xde\xb1\xf4\xb4\xf6\x5a\xf4\xc8\xa0\x8d\x90\x8b\x2f\xee\xb0\x44\x54\x3f\xb9\xe3\xd0\xdf\x72\x5f\x9e\xd0\x75\x3f\xd5\x73\x84\x62\x42\x4f\x03\x40\x94\xf9\xd9\x97\x8e\x8e\x3d\x1c\x04\xe8\x94\xe3\x68\x9e\xeb\x12\xd0\xd1\x05\x39\x23\x11\xf8\x44\x7f\xd3\x49\xfe\x1a\x3f\xc1\x41\xf0\x51\x41\xe6\x5a\x62\xc8\x73\x16\xd5\xda\xfe\xd1\x93\x41\x7a\x3f\x63\x39\xd4\xc9\x03\x95\x02\xff\x1b\x22\x64\x51\x57\x98\x22\xcc\x71\x10\x40\xd0\x15\xbe\x18\x38\xc1\x36\x6c\xe9\x5f\xd9\xf7\x5c\x6c\x9b\x7a\x05\x86\x55\x3e\x94\x94\x85\xc7\x19\x84\x94\xd4\xc0\xa5\xc0\xa5\x36\xd7\xc7\x63\xbc\xd0\x54\x5b\x53\xd5\xac\x27\x2c\x0c\x31\xf5\x11\x4f\xa8\x0c\xa9\x31\xca\x28\xf9\x05\xb1\x73\xe0\x9c\xf8\xea\xf3\x5c\x57\x28\x06\x81\xb0\x50\x9a\xa7\xb3\x9f\x01\x9c\xc3\x72\x06\xae\x66\x86\x70\xcc\x0c\x55\xd2\x9e\x5d\x46\x1c\x62\x15\xf7\x7b\x0c\xb8\x47\xa6\x01\x20\xc1\x16\x9b\xc1\xb6\xcb\xa0\xcd\x86\xc9\x13\x5a\xaa\x85\x8a\xd6\x98\x75\x94\xe1\x8c\xf7\xab\xee\xb1\xec\xc1\x0c\x37\x46\x9b\x86\x5d\xb7\x59\x6c\xbb\x02\x9a\x40\x8c\x08\x55\xd2\xc8\xb5\x6b\xb9\x24\xa1\x51\xa5\xdd\xe0\xcf\x9d\x63\x3d\xc6\x93\x83\xdd\xc3\xe3\xe1\xc7\xd1\x64\x32\x2e\x7c\xde\xeb\xae\x6d\x63\xdc\xbd\x53\xd2\x74\x92\xdc\xb9\x20\x41\x80\xa6\x80\xa6\x2c\xa1\xbe\x92\x30\x0e\xb3\x4b\x5e\x11\x8b\x1a\x5d\x60\x14\x58\xd2\x9f\x96\xc9\xbc\x29\x71\x36\x25\xf4\x09\xd7\x1a\x88\xfe\x6f\xcc\x38\x8a\x3d\x16\xa9\xfd\x00\x45\x3e\x08\x94\x44\x8c\x22\xb8\x24\x8e\x82\xc0\x4d\xe9\x69\xca\x17\xc3\xd3\xe5\xc5\x66\xfd\xac\xdc\x78\x92\xaf\x1a\x04\xd0\xf3\x0f\xb8\x13\x9b\x78\x46\xcf\x09\x67\x34\x04\x2a\xd0\x39\xe6\x04\x4f\x83\x4e\xad\xe3\xf8\xcf\xc7\x1b\x30\x02\x42\x0b\x5a\x75\x31\xd6\x46\x41\x71\xe8\xd8\xf1\xd9\x9c\x39\xd4\xf8\xe4\x14\x59\x13\x8d\xcf\xe4\x55\x2a\x29\xb7\xa6\xd2\x37\xae\xc6\x33\x12\x18\x15\x67\x55\x25\x7e\x4e\xba\x55\xda\xad\x4b\xb7\x11\x67\xd3\x61\x75\x51\xf7\x35\xf8\xf4\xdb\xa6\xe1\x2a\xda\x6b\xa3\xe1\x75\x05\x91\x9a\x57\x4d\xca\x00\x7f\xcf\x36\xda\x05\x53\x57\x8b\x2a\x92\x8c\x8b\xce\xcd\xcd\xaa\xa6\xea\xc2\x28\x09\x82\x27\x1c\x7c\x6b\x6d\xe1\xea\xe4\xe6\x28\x5b\x28\x81\xfb\x4b\x96\x86\xdd\xfc\x93\xee\xf4\xa8\xf2\x11\xdf\xec\xb9\xe9\x63\xbe\x25\x27\x88\x3c\x4c\xa5\x33\xc9\x6a\x1e\x54\xf6\x5f\xdd\xf9\xcf\xc4\x1c\x78\x0e\xb9\x5c\x90\x69\x3f\xd1\x67\x4d\x61\x65\xf5\xa3\x86\xa6\xf4\xa2\x4c\x43\x8b\xb5\xd8\xb3\xd9\xf7\x41\x2b\xbc\x8a\x18\x37\xd6\xa3\x2e\x9d\x1f\x96\x70\xa9\x97\x55\x0c\x29\x33\x4e\x30\xf5\x60\xce\x62\x83\x79\xd6\x9a\x6f\xb3\x99\xe6\x58\x4d\x28\x3b\x43\xfd\xff\xee\xe1\x8e\xf0\xa2\xff\x26\x7e\xb4\x7b\xd8\xd0\xb8\xff\xc9\x62\x81\xe4\x80\x77\xe2\x5d\x49\xf1\x94\xa8\x29\xc3\x59\x5b\x6b\x9e\x9e\x5c\xf7\x8f\xb8\x93\x5f\x8e\x1b\x50\x97\x2c\x00\x95\x73\x81\x4b\x0c\x68\x63\x97\xad\x8d\x4a\x17\xec\xb7\x0a\x40\x9a\xca\xf7\xc0\xfe\xa5\xe7\x0c\x68\x29\xab\xb0\xd0\xc0\xf4\x40\x01\xf6\x7d\x39\xe7\xa2\x10\x47\x11\xa8\x80\x00\x2f\x9a\x6c\x35\x58\xf5\xe5\xe5\xfd\x71\xf5\x82\xf1\xb3\xa7\xe6\x4f\x3c\x55\x46\xfa\x6f\xc6\xcf\x64\xe4\xeb\x17\x3e\xcd\x24\xe6\x68\xa7\x9c\x3b\x29\x6c\x5a\xaa\xa9\x6b\x39\xfc\x71\x8e\xd6\x9e\xdd\xb4\x67\xae\xf4\x14\x59\x78\x56\xfe\xec\x70\x23\xd6\x14\xde\x35\x08\x18\x8b\xaa\xe9\x2b\x57\xfa\x69\xc6\x8c\xec\xab\x8b\x09\x00\x7b\xb6\xed\xb9\x2a\xeb\x3f\xa8\x50\x4d\x3a\x0c\x40\x92\x3a\xf0\x55\xc6\xaa\x8d\xdf\x70\xed\x66\xb8\x0f\xf6\x39\x4a\x67\x0d\x41\x9a\x61\x76\x38\x03\xdb\x51\xed\xe5\xf1\xea\xa5\x2d\x92\x3c\x42\x44\x00\x57\xdb\x00\x31\xc2\x71\xcc\x3c\x82\x85\x3e\x0c\x87\x16\xb1\x77\x88\x7d\x40\xf8\x1c\x93\x40\xf5\x12\x73\xce\x92\xd3\x79\x4b\x5b\x33\x50\xae\x17\x7b\xdd\xd1\xae\x43\xef\x5e\xa8\x5f\x33\x2e\x52\x4a\x59\x1b\x18\xad\x72\xa4\x89\x50\x1f\x2e\x5d\x4e\xdb\xe8\x06\x2a\x48\x12\x2a\x48\xd0\xc0\x47\xbd\x98\x29\x27\x14\x71\xf0\x89\x87\x05\x20\x90\xac\xc6\x02\x62\xb5\xec\x51\x1b\x0f\x2a\x4e\x90\x56\xa4\x05\x20\x57\x39\x23\xf4\x66\xd1\x23\x46\x98\xe7\xdd\x7c\x34\x85\x19\x53\x0f\x80\x5f\xe5\xd2\x1c\x75\x99\x0c\x6e\x12\x18\x93\xd9\x58\x63\xa8\xca\xb7\xa1\xf7\xb4\xee\x0d\x15\xfc\x1e\xa9\xae\x51\x5a\x73\x57\x6d\xd7\xec\x30\x8e\x26\xca\x99\x4e\x06\x39\xe7\x76\x33\xb5\x8f\xcf\x88\x9c\x21\x47\xe6\x3d\x80\x65\x36\x5a\x58\x68\x98\x90\xb5\xbf\x47\x17\x73\xe2\xcd\xcb\x34\x0a\x9e\x00\x62\x5c\xd3\x5a\x11\x92\x35\x21\x60\x5c\xcf\x39\x8c\x48\x9a\xd1\xa7\x65\xd3\xac\x3a\xc9\x25\xd3\x6c\xf3\x9e\x4b\x12\x9b\xf6\x7c\xfb\x78\x17\x35\xac\xda\x7b\x79\x0f\x13\xcf\x9a\x0e\xab\xf4\xfb\xa4\xba\xb2\x70\x4d\xc0\x52\x44\x4d\xb6\x66\xde\x49\x8d\x21\x33\x24\x35\x5a\x7f\xe6\x5b\xf9\x87\x4f\x09\x36\x85\xdf\xb5\x45\x03\xce\x72\x81\xce\x52\x90\x9b\x4c\x0d\xd4\x4e\x40\xa9\xce\xae\xc2\x7c\x3d\x5f\xea\x8e\xe8\x62\x9c\x2f\x8f\x5d\x22\xb0\xae\x82\x8a\x2b\x9a\x3f\x27\x93\xbb\x93\xc9\xce\xe8\xde\x64\xb2\x7b\x77\xc9\xb9\x2e\xd1\x4e\xcd\x51\x80\x55\x69\xe8\x55\xae\x34\xf4\x3b\xb1\xd5\x9b\x35\xf4\x26\xf7\x0b\xab\x88\x80\xb2\x8d\x69\x8f\x33\xa2\xb1\x87\x50\x9f\x8d\x93\xf3\xf2\x92\xa4\xed\x6e\x79\x5a\xb8\xd1\x20\x98\xa2\x44\x10\x1c\x90\xbf\x21\x46\x2f\x5e\xbd\x79\xff\xee\xe3\xab\xa3\x97\xcf\x74\xf2\xe5\xc3\xd1\x6f\xef\x9f\x21\x42\xd3\x82\x61\xf4\x5d\x0e\x70\xa0\x1b\xbf\x1b\xa1\x17\xb3\x05\x5c\x8c\x68\x12\x04\xf7\x11\x11\xe8\xe5\xfb\xb7\xef\xd4\x67\x93\xe2\x38\x09\xc1\x4f\x21\x1e\x3f\x46\x77\x77\x72\x1c\x8e\x55\xe3\xba\x6b\xfc\xa2\x52\xda\x57\xf9\x2d\x37\x07\xae\x33\xa3\x6f\x08\x6f\x70\xdd\xb6\x54\xed\xa5\x5f\xb5\x57\x7d\x39\x52\x86\x19\x88\x35\x75\x98\x41\xd4\xdc\x20\x53\x7b\x8f\x98\xa5\x56\xb8\xd9\xd2\xb4\xd6\xd3\xb4\x4e\xa6\x2c\xd7\x4e\x65\x4d\x0e\xdb\xca\xad\xea\xf5\xfb\x77\x99\x99\x15\x6c\x4b\x5b\x55\xa1\x51\xdb\x56\x09\xda\x61\x61\x0a\xe0\xf1\xe3\x22\xfc\xd6\xbc\x2a\x18\x5d\xea\xb6\x71\xb5\x72\x6d\xad\xac\xba\xf9\x2a\xf1\x00\x95\xba\x17\xa3\x24\x06\x1f\xf9\x89\x0a\x0a\x70\x22\xe6\xf2\xb9\x87\xf5\xda\x89\x88\x79\x7a\x4d\x4a\xc2\x3d\x48\xef\x02\xd1\xd5\x50\x8c\xb7\xba\x14\x24\x89\x81\x5b\x36\x46\x4c\x84\xde\xca\x20\xc6\x18\x2e\xc4\xf1\x05\xe3\xb6\xdb\x9f\xbe\x88\x61\xd7\xaa\xbe\x3b\xe9\x94\x69\x86\x8b\x7d\x95\xa6\x0e\x77\xeb\xcc\x9b\xb3\x75\x1b\xb3\x26\x13\xa8\x09\x1e\x0b\x55\xc7\x09\x27\xc3\x3c\x0e\x2d\x93\x77\xc7\x42\xaa\x23\x05\x5e\xd9\xce\x5c\x33\x01\x9e\x95\x96\xda\x6a\x38\x17\x00\xf6\x12\xce\xe6\xe9\x1b\x9b\xda\xd6\x87\xc4\x69\xc5\xaa\x8d\x48\xdd\xbc\x01\x12\xef\x54\x48\x2d\xd4\xbe\xe2\xc5\xbc\xb5\xa8\x98\x4e\x8f\x7b\x15\xab\xa6\xab\xf9\xe4\xf4\x40\x58\x5a\x0a\x5d\x50\x3b\x33\x5d\x2c\x8a\x23\xf0\x2a\x35\xdd\xfa\x59\xa3\x97\x1c\x6b\xe0\xbc\xce\x5a\xff\x1e\x11\xb6\xab\x27\x8c\xe9\xd5\x12\x21\xc5\xb3\x6f\x1f\x0a\x91\x9f\xcd\x20\x4a\x6b\xa8\x9d\x74\x37\x71\x32\x19\x19\xfe\xdc\x39\x3c\xd8\x99\x4c\xd4\x8e\xe3\xd1\xf0\x0f\x3c\xfc\x7b\x78\x72\x6f\xe7\xf0\x60\x32\x19\x95\x1e\xed\xfe\xff\xee\xee\xa1\x7a\x7e\xaf\xf0\x7c\x32\x19\x4e\x26\xa3\x93\x7b\xbb\x87\xe9\xea\x4c\xca\x46\x31\x65\xd9\x96\x06\x8b\xba\x14\x65\x3c\x6e\xa3\xf9\xfc\xbf\x00\x00\x00\xff\xff\x84\x80\x3e\x94\x66\x98\x00\x00")

func githubComOpctlSpecsOpspecOpfileJsonschemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpctlSpecsOpspecOpfileJsonschemaJson,
		"github.com/opctl/specs/opspec/opfile/jsonschema.json",
	)
}

func githubComOpctlSpecsOpspecOpfileJsonschemaJson() (*asset, error) {
	bytes, err := githubComOpctlSpecsOpspecOpfileJsonschemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opctl/specs/opspec/opfile/jsonschema.json", size: 39014, mode: os.FileMode(420), modTime: time.Unix(1560727591, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"github.com/opctl/specs/opspec/opfile/jsonschema.json": githubComOpctlSpecsOpspecOpfileJsonschemaJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"github.com": &bintree{nil, map[string]*bintree{
		"opctl": &bintree{nil, map[string]*bintree{
			"specs": &bintree{nil, map[string]*bintree{
				"opspec": &bintree{nil, map[string]*bintree{
					"opfile": &bintree{nil, map[string]*bintree{
						"jsonschema.json": &bintree{githubComOpctlSpecsOpspecOpfileJsonschemaJson, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
