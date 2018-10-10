// Code generated by go-bindata.
// sources:
// schema.json
// DO NOT EDIT!

package schema

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xcd\x6e\x9c\x30\x10\xbe\xe7\x29\x90\xd3\x53\xd5\x84\x56\xea\x69\xcf\x95\xd2\x9e\x1a\x65\x7b\xab\xa2\xc8\x81\x59\xd6\x11\xd8\xc4\x1e\x52\xad\xaa\xbc\x7b\xc5\x6f\x60\xb1\xf9\x59\xbc\x29\x48\xec\x61\x0f\xcc\x30\x1e\x8f\xbf\x6f\x3c\x33\xfc\xbd\x70\x1c\xc7\x21\x1f\x94\xb7\x87\x88\x92\x8d\x43\xf6\x88\xf1\xc6\x75\x9f\x94\xe0\x57\xf9\xd3\x6b\x21\x03\xd7\x97\x74\x87\x57\x9f\xbf\xba\xf9\xb3\x4b\xf2\xa9\x78\x53\xc2\x2e\x7d\xed\xd2\xf5\x61\xc7\x38\x43\x26\xb8\x72\xb7\xb9\xb9\x42\xa7\x26\x21\x1b\x27\x5f\x32\x13\x7c\x63\x2a\xa6\xe8\xed\x1b\x4f\x33\x89\x84\xe7\x84\x49\xf0\xc9\xc6\xf9\xdd\x90\x64\x52\x4e\x23\x28\x8c\x37\xdf\x12\x09\x82\x22\x0d\xc1\x7d\x53\x8f\xc4\x52\xc4\x20\x91\x81\x6a\xad\xfa\x66\x5b\x27\xc9\xa4\x78\x88\x53\x29\x51\x28\x19\x0f\x48\x4b\xe9\xd5\xec\x95\xd1\x26\x43\x88\xcc\x62\x67\xd2\xf9\xe8\xad\x69\xcf\xec\x2e\x75\xb3\xbd\x23\xc3\xae\x32\x43\x11\xe3\x3f\x0a\xe7\xbf\x18\x54\x12\xce\x9e\x13\x28\xb5\x50\x26\x60\x50\x2c\x23\x4b\xa5\xa4\x07\x4d\x60\x2f\x3a\x1c\x22\xd4\xf7\xb3\x5d\xd0\xf0\xb6\x7e\xbc\x3b\x1a\xaa\xa3\xf5\xaa\x75\xc4\xe3\x13\x78\xf8\xb6\x50\xcd\x24\xb9\xa1\x08\x7f\xe8\xc1\x16\x2c\x5f\x40\x2a\x26\xb8\x4e\x84\x92\x05\x01\x48\xa5\x93\xf9\x05\x3b\xa6\x22\xda\x07\xe5\x49\x16\xa7\xf1\xb1\x0b\xec\x9a\x83\x73\x07\x77\x95\x69\x66\x8b\x6f\x43\x7c\x43\x7a\x78\x60\x11\x0d\x2c\xe7\xa4\xd2\xb4\xfd\x6c\x67\xdf\xa2\x02\xf9\xc2\xbc\x05\x80\x6c\x9b\x3b\xba\x28\x8c\x55\x09\x68\xee\xc1\xfd\x95\x3b\xba\xa8\xe0\x96\x89\xff\x74\x3a\xbc\xe7\xad\xf7\x9d\x72\x3f\x04\x39\xf6\xd6\xf3\xb5\xa9\x75\xf4\x25\x65\x2a\x05\x87\xc7\x4a\x9b\x3a\x10\x19\x0f\x3a\xd0\x1d\x53\x44\x90\xfc\xb6\xdb\xbd\x4a\xfd\xfa\x63\xa7\xdc\xe9\x38\x16\x33\xb4\x8e\xb7\xd8\x0e\x72\x7b\x8d\x0c\x84\xdd\xe6\x32\xc5\x47\x21\x42\xa0\xba\xda\xa3\xa5\xca\x38\x42\xca\xb0\x01\xaa\x3c\x89\x1e\x87\x6a\x86\xe1\x10\xbd\x02\x98\x03\x34\x4d\xe7\x5f\xff\xdd\x1b\xa5\xaf\xa3\xf2\x87\x89\x36\x7a\x6b\x67\xe6\xe7\xcf\x04\xe3\x04\x47\xd3\x93\x22\x9d\x46\x4d\x4f\xf8\x03\x6e\xf4\x12\x3d\xc3\xca\x8f\xd4\x29\xa3\xc9\xd1\x04\xea\x27\x4e\x2f\x61\x06\x10\x65\x00\x41\xfa\x89\xd1\x47\x88\x7e\x22\x74\x12\xa0\x0d\xfc\x77\x85\xe8\x1d\xa8\x58\x70\xd5\x86\x4b\x0f\x48\x41\x4a\x21\xa7\xa1\x34\x37\xd1\x0b\xd3\xf2\x98\x07\xc1\x94\xed\xec\x5e\x47\x42\xcf\xe0\x4a\x6e\xad\x8e\x32\xd4\x50\x45\x06\xf9\xcf\x18\xc9\xe6\x0c\x23\x01\xa2\x10\xe2\x89\x6d\x30\x55\x07\xee\xcd\x1c\x20\xb2\xe0\xcf\xfc\xcb\xf1\x8a\xe9\x27\xdd\xa7\x23\x6a\xe8\xfc\xe4\xe7\x1e\x8e\x2d\x42\x6c\xbd\x35\x99\xd3\x48\x6c\x5b\xc6\x70\x14\x6d\x23\xaa\xf6\x21\xe3\xc1\x83\xaa\x0f\x84\x1b\x2a\x41\x31\x6b\x9b\xc4\xed\xc0\x30\xb0\xab\x14\xce\x9d\x59\x6f\x74\xbb\x70\x0c\x90\x3e\x8e\xca\x32\xda\xc3\x72\xba\x61\x69\x28\x9a\x2d\x39\xad\x65\x3c\xd7\x5c\xf3\x1c\x23\xac\xb5\x0f\x5d\xfb\xd0\xfa\xcf\x72\x1f\x6a\x22\xd8\x42\x52\x4b\x7a\x7b\x8e\xad\x08\x75\xc3\xd6\xb1\x29\xc4\x76\xfd\xc6\x78\x67\x7d\xbf\x92\xbc\x52\x5d\x49\xde\x5a\xeb\x14\x92\x2b\xc3\xa5\xdc\x32\x3a\x0b\x9e\x97\x33\x7c\x9b\x25\x84\xe6\xf9\x3e\x1f\x64\xcf\xf5\xb3\x69\xe5\xde\xdc\x9b\x9a\xf2\x83\x80\xed\xbe\xe6\x9c\x9f\x5c\xd6\xda\x6d\x4d\xeb\xcb\x4f\xeb\x73\xaf\xdd\x2e\xf2\xff\xd7\x7f\x01\x00\x00\xff\xff\xc1\x1e\x88\x10\x3f\x26\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 9791, mode: os.FileMode(420), modTime: time.Unix(1539185337, 0)}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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
