package models

type Building struct {
	BuildingId             string `json:"building_id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	BuildingName           string `json:"building_name"`
	Halls                  uint   `json:"halls"`
	Rooms                  uint   `json:"rooms"`
	Lighting               string `json:"lighting"`
	Accessibility          string `json:"accessibility"`
	Ramps                  string `json:"ramps"`
	Lift                   string `json:"lift"`
	Washrooms              uint   `json:"washrooms"`
	WashroomsAccessibility string `json:"washroom_access"`
}

type Halls struct {
	HallId   uint   `json:"id"`
	HallName string `json:"hall_name"`
	Capacity string `json:"capacity"`
}

type Rooms struct {
	HallId   uint   `json:"id"`
	HallName string `json:"hall_name"`
	Capacity string `json:"hall"`
}
