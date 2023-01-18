package ds18b20

const (
	SKIP_ROM byte = 0xCC
	READ_ROM byte = 0x33
	MATCH_ROM byte = 0x55
	SEARCH_ROM byte = 0xF0
	CONVERT_T byte = 0x44
	READ_SCRETCHPAD byte = 0xBE
)