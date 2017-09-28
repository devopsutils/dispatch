///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-bindata.
// sources:
// root/go/src/gitlab.eng.vmware.com/serverless/serverless/swagger/function-manager.yaml
// root/go/src/gitlab.eng.vmware.com/serverless/serverless/swagger/identity-manager.yaml
// root/go/src/gitlab.eng.vmware.com/serverless/serverless/swagger/image-manager.yaml
// DO NOT EDIT!

package gen

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

var _functionManagerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4d\x6f\xdb\x46\x10\xbd\xf3\x57\x0c\xd0\x00\x4a\x51\xd8\x52\x8c\xe4\xc2\x53\x5d\x5b\x0e\x04\xb8\x4e\x40\xab\x87\xb6\x6e\x81\x15\x77\x48\x6d\x4c\xee\x32\xfb\x21\x45\x48\xfb\xdf\x8b\x5d\x7e\x8a\x26\x25\x4a\x96\x8d\xea\x24\x92\x33\xb3\x6f\xe6\xbd\x19\xee\x52\xad\x49\x1c\xa3\xf4\x61\x74\x71\x3e\x19\x79\x8c\x47\xc2\xf7\x00\x28\xaa\x50\xb2\x4c\x33\xc1\x7d\xf8\xc7\x03\x00\x98\x2f\x99\x02\xa6\x40\x2f\x11\x2e\x3f\xcf\x40\xa1\x5c\xa1\x84\x48\x48\x77\x2b\xbf\x4c\x50\x29\x88\x0c\x0f\xad\x27\xa4\x84\x93\x18\xa5\x7b\xc6\x42\x3c\xf7\x00\x56\x28\x95\x0b\xfa\xee\x7c\x72\x3e\xf1\x00\x34\xd3\x09\xfa\x70\x53\xfa\xfc\x9a\xfb\x78\x00\xa1\xe0\x9a\x84\xda\x77\xab\x63\x4a\x58\xe2\xc3\xa3\xd2\x98\x71\x86\x6b\xf5\xc8\x7e\x5e\xa5\x6b\x22\xf1\x3c\x14\xa9\xa7\x49\xac\x7c\xef\x0c\x38\x49\xd1\x87\x7b\x2d\x24\xb6\xb3\xb8\x92\x86\x82\xc8\x50\x12\x7b\xad\x40\xf0\x0a\xa8\xaa\x3c\x03\xc3\xb9\x5b\x7c\xcb\x75\xfa\x0d\x43\xe3\xd0\xf5\xf9\xab\x70\x89\x29\x3a\x08\x4b\xad\x33\x6f\x41\x14\x7e\x26\x7a\xe9\xc3\x78\xf5\x6e\x5c\xda\x79\x19\xd1\x4b\x65\x13\x1a\xe7\x59\x65\x42\x15\xf9\x01\xe4\x29\xe4\xff\xcf\xaa\x14\xec\x4f\x99\x34\x25\x72\xe3\xc3\x25\xa5\x40\x20\xc2\x75\xb5\x72\x61\x51\xc1\x9a\x51\x1f\x08\xa5\x37\xdb\x8f\x43\xc1\x95\x71\xf0\xca\xf0\x24\xcb\x12\x16\x3a\x97\xf1\x17\x55\x19\x66\x52\x50\x13\x0e\x31\x24\x92\xa4\xa8\x51\x36\x4c\x19\xf7\x61\x21\xe8\xa6\xb8\x01\x45\x49\xb7\x6e\x6d\x95\xb5\xd2\x89\x58\x7c\xc1\x50\x57\x46\x12\xbf\x1a\x26\x91\xfa\xa0\xa5\xc1\xea\xb6\xab\x31\xf1\xab\x6b\x80\x37\x12\x23\x1f\x46\x3f\x8c\x29\x46\x8c\x33\x47\xc5\xb8\xcc\x7d\xe4\x95\xd1\x54\x26\xb8\xaa\x93\x02\x78\x3f\xf9\xd0\x0c\xb3\x05\x6a\xc6\x57\x24\x61\x14\x18\xcf\x4c\x0d\xe9\x62\x72\xd1\xeb\x51\x49\x97\x84\x21\x66\x1a\xa9\xeb\x89\x50\x22\x69\x50\xd0\x9d\xc0\xa0\x14\x62\x3c\x44\x23\xb7\x4c\x69\x20\x49\x02\xf8\x8d\x29\xcd\x78\xdc\x50\x69\x87\x58\x62\xd4\x37\xad\xe7\xcf\x16\xc1\x57\x83\xb2\xa6\x5c\x6f\x32\xf4\x41\x69\xc9\x78\xdc\x92\x86\xd2\x44\x63\xb7\x36\xaa\xa2\x36\x6d\xfa\xc3\x13\x29\x49\x5b\x78\xb6\x58\x3d\xc1\x59\xa2\x51\x36\x3b\x78\xdb\x98\x69\x4c\x55\x93\xa7\xce\x1c\x42\x91\x24\xe8\xbc\x6f\x84\x4c\x89\xf6\x61\x94\x9a\x44\xb3\x1d\xc2\xbb\x98\x4c\x7a\x65\x74\x6f\xc2\x10\x95\x8a\x4c\x52\x33\xb4\x47\x3c\x5d\xb9\x77\xe2\x1f\xa8\xb4\x1c\x53\x44\x4c\xa2\x7b\x61\x5e\x19\xa5\x45\x0a\x28\xa5\x90\xc7\x68\x7b\x6a\x1d\x47\x9e\x9d\x81\xdf\xcb\xea\xdf\x91\x14\xff\x2d\x26\x62\x4b\x53\x39\xe5\x76\x6e\x7a\x4d\x6e\x9b\x9e\x5e\x07\x4a\x7b\x1f\x44\xd4\xe0\x57\xc0\x5a\xc8\x47\xa8\x6a\xda\x39\x64\x3a\x78\xce\x88\xd6\x28\xb9\x0f\xa3\xbf\xff\x7c\x58\x3f\xd0\x87\xb3\xbf\x7e\x7a\x73\x4c\x63\xde\x30\x4e\x6b\x3c\x8b\x0d\xf4\x61\x0f\x50\x1b\xc9\x15\x10\x50\x8c\xc7\x09\xee\x1c\xf6\x8d\xfe\xfd\x65\xd3\x88\x38\xb8\x89\x5f\x5e\xa5\x03\x85\xf7\x7e\xc7\xaa\xe5\x50\x76\xb4\x2a\x63\x33\x41\xda\xf0\x7c\xbf\x7f\x38\x73\xa1\x21\x12\x86\xd3\x42\x66\x3a\x5c\x1e\x40\xde\x6f\x19\x25\x1a\x2d\x27\xbb\xc8\x30\xce\xaa\x93\x8f\xff\xcf\x2b\xf8\x74\xaf\xd1\x81\x3a\xc9\xab\xf2\x7a\x22\xd9\x7e\x73\x1f\x21\x0e\x8a\x09\x6a\x3c\x40\x1d\xd7\xce\x61\x9f\x3a\xf2\xb0\x2f\xd0\xad\xfd\x09\x36\x58\x70\xab\x37\x9b\xf5\x15\xfb\xad\x35\xe9\xc7\xd2\x70\xf5\xe2\xe3\x5e\x9a\x13\x8e\xfa\xfe\x8d\x7a\x75\x62\xd8\x52\x44\x60\xf8\x1e\x35\x48\xc3\x5b\xbb\xf4\xc1\x22\x18\x3e\x4b\x5e\x76\x44\x04\xe6\x04\xd3\x01\xab\xb3\xd5\xdb\x45\x22\xc2\x47\xbb\x63\x0d\x49\x92\xfc\x78\xcc\xc0\x68\x20\xda\xbd\x67\xaf\x4f\x74\x4a\x13\x69\xf7\xec\x6f\xb9\xe0\x67\x27\x46\x70\xc4\xec\xb1\xbf\x0f\x3b\x2a\x57\x03\x8f\x08\x4b\x2c\xee\x0e\xcc\xbd\x1b\x93\x6e\xb1\x7e\x44\x5d\x77\x8e\xed\x4d\xd0\x4b\xa2\x81\x48\x84\x05\xda\xc8\x39\x47\xd5\x04\x68\x6f\x40\x02\x73\xf8\xd9\xe1\x40\xc1\xb8\x53\x4d\xb3\xc1\x65\xbd\x66\x1f\x3f\xcf\xdf\x1d\x3f\x8f\xcc\xae\xa9\x37\xfe\x2e\xcd\x2b\x6d\x77\x25\x6a\xc9\x70\x85\x40\x6c\xb1\xec\x99\xf4\x74\xe3\xb0\x07\x66\x91\x5b\x17\x42\x5e\x20\xb4\x50\x1a\xe0\x0e\x84\x14\x15\x07\x2d\x63\x18\x3d\x85\xd2\xed\x36\x9c\x69\xe5\xd0\xf5\x8b\xfb\x55\x36\xd6\x95\x84\x02\x73\xd4\x86\xfa\x70\xad\x0a\xe9\x5e\x53\xb5\x64\x1b\xe1\xac\xf7\x9c\xc4\x79\x90\x9c\x88\xc6\x47\x9a\x4c\xda\x2a\x69\x56\xe7\xf6\x88\x9b\x7a\xc1\x0e\xe2\x56\x24\x31\xd8\x6b\x71\x6f\xcf\xf8\xcd\xb5\x1a\xbe\xc8\x4d\x5a\xf6\xc6\xec\x6e\x36\x9f\x5d\xde\xce\xfe\x98\x5e\x17\x77\xae\x82\xe9\xe5\x7c\x76\xf7\xb1\xb8\x0c\xa6\x97\xd7\xbf\x17\xff\xa7\x41\xf0\x29\x28\xfe\x5f\x4f\x6f\xa7\x73\xe7\x74\xdf\xa8\xe7\x80\xbc\x18\x6f\x83\xde\xfa\x56\x25\x8c\xee\x7d\x5e\xd6\xb9\x67\xad\x4a\xf2\x05\xc4\x4a\x82\x67\x10\x0a\x5a\xfe\x65\x29\x89\xb1\x0f\x1b\xdd\x59\xf2\x8e\x6e\x29\xfa\x74\xb7\xd7\x8e\xb6\x87\x1c\xcf\xce\x00\x16\xfc\x4e\x83\xb6\xa2\x3b\xd5\x9c\xd3\x34\xea\x68\xed\xee\xb9\xfe\x64\xa6\x77\x46\x9d\x93\xb8\x0c\xe9\x3e\xcf\x21\x9d\xb3\xa7\xf5\x60\x5c\x63\x5c\xcd\x8e\x54\x50\x16\xb1\x21\x96\xaa\x56\x71\x7f\x5e\xd6\xc6\x62\x08\x4c\x9f\x30\x9e\x12\x3d\x80\xb4\x0e\xaa\xad\xc4\x08\xfd\xc4\x93\xcd\xd6\x54\x2d\xdf\xe5\xfb\xf3\xe9\x09\x60\x93\x51\xcb\x67\x04\x70\xe7\xb3\x7d\x6d\xb5\xcf\xa4\x27\x76\xb9\x19\x6a\xfb\x2e\x84\x48\x90\x94\x83\x55\x61\x28\x51\x1f\x2a\xa9\x56\xe1\xdd\x87\xac\xc1\x1c\x76\xb5\x45\x4b\x68\xa8\x54\x7f\x73\x79\xff\x05\x00\x00\xff\xff\xa2\xa1\x67\x2f\xa9\x19\x00\x00")

func functionManagerYamlBytes() ([]byte, error) {
	return bindataRead(
		_functionManagerYaml,
		"function-manager.yaml",
	)
}

func functionManagerYaml() (*asset, error) {
	bytes, err := functionManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "function-manager.yaml", size: 6569, mode: os.FileMode(420), modTime: time.Unix(1506467560, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _identityManagerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xdb\x8e\x13\x39\x13\xbe\xef\xa7\x28\x0d\xbf\x34\x20\xe5\xc4\xfc\xab\xbd\xe8\xab\x8d\x60\xa5\x8d\x00\x31\x62\x80\x1b\x84\x90\xd3\xae\x74\x1b\xdc\x76\x53\xb6\x93\xc9\x6a\x1f\x7e\x55\xee\x53\x3a\x93\x5e\xe6\xb0\xcc\x5e\x0d\xb1\x5d\xc7\xef\xab\x43\xe3\x76\x22\xcf\x91\x52\x38\xbf\x98\x2d\xce\x13\x65\x36\x36\x4d\x00\x24\xba\x8c\x54\xe5\x95\x35\x29\xfc\x95\x00\x00\x7c\x7c\xb3\x13\x84\x70\x85\xb4\x45\xd2\xe8\x1c\x4c\x61\x25\xd1\x78\xe5\xf7\xf0\x46\x18\x91\x63\x89\xc6\xc3\xf2\x72\xe5\x12\x00\xaf\xbc\xc6\xf4\xf8\x05\x25\x00\x99\x35\x5e\x64\x3e\x8d\x5a\xb1\x14\x4a\xa7\x70\x1d\x70\x2f\x4c\x5e\xfc\xb6\x2d\xd9\xca\x2c\xb3\x65\x02\xb0\x45\x72\xd1\x83\xe7\xb3\xc5\x6c\x91\x78\x91\xbb\x34\x99\x82\x11\x25\xa6\x20\x82\x2f\x58\x77\x26\xd8\xcb\xc4\x65\x05\x96\x18\xef\x0b\xef\xab\xe4\x49\xfc\xe3\xc0\x15\x36\x68\x09\x6b\x04\x17\xaa\xca\x92\x07\x65\xc0\x17\x08\x06\x05\xc1\x26\xf8\x40\x98\x3c\x99\x36\xaf\x93\xcc\x1a\x17\x58\x0f\x24\x53\x10\x55\xa5\x1b\xfd\xf3\xaf\xce\x9a\xa4\x22\x2b\x43\x36\x76\x9b\xac\x85\xc3\x4b\xe1\x8b\x14\xe6\x49\x25\x7c\xe1\x38\xc6\xf9\xf6\xf9\x5c\x89\x32\x05\x88\x01\xe7\xd8\x44\x0e\x10\xe3\x81\x4f\xc3\x40\x3e\x37\x97\x2e\x94\xa5\xa0\x7d\x0a\x6a\x03\xc6\xfa\xc3\x70\x51\x4e\x80\x50\x2a\xc2\xcc\x83\xb7\x70\x3e\xd7\x36\x57\xe6\x1c\xd0\xc8\xca\x2a\xe3\x27\x60\x7d\x81\xb4\x53\x0e\x27\x90\x5b\x7e\xc3\x11\x17\xb6\x44\xa8\x44\x8e\x8d\x09\x5b\x21\x45\x9b\x2b\x99\x42\x54\xd1\x5c\x10\xba\xca\x1a\x87\xae\xf5\x14\xe0\xff\x8b\x8b\xfe\xc7\x11\x3f\x5a\x5f\x0e\xee\x0b\x14\x12\xc9\x1d\x8a\x00\xbc\xb6\x75\x88\x29\x0c\x8e\x47\xb4\x81\x6e\x9e\x1f\x3d\xf6\xfb\x0a\x53\x70\x9e\x94\xc9\x93\x5e\xc3\x46\x04\xed\x47\x5d\x44\x22\x4b\x07\x97\x91\x2d\x62\xe8\xde\xff\x08\x37\x29\x9c\x3d\x99\x4b\xdc\x28\xa3\x58\xd0\xcd\xa3\xe0\x59\x0f\x63\x9d\xea\x79\x25\x9c\xdb\x59\x92\xe9\xfd\x30\x0d\x0e\x89\x33\xee\x98\x8c\x3b\xe5\x8b\x78\xc2\xbc\x06\x61\x24\xb4\xda\x27\x11\xb6\x8c\x33\xc2\xca\x84\x76\xc0\x25\xb8\xb1\xb4\x13\x24\x51\xb6\xc0\xe2\xb5\x67\x69\x0d\xaa\xad\xb6\x8a\xec\x56\x49\x24\x7e\x81\xd7\x59\x21\x4c\x1e\xe5\x22\x8d\xc0\xdb\x6f\x68\x46\x49\x70\xd9\x58\x6f\x1e\x54\x82\x44\x89\xfe\x00\xcd\x29\x28\x93\xc2\xf7\x80\xb4\xef\x12\x58\xd7\x64\x1b\x45\x72\x12\x84\x18\xf5\xe0\xfa\x04\x96\xe3\xca\xab\xa1\x5f\xa7\x94\xdf\x78\x72\xc2\xc0\x09\x72\x5f\x2c\x16\xa3\xcc\x79\xfb\xea\x47\xb4\xbe\x42\x3f\x7d\x61\xed\x37\x85\xe9\x6d\xa8\x7a\x07\xf2\x31\x58\x67\x8f\xc9\xf1\x9a\xdc\x75\x0f\xbe\x27\xb5\x99\x8f\x1f\xde\xbd\x66\xe2\xed\x0a\x95\x15\xf1\xa0\x9b\x02\x97\x2d\x2f\x9f\xe2\x2c\x9f\xc1\x76\xf5\xf2\xcd\x33\xd8\x29\xad\x0f\x1a\x5a\x81\xb0\x26\xbb\x63\x3c\xc5\xc6\x63\xcd\x59\x4b\xea\xcf\x68\x10\x0a\xe1\x60\x8d\x68\x20\x27\x61\x3c\x4a\x58\xef\xa3\x0c\x13\x60\x94\xd3\x1f\x63\x48\xf7\x60\x74\x66\xe5\x08\x9b\x09\xbf\x07\x45\x28\xb9\x41\x1f\x35\x67\x76\x67\xe8\x34\xab\x81\xa7\xae\xb0\xe4\xa7\x5a\x6d\x95\xc9\x9f\x41\x8e\x86\xdd\xec\x23\x78\x5b\xa1\x59\xbd\xec\x72\x34\xe1\x90\x64\xac\xda\xd8\x06\x34\x37\x01\x9e\x3a\x37\x8a\x7a\x99\x65\x3c\x8f\xdf\x73\x59\x73\x03\x99\x5b\x82\x95\xfc\xf2\xfe\xa0\xcc\xef\x58\x6b\xce\x0b\x3f\x12\xf7\xdb\x25\xb7\x90\x8b\xd9\xa2\x7e\x04\x5b\xa1\x03\xce\x06\xe9\x60\x77\x97\x83\xf8\xdf\xe1\xf7\x80\x8e\x67\x6f\xa6\x43\xec\x5c\x05\x36\xf2\x1d\x1a\x33\xae\xa4\xb6\xa5\x45\xad\x40\x98\xa1\xda\x72\x12\xc8\x96\xf1\xfc\x45\x9d\x85\x65\x55\xc1\x53\x35\xc3\x19\x04\xf7\x6c\x76\xbf\x18\x87\xf5\x32\x8a\xed\x89\xe1\x1b\x7b\x2e\x4b\xd7\xb8\x12\xfa\x40\xa6\x0f\xfd\x77\x23\xa7\x1f\x98\xbe\x12\x8d\x42\x17\xcf\xa8\x49\x40\x03\x66\xf7\x64\x58\x4d\xb0\x11\x4a\xbb\x87\xc4\xf3\xe5\x20\x8e\xd3\xb1\xd9\xf8\x57\xe8\x09\xfc\x11\x4a\x61\xa6\x84\x42\x8a\xb5\x46\x58\x5e\xbd\x58\xad\x00\x0d\xc7\x24\xc1\xe3\xb5\x3f\x14\x04\xbb\xe9\xe3\x7e\x90\x87\x81\xd4\x8f\x3c\xfb\xf0\x6e\xc5\xf6\x04\xec\x70\x1d\x37\x16\xf0\x85\xe8\xd8\xe3\x40\x48\xa9\xea\xb7\xc0\x1b\x2b\x95\x75\xf6\xc4\xda\x06\x7f\x4b\x2f\xef\x38\x06\x5c\x88\x35\xb6\x09\x5a\xef\xb9\xa5\xe4\x0c\xb7\xb9\xd5\x7a\x74\x5a\x74\xb8\xc2\x1d\xaf\x66\xf0\x1f\xec\x50\xf0\x53\x67\x59\xeb\xd2\xa3\xce\xb3\x7e\x67\xb3\xc1\xdf\x73\xa0\xd5\xc2\x75\x0b\x0e\x44\xdc\x7d\xe2\xae\xb1\xde\x43\xa6\x51\x70\x3e\xb8\x1d\xb5\x9f\x02\x70\x85\x8e\x3f\x5a\xa0\x4e\xa3\x1b\x19\x49\x36\xb4\xeb\xb2\xc3\x2c\x90\xf2\xfb\x3e\xb0\x29\x64\x51\xf6\x0b\x7b\x96\xc2\xa7\xcf\xff\x1e\x65\x7b\xb3\xf0\xf8\xeb\x4c\x89\xce\x89\x1c\xcf\xe0\xb1\x28\x00\x07\x1c\xe0\x0a\x4b\x6f\x7c\x82\x75\x28\x0b\x03\x95\x16\x19\x16\x56\xcb\xbe\x1a\x27\xf5\x6e\xb2\xc6\xae\xa2\xea\xa5\x5b\x6d\xfe\xb9\x21\x0c\xd0\x66\x6d\x3f\x1b\xeb\x53\xfd\xe3\xce\xb8\x3c\x1e\x2c\x6d\x1e\x5e\xf6\x77\x69\xfc\x9f\x81\x3e\x15\x49\xcf\x38\x51\xa9\x57\x58\xcf\x92\xe3\x95\xbf\x11\x71\xdd\xb7\x4d\x5f\xc8\x13\xd8\x15\x68\xba\xd5\x10\x84\xe6\x49\x77\x8c\x15\x8f\xaa\xba\x0c\xe2\xcf\x7a\x4e\xd5\xf4\x4f\xe4\xd0\xbb\x96\x02\x87\xae\xd9\xf5\xd7\xf6\xc3\xb7\xdd\x19\xfa\x9d\x72\xd0\x7c\x2b\x62\x52\x78\xd5\x23\xda\x5e\xa7\x63\x63\x2a\x66\xeb\xd6\xd6\x1a\x14\x47\x8c\xf1\x50\x3f\x36\xa4\x8c\xc7\x1c\x7b\x20\xeb\x41\x1a\xcf\x7f\xfd\xa5\x39\x6d\xb4\x8e\xfa\x78\x8c\xd5\x81\x8b\xa7\x9c\x18\xb6\x95\x23\x5d\x03\x5b\x0f\x8e\xf8\x47\x9e\xbb\xba\x4f\xdf\xda\x9a\x92\xdd\x3f\xbb\x6f\xd8\x9b\x56\x95\x1c\x35\xd8\xf1\xeb\xf4\x83\x24\xf9\x3b\x00\x00\xff\xff\x0b\x2f\x3f\xb4\x8c\x13\x00\x00")

func identityManagerYamlBytes() ([]byte, error) {
	return bindataRead(
		_identityManagerYaml,
		"identity-manager.yaml",
	)
}

func identityManagerYaml() (*asset, error) {
	bytes, err := identityManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "identity-manager.yaml", size: 5004, mode: os.FileMode(420), modTime: time.Unix(1506549236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imageManagerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\xe3\xb6\x13\xbf\xeb\x53\x0c\xf0\x5f\xc0\x87\x3f\x12\x7b\xb7\xc1\x1e\x74\x6a\x52\x3b\x81\x80\x34\x09\xbc\xce\xa1\x6d\x5a\x80\x96\xc6\x32\x13\x8a\xd4\xf2\xe1\xd4\xd8\xf6\xbb\x17\xa2\xde\xb6\x24\x3b\x89\xb3\x79\xac\x7d\x12\x29\xce\x83\xc3\xdf\x8c\xe6\xc7\x44\xdd\x93\x30\x44\xe9\x42\xef\xd3\xe1\xa0\xe7\x50\x3e\x13\xae\x03\x10\xa0\xf2\x25\x8d\x35\x15\xdc\x85\x7f\x1c\x00\x80\xc9\x9c\x2a\xa0\x0a\xf4\x1c\xe1\xf8\xca\x03\x85\x72\x81\x12\x66\x42\xda\xa9\x74\xc8\x50\x29\xa0\x11\x09\x11\x22\xc2\x49\x88\xd2\xbe\xa0\x3e\x1e\x3a\x00\x0b\x94\xca\x6a\xfc\x78\x38\x38\x1c\x38\x00\x9a\x6a\x86\x2e\x78\x56\xe0\xd7\x54\xc0\x01\xf0\x05\xd7\xc4\xd7\xae\xb5\x8b\x11\xa1\xcc\x85\xe9\xad\xe1\xe1\xcf\x8b\xe8\x9e\x48\x3c\xf4\x45\xe4\x68\x12\x2a\xd7\x39\x00\x4e\x22\x74\x61\x4a\x14\x5a\x35\xab\xce\x5f\xc6\x28\x49\xf2\xa8\x40\x70\xbb\xec\xc0\xfa\xa7\x0a\x51\xba\x59\x6c\x45\x42\x2d\x95\xc6\xe8\x8a\xf8\x77\x1b\x25\xd3\xa5\x10\xa7\x6b\x4b\x15\xd2\x70\x4d\x23\xdc\x4a\x47\xb6\xb6\x54\xa2\xfc\x39\x46\x68\x77\x3f\xd7\x3a\x76\x92\x5d\x5d\x11\x3d\x77\xa1\xbf\xf8\xd8\x4f\xf7\x13\x13\x3d\x57\x49\x04\xfb\xc9\xdb\x34\x94\xb1\x50\x59\x50\x01\xd2\xf0\xa5\xcf\x07\xb5\xf0\x25\x3f\x65\xa2\x88\xc8\xa5\x0b\xc7\x41\x00\x04\x38\xde\xdb\x25\x45\xac\x92\x9f\xc8\x9d\xf4\x02\x17\x48\x10\x9c\xac\xe8\xf0\x05\x57\xc6\xba\x99\x5b\x21\x71\xcc\xa8\x6f\x65\xfa\xb7\x4a\xf0\xec\x45\x2c\x45\x60\xfc\x6d\x16\x12\x49\x22\xd4\x28\x2b\x4b\x29\x77\x61\x2a\x82\x65\x36\x01\x39\x1e\xaa\x53\xb5\xd8\x9e\x14\xfb\x00\x31\xbd\x45\x5f\x17\xcb\x24\x7e\x35\x54\x62\xe0\x82\x96\x06\x8b\x69\x1b\x6d\xe2\x16\x63\x80\x0f\x12\x67\x2e\xf4\xfe\xd7\x0f\x70\x46\x39\xb5\x07\xd5\x2f\xb6\xdf\x73\x72\x75\x2a\x16\x5c\x95\xfb\x02\xf8\x34\xf8\x58\xd5\x53\xf3\xcb\x97\x48\x34\x06\x95\xd7\xeb\x86\xb7\x35\x0d\x70\x34\x18\xb4\x5a\xf2\xf8\x82\x30\x1a\x00\xe5\xb1\xd1\x8f\xb1\x37\x92\x52\xc8\x5e\x25\xbc\x33\x62\x98\x6e\xb5\x77\x86\x1c\x25\xf5\x01\x13\xb1\x22\x2c\x4f\x34\x1c\xe2\x43\xa1\x7c\x4e\x95\x06\xc2\x18\xe0\xdf\x54\x69\xca\xc3\x0a\xa4\x55\x13\xa6\x43\xd4\x45\x64\xd5\xae\xb0\xfa\xd5\xa0\x5c\x05\x6b\x96\xe0\x1b\xf1\x9a\xad\xeb\x33\xc2\x43\x53\xee\x10\x40\x2f\xe3\xa4\x2c\x69\x49\x79\xb8\xd1\x5a\x12\xae\x66\x53\xa7\x94\x69\x94\x79\x9d\xcc\x8c\xd6\x96\xa7\x86\x88\x94\xa4\xd4\x4a\x35\x46\xaa\x7a\x78\x0d\xde\x24\xd5\x80\x31\xf4\x13\x3b\xa7\x42\x46\x44\xbb\xd0\x8b\x0c\xd3\xb4\x33\x5b\xda\x31\xac\x8c\xef\xa3\x52\x33\xc3\xca\x33\xdb\x80\xa8\x26\xdf\x1b\xfd\x6f\x85\x5f\x43\x9e\xbd\x00\xf6\x6d\x41\xef\x7f\x2b\x70\x7e\x41\x22\xfc\x37\x2b\xf0\x2b\xa0\x4b\x41\x90\x7c\x0a\x9c\x2a\x02\x6a\xa2\x4e\x83\xd7\xc9\x3c\x88\x59\x0d\x06\x02\x24\x6a\x23\x79\x71\x5e\x0d\xc5\xb2\xe1\xe4\x63\xa2\x35\x4a\xee\x42\xef\xaf\x3f\x6e\xee\x6f\x82\x9b\x83\x3f\xff\xff\xe1\xb1\x29\x7c\x4a\x79\x50\x75\x6a\xba\x84\xb6\x2d\x8c\xad\xb3\x0a\x08\x28\xca\x43\x86\x1b\xbe\x5f\xd5\x5c\x3f\x59\x56\x94\x6e\x9d\xf0\xcf\x0f\xe0\x5d\xd6\x7e\x6f\x08\xca\x24\x1b\x79\xdc\x17\x67\xe5\x0b\x70\x34\x38\x6a\xb5\x58\xa9\x5f\x5c\x68\x98\x09\xc3\x77\x61\xf2\x45\x3e\x3a\xb1\x79\x28\x62\xaf\xe3\x80\x68\x4c\x60\xd8\x8d\x3f\x63\xd7\x35\x43\xf0\xf5\x34\x52\x3b\x6d\x85\x5e\x6d\x6e\xec\xaa\x2f\xea\xca\x0a\xef\x7d\x24\x44\x80\x0c\x35\x3e\x30\x27\x86\x56\x68\x73\x4e\xa4\xca\xf7\x65\x79\x5f\x96\xbb\xfb\x21\xb7\xbb\x9b\xa0\x4d\x18\x3c\xc3\x94\x0b\xd4\xba\xff\x9a\xa7\x05\x5d\xe8\x24\x08\xcf\x4f\x0e\xd6\x9a\xfd\x9a\x97\x35\x62\xb0\xbe\x76\xe7\xc4\x60\xcf\x09\x56\x60\xf8\xf2\x7c\xa0\xeb\x5e\xa7\x11\xfb\xe5\x9d\x4e\xd7\x75\xce\x1b\xb8\xca\xf1\x76\x7f\x8b\xf3\x42\x37\x38\x3f\xd6\xed\x4d\xff\x1b\x7d\x1c\x71\xa5\xdb\x90\xd6\xef\xcf\x57\x1b\xb3\xcc\x72\xd5\x82\xa6\x7a\xc3\x26\x97\xd7\x48\x6a\x17\x3f\x7d\x7b\x4d\xd0\xab\x6e\x80\xde\x49\x07\xde\x4a\x49\x1b\x41\x59\xd0\x51\xbe\x89\x8a\xbe\x03\x1a\xfa\x14\x0a\x9a\x06\xe1\x2d\xde\xc6\xff\x00\x98\xef\x62\x9d\x8d\xb0\x2f\x18\x67\x07\xec\x53\xa5\xfb\x22\xbb\x07\x5c\xdd\x70\x65\x2a\x11\xb2\xb3\xa9\x74\xda\x3d\x54\xda\xcf\xa2\xc7\x28\xa0\x12\xa1\x52\x39\xde\x62\x99\x9c\xb3\xa6\x25\x40\x7c\x11\x60\xe9\x48\xaa\x8e\x72\x8d\xe9\x5f\xc1\xd3\xdf\x2c\x63\x46\x94\xeb\xcf\x47\xd9\x6c\xa6\x75\x55\xb4\xe8\x63\x26\x24\x6c\xf1\x70\xdd\x87\x3b\x5c\xb6\xea\x49\x7e\x0b\xc2\x4c\xbb\xa5\xf3\x8c\x70\x56\xcd\x55\xc4\x91\x9b\x28\x6f\xe8\xe2\xa5\x9e\x0b\xfe\xa9\x36\xfa\x29\x1b\x71\x11\xe0\xad\xfa\xec\x00\x7c\x89\xd1\xdf\x42\xd9\x2f\x97\x17\xa7\xde\xd9\xf5\x78\x94\x8f\xc7\xa3\xe3\x49\x3e\x18\x8e\xce\x47\x76\xf0\x45\x13\x6d\xd4\x16\xea\xbc\x0b\x6f\xe2\x1d\x9f\x7b\xbf\x8f\x86\x55\x85\xde\xc5\x59\x36\x1c\x8f\x8e\x87\xbf\x65\xcf\xa3\xf1\xf8\x72\x5c\x33\x95\x08\x15\x55\x7f\x2b\x6c\xa4\xff\x06\x90\x3d\x06\xc2\xbf\x43\x79\x2d\x59\x36\xae\xb1\xf8\xf5\x13\xa3\x41\xe7\x81\x95\x90\x31\x86\x06\xd5\xf6\xb9\x5b\xaa\xa3\xf1\x85\xd2\xc5\x4e\x25\xac\x86\x86\xd6\xc4\xca\x41\x93\x2b\x8f\xcd\x94\x51\x7f\x55\xf3\x54\x08\x86\x84\xaf\x25\x3e\xcc\x08\x2b\x12\x5a\x15\x70\x69\xb5\x96\x20\x2a\xb7\xa4\x2a\x80\x68\x17\xb0\x8b\xca\xbe\x81\x28\xc1\x57\x9d\x7b\xf8\x6d\x47\x28\x85\x89\xd5\xd3\xf5\x54\xbf\x76\x5b\x6a\x69\xdc\xe5\x84\x84\xf9\x16\x33\xde\x3a\xa1\xeb\x18\x29\xab\xd1\x23\xb1\xbd\xfe\xa7\xc6\xd7\x01\xe8\x9a\x5f\xcf\x9f\x19\x7b\x9c\x7e\x1f\x9c\xfe\x17\x00\x00\xff\xff\xf5\xb7\x3e\x8d\xcf\x26\x00\x00")

func imageManagerYamlBytes() ([]byte, error) {
	return bindataRead(
		_imageManagerYaml,
		"image-manager.yaml",
	)
}

func imageManagerYaml() (*asset, error) {
	bytes, err := imageManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "image-manager.yaml", size: 9935, mode: os.FileMode(420), modTime: time.Unix(1506467560, 0)}
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
	"function-manager.yaml": functionManagerYaml,
	"identity-manager.yaml": identityManagerYaml,
	"image-manager.yaml": imageManagerYaml,
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
	"function-manager.yaml": &bintree{functionManagerYaml, map[string]*bintree{}},
	"identity-manager.yaml": &bintree{identityManagerYaml, map[string]*bintree{}},
	"image-manager.yaml": &bintree{imageManagerYaml, map[string]*bintree{}},
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
