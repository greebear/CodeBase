package NewStruct

type NewStruct struct {
	BoolKey bool             `json:"bool_key" yaml:"bool_key"`
	Map1    map[int]struct{} `json:"map_1" yaml:"map_1"`
	Map2    map[int]struct{} `json:"map_2" yaml:"map_2"`
}

func GetConfig(isOpen bool) *NewStruct {
	if isOpen {
		c := NewStruct{
			BoolKey: true,
			Map1:    nil,
			Map2:    nil,
		}
		return &c
	}
	return nil
}
