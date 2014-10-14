package bitmath

var byteValues = map[string]float64{
	"MiB": 1048576,
	"GiB": 1073741824,
	"TiB": 1099511627776,
	"PiB": 1125899906842624,
}

type Byte struct {
	ByteValue float64
	StatedValue
}

type StatedValue struct {
	Prefix    string
	BaseValue float64
	Value     float64
}

func generateSizeFunction(prefix string, baseValue float64) func(float64) Byte {
	return func(value float64) Byte {
		return Byte{
			ByteValue: value * baseValue,
			StatedValue: StatedValue{
				Prefix:    prefix,
				BaseValue: baseValue,
				Value:     value,
			},
		}
	}
}

var Mebibyte = generateSizeFunction("MiB", byteValues["MiB"])
var MiB = Mebibyte
var Gibibyte = generateSizeFunction("GiB", byteValues["GiB"])
var GiB = Gibibyte
var Tebibyte = generateSizeFunction("TiB", byteValues["TiB"])
var TiB = Tebibyte
var Pebibyte = generateSizeFunction("PiB", byteValues["PiB"])
var PiB = Pebibyte
