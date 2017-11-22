// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/op.yml.schema.json
// DO NOT EDIT!

package manifest

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

var _githubComOpspecIoSpecSpecOpYmlSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x7b\x6f\xdc\xb6\xb2\xff\xdf\x9f\x82\x70\x8d\xdb\xdd\x1b\xaf\xd6\x4e\x52\xdf\xd6\x45\x61\x18\x79\xf4\xe6\xa0\x79\x20\x4e\x72\x80\xda\x9b\x82\x2b\xcd\xee\xb2\x96\x44\x95\xa4\xfc\x68\x4f\xbe\xfb\x01\x45\xbd\x45\xea\x65\x29\xb6\x93\x18\x38\xa7\x59\x91\x33\xe4\xfc\x38\x1c\x0e\x67\x28\xea\x9f\x2d\x84\xb6\x77\xb8\xbd\x01\x0f\x6f\x1f\xa2\xed\x8d\x10\xc1\xe1\x7c\xfe\x27\xa7\xfe\x4c\x3d\xb5\x28\x5b\xcf\x1d\x86\x57\x62\xb6\xf7\x78\xae\x9e\x7d\xb7\xbd\x2b\xe9\x88\x93\x90\xf0\xc3\xf9\x9c\x06\x3c\x00\xdb\x22\x74\xbe\x67\xed\x5b\x07\x73\x1a\x58\xd7\x9e\x6b\xc5\x5c\x24\x47\x45\x25\x88\x70\x41\x12\xaa\x0a\xea\xa1\x03\xdc\x66\x24\x10\x84\xfa\xb2\xe8\x29\xac\x88\x0f\x1c\x61\x14\x9c\xaf\x55\x8d\x80\xd1\x00\x98\x20\xc0\xb7\x0f\x91\xec\x36\x42\xdb\x3e\xf6\x20\xfd\x55\xe5\xf2\x0a\x7b\x80\xe8\x0a\x89\x0d\xa4\x7c\xa2\x7a\xe2\x3a\x88\x7a\xc0\x05\x23\xfe\x7a\x3b\x7a\xfc\x49\x95\x96\x78\x98\x58\x3f\xcd\x7e\x6a\x5b\xd8\x61\xb0\x92\xf5\xbe\x9b\x3b\x52\x14\x22\x2b\xf2\xb9\x87\xd9\xb9\x43\x2f\xfd\x62\x8b\xc4\xae\x6d\xea\x85\x5d\x68\xc3\x42\x2f\xdf\x9f\xbc\x43\x4b\x40\xd8\x47\x78\xc9\xa9\x1b\x0a\x40\x01\x16\x1b\xb4\x62\xd4\x43\x8c\x52\x21\xab\x07\xe7\x6b\x44\x19\x7a\xff\xf6\x37\xc4\x60\x05\x0c\x7c\x9b\xf8\x6b\x49\x74\x7a\xf2\xe1\x57\xb4\x6f\xed\x2f\x26\xc9\xd0\x5d\x5e\x5e\x5a\x97\x8f\xa2\x91\x7e\xf7\x76\x7e\xf2\xe1\xd7\xfd\xfd\xf9\x14\x11\x0f\xaf\xa1\x1d\x68\xc4\x0f\x42\xc1\xf3\x42\xe8\x01\x08\x30\xc3\x5e\x91\x94\x86\xa2\x37\x2d\x0b\xfd\x66\x3a\x6e\x97\x3a\x7b\x01\x8c\xd7\x8f\xee\x07\x55\x23\x41\x9d\x06\x4d\x03\xcb\xc1\xfb\x00\x2c\x6e\x66\x2b\x6e\x6a\x9b\xc1\x5f\x21\x61\x20\x67\xc9\x69\x4e\x61\xb7\x10\x5a\x44\xe5\xd8\x71\x22\x7a\xec\xbe\xc9\x2b\xf7\x0a\xbb\x1c\xe2\x59\x91\x36\x91\x29\x3d\x66\x0c\x5f\xbf\x89\xc0\xc8\x49\x90\xce\xaa\x5c\xf1\xae\x41\xbc\x63\x59\x05\x45\x78\x82\x00\x26\xc5\xc4\x7e\x41\xca\x64\xac\xe9\xf2\x4f\xb0\x45\xf6\x5c\x33\x09\xb3\x3e\x15\x1e\x99\x2b\xd7\xcc\xb2\xb4\xb8\xcd\xf4\x49\xfe\x3e\xed\x96\x59\xaf\x70\xe8\x0a\x1d\xdb\x44\x2c\xd5\xdd\x5a\x2e\x84\x9f\x80\xcd\x40\xcb\xa6\x3c\x41\x95\x96\x44\x4c\x11\xe1\x88\x2b\xc2\x5d\x53\xeb\x4b\x4a\x5d\xc0\x0d\x52\xd8\xd4\xe7\x82\x61\xe2\x8b\x2a\x7a\x46\x80\xa2\x2e\x3c\xc9\x51\x16\x9b\xd8\x32\x34\x57\xab\x88\x5b\x65\xf2\x94\xb4\x51\x7f\xa3\x4a\x95\x49\x80\x32\x85\x89\x7f\x2f\x0a\xd3\xb3\x22\x84\x51\xcb\xf3\x95\x7a\xab\xae\x87\xaf\x5e\x08\xf0\xca\x28\x97\x07\xf9\x5f\x27\xaf\x5f\xa1\x93\x68\x2d\x43\xa7\x09\x0d\x3a\x87\xeb\x4b\xca\x9c\xcc\x92\x0a\x4a\x5d\x6e\x11\x10\xab\xc8\x9a\x6e\x84\xe7\xc6\x8b\xe7\x25\x23\xeb\x8d\x98\xe5\x56\xd6\xd9\x05\x76\x89\x83\x65\x0b\xb3\xbd\xbd\xef\x38\xd8\xd1\x3f\x7f\xb0\xf6\xf7\xa6\x05\xed\x49\x65\x22\xbe\x80\x35\xb0\x62\xa1\x47\x7c\xe2\x85\xd2\x18\xec\x6d\x69\x86\x57\x96\x77\x17\x30\xa6\x19\x4b\xc0\xfd\x21\x05\x24\x5d\xa5\x23\x03\x8b\xb6\x9f\x8a\x76\x60\xfd\x54\x92\x8c\xfa\xf0\x7a\x55\xd0\x7d\xf9\xd7\x72\x3e\x4b\x58\xcc\xd3\x79\xb7\x9e\xa5\x0e\x96\xfe\xad\x15\xcd\x47\xf1\xd7\x42\x3b\x2c\x99\x75\xe8\xac\x7e\x25\xd2\x91\x86\xaa\x32\xcd\x3a\xc2\x92\x17\x36\xf4\xc9\x5f\x21\x74\x16\x34\x47\x36\x96\x90\x8f\x0c\x53\xad\xb2\x0a\x75\xb3\xef\x45\x9f\x99\x30\xb3\x43\x92\x16\x9a\xdc\x91\xa7\x84\x81\x2d\x28\x1b\xd6\x25\x71\x08\xbb\xfb\x0e\x49\x75\xdb\x23\x6b\xa2\x0b\xec\x86\xf0\x73\x3b\x07\x9f\x81\x8b\x05\xb9\x88\x6b\x48\x03\xca\x02\x06\x02\x1c\x55\x5b\x7a\x26\x0e\x61\x88\xf8\xe8\x72\x43\xec\x4d\xec\xd0\x4a\x3f\x45\x7a\xcf\x46\x27\x25\xef\xe5\x9b\x04\xeb\xee\x23\x39\xe9\x50\xf7\xf6\x93\xee\x90\x13\x23\x55\x4c\xeb\xc2\xac\x88\x0b\xe6\x09\x91\x95\x9a\x66\xc4\x73\xe2\xc2\xa0\x93\x41\x36\xf9\x6d\x36\xdc\xb5\xd9\x20\x47\xe5\x8b\x98\x08\x91\x7a\x69\x67\x42\x10\xba\xee\x13\x06\x4e\xd1\x8b\x37\x68\x6f\x09\x25\x49\x07\xbe\x20\xd8\xe5\x28\xe4\xe0\x20\x27\x64\x51\x08\x23\x14\x1b\xf9\xdc\x8e\x96\x3b\x74\x49\x84\x1a\x47\x4e\x43\x66\x43\x3c\x5b\xa2\xe0\x85\xd4\x88\x42\x58\xc6\x34\x3f\x42\x0e\xac\x14\x4c\xd2\xf5\x08\xae\x02\x06\x3c\x0a\x0d\xd8\x14\x98\x4d\x96\x2e\x20\x41\x91\x52\x0f\xa5\xa8\x6d\x5c\x8a\x8c\x8f\xde\x9b\x08\x30\xe7\xd2\x17\xb8\xcd\xee\x54\xf4\x43\x3f\xf4\x29\x72\xba\xee\x27\x2a\xd1\xdd\xa3\xf0\x43\x6f\x09\xac\x69\x17\x58\xad\xd5\x3f\x82\xe1\xba\x91\x9f\xde\xde\x47\x95\x04\x23\xed\x8f\x1e\x3e\x34\x38\x6d\x6a\xdb\x5c\x28\xd2\xbb\xf9\x86\x91\xae\x02\x96\xb7\x22\x7a\x37\xde\xbf\xee\x08\x8c\x24\x18\x0b\x18\x93\x37\x7b\x0b\xc0\x80\x1f\x7a\x5d\x70\x91\xf5\xc7\x82\xc5\x14\x30\x68\x0f\x4b\x42\xa1\x80\x68\x96\x7e\x45\x99\x87\xcb\x6b\x5d\xdb\xdd\x6e\x3a\x81\x75\xfb\x7d\x1d\x90\x6f\x95\xed\xe1\x91\x9d\x57\x5d\x8c\x23\xe0\x46\x0e\xa5\xa5\xbb\x52\x1e\x0f\xdf\x69\x75\x77\x9c\xb0\x2c\x95\x2c\x3a\xee\x7f\x3d\x7c\x15\x87\x2d\xba\xc4\x95\x24\xc9\x58\x5a\x62\x50\x92\xf2\x90\x97\x82\x47\x9d\x85\x50\x24\x23\x09\xf1\xb8\x8f\x10\xa1\x2b\x48\xe0\x42\x37\x3b\x96\x51\x8d\x15\x05\xeb\x21\x8a\x4f\x2b\x73\xae\x4e\x06\x9f\x8a\xb1\x94\xe9\x87\x56\xc1\x93\x1a\xbb\x9a\x17\x2b\xb1\x1b\xad\x05\x8b\x08\xc6\x12\xcd\xa4\x63\x9f\x6b\x91\xe9\xe4\x9a\x6b\xdc\x26\xf3\xd6\x33\x5f\x6e\x72\xbb\x5f\x29\xf3\x3a\xe4\xf6\x33\xd6\xe8\x3b\xbf\x01\x35\x2f\x82\x43\x6c\xf7\xe2\x75\xeb\x56\x33\x44\xb5\x2a\x78\xb7\x36\x95\xc5\x41\x28\x6e\x2b\x95\xe2\x35\x6d\x0f\xaa\xb5\x6e\x92\x25\x7a\x63\x52\xd7\xa6\x25\x3d\x23\x1c\x6b\x21\x31\x79\xc5\x7d\xf3\x45\x3d\x45\xcd\x13\x8e\x25\xaa\xc9\x36\xf7\x12\x35\xa7\x78\xad\xa5\x4c\x68\xc6\x12\xb0\xbc\xae\xf6\x77\xe5\xab\x41\x2b\xbd\x2b\x6f\xb4\xc3\xb5\x38\x04\xa3\x0f\xf5\x81\x01\x89\xd2\x9c\x45\x35\x76\xa6\xd5\xb2\x5c\x93\xf7\x32\x00\x86\x85\x00\xd6\x73\x96\x54\x88\xc7\x82\xef\xff\xee\x2a\x7c\x2d\x1a\x6b\x99\x1d\x1c\x1f\xc4\x1f\xfb\xa5\x73\xbb\x2c\xe4\x83\xe4\x82\x1b\x77\xa3\x0e\x04\xe0\x3b\xe0\xdb\x1d\xc1\xce\xd3\x8d\x05\xf2\xbd\xc8\x99\xd7\xa5\x02\xba\x27\xc7\xbf\xc8\xe0\x51\x39\x96\xb2\xed\x87\xae\x5b\xf5\x72\x63\x03\x54\x78\xbc\x68\xb6\x1a\x5f\x45\x80\xb6\xea\xb2\x36\x03\xf3\x55\x04\x68\x7b\x00\xf3\x75\x04\x15\x7a\x00\x73\xdf\x82\x48\x35\x22\xde\x20\x56\xa2\xb8\x9a\x63\x25\xf9\x72\x53\xac\xe4\x75\x54\x67\xd0\x58\x49\x5c\xef\xde\xc4\x4a\x74\xd6\xfc\xe6\xb1\x12\xc5\xf5\x76\x63\x25\xb5\x33\xeb\x6e\xc5\x4a\x8a\x83\x50\x8c\x95\x70\x6a\x9f\x43\x8d\x9e\xe7\xcb\x1b\xb5\xb6\x34\x5a\x27\x11\x6d\xad\xfe\x9b\xf4\x5c\x35\x7b\x47\xf4\xbc\xbb\x82\xaa\xee\x7f\x11\xa7\x37\xe2\x91\xd0\x2b\x4f\xe4\xed\x36\x05\xda\xaa\xb5\xbe\xe5\xe1\xe3\xc7\x86\xf7\x45\x2a\x80\x7d\x73\xf3\xfa\x02\xf3\x45\x6e\xa5\x5a\x07\xef\xf4\x79\xf8\xb2\x99\x0e\x03\x60\x1c\xa2\x83\x72\x05\x2c\x14\xf5\x28\x68\x94\xc3\x4f\x5d\x8f\x06\x38\x58\xc0\x4c\x10\x0f\x1a\x0f\x07\x14\xe3\x15\x09\x19\x52\xb2\x0d\x2b\x93\xf5\xa8\x9c\xb5\xd5\x0d\x5a\x87\xb3\x06\x99\x94\xa5\xb2\x45\xdd\x7a\x55\x83\x9a\xb4\xe5\x6c\x16\x9d\x7e\x9b\xc9\x19\xd6\x04\xde\x31\x52\x24\xf1\x81\xb9\xe4\x55\x41\x40\x98\xa3\x68\x62\x82\x83\x96\xd7\xe8\x74\x4d\xc4\x26\x5c\x5a\x36\xf5\xe6\x8a\x60\xee\x10\x29\xee\x32\x94\x9c\xe6\x29\x5d\x86\x77\x03\x85\x60\x00\x49\xc1\xbe\xb5\xff\x28\x63\x31\x2c\xc0\x65\x40\x86\xc1\x19\x3c\x4c\x34\xa1\x8d\x5a\xbb\x23\x49\xc6\xd2\xca\x87\x83\x82\xa6\xa4\x1b\x06\xa9\x0d\xe5\xa2\x74\x40\xb0\x05\x58\x09\xd5\x58\x78\x3d\x1a\x14\xaf\x54\xc6\x61\x20\x23\xc1\xc5\xe3\x6e\x70\x49\x8a\xb1\xa0\x7a\x3c\x28\x54\x91\x6c\x83\xc1\x74\xd0\x19\xa6\x83\xb1\x60\xfa\x61\x68\x98\x0e\x06\x82\x29\x64\xa4\x1b\x4a\x21\x23\x63\x81\x74\x30\x28\x48\x52\xb2\x61\x30\xe2\xe0\x5d\xb4\x38\x86\x78\x8c\x38\x78\xd8\x17\xc4\x46\xf1\x4b\xe4\xe5\x65\x52\x31\x92\x18\x29\xec\x0e\xe7\xf3\xec\xd1\x7c\x50\xe9\xe3\x3e\xd7\x03\xb0\xa5\x2b\x29\x9d\x52\xfc\x0d\xfc\xb5\xd8\x74\xcc\xf4\x2b\xa2\x91\xfc\x68\x53\xe6\xb7\x21\xc9\xbf\xaf\x97\x30\xe9\xec\x5d\x92\xd0\x94\x9c\x6d\x3a\xc6\xb0\x5b\x14\x20\x89\xce\xe9\x8f\x37\xdc\xb7\x38\x6f\xcd\xe6\xef\xeb\x8b\xeb\xf7\xd8\x09\xc7\x47\x0b\x7a\x9c\x46\x18\x09\x9c\x72\xf6\xbc\xc6\xd8\x65\x1b\xd9\x6d\x06\x6b\xb8\x1a\x24\xc4\xaf\xda\xa9\x09\x7d\xe6\xca\x3b\x87\x3e\xd5\x4b\x36\xbd\x42\x9f\x4a\xfc\xbb\x11\xfa\x6c\x11\xe2\x1f\xe7\xed\xb7\xf8\x2d\xa5\x5b\x0d\xf1\xd7\x4e\xb2\x3b\x16\xa5\x2d\x0c\x42\xe9\x2d\xbb\xb2\x86\x97\x10\x7f\xd3\x27\x3f\x55\x7b\xc8\x69\xfb\x74\xf6\x87\x85\x67\x7f\x1f\xcf\x7e\xdf\x9b\xfd\xb4\x78\xd0\xf3\x45\x90\x9a\x6b\x4c\xde\x64\xd7\xfd\x18\x86\xbc\x25\xb7\xf4\xfd\xf3\x01\x78\x65\xaf\xee\x0e\xc0\x2c\x7f\x18\x7b\x00\x76\xf9\x7c\xe5\x00\xec\xf2\x69\xa1\x21\xd8\xe5\x4c\x6d\x1b\x17\xb5\xbf\xd1\x2f\x1f\x00\xd2\x19\xfe\x72\x1d\x93\x91\xcf\x26\x8e\xad\xab\x5d\xd5\xf2\x7f\x9a\xbd\x1c\xf3\x25\x3d\x39\x64\x5b\xf0\x69\xf7\x6a\x45\x0b\x46\x75\xe9\xf5\x4e\x8c\xea\xfc\x37\xad\xdd\xe2\xf6\x5a\xbf\x2e\xdb\x6b\xe3\x90\x9c\x9e\x08\x2c\x37\x60\x36\x76\x5d\xb4\x66\x38\xd8\x64\x4e\x0b\xf8\xd6\x25\x39\x27\x01\x38\x44\xdd\x5b\x27\x7f\xcd\x9f\x60\xd7\xfd\x23\xaa\x39\xd5\xd8\xbf\x5e\x63\x68\x53\x5f\x60\xe2\x03\x93\xbc\x7b\xe3\x1e\xdc\x84\x5a\xda\x7c\xd7\x05\xf7\x26\x3c\x38\x30\x82\xcb\x1c\xb4\x23\x55\x14\x58\x37\x66\xc5\x1a\xbd\xf3\x7f\x29\x9b\x2e\xee\x91\xed\x95\x8f\x70\xeb\x34\xe7\x09\xf5\x3c\xec\x3b\x88\x85\xbe\xdc\xab\x63\x94\xb6\xf5\x33\xa2\x17\xc0\x18\x71\x80\x23\xec\x5f\x23\x0e\x02\x61\x11\x79\x29\x2a\x2c\xee\xc2\x05\x68\xc2\xbd\x66\x5f\x1f\x99\xfd\x7d\x5d\xd7\x3a\xbf\xc9\x5d\x3b\xac\xba\xf7\xb9\x8b\x83\x1b\xff\x2a\xbb\x82\x84\x69\x7d\xa7\x9a\x93\xc3\x3a\x61\x92\xfb\x5c\x08\x70\x44\xfc\x08\xc5\x6c\x54\x2b\xc4\x4d\x87\xa9\xe3\x6a\x1f\x27\xa7\xca\xdd\x58\x1c\x4e\x8f\xa4\xf3\x71\x76\x36\xcf\xf9\x1f\x3b\x5a\x2a\xa3\x23\x92\xfc\xe9\x48\x74\x22\x4d\x2e\x89\xeb\xa2\x25\xa0\x25\x0d\x7d\x27\x1a\x19\xec\xa5\x97\x4c\xa0\xe0\x7c\x5d\x0d\xed\x54\xe0\x8b\x4e\x43\x6a\x2b\x7d\xd2\xd3\xb6\xed\x9d\x49\x7b\x1c\xc2\x94\xea\xa0\xff\x99\x53\x86\xb8\x4d\x83\x28\xc3\x13\xf5\x1f\x04\x0a\x03\xea\x23\xb8\x22\xd5\x21\x4d\x5b\xea\xaa\x60\xb1\x3c\x9a\xa7\x8b\xca\xb3\x72\xad\x0a\x0a\xed\xbc\x6c\x0d\xe9\x36\xf8\x17\x1f\xf0\x20\xca\xfc\xcc\xbf\x20\x8c\xfa\x1e\xf8\x02\x5d\x60\x46\xf0\xd2\x1d\x54\xad\x4f\x3f\xfe\x72\x0b\xda\x4b\xfc\x9c\x36\x5c\xce\x95\x36\xfb\xd8\xd3\xe4\xde\x2a\xc0\x7d\x76\x35\x6e\x30\x82\x31\xb3\x7b\xaa\xa9\x72\x47\x31\x84\x9e\x3e\x27\xc3\xea\xe5\x37\x73\x6b\xec\x9d\x49\x4f\xa3\xeb\x7c\xbe\x4e\x7b\xab\xae\x0a\xee\xa1\xc5\x35\x5e\x1d\x52\xb1\x90\x72\xa4\x37\x2d\xaa\x5c\x11\x91\x9c\x5c\x10\x34\xbd\x00\x48\x0b\x74\x0f\x90\x35\x1a\xa3\xbd\xda\xa8\x45\x33\x19\x59\x8f\x01\xd2\x86\x86\x0a\x60\x95\x9e\x2e\x86\x1b\x63\xcd\xfd\x48\x48\x33\x0e\xd1\xa5\xdb\x05\x23\x84\x6c\xec\xcb\xc9\x9c\x1e\xf0\x88\x52\x65\xd1\x7d\x5c\x54\x6c\xd4\xde\x5a\xd5\xe4\x8d\xf9\xb1\xda\x1e\x06\x94\xe9\x83\x7f\xe5\x5d\xbd\xac\x17\xdb\x96\xf4\x5a\xb0\xac\xbb\x82\x46\x0f\x36\x94\xd7\xc4\x24\x8d\x0a\xdd\xce\xbe\x9e\x46\x66\x74\x32\x53\xff\x9d\x1e\x4d\x84\x1d\xfc\x27\x74\x82\xe9\x51\x4b\x75\xff\x7f\xca\x05\x92\x02\x4f\xf8\x54\xf6\x78\x49\x22\x43\xa9\x57\xf8\x86\x24\x23\x2a\x26\x0f\x2a\x9d\xeb\xa3\xa9\xbd\xd5\x4c\xc5\x9d\x7a\x2d\x89\x6d\xb1\x3f\x34\x07\x2e\xd3\x4a\x95\xbd\x63\xa2\x1d\xf1\xc1\x5f\xec\x38\xd2\x5a\x20\x0f\x07\x01\x44\x4b\x14\x4e\x8a\x74\xc7\xae\x50\x93\x2e\x8f\x8d\xaa\x70\x9e\xb1\xf2\x9e\x7a\x48\x50\x3f\x5a\x66\xa7\xc0\x8c\xa5\x70\x80\x31\x14\x30\x58\x91\xab\x22\x94\xca\xe7\xbb\xa3\x50\xbe\x0e\xdb\xbc\x12\xf1\xb9\xa1\xa4\xa1\xb8\x67\x50\x5e\x52\x76\xfe\xb4\x72\x4f\xab\x4e\xd0\x7f\x53\x76\x2e\xa5\x70\x72\x77\xc5\x8a\x0d\x9a\x14\xe3\x36\xb9\x93\x18\xd1\xf2\xdf\x7c\xde\x62\xcb\x24\x69\x31\xdd\x63\x5c\x77\x63\xcf\x27\xf7\x6c\x31\x44\x9e\x48\x9f\x02\xca\xdc\xfa\xad\x52\x5b\x1d\xde\x7a\x0a\x8c\x11\xbc\xb8\xa8\x77\xe8\x8e\x06\xe5\x98\x5d\xdd\x2b\xe7\x75\xf1\xbc\xe0\xbc\x9c\x1d\x6d\x62\xd7\xc4\x12\xdd\xc4\x99\x2c\xdc\x21\xa9\xed\x51\xdd\xc2\x9a\x25\xb6\x43\x46\x66\xa9\x13\xf4\xcd\xc9\xd4\xb4\x5e\xfd\x90\x48\x5a\x52\x4e\x22\x4b\x51\xb1\x4b\xfe\x06\x8e\x5e\xbc\x7a\xf3\xfe\xdd\x1f\xaf\x8e\x5f\x3e\x53\xee\xdc\x87\xe3\xdf\xde\x3f\x93\x1b\xc5\xf8\x78\xfb\xf7\x59\x85\x43\x55\xf8\xbd\x85\x5e\xac\x92\x7a\x1c\xc9\xad\xe2\x2e\x22\x22\xfb\xb4\x0a\xe7\xa1\x07\x4e\x5c\xe3\x17\xb4\x33\xc9\x58\xd4\x18\x95\x9b\x3a\x26\xb5\x39\xd5\xb4\x5a\xcf\x3d\xf6\xf0\xfb\xe2\x7b\xba\x59\xad\x7e\x71\x26\x2d\xaa\x51\xb2\x4c\xbd\x5e\xbf\x7f\x97\xea\x5b\x4e\xc9\x94\x7a\xe5\x0a\x95\x92\x15\x6a\xd7\xa8\x5a\x54\x41\x6a\x5a\x8e\xe0\x9b\xaa\x95\x39\x9a\x9c\x14\x74\x2b\x8a\x65\xf4\x17\x5a\x1e\x7f\xa8\xb7\xb4\x72\xd5\x29\x5b\xda\x3a\xc5\xee\xc8\xbe\xf9\x65\x88\xcf\xe3\xd5\xd0\xe0\x06\xee\x4c\x21\x0b\xaa\x73\x6a\x0a\x15\x7a\xbb\x36\x09\x17\x93\x83\x73\xe3\xc3\x7e\xb6\xe9\x55\xaf\x96\x20\xa6\x1d\xec\x0f\x65\x2e\x19\xac\xcd\xc9\x67\xc5\xbd\x61\x54\x3c\xee\x2c\x88\x71\xf7\xfa\x43\x98\x5b\xfa\x74\x10\xe6\x8a\x4d\xa7\x1b\x8e\x7d\xf5\x1d\xa9\xdd\xf8\xb6\xc0\xdd\xf8\x4d\xf8\x5d\x44\x59\xbc\x93\xb3\x90\x3a\x7a\xc8\xd3\x25\x44\x5d\x73\x4f\x5d\x2c\xc0\x41\x58\x20\x16\xfa\x82\x78\xb0\x8b\x18\x04\x2e\x8e\x3e\xfb\xb6\x33\x79\xfb\xec\xf9\x54\x6e\x8c\xc4\x26\x09\x52\xd3\x15\x7a\xfb\xec\xb9\x25\xff\x2f\x5b\x8d\xe2\xa8\xf5\xce\x24\xca\x06\xc5\x17\xa8\xa3\x15\x47\x3b\x93\xb9\xdc\x72\x4d\xe5\x46\xd2\x42\xaf\x80\x0b\xc9\x37\x0c\xa4\x87\x4c\xfd\x38\x45\x9e\xf2\xe1\x61\x10\x50\x26\xc0\xd9\x45\xc4\x02\x4b\x52\xc7\x2c\xa7\x15\xf5\x39\x4d\x46\x3e\xbd\xf3\x6f\x37\x53\xa9\xd4\xe4\x97\x4e\x8e\xa8\x0f\xae\xe9\x15\x35\x2a\x32\x9e\x1f\xb9\x78\x68\xed\x59\x7b\x95\x93\xfc\xba\xf3\xfa\x3c\x00\x7b\xae\xea\x5b\x1b\xe1\xb9\xd5\xbe\x97\x3d\xff\x7c\x18\xed\xe3\x24\x0e\xa0\x9d\x9d\x59\x9a\x7f\x4e\x8e\x0e\x27\x67\x67\x51\x90\xed\x78\xf6\x3b\x9e\xfd\x3d\x5b\x3c\x98\x1c\x1d\x9e\x9d\x59\x85\x47\xd3\xff\x9d\x4e\x8f\xa2\xe7\x0f\x72\xcf\xcf\xce\x66\x67\x67\xd6\xe2\xc1\xf4\x68\xa7\xf8\x71\xbb\xf4\xd0\xa7\x0e\x9a\xb4\xd0\x04\xce\xcb\xb8\x82\xf4\x6b\x4e\x2f\xf6\xac\x87\x3f\xa2\x27\xd4\xf3\xa8\x2f\x0b\x10\xbf\xf6\x05\xbe\xca\x80\x0a\xc0\xb6\xec\xa8\x58\x32\x8e\x10\x93\x24\xf3\x29\x22\xbe\xed\x86\x8e\x54\x90\x5f\x9f\xbf\x44\x02\x2f\x5d\x40\x70\x25\xc0\x2f\x2a\xbf\xf6\x93\x82\x5b\xf2\x7f\x9f\xb6\xb6\xfe\x1b\x00\x00\xff\xff\xc3\x20\x9a\xba\x9b\x72\x00\x00")

func githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecOpYmlSchemaJson,
		"github.com/opspec-io/spec/spec/op.yml.schema.json",
	)
}

func githubComOpspecIoSpecSpecOpYmlSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/op.yml.schema.json", size: 29339, mode: os.FileMode(420), modTime: time.Unix(1511328885, 0)}
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
	"github.com/opspec-io/spec/spec/op.yml.schema.json": githubComOpspecIoSpecSpecOpYmlSchemaJson,
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
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"op.yml.schema.json": &bintree{githubComOpspecIoSpecSpecOpYmlSchemaJson, map[string]*bintree{}},
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
