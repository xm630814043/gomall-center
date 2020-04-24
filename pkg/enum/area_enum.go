package enum

type Arealevel int32

const (
	AREA_LEVEL Arealevel = iota
	AREA_LEVEL_PROVINCES
	AREA_LEVEL_CITIES
	AREA_LEVEL_COUNTIES
)
