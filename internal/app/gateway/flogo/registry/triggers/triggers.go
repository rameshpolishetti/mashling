// Code generated by go-bindata.
// sources:
// ext/flogo/trigger/eftl/trigger.json
// ext/flogo/trigger/gorillamuxtrigger/trigger.json
// ext/flogo/trigger/grpc/trigger.json
// ext/flogo/trigger/kafkasubrouter/trigger.json
// ext/flogo/trigger/mqtt/trigger.json
// ext/flogo/trigger/wssub/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/channel/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/cli/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/coap/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/kafkasub/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/lambda/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/mqtt/trigger.json
// vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/rest/trigger.json
// DO NOT EDIT!

package triggers

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

var _extFlogoTriggerEftlTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xdf\x6f\xd3\x30\x10\xc7\xdf\xf3\x57\x58\x7e\x1e\xcd\xc6\x13\x8a\x10\x82\xfd\x40\x9a\x04\x02\xd1\xbe\xa1\x3d\x38\xc9\xc5\x31\x24\xbe\xec\xee\xb2\x52\x4d\xfb\xdf\x91\x9d\x6c\x2d\x2c\x5b\x47\x4b\x5f\x5a\xe7\xec\xef\xe7\x7b\x77\xf2\xf9\x36\x51\x4a\x7b\xd3\x82\xce\x94\x16\x97\x17\xf8\x0a\x2a\x69\xf4\x51\x88\xcb\xaa\x8b\xf1\xaa\x41\x8b\x99\x90\xb3\x16\x68\xd8\x22\xa8\xc2\x8e\x75\x52\xf7\xf9\xac\xc0\x36\x5d\x5c\x9e\x9e\x7d\x99\x63\x25\x4b\x43\x90\xb6\x86\xeb\xc6\x79\x9b\xc2\x2f\x49\xa3\x3e\x1d\xf5\xe9\x9a\x7f\x03\xc4\x0e\x7d\x00\x1d\xcf\x8e\x67\x27\xa3\xab\x93\x26\xda\x7e\x83\x02\xdc\x0d\xa8\x8b\x8f\x8b\x4f\xea\x33\x30\x1b\x0b\xc3\x11\xd3\x4b\x8d\x14\xce\x7c\xf0\x25\xc1\x52\xcd\x3d\x96\x96\x0c\xb3\x7a\x6b\x78\x58\xbf\x8f\xd5\x84\xd4\xde\x0d\xa2\x12\xb8\x20\xd7\xc9\xe8\x18\xa9\x8b\xcd\x9a\x6a\x6c\xa1\x0b\x1e\x99\xd2\xb5\x48\xc7\x59\x9a\x6e\x2f\x50\x08\xe2\x97\x84\xd2\x9e\x29\x96\x41\xc4\x79\xcb\x3a\xfb\x9e\x28\xa5\xd4\x6d\xfc\xdd\xe8\x7e\x4f\xc3\xc9\x18\xbc\x6f\x3d\x0b\x39\x6f\x75\x0c\xdf\x1d\x4d\x0b\x5d\xb9\x9b\xae\xe7\xb1\xf2\x7f\x56\x76\x86\x79\x89\xb4\xa3\x6f\x61\x76\xd3\x09\x99\xe2\x99\x8c\x1f\xe2\x04\xd7\xbd\x23\x28\x75\xa6\x2a\xd3\x30\xbc\x00\x7a\xe1\xcb\x0e\x9d\x97\x83\xc0\x17\xf8\x13\xfc\x41\xc8\xe7\x90\xf7\xf6\x31\x39\x47\x6c\xc0\xf8\xbd\xd0\x73\xd3\xc2\xbc\x33\x13\x79\xff\x0f\xfa\xe5\xf9\xc9\xeb\x37\xa7\x6e\xa2\xdf\x2f\xa4\x27\x4a\x5d\xc5\xb1\xc2\x5e\xba\x5e\x58\x67\xea\x89\xb1\xea\x0c\x99\x96\x1f\x1b\x8d\xf1\x2d\xf7\x5c\xea\xaf\x7b\xe8\xaf\x7b\xa0\xd5\x3e\x80\x02\xbd\xc0\xd4\xb5\xc4\xfc\x07\x14\xb2\x7d\x5e\xfe\xb8\x61\xf7\x62\xe3\x57\xfa\xaf\x36\xd6\xc6\x97\x0d\x84\x27\x75\xc0\x6c\x3c\x57\x63\x63\xd7\x06\x1b\x16\x25\xf0\x3a\xb9\x27\xc6\xf9\x21\xc1\x49\xc2\x19\xfa\xd2\xc5\x37\x79\x3b\x26\xfe\x5f\x25\x61\x75\x97\xfc\x0e\x00\x00\xff\xff\x0f\x7b\x3c\xc4\xba\x06\x00\x00")

func extFlogoTriggerEftlTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerEftlTriggerJson,
		"ext/flogo/trigger/eftl/trigger.json",
	)
}

func extFlogoTriggerEftlTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerEftlTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/eftl/trigger.json", size: 1722, mode: os.FileMode(436), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _extFlogoTriggerGorillamuxtriggerTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\xcf\x6f\xea\x38\x10\xc7\xef\xfc\x15\x56\xce\xb4\xb4\x7b\x5a\xa1\xd5\x6a\xcb\x8f\x96\x6a\xdb\x57\x44\xd2\x53\xd5\xc3\x90\x0c\x89\xf5\x1c\x3b\x1d\x8f\x4b\x51\xc5\xff\xfe\x94\x38\xbc\x82\x1a\x28\x22\x7d\x17\x02\x9e\x7c\x3f\xfe\x7a\xc6\xf6\xf0\xde\x11\x22\xd0\x90\x63\xd0\x17\x01\xcb\x79\x6c\xce\x52\x60\x5c\xc2\xea\x8c\xd0\x72\xd0\x2d\xe3\xbc\x2a\xaa\xf8\x42\x99\xd4\xf4\x99\x64\x9a\x22\xf9\x10\xe1\xa2\x8c\xa4\x92\x33\x37\x3f\x8f\x4d\xde\x8b\x6e\x07\xc3\x87\xd0\x2c\x78\x09\x84\xbd\x1c\x6c\xa6\xa4\x4e\x7b\xf8\xc6\xbd\x4a\xdf\xab\xf5\xbd\xd4\x90\x54\x0a\x72\xf7\xb6\x43\x7c\x45\xb2\xd2\xe8\x92\x7a\x71\x7e\x71\x7e\x59\x5b\x90\xac\x2a\x0f\x33\x8c\x51\xbe\xa2\xb8\xaf\xc1\x62\x12\x45\x53\x71\x8f\xd6\x42\x8a\xfe\x5d\x70\x9c\x19\xaa\x5e\x86\x1c\x6d\x26\xa6\x46\x49\x9b\x21\xb3\x14\xff\x50\xe1\x7f\xfc\x57\x2d\xb6\x74\xfc\xaf\x57\x25\x68\x63\x92\x05\xd7\x73\x87\x32\x2f\x14\x8a\xd9\x38\x8c\x44\xe4\xfd\x89\x85\x21\xb1\x59\x90\xa8\xd3\xe4\xc5\xb6\x84\xeb\xd4\x06\x7d\xf1\xd4\x11\x42\x88\xf7\xea\x73\x2b\xb7\x85\x21\x9f\xcd\x6a\x74\x93\x51\xa9\x19\x37\x2b\xaf\x02\x84\x2f\x4e\x12\x26\x41\x5f\x30\x39\xac\x86\xd7\xdd\x66\x24\x13\xc4\xdb\xda\x0d\xd4\x32\x49\x9d\x36\x32\x17\xa0\xec\x31\xd0\xb1\x4e\x0a\x23\x75\x83\xe3\x6f\x80\x47\xe6\x27\xea\x3f\x42\x1e\xe1\xdc\xa5\x9f\xc9\x73\x63\x14\x82\x6e\x85\x0e\x21\xc7\xb0\x80\x06\xdf\xdf\x41\xbf\x1d\x5d\xfe\xf5\xf7\x40\x36\xe4\xbb\x1d\x1d\x35\xcc\x15\x46\x77\xe1\x7e\xf0\x41\xbd\x45\x7a\x45\x1a\x62\xd3\xde\xad\xeb\x75\x84\xfe\x7f\x5c\x9d\x26\xf7\xf6\x87\x4a\xa2\xe6\x2b\xc7\xd9\x69\x14\x26\x67\x39\x64\x43\x78\x9a\x7e\x0e\x56\xc6\xe5\xf4\xd7\x52\x9d\x88\x50\x09\x14\x13\x63\x4f\xcc\x62\xa9\x9e\x36\xde\x1f\xc7\xaa\x07\x60\x5b\x38\x1f\x48\x9d\x8c\x7e\xb4\xd3\x4f\xc1\xda\xa5\xa1\xe4\x74\xca\xa3\x45\xba\x96\x8a\x0f\x5c\x79\x5f\x32\x6e\xc8\xb8\xa2\x0d\x64\x69\x1f\x8b\x94\x20\xc1\xd9\xe6\x18\x1e\xc6\x74\x84\x78\xae\x3a\x84\x71\x5c\x38\x3e\xd4\x20\x80\x20\xb7\x9f\x71\xf5\xf8\x41\x57\x05\x70\x36\x6d\xa1\x7f\x71\x48\xab\x36\x80\x0c\x21\x69\x4a\xe9\x31\xda\xd8\x68\xc6\xa6\x56\x03\x7a\xf5\xd5\xc1\x86\x78\xa7\x65\x1c\xad\x5c\xda\xd8\x68\x8d\x71\xd5\xed\x0f\xc8\x7f\x97\x2f\x03\x9d\x28\x2c\xff\x54\x78\xd6\xe7\x8e\xff\x31\xcb\xd6\x3c\x39\x72\x66\x3e\xf6\xc8\xfe\x6e\xb7\x7d\xb1\x0b\xdf\xfb\xb7\x42\xa0\x94\x59\xfa\xc8\x53\x70\x33\x8e\x82\xae\x08\xa6\x0f\xa1\x7f\x3e\xfa\xc7\x55\x34\x9c\x94\x5f\x46\xe3\xbb\x71\x34\x0e\x9e\x6b\xf5\xba\xbb\xdf\x5d\xb9\x6d\x4e\xf3\x76\x04\x1c\x1c\x9b\xdb\x64\x86\x85\x5a\x35\xcd\xb1\xd3\x7e\x0e\x93\x9c\xc5\x0a\x33\xa9\x6b\xd0\x8e\x36\x34\x3a\x91\x3b\x75\xdf\x73\x76\x7d\xf9\x85\x28\x33\xb9\xee\xac\x3b\xbf\x02\x00\x00\xff\xff\x86\x68\x16\x0b\x2b\x0b\x00\x00")

func extFlogoTriggerGorillamuxtriggerTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerGorillamuxtriggerTriggerJson,
		"ext/flogo/trigger/gorillamuxtrigger/trigger.json",
	)
}

func extFlogoTriggerGorillamuxtriggerTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerGorillamuxtriggerTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/gorillamuxtrigger/trigger.json", size: 2859, mode: os.FileMode(509), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _extFlogoTriggerGrpcTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x6f\xd3\x40\x10\x85\xef\xfe\x15\x23\x9f\xdb\xb8\x5c\x23\x84\x10\xe1\x40\x55\x54\x50\x13\x4e\xa8\x87\xb1\x3d\x59\x2f\xb2\x77\xcc\xec\x6c\x8b\x55\xf5\xbf\xa3\x5d\x3b\xa4\xc5\x49\x8a\x80\x4b\x62\xef\xec\xfb\xe6\xed\x9b\xf5\x43\x06\x90\x3b\xec\x28\x5f\x42\xae\xb6\xac\xf8\xdc\xa0\xd2\x3d\x0e\xe7\x46\xfa\x2a\x3f\x8b\x75\x1d\xfa\x54\xdf\xb6\x6c\x78\xa9\x62\x8d\x21\x19\x4b\x42\xdb\x58\x31\x56\x9b\x50\x2e\x2a\xee\x8a\xcd\xe5\xbb\xd5\xa7\x35\x6f\xf5\x1e\x85\x8a\x0e\x7d\xd3\x5a\x67\x0a\xfa\xa1\x45\xd2\x17\x93\xbe\xd8\xf3\xef\x48\xbc\x65\x17\x41\x17\x8b\x8b\xc5\xab\xa9\xab\xd5\x36\xb5\x35\x37\x9f\x57\xb0\x19\x55\xf0\xc5\xa3\xa1\x71\x03\x06\x6d\x58\xe2\x8e\x6b\x14\xf2\x0d\x5c\x85\x0e\x05\x36\x0d\x2b\xc2\x6b\xa7\xf1\xff\x6d\x3a\x53\x34\xf6\x66\x14\xd5\xe4\x2b\xb1\xbd\x4e\xfd\xd6\xb6\xeb\x5b\x82\x67\x2d\xb6\x2c\xb0\xf3\x0d\x53\x1a\xa3\xd8\x93\xaa\x75\xc6\xe7\x4b\xf8\x9a\x01\x00\x3c\xa4\xdf\x27\x11\xf6\x2c\x9a\xf6\xa6\xd5\x5d\x70\xd6\x29\xed\x22\x4b\x05\xa1\xef\xc1\x0a\xd5\xf9\x12\x54\x02\xa5\xe5\xc7\xb3\x23\x48\x61\xe5\xeb\xf8\x32\xe3\x7a\x15\xeb\xcc\xdf\x61\x3d\xc9\x9d\xad\xe8\xff\x83\xc9\x61\xd9\xd2\xe6\xe3\x7a\x8e\x2d\x99\x5b\x42\x97\xbf\x68\x8c\x64\x45\x87\x82\x9c\x7c\xfd\x81\xfe\x8a\x86\x17\xe4\x19\xc0\x6d\x9a\x2a\x07\xed\x83\x9e\x1a\x2a\x0a\x76\x7e\x8e\x9b\xd6\x4f\xba\x89\xb7\xfc\x3d\x2a\xce\xd5\x5c\x7e\xa3\x4a\x4f\xab\x2b\x76\x4a\xee\x40\x10\xe8\x86\xdf\x8f\xd1\xa0\xab\x5b\x8a\xdf\xc3\x88\x99\xdf\xd6\x7d\x83\x27\x2d\x30\x28\x5f\xd6\x37\xd4\xb7\xfb\xc0\x8e\x8d\xec\x97\xd1\x83\xa4\xe0\x29\x61\x3e\x4c\x46\xfe\x8d\xd6\x91\x36\x5c\x3f\xbb\x9e\x47\x26\x39\x86\x00\x70\x9b\xc5\xa7\xc7\xec\x67\x00\x00\x00\xff\xff\x6d\x7f\xea\x08\xd4\x04\x00\x00")

func extFlogoTriggerGrpcTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerGrpcTriggerJson,
		"ext/flogo/trigger/grpc/trigger.json",
	)
}

func extFlogoTriggerGrpcTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerGrpcTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/grpc/trigger.json", size: 1236, mode: os.FileMode(436), modTime: time.Unix(1536643220, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _extFlogoTriggerKafkasubrouterTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x41\x8f\x94\x40\x10\x85\xef\xfc\x8a\x4a\x9f\x67\x86\xf5\x4a\x8c\x71\x77\x4f\xc6\x18\x13\x1d\x4f\x66\x0f\x05\x14\xd0\x19\xe8\xc6\xaa\xea\x1d\xcd\x66\xfe\xbb\xe9\x06\x94\xc4\x1d\x9d\x71\x2f\xd0\x50\xf5\xde\x57\xfd\x1a\x9e\x32\x00\xe3\x70\x20\x53\x80\x51\x5b\x56\x7e\x7b\xc0\xe6\x80\x12\xca\x2d\xfb\xa0\xc4\x66\x13\x5b\xf4\xc7\x98\x5a\x9a\xde\xb7\xbe\x50\xb6\x6d\xbb\x94\x98\x9a\x58\x69\xad\x76\xa1\xdc\x55\x7e\xc8\xf7\xef\xee\xee\x3f\x7e\xf6\x8d\x1e\x91\x29\x1f\x50\xba\xde\xba\x36\xa7\xef\x9a\x27\x7d\x3e\xeb\xf3\x05\xb5\x26\x3d\x12\x8b\xf5\x2e\x5a\xde\xec\x6e\x76\xaf\x66\xbe\xd5\x3e\x0d\xf0\x89\x2a\xb2\x8f\x04\xef\xa3\x14\x3e\x90\x08\xb6\x24\x53\x13\x06\xed\x3c\xc7\xae\x5b\x87\x4e\xe1\x76\x28\x3b\xb4\x07\x64\x78\x8d\x38\xad\xdf\xa6\x3d\xc6\x29\xdf\x4c\x9a\x9a\xa4\x62\x3b\xea\x8c\x9c\x6c\xf7\xd3\x7c\x70\xb4\xda\x41\xe5\x5d\x6d\x63\x7d\x5b\xa2\x50\x0d\x71\x58\xeb\xda\x49\x2e\xa4\xf1\x41\x4c\xf1\x35\x03\x00\x78\x4a\xd7\x55\xa6\x77\xec\x0f\xc4\x5f\xb8\x4f\xfd\xa9\xb4\x64\x29\xca\xd1\x27\xbd\x3e\x6d\x9e\x97\x07\x99\x73\xb9\x5a\x39\xa2\xc8\xd1\x73\xfd\x7f\x6a\xe5\x20\x2a\xea\x99\xfe\xa1\xcf\x00\x1e\x52\x12\x3e\xe8\x18\x54\x4c\x01\x67\x92\x18\xa6\xb3\xba\xd4\xaf\x43\x57\xf7\x14\x4f\x73\x72\x5a\x45\x3d\x13\x7e\x33\x56\x94\xbd\x1f\x6d\xf5\x8b\xf1\x27\x65\x55\x61\xfa\x16\x2c\x53\x6d\x8a\xb8\x5d\x32\x73\xe5\xb4\x39\xef\x3e\x22\x6b\xfa\x16\xe4\x5a\x04\x98\x06\x7b\xb9\x08\xd2\xb2\x0f\xe3\xd5\x5b\xb8\xd8\xde\x37\x8d\x90\x3e\xe7\x6f\x9d\xbe\xd4\xfc\x7e\xf9\x59\xfe\x32\xff\x62\x93\xee\x0f\x59\x5c\x9d\xb2\x9f\x01\x00\x00\xff\xff\x58\xd0\x2d\x1e\x89\x04\x00\x00")

func extFlogoTriggerKafkasubrouterTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerKafkasubrouterTriggerJson,
		"ext/flogo/trigger/kafkasubrouter/trigger.json",
	)
}

func extFlogoTriggerKafkasubrouterTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerKafkasubrouterTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/kafkasubrouter/trigger.json", size: 1161, mode: os.FileMode(509), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _extFlogoTriggerMqttTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc1\x4e\x1b\x31\x10\xbd\xe7\x2b\x2c\x9f\x69\x16\x7a\xaa\x56\x55\x55\x01\x3d\x70\x88\xda\x92\xdc\x2a\x0e\x5e\xef\xc4\xeb\xb2\xf6\x38\xe3\x59\x68\x84\xf8\xf7\xca\xce\x12\xd2\x76\x93\xd0\x04\x2e\x89\xf7\xd9\xef\xbd\x99\xb1\xc7\x7e\x18\x09\x21\xbd\x72\x20\x4b\x21\xd9\x56\x1a\xdf\xb9\x05\xb3\x3c\x49\x38\x2f\x43\xc6\xe7\x2d\x1a\x2c\x99\xac\x31\x40\xab\x29\x82\x79\x9a\x31\x96\x9b\xae\x1a\x6b\x74\xc5\xec\xea\xfc\xe2\xeb\x14\xe7\x7c\xaf\x08\x0a\xa7\x62\xd3\x5a\x6f\x0a\xf8\xc5\x45\xe6\x17\x3d\xbf\x78\xd6\xbf\x03\x8a\x16\x7d\x12\x3a\x1d\x9f\x8e\xcf\x7a\x57\xcb\x6d\xb6\xbd\x06\x0d\xf6\x0e\xc4\xe4\xfb\x6c\x26\x26\x10\xa3\x32\xb0\x5a\xa2\x3a\x6e\x90\xd2\x9a\x89\xd5\x8d\x82\x56\x5c\x83\xb1\x91\x81\xc4\x47\x47\xab\xe1\xe7\x9c\x4d\x0a\xed\xd3\x8a\x54\x43\xd4\x64\x03\xf7\x8e\x59\x75\xb6\x99\x53\x83\x0e\x42\xf2\x28\x85\x6c\x98\x43\x2c\x8b\x62\x7f\x82\x4c\x90\xbf\x18\x68\x67\xb2\x11\x98\xad\x37\x51\x96\x3f\x46\x42\x08\xf1\x90\x7f\x37\xaa\x5f\x11\xde\xf6\xa1\x64\xfc\xa9\xfa\x91\xc9\x7a\x23\x33\xfc\x78\x32\xcc\xb5\xf5\x61\xbc\x2e\x1e\xea\x18\x54\x8c\xf7\x48\x07\xfa\x46\x46\x82\xc3\xa8\x0b\x8c\xff\x12\x7d\xe7\x2a\xa0\xdd\x44\xdd\x82\xf2\x11\xe2\x00\xbd\x42\x4c\x93\xbb\xf9\x4c\x4a\xef\xa8\xd6\x1a\x27\x58\x74\x96\xa0\x96\xa5\x98\xab\x36\xc2\x0b\x44\xbf\xf8\x3a\xa0\xf5\xfc\x26\xe2\x33\xbc\x05\xff\x26\xca\x97\x50\x75\x66\x7b\x35\x8f\x91\x9e\x2a\x07\xd3\xa0\x06\xe2\x7e\x0d\xf5\xab\xcb\xb3\xf7\x1f\xce\xed\x40\xbd\x5f\xa8\x3e\x12\xe2\x26\x77\x35\x76\x1c\x3a\x8e\xb2\x14\x5b\xba\x3a\x28\x52\x6e\xe0\xc8\xf5\xf8\x9e\x1e\xe3\xe6\xdb\x11\xfc\x45\x07\xb4\x3c\x46\x40\xa3\x67\x18\x3a\x96\x58\xfd\x04\xcd\xbb\xc9\x6e\xe3\xc6\xfe\xef\x2e\x4f\xfb\xf4\xc7\xf1\x7c\x22\x2b\xbf\x94\x7f\xed\x41\xa3\x7c\xdd\x42\x7a\x0e\x56\x32\x1b\x57\x6d\xbf\x2b\xcf\x06\x9b\x16\x18\xac\x5e\x1b\x6c\x89\x6f\x1d\xe1\xa0\xc4\x05\xfa\xda\xe6\x07\x65\xbf\x4c\xfe\xbf\x19\xa5\xd1\xe3\xe8\x77\x00\x00\x00\xff\xff\x2d\xbc\x86\x5d\x77\x07\x00\x00")

func extFlogoTriggerMqttTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerMqttTriggerJson,
		"ext/flogo/trigger/mqtt/trigger.json",
	)
}

func extFlogoTriggerMqttTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerMqttTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/mqtt/trigger.json", size: 1911, mode: os.FileMode(436), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _extFlogoTriggerWssubTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4f\x4b\x03\x31\x10\xc5\xef\xfb\x29\x86\x9c\xb5\xa9\xd7\x45\x44\xf4\xe4\x49\x51\xc1\x43\xe9\x21\x9b\x4e\x93\x60\x36\x59\x32\x93\xae\xa5\xf4\xbb\x4b\xb2\x2b\xac\xff\x2e\x81\xcc\x7b\xf3\x7b\x33\x73\x6a\x00\x44\x50\x3d\x8a\x16\x04\xbb\x4e\xc7\xcb\x91\x28\x77\xe2\xa2\x08\x7c\x1c\xaa\xb0\xf7\xd1\xc4\x96\x93\x33\x06\xd3\x24\x25\xdc\x17\xc5\x38\xb6\xb9\x5b\xe9\xd8\xcb\xd7\x87\xbb\xfb\xc7\x97\xb8\xe7\x51\x25\x94\xbd\x22\xeb\x5d\x30\x12\x3f\x58\xd6\x7e\x39\xf7\xcb\x45\xc0\x01\x13\xb9\x18\x0a\x69\xbd\x5a\xaf\xae\xe6\x58\xc7\xbe\xe6\x3e\xa3\x46\x77\x40\x18\xb1\xa3\xa8\xdf\x91\xa1\x47\x22\x65\x10\xf0\x80\x81\x69\xb2\xab\xcc\x36\xa6\xea\x57\x3d\x92\x85\xa7\xe8\x1d\x59\x64\x76\x70\x9d\x86\xe9\x73\x5b\x97\x2b\x83\xde\x4c\x5d\x3b\x24\x9d\xdc\xc0\x73\xfc\xdb\xdf\x11\xf0\x6d\x69\x2a\xcc\x60\x48\xb4\x9b\x06\x00\xe0\x54\xdf\xc5\x05\x73\xf2\xd5\x59\x8b\x5f\xd7\x23\x4e\x2e\x18\x51\xcb\xe7\x06\x60\x5b\x59\x31\xf3\x90\x99\x44\x0b\xff\xb0\x74\x0c\x8c\x81\x7f\xf3\x54\x38\xfe\x84\x59\x15\x76\x1e\xcb\x0d\x26\xcc\x62\x52\xd8\x6c\x9b\x62\x3d\x37\x9f\x01\x00\x00\xff\xff\xb7\x09\x8b\x0f\xec\x01\x00\x00")

func extFlogoTriggerWssubTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_extFlogoTriggerWssubTriggerJson,
		"ext/flogo/trigger/wssub/trigger.json",
	)
}

func extFlogoTriggerWssubTriggerJson() (*asset, error) {
	bytes, err := extFlogoTriggerWssubTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ext/flogo/trigger/wssub/trigger.json", size: 492, mode: os.FileMode(509), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x31\x4f\xfb\x30\x10\xc5\xf7\x7c\x8a\x93\xe7\xff\x3f\x2e\xab\x47\x3a\x21\x21\x31\xb4\x1b\xea\xe0\x26\x17\xc7\x92\x63\x9b\xf3\x05\x14\x55\xf9\xee\x28\xc6\x05\xb7\x62\x60\x49\xa2\xf7\xf2\x7e\xef\xe9\x2e\x0d\x80\xf0\x7a\x42\xa1\x40\x0c\x2e\x98\xf0\xbf\x1b\xb5\xf7\xe8\xc4\xbf\xcd\xe2\x25\xfe\x58\x8a\xc9\x1a\x83\xf4\x65\x11\x0e\x9b\x63\x2c\x8f\xf3\xb9\xed\xc2\x24\x8f\x4f\x8f\xfb\x97\x43\x18\xf8\x43\x13\xca\x42\x0b\x9e\xc9\x9e\x65\x89\xca\x1b\xfa\x3b\x52\xb2\xc1\x6f\x98\x5d\xbb\x6b\x1f\x4a\xa7\x65\x97\x4b\x9f\x6d\x62\xf4\xc0\x01\xf6\x75\xaa\xc7\xd4\x91\x8d\x5c\x92\x07\x3b\x45\x87\xd7\x5f\xe0\x58\x6f\x1c\xc3\x84\x51\x9b\x4c\x1b\x99\x63\x52\x52\xfe\x79\x30\xa2\x9c\x74\x62\xa4\xdf\xc7\x27\x64\xb6\xde\x24\xa1\xe0\xb5\x01\x38\x65\x31\xcc\x1c\x67\x2e\x12\xc0\x25\x3f\xab\x13\xf7\x9a\x75\x8e\x67\xf5\x7a\x5d\xed\x17\x91\xb5\xf5\x1b\x44\x18\xdd\x72\x83\x1e\xb5\xef\x1d\x92\x50\x85\x7a\x3f\xa0\xee\xab\x1a\xeb\xcd\x77\xb5\x89\xc9\x7a\x53\x3b\x84\x6f\xb3\x25\xec\x05\x28\x60\x9a\xb1\x38\x6b\x7e\x9f\x9a\xed\x6b\x6d\x3e\x03\x00\x00\xff\xff\x1e\x6e\xa2\xd3\x34\x02\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/channel/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/channel/trigger.json", size: 564, mode: os.FileMode(436), modTime: time.Unix(1536041955, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x3f\x6b\xf3\x30\x10\x87\x77\x7f\x0a\xa1\x39\x6f\x94\x77\xf5\xd8\x4c\x81\x42\x87\x64\x2b\x19\x2e\xf6\x59\x16\xe8\x1f\xa7\x73\x4b\x08\xfe\xee\xc5\x8a\xea\x18\x9c\x40\x17\x49\x3c\x77\x7a\xee\xc7\xdd\x2a\x21\xa4\x07\x87\xb2\x16\xb2\xb3\x41\x87\x7f\x8d\x35\x72\x33\x61\xbe\xc6\x07\xae\x99\x8c\xd6\x48\xf7\x52\xea\x8d\x9b\x4a\x0e\x8c\xbf\x13\xc2\x6e\x02\xda\x70\x3f\x5c\xb6\x4d\x70\xea\x74\x78\xdb\x7f\x1c\x43\xc7\xdf\x40\xa8\x8a\x3b\x78\x26\x73\x51\x45\xa6\xe6\x59\x5f\x48\xc9\x04\x3f\x29\x76\xdb\xdd\xf6\x7f\x49\x60\xd8\xe6\x08\xfb\xf7\x83\x38\x2d\x03\xb4\x98\x1a\x32\x91\xcb\x9f\xa3\x71\xd1\xa2\x58\xb5\xf5\xc1\x61\x04\x9d\x1d\x3d\x73\x4c\xb5\x52\x7f\x8e\x88\xa8\x1c\x24\x46\x5a\xc7\x0d\x03\xc7\x81\x65\x2d\x3e\x2b\x21\x84\xb8\xe5\x73\xb1\x49\x20\x9d\x72\x67\xa6\xbf\x8b\x04\x22\xb8\xca\x4c\xc7\x4a\x88\x73\x59\x5c\xb4\xd7\xd7\xa6\x16\x18\x9e\x98\xfc\xca\xd3\x83\x6f\x2d\x92\xac\x8b\x43\x26\x64\x36\x5e\xa7\xd9\xfd\xb0\x2f\xfc\x4d\x70\x0e\x7c\x3b\x8f\x58\x0c\x49\x4c\xc6\x6b\x59\x0a\xe3\xe6\xb5\xa4\xc5\x0e\x06\xcb\xcf\x24\x97\x10\x2c\x82\x9f\x2d\xf9\x3e\x57\xd3\x6b\xac\x7e\x02\x00\x00\xff\xff\x8e\x26\xfb\x7b\x7e\x02\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/cli/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/cli/trigger.json", size: 638, mode: os.FileMode(436), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8f\x9b\x30\x10\x85\xef\xfc\x8a\x91\xcf\x34\xa4\x57\x6e\x69\x8a\xda\x4a\xad\x12\x35\xec\x29\xca\xc1\x81\x09\x58\x02\xec\x8c\x87\x44\x51\xc4\x7f\x5f\x61\xc3\x82\x16\xe5\xb2\x17\x06\xcd\x9b\x79\xdf\xf3\x3c\x03\x00\xd1\xc8\x1a\x45\x0c\xe2\x52\xe9\x42\x7f\xcb\xb4\x34\x22\xec\xfb\xfc\x30\x53\x3f\x66\x52\x45\x81\xe4\x25\xc2\x4b\xaf\x14\x8a\xcb\xf6\xbc\xca\x74\x1d\xa5\x7f\x7e\x6c\x77\x07\x7d\xe1\xbb\x24\x8c\x46\xab\x86\x49\x9d\xa3\x61\x35\x9a\xac\x6f\x48\x56\xe9\xa6\xf7\x58\xaf\xd6\xab\xef\x03\x50\x71\xe5\x88\xff\x31\x43\x75\x43\xd8\xea\xcd\x1e\xfe\xa1\xb5\xb2\x40\x3f\x92\xa3\xcd\x48\x19\x1e\x96\x0f\xaa\x36\xd5\x30\x97\xce\x03\x5a\x64\x56\x4d\x61\x45\x0c\xc7\x00\x00\xe0\xe9\xbe\xb3\xd7\x1a\x4d\xec\x66\x5d\x77\x7c\xab\x6a\x18\x47\x13\x27\x10\x5e\x5b\x45\x98\x8b\x18\x98\x5a\x74\xed\x2e\x00\x38\x39\x8c\x6e\xd9\xb4\xfc\x1a\x72\x6d\x91\x1e\x7b\x49\xb2\xb6\x4b\x96\xf1\x7d\x6f\x19\xbe\x48\x29\x1f\x95\x96\xf9\x72\xd9\x32\xa9\xa6\x10\x9f\xf2\x94\xb2\xc9\x2b\x24\x11\x0f\x4e\xcb\x3b\x4c\x8c\x19\xa5\x46\x2e\xf5\x04\x59\x62\x66\xca\xc7\x41\xc0\x5f\x64\x26\xc9\xaa\xd2\x77\xaf\x1c\xc5\xaf\x24\x15\x21\x88\xfd\xee\xe0\xeb\x9b\x2f\x9b\x74\xfb\xbb\xff\xf9\x99\xfc\x4d\xd2\x44\x9c\x86\xed\x2e\x7c\x9d\xce\x48\x2e\xbf\x96\x6d\x34\x77\xb5\x47\x75\x41\x17\xbc\x07\x00\x00\xff\xff\x90\x30\x39\x84\xf4\x02\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/coap/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/coap/trigger.json", size: 756, mode: os.FileMode(509), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x41\x6f\xd4\x30\x10\x85\xef\xf9\x15\x96\xcf\xcb\x6e\xb9\x46\x08\xa1\x72\x42\x08\x90\xe8\x22\x0e\xa8\x07\x27\x99\x38\xa3\x75\x3c\x66\x66\xdc\x15\xaa\xf6\xbf\x23\x27\x69\x1b\x44\x0b\x0b\xbd\x24\x56\xde\xbc\xef\xc5\xcf\xbe\xad\x8c\xb1\xd1\x8d\x60\x6b\x63\xfb\x40\x9e\x5e\x1c\x5c\x7f\x70\x92\x1b\xbb\x29\x9a\xfe\x48\x0f\x5a\xad\x8c\xde\x03\xcf\x12\x43\x5f\x14\x8f\x3a\xe4\x66\xdb\xd2\xb8\xdb\xbf\xbb\x7c\xfb\xe9\x8a\x7a\x3d\x3a\x86\xdd\x8c\x6b\x29\x2a\x63\xb3\x5b\xac\xbb\x5f\xf1\x37\xc0\x82\x14\x0b\xe7\x62\x7b\xb1\x7d\xb9\x84\xa2\x86\x29\xf5\x33\xb4\x80\x37\x60\xde\x17\x93\xf9\x00\x22\xce\x83\xcc\x43\x2e\xeb\x40\x5c\xa6\xbe\x42\xec\x20\x04\xf3\x11\xdb\x81\x82\x98\x57\xc7\x38\xaf\xde\x28\x36\x2d\x95\x5f\x7b\x3d\x7b\x3a\x90\x96\x31\xe9\x12\x79\x85\x63\x0a\x77\xf4\xfd\x7a\x6f\x02\xaa\x18\xbd\xd8\xfa\x5b\x65\x8c\x31\xb7\xd3\x73\xd5\xd5\x25\xd3\x01\xf8\x0b\x87\x69\x7e\x92\xee\xaa\x12\x65\x8c\xde\x4e\x9f\x4f\x9b\xc7\xed\x59\x96\xa4\x7f\x76\x26\x27\x72\x24\xee\xfe\xcf\xad\x9c\x45\x45\x89\xe1\x2f\xfe\xca\x98\xeb\xa9\x09\xca\x9a\xb2\xda\xda\x3c\x51\xc4\x38\x1f\xca\xb9\xb8\xc1\xc5\x2e\x40\x39\xb6\x99\xb4\x6a\x7a\x49\x78\xc8\x58\xa5\xec\x29\x61\x7b\x9f\xf1\x7b\xca\x4a\x61\xf8\x9e\x91\xa1\xb3\x75\xd9\x2d\xd8\x45\x39\x6d\x9e\xa6\x27\xc7\x8a\xe5\x52\xc8\x1f\x22\xce\xe0\x78\xa6\x9c\x9e\x87\xa0\xbe\x17\xd0\xc7\x18\x18\xf5\x1e\x30\xbd\xaf\xab\xb2\x3a\x55\x3f\x03\x00\x00\xff\xff\x78\x90\x4d\x6c\xc4\x03\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/kafkasub/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/kafkasub/trigger.json", size: 964, mode: os.FileMode(509), modTime: time.Unix(1535605435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4f\x6b\xe3\x30\x10\xc5\xef\xfe\x14\x83\xcf\x59\xcb\xbb\x47\xdf\x76\xf7\x54\x28\xf4\x90\x40\x0f\xa5\x87\xb1\x3d\x96\x55\x64\x49\x48\xa3\xa4\x21\xe4\xbb\x17\xc9\x4e\x93\x92\xf4\xcf\xc5\x98\xf7\xf4\xde\xfc\x86\x39\x14\x00\xa5\xc1\x89\xca\x06\xca\x41\x5b\x69\x7f\x69\x9c\xda\x1e\xcb\x55\x72\x78\xef\xce\x4e\xc3\x5e\x49\x49\x7e\xb6\xc2\xa8\xa6\x64\x39\x1d\xa5\x32\xb3\xe6\x69\x48\x92\x54\x3c\xc6\xb6\xea\xec\x24\x36\x77\xff\xfe\x3f\xac\xed\xc0\x3b\xf4\x24\xe6\x01\x9d\x35\xec\x55\x2b\x96\x3a\x71\x39\x70\x4b\x3e\x28\x6b\x52\x4b\x5d\xd5\xd5\xef\x05\x43\xb1\xce\x1c\xf7\xf9\x29\x6c\x2e\x41\x7a\x0a\x9d\x57\x8e\x97\xd8\xdf\xc7\x35\x7c\x7c\x06\x31\x50\x0f\x6c\x21\x30\x7a\x06\x84\x41\xdb\x1d\x60\x48\x7f\xd1\x74\x29\x58\xcd\x55\xa3\x9d\xc8\xa1\xcc\xa3\x46\x66\x17\x1a\x21\x7e\xbc\x0c\x91\x98\x30\x30\xf9\x9b\x8b\x05\x62\x56\x46\x86\xb2\x81\xa7\x02\xe0\x39\x8b\x36\xb2\x8b\xbc\x48\x00\x87\xfc\xbd\x38\x48\x2a\xa7\x57\xce\x0d\xd9\x38\xdd\xc3\xb6\x2f\xd4\x71\x99\xe5\xe3\xea\x76\x98\xb6\xdf\x05\xdf\x39\x3c\x39\xbd\xff\x1c\xa3\x47\xc6\xeb\x2a\x34\xfb\xaf\x01\x02\x23\xc7\x70\x1d\x54\x86\xe9\x74\xbd\x6c\x6c\x51\xc7\xe4\xfc\xa9\xeb\x33\x58\x71\x2c\xde\x02\x00\x00\xff\xff\xa7\xcd\xe6\xf0\x9d\x02\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/lambda/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/lambda/trigger.json", size: 669, mode: os.FileMode(436), modTime: time.Unix(1536041955, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x6f\x9c\x30\x10\x85\xef\xfc\x8a\x11\xe7\x74\x49\xaf\xa8\xaa\xaa\xf6\xd4\xc3\xaa\x6a\xb2\xb7\x2a\x07\x63\x66\x61\x5a\x60\xbc\x33\xe3\x44\x51\x94\xff\x5e\x19\x68\x16\x35\x0b\xaa\xaa\x5c\x90\x35\xe3\xf7\xbd\x37\xb6\x79\xca\x00\xf2\xc1\xf5\x98\x97\x90\x1f\x3b\x6e\xf8\x5d\x7f\x32\xcb\xaf\x52\xdd\x1e\xc3\xb9\x5e\x9a\x50\xd3\xa0\x4c\x2d\xc1\x63\xea\x34\x64\x6d\xac\x76\x9e\xfb\xe2\xf0\xf5\xf3\x97\x6f\xb7\x7c\xb4\x07\x27\x58\x4c\x28\xcf\x83\x09\x55\xc5\x2c\x2d\xce\xe8\x7b\x14\x25\x1e\x12\xe3\x7a\x77\xbd\x7b\x3f\x1b\x92\x75\xa3\xe3\x0d\x7a\xa4\x7b\x84\xfd\xf7\xc3\x01\xf6\xa8\xea\x1a\x9c\xb6\xb8\x68\x2d\x4b\xda\xb3\x27\xdf\x3a\xec\xe0\x06\x1b\x52\x43\x81\x0f\xbd\x4c\xcb\x4f\x46\x95\xe7\x94\xea\xe3\x24\xaa\x51\xbd\x50\xb0\xd9\xf1\x96\xfa\xd0\xcd\xf0\xc3\x72\xaa\x96\x7b\x0c\xc9\xaa\x84\xbc\x35\x0b\x5a\x16\xc5\x3f\x8f\x88\x58\xf4\x2e\x05\xb9\x30\xae\xa2\x19\x0d\x8d\xe6\xe5\x8f\x0c\x00\xe0\x69\xfc\x2e\x8e\xbe\x12\xfe\x35\xa7\x18\xeb\x7f\x8e\x5e\x4d\x68\x68\xce\x75\xc1\x53\x24\xc1\x3a\x2f\xc1\x24\xe2\x58\x7e\xbe\xba\xcc\xa4\xfa\x6d\x79\x51\x37\x12\x6e\x2a\x83\x53\x7d\x60\x59\xcf\xb3\xa9\x56\x63\xc1\xff\x93\x9e\x58\x5f\x0b\x6b\x8e\x55\x87\xdb\x42\xdf\xa1\x1b\x14\xf5\x82\xbc\x62\x4e\xcd\x59\x9f\x01\xdc\x8d\x57\xcc\xd1\x42\xb4\xbc\x84\x95\x1b\xee\x17\xaf\x78\x7d\x8e\x17\x9c\x60\xe8\x1e\xd7\x69\xb5\x33\xf7\x1a\xc5\xd5\x4f\xf4\xf6\x37\xaa\x75\x43\xdd\x61\xfa\x67\x26\xcc\xe2\x35\xce\xf8\xb3\xc1\xc2\xc2\x38\x90\x7f\xf1\x58\x09\x3c\xf9\x00\xdc\x65\x69\xf5\x9c\xfd\x0e\x00\x00\xff\xff\x61\x25\x72\x0f\x50\x04\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/mqtt/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/mqtt/trigger.json", size: 1104, mode: os.FileMode(509), modTime: time.Unix(1536041955, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x52\x4d\x6f\x9c\x30\x10\xbd\xf3\x2b\x46\x3e\xd3\x25\xbd\x72\x6b\x53\xd4\x54\x6a\x95\x55\x70\x4f\x51\x0e\x0e\xcc\x82\x25\xb0\x9d\xf1\x90\x08\x45\xfc\xf7\xca\x98\xcd\xa2\x26\x44\x69\xf7\x82\xd1\x3c\xbf\x0f\x3d\xcf\x73\x02\x20\x8c\xea\x51\xe4\x20\x0e\x9d\x6d\xec\x27\x42\xcf\x22\x0d\x73\x1e\xdd\x69\x9e\x33\xe9\xa6\x41\x8a\x10\xe1\x21\x20\x8d\xe6\x76\xb8\xdf\x55\xb6\xcf\xe4\x8f\xaf\x97\xd7\xa5\x3d\xf0\x93\x22\xcc\xa2\x54\x65\x0d\x93\xbe\xcf\x16\x6a\x76\x92\x7e\x44\xf2\xda\x9a\xa0\x71\xb1\xbb\xd8\x7d\x5e\x0c\x35\x77\xb3\xe3\x0d\x56\xa8\x1f\x11\xae\xa4\xdc\xc3\x2f\xf4\x5e\x35\x18\xaf\xd4\xe8\x2b\xd2\x8e\x17\x72\xa9\x7b\xd7\x21\xdc\x14\xa5\x04\xb9\x0e\xd8\xda\x1e\x5d\x60\xe5\x20\x5a\x66\xe7\xf3\x2c\xfb\x70\x5a\xc4\xac\x57\x9e\x91\xde\x48\xee\x91\x59\x9b\xc6\x8b\x1c\x6e\x13\x00\x80\xe7\xf9\xbb\xaa\xd1\x59\x8a\x77\xe7\xe9\xb1\x44\x6d\x18\x8f\xe9\x66\x80\xf0\x61\xd0\x84\xb5\xc8\x81\x69\xc0\x79\x3c\x25\x00\x77\xb3\x8d\x1d\xd8\x0d\xfc\x8e\x89\x22\xd5\xfb\xd7\x36\xcb\x3c\xaa\xa5\x5b\x5c\x6e\xf7\x67\xf0\x1f\x06\xa4\xf1\x1c\x81\x16\x55\xbd\xae\xe2\x5f\xb8\xe1\x95\xd0\xbc\x51\xb0\x32\xa3\xf8\xab\x44\x42\xd7\x8d\xdb\x1d\x56\xb6\xc6\xed\x87\x7a\x37\x45\xad\x58\x7d\x28\x42\xab\x4c\xdd\x21\x89\x7c\xd1\x78\xbd\x3f\x27\xf5\x95\x7e\x8f\xdc\xda\xfa\xc5\x61\xe5\xe1\x99\xb4\x69\xd6\xc8\xcb\x22\x41\xdc\xa4\x15\xa4\xba\xce\x3e\x45\xe4\x56\x7c\x2f\xa4\x48\x41\xec\xaf\xcb\x78\xfe\x8e\xc7\x17\x79\x79\x15\x7e\xbe\x15\x3f\x0b\x59\x88\xbb\x85\x3d\xa5\xdb\xe9\xc2\x02\xfd\x5f\xb6\xa3\xf8\x7c\x06\xab\x29\x99\x92\x3f\x01\x00\x00\xff\xff\x37\x69\xda\x15\x85\x04\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/rest/trigger.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/rest/trigger.json", size: 1157, mode: os.FileMode(436), modTime: time.Unix(1535605435, 0)}
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
	"ext/flogo/trigger/eftl/trigger.json":                                         extFlogoTriggerEftlTriggerJson,
	"ext/flogo/trigger/gorillamuxtrigger/trigger.json":                            extFlogoTriggerGorillamuxtriggerTriggerJson,
	"ext/flogo/trigger/grpc/trigger.json":                                         extFlogoTriggerGrpcTriggerJson,
	"ext/flogo/trigger/kafkasubrouter/trigger.json":                               extFlogoTriggerKafkasubrouterTriggerJson,
	"ext/flogo/trigger/mqtt/trigger.json":                                         extFlogoTriggerMqttTriggerJson,
	"ext/flogo/trigger/wssub/trigger.json":                                        extFlogoTriggerWssubTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/channel/trigger.json":  vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/cli/trigger.json":      vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/coap/trigger.json":     vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/kafkasub/trigger.json": vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/lambda/trigger.json":   vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/mqtt/trigger.json":     vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJson,
	"vendor/github.com/TIBCOSoftware/flogo-contrib/trigger/rest/trigger.json":     vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJson,
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
	"ext": &bintree{nil, map[string]*bintree{
		"flogo": &bintree{nil, map[string]*bintree{
			"trigger": &bintree{nil, map[string]*bintree{
				"eftl": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerEftlTriggerJson, map[string]*bintree{}},
				}},
				"gorillamuxtrigger": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerGorillamuxtriggerTriggerJson, map[string]*bintree{}},
				}},
				"grpc": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerGrpcTriggerJson, map[string]*bintree{}},
				}},
				"kafkasubrouter": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerKafkasubrouterTriggerJson, map[string]*bintree{}},
				}},
				"mqtt": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerMqttTriggerJson, map[string]*bintree{}},
				}},
				"wssub": &bintree{nil, map[string]*bintree{
					"trigger.json": &bintree{extFlogoTriggerWssubTriggerJson, map[string]*bintree{}},
				}},
			}},
		}},
	}},
	"vendor": &bintree{nil, map[string]*bintree{
		"github.com": &bintree{nil, map[string]*bintree{
			"TIBCOSoftware": &bintree{nil, map[string]*bintree{
				"flogo-contrib": &bintree{nil, map[string]*bintree{
					"trigger": &bintree{nil, map[string]*bintree{
						"channel": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerChannelTriggerJson, map[string]*bintree{}},
						}},
						"cli": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerCliTriggerJson, map[string]*bintree{}},
						}},
						"coap": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerCoapTriggerJson, map[string]*bintree{}},
						}},
						"kafkasub": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerKafkasubTriggerJson, map[string]*bintree{}},
						}},
						"lambda": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerLambdaTriggerJson, map[string]*bintree{}},
						}},
						"mqtt": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerMqttTriggerJson, map[string]*bintree{}},
						}},
						"rest": &bintree{nil, map[string]*bintree{
							"trigger.json": &bintree{vendorGithubComTibcosoftwareFlogoContribTriggerRestTriggerJson, map[string]*bintree{}},
						}},
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
