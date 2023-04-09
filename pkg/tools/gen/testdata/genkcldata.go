package testdata

// Person Example
type Person struct {
	name         string            `kcl:"name=name,type=str"`           // kcl-type: str
	age          int               `kcl:"name=age,type=int"`            // kcl-type: int
	friends      []string          `kcl:"name=friends,type=[str]"`      // kcl-type: [str]
	movies       map[string]*Movie `kcl:"name=movies,type={str:Movie}"` // kcl-type: {str:Movie}
	MapInterface map[string]map[string]interface{}
	Ep           *employee
	Com          Company
	StarInt      *int
	StarMap      map[string]string
	Inter        interface{}
}

type Movie struct {
	desc     string      `kcl:"nam=desc,type=str"`                                    // kcl-type: str
	size     int         `kcl:"name=size,type=units.NumberMultiplier"`                // kcl-type: units.NumberMultiplier
	kind     string      `kcl:"name=kind?,type=str(Superhero)|str(War)|str(Unknown)"` // kcl-type: "Superhero"|"War"|"Unknown"
	unknown1 interface{} `kcl:"name=unknown1?,type=int|str"`                          // kcl-type: int|str
	unknown2 interface{} `kcl:"name=unknown2?,type=any"`                              // kcl-type: any
}

type employee struct {
	name        string            `kcl:"name=name,type=str"`           // kcl-type: str
	age         int               `kcl:"name=age,type=int"`            // kcl-type: int
	friends     []string          `kcl:"name=friends,type=[str]"`      // kcl-type: [str]
	movies      map[string]*Movie `kcl:"name=movies,type={str:Movie}"` // kcl-type: {str:Movie}
	bankCard    int               `kcl:"name=bankCard,type=int"`       // kcl-type: int
	nationality string            `kcl:"name=nationality,type=str"`    // kcl-type: str
}

type Company struct {
	name      string      `kcl:"name=name,type=str"`             // kcl-type: str
	employees []*employee `kcl:"name=employees,type=[employee]"` // kcl-type: [employee]
	persons   *Person     `kcl:"name=persons,type=Person"`       // kcl-type: Person
}
