package csmrand

import _ "unsafe"

//go:linkname fastrand64 runtime.fastrand64
func fastrand64() uint64
