package sdhash

/*
#cgo CXXFLAGS: -Wall -std=c++11 -I/usr/local/include/
#cgo LDFLAGS: -lstdc++ -lsdhash
#include <stdint.h>
#include <libsdhash/helper.h>
*/
import "C"
import "unsafe"

func SDHash_From_Buffer(fileName string, buffer []byte) string {
	cFname := C.CString(fileName)
	cBuffArray := (*C.char)(unsafe.Pointer(&buffer[0]))
	cBuffSize := C.ulong(len(buffer))
	cSDHash := C.sdhash_from_buffer(cFname, cBuffArray, cBuffSize)
	return C.GoString(cSDHash)
}

func SDHash_From_FPath(filePath string) string {
	cFname := C.CString(filePath)
	chash := C.sdhash_from_path(cFname)
	return C.GoString(chash)
}

func SDHash_Compare_Hashes(hash1, hash2 string) int {
	chash1 := C.CString(hash1)
	chash2 := C.CString(hash2)
	cSim := C.sdhash_compare_hashes_helper(chash1, chash2)
	return int(cSim)
}
