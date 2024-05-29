// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__vunsigned = 0
)

const (
    _stack__vunsigned = 8
)

const (
    _size__vunsigned = 324
)

var (
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {78, 8},
        {79, 0},
        {90, 8},
        {91, 0},
        {114, 8},
        {115, 0},
        {273, 8},
        {274, 0},
        {312, 8},
        {313, 0},
        {320, 8},
        {324, 0},
    }
)

var _cfunc_vunsigned = []loader.CFunc{
    {"_vunsigned_entry", 0,  _entry__vunsigned, 0, nil},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}