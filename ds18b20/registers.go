package ds18b20

const (
	SKIP_ROM      uint8    = 0xCC
	READ_ROM         byte = 0x33
	MATCH_ROM        byte = 0x55
	SEARCH_ROM       byte = 0xF0
	CONVERT_T        byte = 0x44
	READ_SCRETCHPAD  byte = 0xBE
	WRITE_SCRETCHPAD byte = 0x4E
	RESOLUTION_9BIT  byte = 0x1F
	RESOLUTION_10BIT byte = 0x3F
	RESOLUTION_11BIT byte = 0x5F
	RESOLUTION_12BIT byte = 0x7F
)
