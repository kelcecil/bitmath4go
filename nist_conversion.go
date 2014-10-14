package bitmath

func (b Byte) AsMebibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "MiB",
			BaseValue: byteValues["MiB"],
			Value:     b.ByteValue / byteValues["MiB"],
		},
	}
}

func (b Byte) AsGibibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "GiB",
			BaseValue: byteValues["GiB"],
			Value:     b.ByteValue / byteValues["GiB"],
		},
	}
}

func (b Byte) AsTebibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "TiB",
			BaseValue: byteValues["TiB"],
			Value:     b.ByteValue / byteValues["TiB"],
		},
	}
}

func (b Byte) AsPebibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "PiB",
			BaseValue: byteValues["PiB"],
			Value:     b.ByteValue / byteValues["PiB"],
		},
	}
}
