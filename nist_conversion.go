package bitmath4go

func (b Byte) AsKilibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "KiB",
			BaseValue: byteValues["KiB"],
			Value:     b.ByteValue / byteValues["KiB"],
		},
	}
}

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

func (b Byte) AsExbibytes() Byte {
	return Byte{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "EiB",
			BaseValue: byteValues["EiB"],
			Value:     b.ByteValue / byteValues["EiB"],
		},
	}
}
