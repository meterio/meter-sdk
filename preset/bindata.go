// Code generated by go-bindata.
// sources:
// shoal/delegates.json
// mainnet/delegates.json
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

var _shoalDelegatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x4b\x6f\x9b\x4c\x18\x85\xf7\xfe\x15\x23\xb6\xe4\x73\x66\x86\x81\x19\x23\x65\xc1\x70\xb3\x13\xdf\x2f\x21\xf6\xa7\x2a\xc2\x80\x0d\x06\x33\x04\x70\x80\x56\xf9\xef\x95\x5b\x67\x97\xaa\x9b\x6e\xdf\xa3\xf3\x4a\xe7\x79\xfe\xef\x01\xf0\xa3\x07\x00\x00\x52\xee\x9f\x23\x49\x07\x52\x15\x0b\x3f\xfb\x0f\x22\xe9\xee\xf7\xdd\x0f\xc3\x32\xaa\xaa\x6b\x04\x5b\xc3\x34\x54\xc6\x10\xd4\xb0\x83\xb0\xc6\x15\xcb\xa4\x16\x32\x4c\xcc\x4c\x62\x0d\x08\xc7\xda\x80\x52\x13\x22\xe7\xb3\x5c\x5c\xf6\xaf\x69\xd4\x5d\xcb\x7c\xce\x9d\xed\x50\xe9\xe4\x8d\xb2\x2c\x57\x53\x6f\xdf\xbc\xb5\xa2\x8b\xf0\x8b\x5d\x9b\x78\xf7\xb4\x8e\xd9\x25\xe8\x9a\xc5\x82\x5d\x9e\xe7\x9b\x38\x3b\x53\x51\xc4\x71\x08\x51\xcb\x8f\xa9\xcb\x1f\xb3\x15\x1c\x85\xe9\xf1\x59\xa5\x62\xbb\x6a\xe6\x51\xe7\x2e\xd3\xd1\x70\x32\x79\xd0\x75\x3d\x5b\x9e\x6d\xb3\x0c\xba\x28\x69\x1f\x4f\x6d\x98\x3e\x65\x7c\xe7\xd9\xe3\xae\x32\x4e\x7b\x85\xd7\x48\xcd\xef\xdd\xa0\xbd\x8c\xcf\x14\x6a\x72\x7c\x52\xa7\xf2\x24\xbc\x9f\xc7\xb8\xd9\xec\x8a\x82\xd6\x96\xe1\xaf\x0f\x47\x67\x1f\x2d\x82\xca\x1e\x7a\x6b\x77\x30\xcb\xa6\xb0\xb1\x1f\x3e\x57\xbc\x8b\x3a\xc9\x8f\xaf\x85\x68\xa2\x52\xd2\x01\x82\xf0\x16\xe4\x51\xdd\x88\x32\x7d\xbd\x32\x92\xf4\x1b\x49\x00\xa4\xa4\xb8\x2e\x46\x4a\x1f\x23\xd2\x57\xb5\x3e\xd2\xe8\xed\xd7\x95\x89\x28\x6b\x49\x07\x4c\xa3\xf0\xd7\xe9\xa3\x07\xc0\xc7\xdd\x1f\x45\xe0\x2f\x45\x30\x4e\x0c\xc3\x51\x20\x37\x4d\x1b\x2a\x54\x25\x64\xc0\x09\xc2\xcc\x56\x6d\xcd\x72\x34\xaa\x20\xcd\x32\x6d\x64\x7f\x25\x82\x7b\xf2\xc8\x9b\x1e\x9e\x65\xb5\x5d\xd4\x59\x32\x62\x76\x11\x2a\xcd\x84\x11\x1e\x5c\xf0\xcb\xf2\x6d\x96\x8d\x83\xa1\x21\x37\x5d\xe6\x7a\x3b\xb1\x49\x5e\xb6\x7b\x3e\x9a\x2d\x9b\xda\xe4\x0b\xcb\x1b\x9f\x17\x04\x05\xde\x86\x89\xb0\x79\xaf\xf9\x81\xee\xc7\xa9\x27\x48\x75\x15\xb1\x1e\xde\xbb\x6f\x75\x19\x08\xf9\x7b\x10\x26\x36\x1d\x06\x38\xd8\x65\x28\xdd\x56\x48\xcc\xcb\x83\xe1\x11\x9f\xb4\xee\xf9\x90\x4f\x54\xfe\x34\x39\x59\x1a\xaa\xad\x7c\x59\xb0\x8d\x97\x1c\x73\x42\xaa\x6d\x42\x1a\x7f\x90\x7a\x81\xe9\xc7\xef\xce\x6c\x25\x43\x77\xe5\x1c\x8d\x7f\x24\x02\x61\xd4\x47\x4c\xfb\x8b\x89\xde\xb7\x9f\x01\x00\x00\xff\xff\x9f\xd0\xda\x04\x15\x03\x00\x00")

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

	info := bindataFileInfo{name: "shoal/delegates.json", size: 789, mode: os.FileMode(436), modTime: time.Unix(1640761524, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainnetDelegatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xc9\x6e\xda\x50\x18\x85\xf7\x3c\xc5\x95\xb7\xa4\x70\x27\x5f\x0f\x52\x16\x9e\x81\x18\x62\x03\x4e\xc0\x55\x15\x19\xb0\xc1\x01\xcf\x80\x6d\xaa\xbc\x7b\xe5\x96\xec\x52\x75\xd3\xed\x7f\x74\x7e\xe9\x7c\xdf\xf7\x1e\x00\x3f\x7b\x00\x00\xc0\xa5\x41\x12\x72\x32\xe0\x92\x20\x4e\xbf\x41\xc4\x3d\xfc\x39\x07\xbb\x5d\x19\x56\x55\x97\xc0\x46\xd1\x14\x5e\x14\x11\x64\xd8\x44\x98\xa9\x44\xd7\x04\x1d\x29\x1a\x16\x35\xaa\x4b\x54\xc5\x4c\x12\x04\x0d\x22\xf3\xb3\x9c\x5f\x36\x6f\xc7\xb0\xed\xca\xea\x94\x96\x12\xec\x5b\x6d\x61\x8c\xd2\xa1\x18\x32\xd5\xdb\xf3\xe4\x79\xd2\xf8\x02\x59\xd4\x96\x6d\xbb\xde\x5a\xec\xfb\x75\x5b\x1d\xdd\xa9\x7f\x3e\x5e\x8b\xc6\x9d\xcd\xe0\x76\x9c\x9e\xa6\xd1\x81\x38\xd1\x7e\xb9\xb4\x5e\x8a\x7c\xef\x19\xeb\xab\x3d\xb6\xa2\xa1\x55\xd4\xfc\xf6\x51\x96\xe5\x6d\xb2\xe3\xb7\x4e\x92\xcd\x91\xa6\x9e\x62\x2b\xaa\x60\x21\x59\x26\x7e\x72\xd3\x66\xda\xd2\xb6\x5f\xa6\x7e\xeb\xb1\xd0\xe6\x2f\xab\xdb\xf4\x10\x84\xe9\x8a\xed\x94\xa7\x0b\x49\xb4\x7a\xb3\x6e\x90\x54\xb0\x23\xa9\x87\xd1\xfb\x9c\x04\xb1\xea\x8d\x5d\x6d\xe7\xf5\x5f\x5d\xe5\xf1\x73\xc5\x35\x3b\xc7\xe9\xfe\x2d\xcf\xea\xb0\xe4\x64\x80\x20\xbc\x07\x69\x78\xae\xb3\xf2\xf8\xd6\x31\xe2\xe4\x3b\x48\x00\xb8\x38\xef\x16\x23\x32\xc0\x3c\x1a\x20\x2c\x0d\x08\xbe\xff\xea\x98\x64\xe5\x99\x93\x81\xc8\x04\xf8\xfb\xf4\xd1\x03\xe0\xe3\xe1\x6f\x1e\xf0\x97\x1e\x44\x95\x2a\x8a\x49\xa0\xaa\x69\x06\x24\x02\x4f\xa9\xa4\x52\x84\x45\x83\x37\x98\x6e\x32\x81\x20\xa6\x6b\x06\x32\xbe\xf2\xf0\x5c\x18\xa3\x6b\x74\x98\x2f\x33\x88\x74\xd1\x71\xa7\xaf\xf9\x12\xae\xd4\xec\x4a\x23\x23\xaf\x86\x81\xde\xfa\xe3\xc5\x53\x36\x5a\x07\xbb\xc0\xcb\x27\x81\x90\x2b\xc6\x32\xaf\x54\x73\x4c\xd8\x68\x26\x62\x9c\x2e\xb4\xeb\xb0\xf1\x5f\x7c\x41\x3a\xdd\x5c\x73\x5f\x3a\xcc\x81\x9d\x87\x39\x65\x34\x31\x92\xe2\xbc\x5c\xc1\x67\xe7\x70\x78\xaf\x48\x3f\x54\x4e\xcd\x29\x9b\xd0\x0d\x3f\x37\x2a\x65\xb6\xc1\xd2\x42\x6c\xb6\x28\xba\xd8\xce\xe5\xf6\x92\x13\x7b\x81\x61\x76\xb3\x2f\x0b\x91\x0e\x05\x65\x72\x7c\x1f\x15\x99\x6b\x09\x8d\x74\xa3\x71\x11\x07\x5b\x7e\xff\x5f\x3c\x20\x3c\x40\x54\x18\x20\x84\xfe\x21\xa2\xf7\xa3\xf7\x2b\x00\x00\xff\xff\x65\xd3\x03\x37\x14\x03\x00\x00")

func mainnetDelegatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_mainnetDelegatesJson,
		"mainnet/delegates.json",
	)
}

func mainnetDelegatesJson() (*asset, error) {
	bytes, err := mainnetDelegatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mainnet/delegates.json", size: 788, mode: os.FileMode(436), modTime: time.Unix(1640761671, 0)}
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
	"mainnet/delegates.json": mainnetDelegatesJson,
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
	"mainnet": &bintree{nil, map[string]*bintree{
		"delegates.json": &bintree{mainnetDelegatesJson, map[string]*bintree{}},
	}},
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

