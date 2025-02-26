package matf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var (
	verySimpleMatrix = []byte{0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
		0x4d, 0x61, 0x54, 0x72, 0x49, 0x78, 0x00, 0x00, 0x09, 0x00, 0x00,
		0x00, 0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xf0, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
		0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f}
	verySimpleStruct = []byte{0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00,
		0x74, 0x65, 0x73, 0x74, 0x69, 0x6E, 0x67, 0x5F, 0x73, 0x74, 0x72,
		0x75, 0x63, 0x74, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x40, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x66,
		0x69, 0x65, 0x6C, 0x64, 0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x69, 0x65,
		0x6C, 0x64, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00, 0x38,
		0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
		0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xF0, 0x3F, 0x0E, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x06, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x08,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x40}
	verySimple3DMatrix = []byte{0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x06, 0x08, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x6d, 0x61, 0x74,
		0x72, 0x69, 0x78, 0x33, 0x64, 0x09, 0x00, 0x00, 0x00, 0xe0, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x45, 0x40, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00,
		0x00, 0x00, 0xe0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xf0, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00}
	verySimpleCell = []byte{0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x04, 0x00, 0x63, 0x65, 0x6c, 0x6c, 0x0e, 0x00, 0x00, 0x00,
		0x30, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00,
		0x00, 0x30, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00,
		0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00,
		0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08,
		0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x05, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	verySimpleChar = []byte{0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0d, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x63,
		0x6f, 0x6c, 0x6c, 0x32, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x4e, 0x00, 0x00, 0x00, 0x73, 0x00, 0x53, 0x00, 0x23, 0x00, 0x74,
		0x00, 0x54, 0x00, 0x20, 0x00, 0x72, 0x00, 0x52, 0x00, 0x61, 0x00,
		0x69, 0x00, 0x49, 0x00, 0x20, 0x00, 0x6e, 0x00, 0x4e, 0x00, 0x62,
		0x00, 0x67, 0x00, 0x47, 0x00, 0x20, 0x00, 0x31, 0x00, 0x32, 0x00,
		0x63, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20,
		0x00, 0x64, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00,
		0x20, 0x00, 0x65, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20, 0x00, 0x20,
		0x00, 0x20, 0x00, 0x66, 0x00, 0x00, 0x00}
)

func TestExtractDataElement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		data          []byte
		order         binary.ByteOrder
		dataType      int
		numberOfBytes int
		step          int
		ele           interface{}
		err           string
	}{
		{name: "TooLessData", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: 42, numberOfBytes: 42, step: 0, err: "unexpected EOF"},
		{name: "Unknown", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: 42, numberOfBytes: 4, step: 0, err: "is not supported"},
		{name: "MiInt8", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiInt8, numberOfBytes: 1, step: 1, ele: []interface{}{17}},
		{name: "MiUint8", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiUint8, numberOfBytes: 1, step: 1, ele: 17},
		{name: "MiInt16", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiInt16, numberOfBytes: 2, step: 2, ele: 8721},
		{name: "MiUint16", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiUint16, numberOfBytes: 2, step: 2, ele: 8721},
		{name: "MiInt32", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiInt32, numberOfBytes: 4, step: 4, ele: 1144201745},
		{name: "MiUint32", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiUint32, numberOfBytes: 4, step: 4, ele: 1144201745},
		{name: "MiInt32", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.BigEndian, dataType: MiInt32, numberOfBytes: 4, step: 4, ele: 287454020},
		{name: "MiUint32", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.BigEndian, dataType: MiUint32, numberOfBytes: 4, step: 4, ele: 287454020},
		{name: "MiSingle", data: []byte{0x11, 0x22, 0x33, 0x44}, order: binary.BigEndian, dataType: MiSingle, numberOfBytes: 4, step: 4, ele: 1.2795344e-28},
		{name: "MiInt64", data: []byte{0x11, 0x22, 0x33, 0x44, 0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiInt64, numberOfBytes: 8, step: 8, ele: 4914309075945333265},
		{name: "MiUint64", data: []byte{0x11, 0x22, 0x33, 0x44, 0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiUint64, numberOfBytes: 8, step: 8, ele: 4914309075945333265},
		{name: "MiDouble", data: []byte{0x11, 0x22, 0x33, 0x44, 0x11, 0x22, 0x33, 0x44}, order: binary.LittleEndian, dataType: MiDouble, numberOfBytes: 8, step: 8, ele: 3.529429556587807e+20},
		{name: "MiMatrix", data: verySimpleMatrix, order: binary.LittleEndian, dataType: MiMatrix, numberOfBytes: 1, step: 144, ele: []interface{}{MatMatrix{Name: "MaTrIx", Flags: 0x6, Class: 0x6, Dim: Dim{3, 3, 0}, Content: NumPrt{RealPart: []interface{}{1, 0, 1, 0, 1, 0, 1, 0, 1}, ImaginaryPart: interface{}(nil)}}}},
		{name: "MiStruct", data: verySimpleStruct, order: binary.LittleEndian, dataType: MiMatrix, numberOfBytes: 1, step: 344, ele: []interface{}{MatMatrix{Name: "testing_struct", Flags: 0x2, Class: 0x2, Dim: Dim{1, 1, 0}, Content: StructPrt{FieldNames: []string{"field1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "field2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"}, FieldValues: map[string][]interface{}{"field1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00": {MatMatrix{Name: "", Flags: 0x6, Class: 0x6, Dim: Dim{1, 1, 0}, Content: NumPrt{RealPart: []interface{}{1}, ImaginaryPart: interface{}(nil)}}}, "field2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00": {MatMatrix{Name: "", Flags: 0x6, Class: 0x6, Dim: Dim{1, 1, 0}, Content: NumPrt{RealPart: []interface{}{2}, ImaginaryPart: interface{}(nil)}}}}}}}},
		{name: "Mi3dMatrix", data: verySimple3DMatrix, order: binary.LittleEndian, dataType: MiMatrix, numberOfBytes: 1, step: 1048, ele: []interface{}{MatMatrix{Name: "matrix3d", Flags: 0x806, Class: 0x6, Dim: Dim{3, 4, 5}, Content: NumPrt{RealPart: []interface{}{42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, ImaginaryPart: []interface{}{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}}}},
		{name: "MiCell", data: verySimpleCell, order: binary.LittleEndian, dataType: MiMatrix, numberOfBytes: 1, step: 128, ele: []interface{}{MatMatrix{Name: "cell", Flags: 0x1, Class: 0x1, Dim: Dim{1, 1, 3}, Content: CellPrt{Cells: []MatMatrix{{Name: "", Flags: 0x6, Class: 0x6, Dim: Dim{0, 0, 0}, Content: NumPrt{RealPart: []interface{}(nil), ImaginaryPart: interface{}(nil)}}, {Name: "", Flags: 0x6, Class: 0x6, Dim: Dim{0, 0, 0}, Content: NumPrt{RealPart: []interface{}(nil), ImaginaryPart: interface{}(nil)}}, {Name: "", Flags: 0x6, Class: 0x6, Dim: Dim{0, 0, 0}, Content: NumPrt{RealPart: []interface{}(nil), ImaginaryPart: interface{}(nil)}}}}}}},
		{name: "MxCharClass", data: verySimpleChar, order: binary.LittleEndian, dataType: MiMatrix, numberOfBytes: 1, step: 160, ele: []interface{}{MatMatrix{Name: "coll2", Flags: 0x4, Class: 0x4, Dim: Dim{3, 13, 0}, Content: CharPrt{Chars: []string{"string1      ", "STRING2      ", "# a b c d e f"}}}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewReader(tc.data)
			ele, step, err := extractDataElement(r, tc.order, tc.dataType, tc.numberOfBytes)
			if err != nil {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); !matched {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				} else {
					return
				}
				t.Fatalf("Expected no error, got: %v", err)
			} else if len(tc.err) != 0 {
				t.Fatalf("Expected error, got none")
			}
			if step != tc.step {
				t.Fatalf("Step\tExpected: %d \t Got: %d", tc.step, step)
			}
			fmt.Printf("Expected: %#v\tGot: %#v\n", tc.ele, ele)
		})
	}
}

func TestExtractNumeric(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		data  []byte
		order binary.ByteOrder
		step  int
		ele   interface{}
		err   string
	}{
		{name: "[1,2]", data: []byte{0x09, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40}, order: binary.LittleEndian, step: 24, ele: []int{1, 2}},
		{name: "SmallData", data: []byte{0x06, 0x00, 0x04, 0x00, 0x01, 0x03, 0x03, 0x07}, order: binary.LittleEndian, step: 8, ele: []int{117637889}},
		{name: "TooFewBytes", data: []byte{0x01, 0x10}, order: binary.LittleEndian, step: 1, err: "Unable to read"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewReader(tc.data)
			ele, step, err := extractNumeric(r, tc.order)
			if err != nil {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); !matched {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				} else {
					return
				}
				t.Fatalf("Expected no error, got: %v", err)
			} else if len(tc.err) != 0 {
				t.Fatalf("Expected error, got none")
			}
			if step != tc.step {
				t.Fatalf("Step\tExpected: %d \t Got: %d", tc.step, step)
			}
			fmt.Printf("Expected: %#v\tGot: %#v\n", tc.ele, ele)
		})
	}
}

func TestExtractFieldNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		data            []byte
		order           binary.ByteOrder
		fieldNameLength int
		numberOfFields  int
		fields          []string
		err             string
	}{
		{name: "['abc']", data: []byte{0x61, 0x62, 0x63}, fieldNameLength: 3, numberOfFields: 1, fields: []string{"abc"}},
		{name: "0", data: []byte{0x00}, fieldNameLength: 0, numberOfFields: 0},
		{name: "UnableToRead", data: []byte{0x61, 0x62, 0x63}, fieldNameLength: 9, numberOfFields: 1, err: "unexpected EOF"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewReader(tc.data)
			fields, err := extractFieldNames(r, tc.order, tc.fieldNameLength, tc.numberOfFields)
			if err != nil {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); !matched {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				} else {
					return
				}
				t.Fatalf("Expected no error, got: %v", err)
			} else if len(tc.err) != 0 {
				t.Fatalf("Expected error, got none")
			}
			if !reflect.DeepEqual(tc.fields, fields) {
				t.Fatalf("Fields\tExpected: %#v\tGot: %#v\n", tc.fields, fields)
			}
		})
	}
}

func TestExtractArrayName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		data      []byte
		order     binary.ByteOrder
		arrayName string
		step      int
		err       string
	}{
		{name: "ThisIsALongerName", data: []byte{0x01, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x54, 0x68, 0x69, 0x73, 0x49, 0x73, 0x41, 0x4c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, order: binary.LittleEndian, step: 25, arrayName: "ThisIsALongerName"},
		{name: "TooFewBytes", data: []byte{0x01, 0x10}, order: binary.LittleEndian, step: 1, err: "Unable to read"},
		{name: "ZeroLength", data: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, order: binary.LittleEndian, step: 8, arrayName: ""},
		{name: "UnableToRead", data: []byte{0x01, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00}, order: binary.LittleEndian, err: "Unable to read"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewReader(tc.data)
			name, step, err := extractArrayName(r, tc.order)
			if err != nil {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); !matched {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				} else {
					return
				}
				t.Fatalf("Expected no error, got: %v", err)
			} else if len(tc.err) != 0 {
				t.Fatalf("Expected error, got none")
			}
			if step != tc.step {
				t.Fatalf("Step\tExpected: %d \t Got: %d", tc.step, step)
			}
			if strings.Compare(name, tc.arrayName) != 0 {
				t.Fatalf("Fields\tExpected: %#v\tGot: %#v\n", tc.arrayName, name)
			}
		})
	}
}
