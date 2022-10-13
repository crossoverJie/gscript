package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _internal_internal_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func internal_internal_go() ([]byte, error) {
	return bindata_read(
		_internal_internal_go,
		"internal/internal.go",
	)
}

var _internal_internal_gs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xcd\x6e\x1c\xb9\x11\xbe\xcf\x53\x54\x06\x88\x31\xb3\xd2\x8e\xa4\x3d\x24\xc0\x6a\x7b\x01\xcb\x59\x5b\x58\x78\x6d\x45\x52\x92\x83\x30\x07\x4e\x77\xcd\x0c\xad\x6e\xb2\x4d\xb2\x47\x1a\x09\x03\x04\xc9\x29\xa7\x05\x12\xe4\x01\x02\xe4\x90\x5c\x72\x49\x2e\xb9\xe4\x6d\x6c\x6c\xde\x22\x28\x92\xdd\xcd\x66\xb7\x6c\x39\xb9\x44\x07\xbb\x87\xac\x9f\xaf\xaa\x3e\x16\x7f\x0e\x0e\x40\xa1\xa9\x94\x00\xa6\x14\xdb\x42\x8e\x62\x65\xd6\x23\x2e\x0c\x7d\x4e\x98\xd8\x5e\xcd\x81\x4d\xef\x77\xa3\x51\x2c\x9a\xb2\xd2\xca\xa5\xac\x8c\xe5\x52\x59\x72\xd4\x80\x39\x16\x28\x8c\x86\xa5\x92\x05\x30\xd0\xb2\x52\x29\x7a\x75\x2e\x8c\x04\x06\x19\x6a\xc3\x05\x33\x5c\x7a\xbb\x33\x67\x54\x96\xdb\xc9\x62\x6b\xf0\x6a\x0e\x99\x36\xfb\xe0\xbf\xb5\x4a\x23\x30\x6b\xa6\xd7\xa9\xcc\xd0\xaa\xd1\x0f\x02\x03\xda\x4a\x95\x8a\x0b\x93\xdb\x30\x1a\x70\x4b\xa9\x0a\x66\x34\xb0\x34\x95\x2a\xe3\x62\x05\x16\x87\x1b\x06\x5d\x62\xca\x97\x1c\x15\x30\x91\xc1\x8d\xe2\x06\x35\x09\x68\xc3\x44\xc6\x54\x06\xb2\x32\x65\x65\x66\xce\xf4\x72\xa2\x8d\x22\x13\x4e\x7b\x1f\xc8\xd1\x6c\x36\xfb\x6f\x7c\xb9\x70\x34\x98\x35\x82\x42\x5d\xe5\x86\xe4\x9d\xfd\xd9\xc8\xfb\xd1\x8f\x72\x0b\xb5\xdf\x4a\x5b\x9f\x6b\x84\x0c\x97\xac\xca\x4d\x33\xb3\x94\x0a\xb8\xd1\x20\x4b\x54\x4c\x64\xfa\x51\xe1\x4e\x3a\x7e\x98\xd6\xa8\xcc\x37\x6f\x2b\x96\xbb\x04\x1f\x39\x20\xec\x8b\x96\x06\x1b\x54\xc6\xc7\x40\x66\xa9\x88\x75\x95\x7d\x41\x8d\x3c\xd9\x1a\x7c\x4a\x63\x75\x54\xb6\x76\xad\x7a\xab\xe4\x90\x75\x12\x62\xe4\x85\xfd\xa8\xb9\xb2\x20\xdd\x34\x67\x5a\x83\x9b\x38\xa9\x78\x9e\xa1\xba\x1f\x01\x40\x4d\xa2\x45\xb5\x84\x04\xae\x0e\xe7\xf7\xbb\xe3\x91\x9d\x38\x38\x00\x56\x96\x28\x32\x48\xa5\x30\x28\x6c\x12\x16\xd5\x72\x1f\xb8\xe9\x54\xc6\xad\x0f\x90\x4b\xd0\x56\x91\x48\x67\xd3\xe6\x61\xb4\x21\xd8\xe9\xc0\xa9\xc1\xa2\x84\xa4\x1b\xef\xf4\xb8\x91\x72\xee\x27\xd6\x27\x89\x06\x53\x9e\xe8\xb4\x1e\x83\x99\x9d\xfd\x37\x42\xbf\xf8\x38\xee\xc5\xac\x01\xfe\x2b\x02\x4e\x70\x74\x90\xbe\x41\x44\x8b\x61\x38\x8b\x06\x8b\xfd\x6f\xa5\xe4\xcd\x84\x0c\x8b\xc0\x0c\x5f\xc2\x44\xc0\xd7\x70\x38\x85\x76\xd0\xa3\xbe\x59\xa3\x20\x78\x0a\x81\x6b\x10\xd2\x00\x0a\x59\xad\xd6\xa0\x4b\x96\x12\xe8\xa5\x99\x75\x74\xc8\x18\xf5\x9b\x45\xb5\x9c\xc2\xe7\x0e\x03\x7d\x7e\x05\x22\x36\x1f\xa4\x5e\xe0\xcd\x89\x2b\x79\xad\xb0\x0f\x5f\x7c\x56\xdb\xd9\x13\x96\x07\xb1\xae\xed\x40\x4e\x73\x9f\x52\x3a\xed\x8b\x38\x1e\x39\x99\xee\xec\xae\xf9\xb5\xab\x8b\xe4\x73\xe4\x09\xe2\xe9\x12\xe4\xc9\xa7\xb5\x25\x74\xe3\x72\x37\xda\x8d\x3a\x9c\xd6\xf7\xa1\xa9\x37\x92\x0b\xcf\xbb\xab\xb9\xed\xb9\x7a\xbf\x9e\xd3\x58\x46\xa5\xa0\x0c\x58\x99\x29\x24\x49\xbf\x28\x1e\xc5\x78\xdc\xc6\xb3\xfb\x80\xfe\xd1\x03\xfa\x56\xe2\xea\x70\x3e\x64\x25\x5e\x18\x0b\x1d\x2f\x0b\x0c\xd9\x6f\xf9\x04\x89\xab\xb5\x9e\xc2\x67\x1d\x0c\x9f\x1f\x05\xa2\xd4\xd4\x2c\xff\x78\x72\x78\x0c\x1c\xbe\x82\x56\xf2\x18\xf8\xde\x5e\x8c\xb6\xce\x12\x24\x1e\x31\x9f\x77\xeb\xf8\x20\xc4\x88\x0d\x04\x50\xc0\x5e\x0d\xf2\x83\x61\x77\xfa\x12\xe8\x05\x24\xdd\xa1\x49\xa0\xae\x17\x33\xbb\xa8\x44\x38\xe6\x7b\x3f\x57\xda\x34\xc0\xc3\x54\xeb\xc5\x2c\x6c\x49\x56\x70\xea\x1b\x5d\x6b\xe0\x6a\x0e\x0a\x0b\xc6\x45\x63\xe3\xe8\xcb\x36\x5d\xf3\x4e\x56\xfb\x49\x75\xaa\x3e\xab\x51\x52\xbb\xee\xbb\xd5\x0c\x02\x50\x90\x78\x04\xbd\xb4\x47\x26\xd4\x60\x42\x3d\xd5\xf4\x62\x56\x2f\x26\x1f\xa3\x13\x59\x48\x99\xd3\x79\xe0\x4c\xe1\x92\xdf\x36\x8d\xb9\x59\x1a\xa5\x1d\xef\xf7\xe9\x0f\xd6\xba\x96\x89\x3b\xb9\x37\x76\xdc\x29\x74\xd8\x28\xf5\x14\xbe\xf6\x1c\x2e\xa7\xf0\xe4\x49\xb0\xce\xf5\xd5\xe1\x97\x7e\x62\x6e\x17\x95\x33\x16\xac\xfd\xa6\xbf\x6b\x18\x6f\xc6\xd4\xe2\xa9\xa5\x53\xbf\x97\x4b\x60\x7e\x6f\x1c\xb3\xf1\xc8\xb7\x6c\x7f\x1a\x73\xdb\xf1\x26\x3a\x2f\x7d\x7b\xf1\xfa\x95\xcf\x41\xbd\x89\xd2\x50\xe7\x8c\xe4\x84\xde\x56\xa8\xb6\x70\xc3\xcd\x1a\x4a\x66\xd6\x23\x92\xa0\x89\x17\x68\xea\x74\xbe\xd1\x52\xb4\x19\x65\x66\x6d\x0d\x90\x85\xd3\xcb\xcb\x33\xc8\xf9\x82\xbe\xcf\x51\x97\x52\x68\xb4\xe2\xa3\xe7\xee\x0c\x63\x9d\xba\xa3\x5e\x86\x1d\x1b\xcd\x0f\x12\xf7\x67\x81\x73\xd4\xd6\xc2\xda\x14\x79\x6d\xe1\xf4\xf2\xbb\x97\x1f\xb1\x40\xe2\x75\x4c\x34\x01\x13\x85\x39\x33\x7c\x83\xf6\xa7\x86\x82\x6d\x41\x16\x9c\x4e\xbb\xcc\x1e\xd0\x74\xce\xf4\x7a\x5a\x67\xe6\xe7\x94\x82\x33\x66\xd6\x93\x38\xc6\x26\xa1\x6e\x7f\x75\xab\x71\xc3\xf2\x0a\x6d\x1b\xa2\x31\xc1\x0a\xa4\x03\x45\x51\x4a\x81\xc2\x50\xb1\x68\xd8\xa6\xb5\x39\xc0\x3c\x97\xaa\xf8\x25\xa9\x4d\x86\xf0\x5f\xe3\xf6\x7f\xf4\x76\xf6\xfa\xe2\xb2\xf6\x75\x26\xb5\xf9\x44\x7f\xa0\xf0\x6d\x85\xda\x80\x62\x37\xb0\x90\xd9\xb6\xb6\x75\xee\xc6\x4f\x64\xb6\xed\xe5\xc6\xed\x55\xa7\xc6\x94\xcf\xe8\x34\x75\x6b\x3a\xfb\x15\x89\x39\x72\x47\x14\xf0\x6c\x8d\x1b\x1d\xb1\x00\x12\x27\xbc\x09\xd6\x62\xc8\x23\x67\xc0\x85\x62\x59\x13\x9e\x91\x22\x9e\x38\x37\xd0\xf3\x43\x5c\x81\x04\x36\x3d\x0f\x56\x3f\xf4\x60\x59\x15\x7a\xf0\x16\xde\x36\x6c\x19\x30\x4f\x2d\xa3\xa5\x93\xcd\x55\xef\x4c\x55\x1e\x0f\x1c\x15\x96\x71\xc5\x6c\x91\x62\xf3\x1b\x48\x02\x2e\x39\x9c\x24\xd8\xf3\xb1\x19\xf2\x51\x0e\x31\xe3\x41\x3f\x5d\x1e\x3d\xc2\x97\x6d\x8b\xa1\x3f\x15\xb0\xa7\xef\x83\x78\x06\x49\x87\x62\xc3\xf9\x22\xc1\xb6\x4f\x1e\x1c\xc0\x09\xa7\x9b\x94\xac\x0c\x8e\xd6\xc6\x94\xa7\x4c\x64\x79\x13\x4f\x81\x66\x2d\xb3\xa8\x51\x2c\x2b\x91\xc2\x24\xe0\xea\x14\xd6\x56\xcb\xc3\xa2\xc6\xe1\x6f\x91\x63\xd2\x48\xc6\x7b\x01\x96\x40\x0f\x52\x73\x0b\x49\x38\x52\x6f\xe4\xce\xde\x24\x35\xb7\xd3\x63\x07\xf3\xbc\x12\x40\xf8\x40\xa3\xda\xa0\x9a\x59\xb0\xe7\x55\x7d\x8c\x03\x96\x65\xaa\x5e\x87\x17\x5b\x6d\xb0\x00\xa9\x81\x95\x7c\xd4\xec\xdd\x2f\xd0\xbc\xbe\x78\xaa\x56\x7a\x42\x82\x5e\xef\x99\x2c\x0a\x26\xb2\xda\x0c\xb5\x84\x26\x5c\xba\xb7\xa9\x15\x09\xdb\x63\xff\x73\xde\x26\x66\xc9\x73\x7c\x15\xca\xda\xce\xb2\x6f\x8f\x5e\x25\xaa\x82\x94\xce\xb1\x90\x9b\x9e\x46\xe0\xfb\x05\x9a\x9b\x6c\xd2\xf4\x7f\xdd\x81\xed\x4f\xaf\x76\xec\xbe\xb9\x72\xbd\x40\x63\x05\xd4\x4a\xcf\x02\x76\x5c\xcd\x61\x15\x44\x17\x17\x3d\x88\x7c\x60\x0d\xa6\x8f\x48\x41\x6c\xb1\x4e\x9b\x13\x26\x91\xd0\x30\xed\xbe\xc6\xa0\x30\x5c\x8a\x2f\x6d\x3e\xe8\xb2\xe2\x2e\xd5\x29\x2f\x58\x6e\xa5\x6e\x3e\x3d\xa9\x0d\x8c\xb6\x20\xad\x92\x97\xb6\x92\x11\x1a\x65\x4b\xa1\x83\xae\x4f\x6a\x23\x17\xd0\x70\x95\x1a\x57\xbe\x8c\xcd\x4c\xcf\xb4\xdb\x64\x18\x28\x29\x0d\x66\x6e\xd3\x24\x27\x90\x4a\xa5\xec\x3e\x5e\xbf\x63\x90\xfb\xb4\x52\x8a\x36\x9b\x8c\x2b\x4c\x8d\xa4\x7d\x2d\xa8\xc5\xca\x53\x62\xa0\x82\x34\xde\x3d\xe1\xfc\x8c\x19\x0c\xb8\xf4\xcc\x59\xbe\xe4\x45\x13\x8e\xb9\x6b\x72\x99\xb3\xad\xac\x8c\x25\x1b\xe5\xf3\x17\xa2\x3d\xe1\x99\xbb\x60\x0f\x22\xa3\x64\xe2\x3e\x82\xf5\x38\xeb\x03\xb8\x43\x45\xd2\xf0\xa2\x61\x1a\x09\x50\x15\x01\x8a\x2d\x59\xc0\xe6\x2e\x4a\xc1\x33\x59\x94\x3c\x47\x15\xac\x74\xc8\xaa\xa2\x7c\x7a\x71\x39\x69\xe8\x9d\x85\x8b\x8e\x66\x2f\xb6\xc5\x42\xe6\xb1\x80\x4f\xc0\x37\xc2\xa8\xad\x73\x4f\xfb\xde\x35\x6e\xf7\x2d\xb3\x9c\x63\x3b\x0b\x02\x6f\x4d\xf0\xdb\x9e\x04\xaf\xfd\x3e\xb9\x5f\xcb\x04\x31\x5c\xe3\x36\xb9\x6e\x7b\xb1\xb5\x97\x04\xdb\x26\xd9\x4b\x44\x1b\x9a\x43\xf2\x1d\x2b\xef\x5b\x27\x57\x73\x30\x6c\x91\x23\x5d\xca\x8f\x7e\xd2\x5c\xc0\x29\x77\x9a\xdf\x61\x72\xe8\xcf\xf2\x65\xe5\x5e\x9d\x08\xb8\x43\x44\xde\x02\x30\x07\x07\x60\x64\x26\x21\x55\x52\x6b\xb9\x41\xf5\x2d\x47\x78\xff\xbb\xbf\xbe\xfb\xdb\x3f\x83\x7b\xeb\x84\x30\x27\xe3\xf1\x74\xe8\xb6\x3a\x78\xd5\xf5\x8f\x89\x94\x4e\x48\xdc\xbb\x62\x77\x8b\xb3\x57\x22\x3f\x65\xa5\x7e\xec\x5e\x68\x28\xaa\x29\xb4\x72\x2e\x7f\x64\xc4\x4e\x75\xee\x3a\x74\xa1\x46\xf8\x51\x02\x82\xe7\x11\x34\xa7\x66\xb0\x28\xdd\x57\x02\xd8\xbd\x23\xd9\xcb\x6e\x3b\xef\x8d\x0c\x3c\x82\xd8\x93\x8f\xa3\xed\x29\xd3\xeb\x3a\x9a\x46\x75\xd6\x8d\x2b\xc4\xd6\x51\x0b\x42\x7d\xf2\x84\x0a\x42\x43\x5d\x2b\x7d\xe7\xf4\xd7\xca\x38\xaa\x40\xc0\xc0\xf8\x2f\xae\x48\xbf\x32\x3d\xab\x10\xa2\x68\xb9\x3c\xac\x59\xd7\x00\x12\x4f\x76\x4b\x2c\xdf\x6c\x31\xbe\xa8\xf2\x3b\xdc\xdb\x0b\x6e\xce\x3b\xc0\x5c\x63\xff\x1d\xeb\xdd\xf7\xbf\xf9\xe1\xd7\xbf\x85\x37\xd9\xf5\xd1\xec\xa7\xf0\xef\x3f\xfc\xeb\x87\x3f\xfd\xe5\xdd\x9f\xff\xf1\xfe\xfb\xdf\xbf\xff\xfb\x1f\x3f\xee\xdf\xb9\xa7\xf2\x0d\xfb\x6f\x43\x09\x8f\x6e\xb4\x1a\x56\xd8\x2c\x8f\xce\x3b\xcf\xff\x29\xdf\x2d\x67\xd1\x16\x69\x98\xf5\x0f\x70\x15\x07\x38\x4a\xfc\xf4\x24\xb4\xd3\x44\xca\x07\xf8\x3a\xc0\xcb\xfa\x9d\x6a\x36\xc0\xc5\x2e\x63\x28\x12\x8c\x78\xd5\x7d\x12\x0b\x20\x76\x56\x49\xb0\x4c\x70\x60\x79\x3c\x04\x61\x37\xc0\x31\x2f\x2b\x78\x1e\x83\xf0\x64\xa0\xc4\xad\xd0\x5c\xf0\x3b\x1c\xd8\x70\x89\x48\xc1\x56\x43\x7b\xcd\x4b\x2e\xae\x31\x7b\xc9\xb5\xf1\xfd\xd9\x0d\xbc\x92\x19\xb6\xdb\x45\x80\xac\x9d\x0e\xf6\x8b\x76\x70\x12\x5f\xde\xdc\xe5\xb4\x73\x9d\xb2\x65\x4f\xda\x20\xda\xbd\xa1\x05\x73\x1f\x3b\xb3\x57\x5d\xda\x67\xb5\x79\x60\x8b\x60\x59\xd6\xf3\x4e\x65\xd9\x50\xe2\x3f\x61\x09\x04\x4e\xe9\x1e\x18\xc4\x16\xde\x3a\xc9\xb2\x7f\x7a\x1b\x22\x70\xfd\x2a\x97\x77\x29\x45\xf0\xbb\xa3\xae\x95\xdc\xf7\xa4\x66\x3e\x4b\x1f\x37\xd0\x6b\x15\x30\xd0\x2b\xe8\x80\x10\x12\x22\x08\x52\x1b\xa6\xc8\xa6\x85\x1c\xad\x51\x37\x37\xb8\x44\xeb\xdb\x90\x15\x71\xdc\xed\xbd\xef\x39\xc3\x4e\x62\x68\xe5\x7c\x3a\x69\xff\x13\x00\x00\xff\xff\xf0\xc4\x37\xff\xa8\x1c\x00\x00")

func internal_internal_gs() ([]byte, error) {
	return bindata_read(
		_internal_internal_gs,
		"internal/internal.gs",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"internal/internal.go": internal_internal_go,
	"internal/internal.gs": internal_internal_gs,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"internal": &_bintree_t{nil, map[string]*_bintree_t{
		"internal.go": &_bintree_t{internal_internal_go, map[string]*_bintree_t{}},
		"internal.gs": &_bintree_t{internal_internal_gs, map[string]*_bintree_t{}},
	}},
}}
