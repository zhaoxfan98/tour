// Code generated by go-bindata. DO NOT EDIT.
// sources:
// configs/config.yaml
package configs

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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4f\x6f\xd3\x4c\x10\xc6\xef\x95\xfa\x1d\x46\x7a\xcf\x75\x6c\xa7\x4d\xfc\xee\xa9\xff\x52\x51\xd4\x40\x84\x5d\xf5\x88\x36\xf1\xc4\x31\x5a\x7b\x37\xbb\xe3\xd4\xed\x0d\x09\xf5\x86\x38\x20\x0e\x70\x2a\x17\x2e\x48\xc0\x01\x09\x09\x89\x6f\x43\x5a\x3e\x06\xf2\xda\x34\x26\xc0\xcd\xfe\xcd\xce\xec\xf3\x3c\x3b\x21\xea\x05\x6a\xb6\xb9\x01\xf0\xa8\xc8\x87\x32\x46\x06\x31\x8e\x8b\xa4\x22\xf7\x88\xd4\x48\x6a\x62\x10\xb8\x81\x6b\xcf\x20\x8f\xa3\x34\x43\x59\x10\x83\x9e\x45\x67\x3a\x25\xfc\x8d\xed\x29\x65\x07\x1e\xe2\x94\x17\x82\x0e\x64\x4e\x58\xd2\x5a\x5b\x53\x1c\xf1\x04\xc3\xf4\x12\x19\x78\x16\x0f\x79\xd9\x46\x96\x9d\xc8\x24\xe4\x0b\x1c\x71\x9a\x31\x30\x24\x35\x4f\xb0\x23\x64\x62\x9a\xe2\x51\x2a\xf0\x01\xcf\x90\x01\x57\xaa\xc5\x06\x25\x31\x70\x84\xb4\x5e\x4e\x95\x90\x3c\xfe\x73\x4e\x61\xb9\x69\x1d\xb1\x89\x9c\x6a\xc1\x60\x46\xa4\x58\xa7\xe3\xf9\x7d\xc7\x75\x5c\xc7\x63\x55\x0c\x1d\x43\x9c\xd2\xc9\xaa\xe1\x38\xe3\x09\x0e\x79\x59\x6b\xde\x01\xf8\x0f\x86\xfb\x6b\xe5\x3d\x21\xe4\xf9\xa0\x24\x63\x83\x01\xd8\x02\xe7\x89\x4a\x5a\xdf\xb8\xfa\x51\x79\xb2\xb9\x71\xc8\x89\x8f\xb9\xc1\x3a\xc8\xfd\xe8\x42\x21\x83\xec\xc2\xcc\x85\x9d\x6c\x50\xe7\xd6\xb1\x96\x92\xaa\x1b\x97\x6f\xdf\x2f\xaf\x5e\x7f\xff\x76\x7d\xfb\xe6\xd9\xcd\xab\x4f\x37\xcf\x3f\x2c\xbf\xbe\xfc\xf1\xf9\xdd\xf2\xc5\x97\xaa\x61\xc4\x8d\x39\x97\x3a\x66\xd0\x75\xbd\x6d\xdf\xfb\x67\xcb\xf2\xe3\xd5\xed\xf5\x53\xfb\xf8\xd2\x10\x83\x95\xf7\x6e\xd7\xed\xd5\x62\xea\xac\xc7\x42\x26\x8f\x0d\xea\x45\x3a\xc1\x8a\x47\x7c\x2c\x70\xa4\x71\x9a\x96\x4d\xb1\xa2\x07\x33\xae\x0d\x12\x83\x82\xa6\x41\xad\x44\x1b\xbb\x2e\x0c\x22\x5d\x60\xf3\xe6\xc7\xb1\xc0\x03\x99\xe7\xa6\xb5\x07\x0f\x15\xe6\x0d\xeb\xba\x9b\x1b\xf7\xcf\x22\x1b\x46\x88\x13\x5d\x0d\xbc\x9c\x71\x59\xfd\x1f\x1b\x53\xa0\xae\xaf\xdc\x6a\xe9\x19\x94\x2a\xd5\xc8\xa0\xef\x57\x5b\x34\xc8\x78\x2a\xd8\xca\x97\xc9\x48\x39\xf3\xb9\x33\x91\x99\x55\x65\xb7\x7c\xbb\xb7\xf3\x2b\xdd\xda\xa3\xf7\x7f\x2f\xd8\xe9\xfb\x7e\xb0\xdb\x3a\x7a\x17\x65\xa5\xc0\x0d\xfc\xfe\xee\x7c\x5e\x0b\x09\xc3\x13\x06\xd4\xb8\x3a\xd2\x32\xfb\xfb\x84\x48\xde\x6d\xc1\x7a\xf9\x67\x00\x00\x00\xff\xff\xf9\x52\x8c\xc7\x8f\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 911, mode: os.FileMode(438), modTime: time.Unix(1629956281, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
