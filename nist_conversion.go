package bitmath4go

func (b BitmathBase) AsKilibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "KiB",
			BaseValue: byteValues["KiB"],
			Value:     b.ByteValue / byteValues["KiB"],
		},
	}
}

func (b BitmathBase) AsMebibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "MiB",
			BaseValue: byteValues["MiB"],
			Value:     b.ByteValue / byteValues["MiB"],
		},
	}
}

func (b BitmathBase) AsGibibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "GiB",
			BaseValue: byteValues["GiB"],
			Value:     b.ByteValue / byteValues["GiB"],
		},
	}
}

func (b BitmathBase) AsTebibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "TiB",
			BaseValue: byteValues["TiB"],
			Value:     b.ByteValue / byteValues["TiB"],
		},
	}
}

func (b BitmathBase) AsPebibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "PiB",
			BaseValue: byteValues["PiB"],
			Value:     b.ByteValue / byteValues["PiB"],
		},
	}
}

func (b BitmathBase) AsExbibytes() BitmathBase {
	return BitmathBase{
		ByteValue: b.ByteValue,
		StatedValue: StatedValue{
			Prefix:    "EiB",
			BaseValue: byteValues["EiB"],
			Value:     b.ByteValue / byteValues["EiB"],
		},
	}
}
