package fixtures

import "time"

type Types struct {
	TheString string `json:"the_string"`

	TheBool bool `json:"the_bool"`

	TheInt8    int8    `json:"the_int8"`
	TheUint8   uint8   `json:"the_uint8"`
	TheInt16   int16   `json:"the_int16"`
	TheUint16  uint16  `json:"the_uint16"`
	TheInt32   int32   `json:"the_int32"`
	TheUint32  uint32  `json:"the_uint32"`
	TheInt64   int64   `json:"the_int64"`
	TheUint64  uint64  `json:"the_uint64"`
	TheInt     int     `json:"the_int"`
	TheUint    uint    `json:"the_uint"`
	TheUintptr uintptr `json:"the_uintptr"`

	TheFloat32 float32 `json:"the_float32"`
	TheFloat64 float64 `json:"the_float64"`

	TheComplex64  complex64  `json:"the_complex64"`
	TheComplex128 complex128 `json:"the_complex128"`

	TheByte byte `json:"the_byte"`

	TheRune rune `json:"the_rune"`

	// TheTime represent a string format date-time
	TheTime time.Time `json:"the_time"`

	TheArrayBool   []bool      `json:"the_array_bool"`
	TheArrayInt    []int       `json:"the_array_int"`
	TheArrayInt64  []int64     `json:"the_array_int64"`
	TheArrayString []string    `json:"the_array_string"`
	// TheArrayTime   []time.Time `json:"the_array_string"`
	// TheArrayStruct []Color `json:"the_array_struct"`

	// TheMapStringInt represent an object
	TheMapStringInt map[string]int `json:"the_map_string_int"`
}
