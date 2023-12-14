package constant

var (
	BitRate = map[string]string{
		"avc":  "5000K",
		"hevc": "1800K",
	}
)

const (
	Type      = iota + 1
	Kilobyte  = 1000 * Type
	Megabyte  = 1000 * Kilobyte
	Gigabyte  = 1000 * Megabyte
	Terabyte  = 1000 * Gigabyte
	Petabyte  = 1000 * Terabyte
	Exabyte   = 1000 * Petabyte
	Zettabyte = 1000 * Exabyte
	Yottabyte = 1000 * Zettabyte
)
