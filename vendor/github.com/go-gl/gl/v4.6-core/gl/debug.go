// Code generated by glow (https://github.com/go-gl/glow). DO NOT EDIT.

package gl

import "C"
import "unsafe"

type DebugProc func(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message string,
	userParam unsafe.Pointer)

var userDebugCallback DebugProc

//export glowDebugCallback_glcore46
func glowDebugCallback_glcore46(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message *uint8,
	userParam unsafe.Pointer) {
	if userDebugCallback != nil {
		userDebugCallback(source, gltype, id, severity, length, GoStr(message), userParam)
	}
}