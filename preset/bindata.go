// Code generated by go-bindata.
// sources:
// shoal/delegates.json
// DO NOT EDIT!

package preset

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

var _shoalDelegatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x99\x4f\xcf\xaa\xce\xb6\xe7\xe7\xf7\x55\xec\xec\xa9\xa7\xb5\xfe\x51\x55\xfc\x92\x3b\x00\x04\x05\x54\x14\x41\xd0\x4e\xe7\x04\x28\x40\x11\x04\x14\x45\xb8\xb9\xef\xbd\xf3\x3c\xe9\xa4\xcf\x19\x78\x4e\x72\xf7\xe3\xc0\x49\x25\x98\xe2\x93\xb5\xbe\xeb\xb3\xfc\xdf\xff\xf1\xeb\xd7\xaf\x5f\xff\xf5\xfd\xfd\xf5\xf9\x7d\x8b\xaa\xf4\xf7\xaf\xbf\x7e\xfd\x7e\x9c\xeb\xa8\xfc\x5f\x00\xfe\xfe\xdb\xff\x3f\x8c\x84\xb8\xa7\x8f\xc7\xef\xbf\x7e\xfd\x06\x6f\x12\x41\x46\x05\x24\x29\xc1\x02\x44\x8c\x20\x21\x23\x20\xa7\x28\x8e\x64\x81\x29\x07\x4c\xe6\x51\x06\x31\xff\xc7\x07\x34\xcf\xf8\xef\xd7\x74\xf8\xfe\x01\xd5\x96\x88\xad\x6f\x45\xf2\x5e\x77\x2e\x78\x28\xb7\x55\x0d\x8c\xb2\x0c\x34\xe1\x3d\xe6\xca\x7e\xb1\x2a\xdb\xd5\xb8\x76\x1e\x46\xe7\x19\x8e\x8f\x24\x08\x46\x4f\x21\xa4\x23\xb1\x75\xbe\xde\x1f\x15\x67\x0b\x23\xb8\xae\xdb\x71\x19\xb5\x2a\x56\x6c\x4d\xb4\xce\x88\xd7\xff\xf9\xd7\x5f\x7f\xad\x0f\x8f\xf7\x69\x9e\x69\x2f\x2f\xe2\xc2\x88\xcc\xd1\xef\x8c\x2c\x99\xc0\xfd\x5c\x23\xc1\x2c\x3e\xe9\x0d\x3b\x84\xca\x69\x39\xdf\xeb\xd5\x8c\x67\x8a\x47\xeb\xea\x6a\xcd\xf7\x87\xb3\x64\x15\xf7\xc4\xeb\x2c\xc3\xd0\x8e\xea\xf3\x98\x26\x2f\xf7\x75\x9a\xd4\x31\xdd\xab\x3b\xe5\x3f\xff\xf1\x2a\xaf\xba\xbb\xdc\xf2\xbf\x37\x75\x9f\xde\xbf\xee\x03\x01\xf8\x87\xd3\x5b\xda\xf5\xf5\xfd\xfa\xf7\xaf\x37\xf6\x75\xfa\x5f\xbf\x7e\x5f\x9a\xef\x6b\x4b\x64\x0a\xb9\x34\xc5\xd2\x14\x12\xf6\xfb\x6f\xbf\x7e\x37\xf5\xbd\xfb\x3a\xe1\x94\x81\x5f\xff\xfd\xfd\x88\xff\xfe\xf5\xb7\x5f\xff\x0e\x0c\xfa\x08\x86\x49\x90\x0a\x81\x23\x9c\x88\x04\x42\x90\x90\x28\x21\x24\x8e\x33\x9a\x88\x98\x64\x02\xf2\x04\x41\x91\x49\xe2\x23\x18\xe7\x92\x54\x5e\xd4\x19\x35\x7b\x5b\xc0\x75\x36\x17\xe3\x18\xbf\xb4\x89\x12\x92\xed\x45\x4e\xc6\x28\x0a\x51\xe9\x4f\x34\x63\x65\x3d\x8c\x77\xdc\x49\xe6\xc6\x34\x97\xbb\x95\x89\x22\xdb\x18\x1b\x1d\x3c\x05\xb0\xaf\xd6\x5b\x5d\x6d\xed\xe4\x58\x96\x43\x2a\xfb\xbe\xfe\xf8\x02\xe3\x65\x33\x4c\x77\x02\xb5\xe6\x52\xb5\x5b\x53\x35\xc0\xba\xdb\x65\xe9\x35\xf7\x03\xba\xbe\x81\x42\x65\x5a\x76\xf6\x22\x05\xf5\xaf\x3c\x05\x2e\x3c\x5f\xd5\x84\x01\xaa\xe6\x4b\xe9\x18\xb6\xc1\x39\x56\xec\xa1\xda\x48\x78\x3b\xb4\xcb\x35\x63\xfb\x5d\x6b\x66\xf1\x4e\xff\x19\x30\x68\x0a\xe1\x14\x61\x3e\x85\x9c\xff\x01\x18\xfc\x11\x0c\x60\x29\x86\x1c\x67\x29\xc9\xb0\xa0\x29\x49\x23\xc6\x00\x64\x52\x44\x78\x8a\x78\xc2\x32\x9c\x8a\x04\x09\xf9\x23\x98\x8d\xbc\xf1\x7b\xe5\x72\x96\x37\xaa\x94\x0d\xe2\xbe\x59\x0c\xa6\x66\x8d\x60\xa9\x2c\xd1\x31\x37\x0d\x30\xef\x63\x75\xde\xcc\x0f\x9b\x74\x0d\xac\xdb\xc4\x72\x35\xb3\x91\x78\xda\xdb\x47\x97\x44\x93\x90\x98\xf3\xe6\xa5\x24\x63\xe7\xa2\x3d\xbd\x1b\x9d\xfc\xbe\x39\x3a\xff\x02\xe3\xea\x9b\x26\xb9\x5b\xe9\xac\xf2\xc6\x9e\x87\x79\xa8\xba\xae\xee\xae\x94\xdd\x61\x87\xcc\x45\x05\x4a\x62\x5e\x82\x16\x8c\xae\x17\x38\xe5\x6e\xe9\x05\xb3\xb5\xb1\x00\xde\xf3\xed\xaa\xbe\x7e\x0a\x9f\x77\x7f\xa2\xe5\xeb\x73\x73\x14\xed\x76\x3c\x67\xfa\xab\x1e\x97\xb7\x1f\x02\x03\xf9\x14\x61\x3a\x95\xc1\x14\x01\xf9\x33\x98\x7f\xc7\x85\x7c\xee\x64\x52\x22\xe3\x58\x80\x38\x63\x69\x2c\x61\x20\xd0\x17\x10\x42\x12\xc4\x29\x95\x38\x17\x51\x22\xa4\x14\x7e\xee\x64\xc6\x30\xb6\xb0\xdc\x6d\x77\x12\x10\xa6\x7b\xb2\x41\x04\x1f\x1e\x89\x9c\xd6\x06\xac\xc5\x6d\x1e\x24\x5e\x1f\x69\xc3\xfb\x5a\x3e\x46\xb9\x0a\x2f\x97\xd4\xaf\xea\x7b\x12\x56\xaa\x1d\x3f\x67\x55\x71\xec\xb5\xdb\x1e\xd8\xb3\x22\xd6\xf7\xaf\xf7\xed\x72\xc4\xeb\x86\x91\x2f\x2e\x99\x7a\xc4\xbb\xa8\x20\x8b\x78\xad\x6d\x7a\x78\x4e\xde\x93\xfb\x36\x39\x48\xd7\xb3\x71\x59\xe1\xa5\x7a\x31\x55\x49\xcd\xd1\x80\x5c\xb7\x4a\x6e\xa1\x29\xf5\x6d\x73\xb0\xed\x2b\x79\xd4\xf5\x02\x2c\xe2\xd7\x49\x7b\x3c\x9a\xc9\x4c\xbf\xb7\x38\x5d\x9d\xe3\x2e\x8e\x9c\xa3\xf2\x33\x5c\x30\x99\x22\x48\xa7\x0c\x4f\x11\x00\x7f\x50\x30\xd2\xe7\x82\x41\x32\x07\x2c\xce\x38\x26\x49\x2a\xc5\x50\x06\x5c\x8e\x25\x06\x01\xa3\x52\x4c\x41\x9c\xca\x34\x8e\x58\x1c\x7f\x04\xa3\xcf\xac\xa7\xba\xba\x05\xb8\x5d\x9f\x61\xb0\xcf\x54\x34\xe8\x3b\xa7\xf7\xa5\x7a\x67\xe1\x00\x55\x93\xa2\x71\xf9\x62\x5d\x6f\x06\x73\x81\xaa\xc5\xf6\x15\x3b\x4f\x47\x79\x49\x7e\x60\x4e\xfc\x8c\x92\x6e\x39\x21\x59\x53\xd3\xa3\x8b\xf5\x63\x6f\x99\x70\xff\x28\xf4\xfe\x0b\xcc\xfd\xa4\xf4\x62\xd4\x26\xcd\x9a\xd8\x82\x99\xb3\x7b\xfd\x3e\xf6\x5e\x7d\xba\xa9\xa0\xa6\xa1\xb3\x0f\x1e\xc1\x44\xf8\x70\xbe\xb0\x82\xe4\x38\x0b\x9b\xe1\xda\xe4\xce\xf1\xa1\x64\x07\x65\x28\x64\x59\x48\x27\xaf\xee\xad\x79\xda\xbe\xee\xce\x68\x88\x16\x6f\xc8\xf5\xa6\xfc\x58\xc4\x20\x80\xa6\x9c\x4e\xe9\xbf\xa8\x97\x7f\xcb\x85\x7e\xe4\x92\x00\x29\x93\xa9\x24\x81\x8c\x08\x91\x45\x22\x86\x82\x52\x92\x45\x98\x7d\x25\xff\x57\xf8\x64\x14\x24\x42\xfa\xcc\x05\xb6\xe7\xd0\xf6\x76\x5e\xd7\x59\xf6\x89\xee\x9f\xc5\xe6\x4d\xf6\x0a\xb9\xba\x21\xf4\xac\xd8\x3f\x2c\x1d\xd4\xca\xa5\xa5\x2f\x0e\xbc\xda\x1a\xc7\xd9\xfe\xf2\x3e\x5f\x8a\x45\x99\xab\x97\xba\x80\x05\x1b\x03\xb9\x8d\xf4\xa2\x0f\xc3\xc3\xc2\xbf\xeb\x07\x24\x56\x7b\xf0\xc5\xc5\x59\x57\xd0\x2e\x57\x69\x50\x90\x07\x88\x86\x45\x76\x3c\x6b\xf4\x79\xe8\x17\xa7\xf5\xe2\x76\xc4\xcb\xa8\x74\x57\x73\xbd\x27\x38\x1c\xea\xe3\x5d\x28\xf1\xc8\x79\xba\xb0\x4b\x78\x20\x3e\x94\xf1\xbb\x8d\xb4\xb5\x42\x6f\x43\x66\xb9\x77\xc6\xbb\x5a\x22\x50\x9d\xf4\x3f\xc7\x85\x48\x53\xc8\xc9\x94\x49\x7f\x00\x86\x7d\x04\x43\xb1\x1c\x03\x59\x16\x94\x8b\x54\xc0\x28\x01\x14\x64\x88\xb0\x84\x64\x1c\xa5\x92\x04\x04\x43\x3c\x01\x31\xfc\x08\xc6\xcc\x4e\xee\xac\x6d\xec\xc3\x06\xb2\x45\xd9\x75\x9d\x49\x87\x70\x21\x0d\x45\x24\x17\x2c\x7d\xbb\x59\x83\xba\xe8\xf6\xc4\x60\xb7\xdb\xf4\x6e\xfb\x68\x4e\x24\x91\x82\x9b\x2f\x4e\xfb\xf3\xb3\x5b\x03\xc9\x4b\xd9\xc3\x68\xd2\x5e\xde\xe7\x0d\x2c\xec\xd7\x7e\xcb\x60\xfd\x05\xa6\x48\x4f\xe6\xd8\xd8\x0b\x5d\x59\x94\x5a\x89\x15\x3b\xc7\xe4\x79\x43\xd1\xe1\x84\x9a\xa7\xb8\x90\xb9\x99\x2b\xd6\x5b\x73\xca\xe7\xfe\x7a\x4a\xd5\x7a\x01\xaa\x60\xd5\x56\xa7\xdb\x25\x2d\x4e\xc9\x0d\x29\xfb\xfa\x1e\x98\xef\xe5\xe4\x9c\x70\xa3\xcc\xad\x83\x6c\x88\xf1\xa7\x66\x32\x34\xe5\x7c\x4a\xe5\x29\x84\x7f\x80\x85\x7f\xc4\x22\x58\x4a\x60\xc6\x21\x20\x09\x17\xb2\xc4\x00\x45\x00\x67\x11\xe5\x4c\x26\x49\x86\xb3\x14\x44\x51\x04\x53\xf2\x11\x8b\x72\x28\xcd\xa7\xa5\x8c\x4b\x89\xbf\x6c\x47\xaa\x72\xdb\x76\xdf\x9a\x74\xb1\x9e\x70\xd2\xbe\x26\xa2\xef\xf6\x77\x76\xb9\x16\x34\x00\x6b\x7e\xea\x92\x49\x3f\xe6\x2c\x4e\x7b\xba\x91\xbc\x97\x04\x13\x7c\xa9\x9e\xaf\x07\xdf\xce\x5c\xe5\xf2\x90\xce\xf9\xad\x18\x7a\xf9\x1b\xcb\x7d\xd5\xab\xcf\xb9\xf1\xde\x0f\x13\xd1\x1f\x19\xa0\x01\xb8\xca\xa7\x0b\x52\x1d\xc7\xf3\xbc\x83\x58\xc8\x04\x54\xda\xce\xb1\x8f\xef\xa7\xe8\xfd\xda\x5e\x0c\x1c\x6c\xb3\x6d\x1b\x04\x94\xb4\xb7\xd5\xf2\x86\x6b\x54\x52\x2f\xde\x21\xdd\x32\xfc\xfb\xfa\x36\xa9\xa4\x1f\x0a\xfe\xaf\x80\x41\xf2\xd7\x9c\x3c\x95\xff\x04\x8c\xfc\x11\x4c\x14\xa1\x48\x70\x1c\xa7\x10\x45\x14\x24\x3c\x4d\xa8\x20\x80\xa6\x2c\x4a\x22\xca\x65\x10\xa5\x22\x8a\x13\xf4\x2f\x1c\xa6\x07\x72\xba\xcd\x2b\x69\xd5\x6f\xbc\x70\xe0\x79\x34\x73\xd7\x45\xeb\xe1\x2a\x54\xb7\x4a\xde\xf0\xf9\xd8\xbc\xad\xa3\x25\xf2\xec\x3a\xc8\xfb\x44\x9c\x4f\xf1\xfb\x15\x37\x7e\x39\x0b\x94\xca\x35\xb3\x74\x67\x79\xa7\x5e\xad\x67\x93\xac\x2a\x01\xbd\x5a\xd1\x7e\xa1\x7c\x81\xa9\x69\x74\x19\x9e\xa9\x05\x60\xb7\xc8\x37\xc6\xdd\x0f\xe2\x36\x46\xbc\xf1\x6d\x3d\xa7\xf1\x42\xdd\x5a\xeb\xd4\x3d\xaa\x8d\x33\x1b\xaf\xe1\x9a\xe8\xaa\x59\x38\x65\xd0\xda\xdd\x7e\x71\x85\xfd\x36\x8b\xb7\x8a\x6c\x2e\x16\xad\x2a\x6d\x0b\x71\x58\xbd\xe6\x3e\x9b\xdc\x7f\xa8\x5e\xf0\x94\x83\x29\x92\xa6\xec\x0f\xe2\x05\x82\xcf\xf3\x98\xe0\x42\x22\x69\x24\xa7\x50\xe2\x48\x64\x19\x86\x94\x50\x84\xe3\x04\x72\x98\xe1\x28\x22\x28\x4e\x53\x9c\x7c\xa4\xa2\xf5\xed\xf5\x29\x8a\x60\xa3\xab\x06\x54\x77\x2e\x52\xc2\x4a\x2b\xbc\xfc\xb2\x32\x1b\xc0\x7b\x7f\xbe\x5f\x9d\xd8\xb9\xdf\xc6\x4d\x54\x9e\x42\x64\x55\x3b\xeb\xb5\x75\x7c\x1d\x1d\x88\x5b\xed\x77\x17\x58\xc9\xa1\xbd\x77\xe6\x8f\xe3\xc9\xd3\xab\xe0\x36\x1c\x33\xff\xfe\x6d\x96\xf7\x61\x0d\x09\xde\x89\x1d\x0f\xb2\x62\x3d\x7b\xd3\x25\x8f\xcd\x04\x82\xf3\x53\xa7\xa5\x3f\x39\xbe\x9c\x89\x13\x65\x4a\xf3\xaa\xf3\xd0\xb5\x12\xeb\x7a\x38\x97\x59\xe1\x2f\x2f\xf8\x99\xaf\x44\xf1\x36\xe3\xab\x17\xad\xb5\x33\x13\x8b\x6d\x34\x3b\xe1\x9d\x67\x77\xd1\x4f\x09\xcc\x57\xbc\xa0\x29\xa6\x53\x04\xd1\x1f\x80\xf9\xac\xfc\x5c\x44\x29\x4f\x28\x89\xb1\xc4\x23\x06\x13\x28\x18\xe7\x31\x8f\x25\x0a\x01\xc1\x44\x66\xb1\x20\x09\x42\xf8\xf3\xa0\x7c\x36\x05\x81\x68\xce\x68\xd7\x6f\x54\xc1\x7c\x68\x96\x9b\xd5\xbe\x56\x27\xd9\x2c\x38\xc4\xcb\x5b\x89\x54\xab\xa9\xa3\xfa\x14\x5c\xc2\x55\xb9\xc7\x12\xd5\x6a\xfb\xb8\x10\x82\x36\xd4\x13\xf1\x6c\xde\x6d\x8f\xbc\xca\xc0\x26\x6e\x41\x57\x9f\xc8\xba\x0a\xb6\xdf\xe5\x52\x80\x77\x5f\x6f\x8f\xe3\x36\xe8\xd6\xfb\x9b\xb8\x0f\x47\xc7\x2d\x65\x61\xe4\x2b\x4b\xad\x9f\x96\xaf\x5e\x0b\x37\x96\x4b\xb5\x6a\xb5\x53\x66\xd4\x90\xf5\xaa\xf0\x55\x3a\x31\xcd\x62\xac\xef\x7c\xeb\xd8\x15\x6f\x0e\xdb\xc5\x30\x9e\xca\xa7\xdc\xea\xf3\x93\x61\x97\x3f\x98\xfb\x88\x4d\xa1\xc4\xa6\x08\xfd\x41\x23\x83\x9f\x9d\x9f\xf2\x88\x50\x59\x8e\x33\x44\x53\xc6\x41\x2a\xcb\x58\x48\x30\x96\x63\x10\xa3\x34\xe5\x19\x4f\xa1\x90\xa1\xfc\x39\xf8\x8d\xd9\x70\x75\x8c\xf0\x62\x7a\x59\xbc\xb3\xf2\xe7\x1a\x35\x96\xba\xef\xa2\xc3\xa3\x8c\xca\x3a\x91\xc5\x32\x70\x35\x19\xef\x02\x47\xed\x2b\x66\x15\xef\x73\x67\x53\x7d\x57\x52\xca\x7c\xdc\x2a\x67\xec\xfb\xe3\x81\x25\xf8\xae\x1c\xe0\xb1\xa2\xae\xd9\x5d\xbb\xc3\xf5\xdb\xf9\x0b\x2b\xab\xab\xcb\x0a\x1b\x0b\x75\xe7\x1c\x0d\x99\x78\x64\x8c\x5a\xb4\x19\x79\x7f\xd9\x35\x2d\xcb\x46\x88\x90\x96\x26\xa1\xaa\xf5\xcf\xa7\xe8\xf0\x52\x9f\x24\x6b\x67\xa5\x4f\xdc\x59\x7d\x2a\x46\xac\x48\xa6\xb2\xe0\xf3\x59\xd7\x79\x73\xe8\x78\xc4\x58\x86\x3f\xa5\x30\x53\x2e\x4f\x21\xa0\x53\x84\xc8\x1f\x70\xf9\xac\xfc\x04\x27\x52\xc2\xb9\xa0\x88\xf1\x94\x02\x94\x26\x29\xa3\x31\xc7\x3c\xc1\x99\x88\x23\x91\x12\x42\x32\x4a\x3e\xef\x62\xec\x91\x3d\xd8\x5a\x3e\x9a\xf7\xc1\xcd\xaa\x20\xb1\x1f\xde\x4d\x52\x02\x10\x35\xcb\x73\x02\x9e\x50\x07\xca\xa5\xf4\xe7\x26\x2b\x6c\x4c\xc1\xea\x58\x9b\x72\xe6\x54\x6b\xbf\x77\x81\x50\x9e\xf9\x0b\x56\x51\xe7\x67\xd8\x6e\x42\xe6\x14\xd9\xfc\x78\xa0\x9b\xa1\xfb\x9e\x94\x5d\xe9\xb8\x94\xb5\xb7\x38\xac\xc2\xb8\x0a\xfc\x75\x72\xf5\x1b\xea\xaa\xf0\xb9\x5a\x5d\xdf\x2d\xc5\x33\xfd\xa6\xde\x0f\xba\xca\xf3\x4a\x8a\xf3\xd7\x51\x8e\xae\x75\xf8\x26\x13\xfd\x64\x79\xf6\x4c\x76\x7d\xbb\x0c\x76\x59\xe8\x5c\xb7\x23\x89\x4f\x4b\x72\xaf\xc7\x71\xf2\x73\x4b\x32\x84\xa5\x29\x81\x53\xf8\xf5\x72\xff\xc7\x60\x3e\x3b\x3f\xfd\x22\x91\x64\x08\x46\x14\x26\x82\x21\x39\x95\x62\x59\x22\x18\xc3\x18\x45\x58\x8e\xe5\x84\x49\x52\xfc\xcf\x15\xf7\x4f\x60\x16\x15\xaf\xf6\x4f\x56\x8c\xc0\x79\x09\x6d\x26\x6e\x0d\x2d\x77\x5a\xbc\xc7\x62\x13\x68\x3e\xed\x2a\xd5\x73\x6a\x70\x66\x61\x30\x58\x7c\x4f\xc3\x66\x7e\x6b\xc7\x4c\xdb\x78\xaa\xbe\x9e\xcd\xce\xcd\xe3\xa1\xd8\x7a\xba\x75\x86\xe7\xb8\xc8\xb3\xf0\x02\x03\xf3\x5d\x7f\x83\x69\xdf\x67\x6a\x9e\xfa\x44\x45\xa7\xd9\x61\xd5\x04\xb3\x17\x87\x56\x69\xbe\xd7\xf0\x18\x37\xa1\x66\x9c\x17\xe5\x24\x19\xee\xc3\x55\x65\xc9\xf8\x20\xc2\xe8\x10\xca\xe7\xcb\x5c\xd6\x7a\x0d\xa9\xb0\x5c\xf9\xae\x96\x45\x47\x29\x64\xa7\x75\x06\xe3\xb3\xad\xc1\x67\xf9\x43\x6a\x89\xa7\x32\x9e\x42\x59\xfe\xb3\x1d\x19\xfc\xac\xfc\x8c\x64\x29\xc3\x49\x94\xa2\x94\x4a\x08\x48\x58\x70\x29\x82\xb2\x48\x45\x92\xf0\x18\x31\x9e\xc9\x48\xce\xc4\xbf\xd8\xc5\xbc\xda\xfc\x58\xbf\x27\x42\xed\x34\xa3\x38\xc0\x18\xcf\xe7\xf3\x78\xc8\xe4\xed\xc6\x3b\x92\xb6\xbe\x35\xa3\x8d\x94\xeb\xe4\xe1\x66\xe4\x54\xc6\xa2\xaa\x86\xe7\x61\x59\xac\x12\x6b\xe7\x90\x3e\xba\x6f\x5f\xaa\x8e\x79\xba\xbc\x14\x68\x13\x14\xb2\xb3\x11\x9e\xf7\xf8\xde\xc5\xc4\xbd\xf6\xb4\x3a\xd5\x2c\x46\xed\x10\x5e\x13\xc9\x51\xd2\xa5\x6b\x1d\xae\xba\x5b\xc5\x9d\x4d\xfb\x24\xbc\x3f\x9e\x6f\xae\x6e\x4e\xce\x00\xad\xb4\x95\x95\x82\xf9\x7a\x16\x05\xc3\x58\xb5\x59\xa5\x77\xb7\x83\xc5\x7b\xfd\x60\xd7\x33\x55\x7a\xee\xa3\x95\x70\xa3\xfe\x87\x0a\x06\xf2\x29\xa4\x78\x0a\x01\x98\x42\xfc\x07\x6e\x09\x3f\x4b\x3f\xe6\x00\x41\x98\xa0\x48\xc0\x84\xca\xb1\xc4\x52\x94\xc1\x58\x8a\x39\x89\x30\x48\x18\x4b\xa8\xe0\x71\xca\x3e\x2f\x63\x56\x66\xd0\x2d\xdc\xfc\x65\xc0\x19\xb0\xec\x32\x7e\xef\xde\x47\xf7\x64\x30\xd4\x3f\x0a\xff\xd9\x21\x02\x82\xde\x5d\xa7\x78\xab\xbe\x57\xd7\xdd\x7d\x10\x95\x2c\xbc\xca\xef\xf9\xea\xa4\x6f\x5c\xcd\x08\x5f\x5b\xe5\x88\x9d\xc7\x65\xab\x6b\xd0\xee\xf5\x9e\xed\x82\xfb\x77\xc5\xd4\xe6\x7d\x9c\x97\x9b\x3e\x7e\x67\xbb\x79\x4b\x66\xc6\xb1\x9f\x55\x79\x75\xd5\xac\x88\xac\xd8\xd6\x98\x8d\xdd\xb0\xa9\x73\xec\x16\xf5\x1b\xa3\xa1\x9b\xb0\x4d\x65\xab\x87\x21\x98\xe4\x97\x8b\xa9\x55\xb7\xbd\x76\xa8\x2e\xab\xee\x0a\x8e\xeb\x48\xd4\x2e\x84\xfe\x6e\xf3\x43\x11\xf3\x4d\x06\x4d\x09\x9b\x22\xfc\x27\x53\xd9\x67\xe9\x87\x5f\x8a\x8f\xb3\x94\x63\xc8\x98\x44\x10\x65\x19\xc2\x10\x01\x29\x05\x59\x82\x05\xc9\x90\x84\x65\x96\x7c\xb6\xcb\x55\x68\xe2\xeb\x76\x2b\x85\xb5\xae\xe8\x60\xac\x35\xd1\xa6\xf8\xc1\x1f\xe6\xfd\xba\x62\xea\x3b\xcd\x34\x43\x7a\x10\x96\x95\xfb\xb9\x7e\xd1\x17\x44\x5f\x3d\x0a\x76\x0b\x9b\xce\x4e\xf7\xc7\x2c\x1b\x54\xc3\x3f\x80\x11\xeb\xed\x6a\xe9\xfa\x34\x9f\xec\xd3\x62\xfe\xfe\x5e\x2b\x5f\x4f\x4c\xbd\x1d\xbb\xfe\x3c\x2c\x5d\xdb\x25\xf9\x41\x15\x0d\xd9\x1e\x43\x43\x53\xe9\x49\xe8\xf1\xdc\x5f\x1e\x36\x17\xec\xeb\xcb\xd7\x6a\x73\xf1\xe6\x70\xee\x7b\x0d\xab\x13\x2f\x2b\xef\x7d\x5a\x9c\x45\x78\x3e\x4d\x30\x15\x1a\x35\x3c\x30\x19\x67\xe8\xf2\x96\xfa\xfc\x07\xc1\xe0\xe9\x57\x2f\x93\xff\x24\xfc\x3f\x6b\x7f\xc2\x30\x67\x42\x82\x22\x49\x12\x9c\x22\x09\x46\x54\x66\x51\x1a\x71\x8a\x98\x8c\xb2\x58\xe6\x84\x66\x18\xd1\xcf\xbd\xac\x8c\x56\xac\x6a\xe0\x70\xaf\xe6\x96\x9d\x59\xda\x06\xb0\xb5\x24\xe6\xef\x70\x5f\x3c\x75\xf9\xb2\x6e\x5a\x6d\xfe\xbe\x86\xa4\xaf\x96\x03\x9b\x4d\x1a\x41\xdf\xca\xb9\x59\x5a\xdd\x72\x24\xf3\x85\x64\x68\xb7\x7b\xb4\x20\xaf\x68\xfe\xba\xe8\x26\xca\xf0\x61\xdf\xec\xcc\x6f\x8f\x19\x1f\x89\xdc\x8e\xc1\x78\x0b\x79\x31\x3b\x0f\x38\x71\xc3\x1d\x70\x60\x2e\x6e\xd2\x50\x5e\xaf\xbb\xc5\xa0\x4c\x9c\xd9\x70\x3c\x27\x15\x2e\x57\xf5\x0e\x84\xf4\x5d\xe8\xda\x2c\x99\x34\xef\x67\xbf\xc8\x56\x6f\x2e\xb7\xc6\x11\xa7\xc5\xcb\x4a\x66\x91\x13\x9a\xde\xe3\x67\xc1\x40\x19\x4f\x91\xf4\x27\x64\x3e\x7b\xbf\x84\x78\xc4\x58\x2c\x45\x11\xc9\x32\x01\x13\x04\x63\x99\x72\xc4\xe4\x48\x12\xa9\x0c\x01\x95\x23\x89\x8a\x7f\x5e\xb4\xfd\xf3\x42\x66\xeb\xb1\x37\x0c\x97\xca\x2e\x9d\x78\xd9\xdd\xf3\x76\xf9\xa9\x2d\x86\x85\xac\x47\xb3\xb0\xc6\x5b\xbb\xda\x3e\x4f\x0a\x75\xb0\x22\xd7\x63\x7b\xc7\x65\xab\x06\x13\xd4\x0c\x79\xeb\xf2\xd7\xad\x2b\x96\x70\x77\xf2\xe3\xcb\x0e\x87\xf3\x41\x9c\x1a\x24\x8d\xa3\x66\x7f\x93\x89\xaf\xb0\xea\x14\x4b\x34\xf1\x86\x4f\x76\x72\x79\x85\x79\xf9\xd6\x33\xa2\xe4\xd6\x21\x57\x26\x39\xf5\x57\xad\x36\x53\xf4\xbc\x15\x79\x59\xf1\xda\x8b\x5e\x81\x1b\x9a\x33\x75\x58\x0d\x71\x91\x5b\x7d\xe4\xf0\x2a\x5b\x38\xaa\x9b\xc1\xc2\x51\xab\xfb\xce\x79\xd1\x9f\x24\x83\xa6\x10\x91\x29\xf9\x83\x8d\x3f\xfa\xac\xfe\x58\xb0\x34\x45\x11\xcd\x12\x21\xcb\x19\xc6\x28\x8a\x98\x9c\xca\x59\x16\xd3\x34\x81\x12\xc6\x82\xca\x69\x82\x3f\x87\xcc\x86\xa9\x8b\xd9\x7e\x6e\x58\xc6\x95\xab\x07\x86\x0e\xdb\x43\x71\x10\x02\x2c\x55\xcf\x07\xe9\xaa\xb3\xc2\x93\x57\x46\x26\x01\x97\x90\xa4\x34\x73\x66\xd5\x30\xdb\xda\xdd\x15\x1e\xf1\x4c\xab\xce\x85\xaf\x87\xf3\x9e\x9d\xfc\xc3\x98\xde\x33\x65\x84\xde\x5b\x7e\xe3\xef\xff\x2e\x9d\xa6\x58\x2f\x5b\x43\xa5\x80\x8e\xca\xa1\x85\x15\x36\xd7\xd5\x36\xdd\x4e\x4e\xe0\xe9\x55\x37\xec\x9d\xbb\x87\x9e\xab\xb3\x75\x62\x5c\x76\x13\x79\xd4\x75\x07\x6e\x17\xb5\x73\x3b\x24\x38\x7e\xed\xb4\xac\xf6\xd6\x11\x54\x9c\xe2\xbe\x98\x9c\x1d\x27\xf6\xe1\x79\xe9\xff\x90\x61\xfe\x3f\x30\x88\xe0\x29\x64\x7f\x60\x98\xe8\xb3\xfb\xa7\x48\xc2\x59\x8a\x64\x11\x53\x2c\x27\x42\x44\x69\x14\x71\x28\x12\x04\x44\xc4\x11\x46\x72\xcc\xa9\x04\xc1\xbf\xd8\xf9\x5f\x75\x32\x6f\xe2\x47\xe3\x1b\x0a\x39\xe4\x74\x2b\x46\x27\xdb\xc0\x61\x11\xbe\xb8\x08\xfc\xea\x34\xd9\x6d\x8b\x0e\xd9\xfb\xdd\x4b\x99\x9d\xc9\xe6\xd8\x0f\xc7\xd2\x47\x52\xf0\xce\x7d\x6a\x4a\xd1\x65\x5b\xde\xe2\x58\x7f\xef\xb7\x2c\x1f\xa2\xed\x4b\xc1\xee\xe1\xf1\x9d\x32\x4a\xb3\xa1\xdb\x57\xae\xbd\xd8\x6d\xdd\x63\xbf\x59\x63\x79\x46\x52\x69\x63\xf2\x7c\x8b\xf7\x7d\x16\x36\x6b\xdf\x3d\x47\x71\x3e\xcb\xea\x78\xeb\x92\xb4\x7c\xbe\x8e\xa9\x97\xbe\x64\x7e\x9a\xf3\x4a\xe2\x22\x7f\x1a\xe3\x6a\xd4\x2f\xc1\xe3\xa9\x58\x85\x31\xef\x27\xf8\x87\x06\xe6\x6f\x32\x74\x8a\xe1\x14\x7f\x16\x99\xff\xf3\x1f\xff\x37\x00\x00\xff\xff\x37\x7b\xda\xd2\x88\x21\x00\x00")

func shoalDelegatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_shoalDelegatesJson,
		"shoal/delegates.json",
	)
}

func shoalDelegatesJson() (*asset, error) {
	bytes, err := shoalDelegatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shoal/delegates.json", size: 8584, mode: os.FileMode(420), modTime: time.Unix(1589509583, 0)}
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
	"shoal/delegates.json": shoalDelegatesJson,
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
	"shoal": &bintree{nil, map[string]*bintree{
		"delegates.json": &bintree{shoalDelegatesJson, map[string]*bintree{}},
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

