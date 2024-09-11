// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__parse_with_padding = 688
)

const (
    _stack__parse_with_padding = 192
)

const (
    _size__parse_with_padding = 48776
)

var (
    _pcsp__parse_with_padding = [][2]uint32{
        {0x1, 0},
        {0x6, 8},
        {0x8, 16},
        {0xa, 24},
        {0xc, 32},
        {0xd, 40},
        {0x14, 48},
        {0xbe2, 192},
        {0xbe3, 48},
        {0xbe5, 40},
        {0xbe7, 32},
        {0xbe9, 24},
        {0xbeb, 16},
        {0xbec, 8},
        {0xbf0, 0},
        {0xbe88, 192},
    }
)

var _cfunc_parse_with_padding = []loader.CFunc{
    {"_parse_with_padding_entry", 0,  _entry__parse_with_padding, 0, nil},
    {"_parse_with_padding", _entry__parse_with_padding, _size__parse_with_padding, _stack__parse_with_padding, _pcsp__parse_with_padding},
}
