package bitmath4go

import "math"

var byteValues = map[string]float64{
	"KiB": math.Pow(2, 10),
	"MiB": math.Pow(2, 20),
	"GiB": math.Pow(2, 30),
	"TiB": math.Pow(2, 40),
	"PiB": math.Pow(2, 50),
	"EiB": math.Pow(2, 60),
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

var (
	Kilibyte = generateSizeFunction("KiB", byteValues["KiB"])
	KiB      = Kilibyte
	Mebibyte = generateSizeFunction("MiB", byteValues["MiB"])
	MiB      = Mebibyte
	Gibibyte = generateSizeFunction("GiB", byteValues["GiB"])
	GiB      = Gibibyte
	Tebibyte = generateSizeFunction("TiB", byteValues["TiB"])
	TiB      = Tebibyte
	Pebibyte = generateSizeFunction("PiB", byteValues["PiB"])
	PiB      = Pebibyte
	Exbibyte = generateSizeFunction("EiB", byteValues["EiB"])
	EiB      = Exbibyte
)
