package bitmath4go

import "math"

var byteValues = map[string]float64{
	"Byte": math.Pow(2, 0),
	"KiB":  math.Pow(2, 10),
	"MiB":  math.Pow(2, 20),
	"GiB":  math.Pow(2, 30),
	"TiB":  math.Pow(2, 40),
	"PiB":  math.Pow(2, 50),
	"EiB":  math.Pow(2, 60),
}

type BitmathBase struct {
	ByteValue float64
	StatedValue
}

type StatedValue struct {
	Prefix    string
	BaseValue float64
	Value     float64
}

func generateSizeFunction(prefix string, baseValue float64) func(float64) BitmathBase {
	return func(value float64) BitmathBase {
		return BitmathBase{
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
	Byte = generateSizeFunction("Byte", byteValues["Byte"])
	KiB  = generateSizeFunction("KiB", byteValues["KiB"])
	MiB  = generateSizeFunction("MiB", byteValues["MiB"])
	GiB  = generateSizeFunction("GiB", byteValues["GiB"])
	TiB  = generateSizeFunction("TiB", byteValues["TiB"])
	PiB  = generateSizeFunction("PiB", byteValues["PiB"])
	EiB  = generateSizeFunction("EiB", byteValues["EiB"])
)
