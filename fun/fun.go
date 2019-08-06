package fun

import (
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
	"reflect"
	"strconv"
)

// GetMd5Str computes md5 and returns the hex string
func GetMd5Str(str string) string {
	b := []byte{}
	for _, v := range md5.Sum([]byte(str)) {
		b = append(b, v)
	}
	return hex.EncodeToString(b)
}

// Str2Uint32 converts string to uint32
func Str2Uint32(str string) uint32 {
	if i, err := strconv.Atoi(str); err == nil {
		return uint32(i)
	}
	return 0
}

// GetUint32ByName returns o.name as uint32
func GetUint32ByName(o interface{}, name string) uint32 {
	if fd := reflect.ValueOf(o).Elem().FieldByName(name); fd.IsValid() {
		return uint32(fd.Uint())
	}
	return 0
}

// SetUint32ByName sets o.name to uint32
func SetUint32ByName(o interface{}, name string, val uint32) {
	if fd := reflect.ValueOf(o).Elem().FieldByName(name); fd.IsValid() {
		fd.SetUint(uint64(val))
	}
}

// GetStrByName returns o.name as string
func GetStrByName(o interface{}, name string) string {
	if fd := reflect.ValueOf(o).Elem().FieldByName(name); fd.IsValid() {
		return fd.String()
	}
	return ""
}

// SetStrByName sets o.name to string
func SetStrByName(o interface{}, name string, val string) {
	if fd := reflect.ValueOf(o).Elem().FieldByName(name); fd.IsValid() {
		fd.SetString(val)
	}
}

// SetByObj sets o2 to o1
func SetByObj(o1 interface{}, o2 interface{}) {
	t1 := reflect.TypeOf(o1).Elem()
	t2 := reflect.TypeOf(o2).Elem()
	v1 := reflect.ValueOf(o1).Elem()
	v2 := reflect.ValueOf(o2).Elem()

	if t1 == t2 {
		for i := 0; i < v2.NumField(); i++ {
			v1.Field(i).Set(v2.Field(i))
		}
	}
}

// ToAbs converts to absolute path
func ToAbs(p string) string {
	var err error
	p, err = filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	return p
}
