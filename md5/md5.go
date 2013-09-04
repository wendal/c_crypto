package md5

// #include <stdlib.h>
// #include <openssl/md5.h>
// #cgo LDFLAGS: -lcrypto
import "C"
import (
        "fmt"
        "unsafe"
)

type M struct {
        ctx *C.MD5_CTX
}

func New() *M {
        m := &M{}
        m.ctx = new(C.MD5_CTX)
        C.MD5_Init(m.ctx)
        return m
}

func (m *M) Write(data []byte) (n int, err error) {
        n = len(data)
        C.MD5_Update(m.ctx, unsafe.Pointer(&data[0]), C.size_t(n))
        return
}

func (m *M) Final() string {
        re := (*_Ctype_unsignedchar)(C.malloc(16))
        C.MD5_Final(re, m.ctx)
        dst := fmt.Sprintf("%02X", C.GoBytes(unsafe.Pointer(re), 16))
        C.free(unsafe.Pointer(re))
        return dst
}