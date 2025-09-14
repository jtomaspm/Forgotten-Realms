package game

import (
	"fmt"
	"time"
)

type VillageCapacity struct {
	Warehouse  float32 `json:"warehouse"`
	Population int     `json:"population"`
}

type Village struct {
	Coordinates Coordinates        `json:"coordinates"`
	Resources   Resources          `json:"resources"`
	Production  ResourceProduction `json:"production"`
	Capacity    VillageCapacity    `json:"capacity"`
	Population  int                `json:"population"`
}

func NewVillage(coordinates Coordinates) *Village {
	return &Village{
		Coordinates: coordinates,
		Population:  1,
		Resources: Resources{
			Wood: 750,
			Clay: 750,
			Iron: 750,
		},
		Production: ResourceProduction{
			Base: Resources{
				Wood: 3600,
				Clay: 3600,
				Iron: 3600,
			},
			Buildings: Resources{
				Wood: 0,
				Clay: 0,
				Iron: 0,
			},
			Multipliers: Resources{
				Wood: 1,
				Clay: 1,
				Iron: 1,
			},
		},
		Capacity: VillageCapacity{
			Warehouse:  10000,
			Population: 200,
		},
	}
}

func (v *Village) Update(deltaT time.Duration) {
	hours := deltaT.Hours()
	if v.Resources.Wood < v.Capacity.Warehouse {
		if n := float32(float64((v.Production.Base.Wood+v.Production.Buildings.Wood)*v.Production.Multipliers.Wood) * hours); n <= v.Capacity.Warehouse {
			v.Resources.Wood += n
		}
	}
	if v.Resources.Clay < v.Capacity.Warehouse {
		if n := float32(float64((v.Production.Base.Clay+v.Production.Buildings.Clay)*v.Production.Multipliers.Clay) * hours); n <= v.Capacity.Warehouse {
			v.Resources.Clay += n
		}
	}
	if v.Resources.Iron < v.Capacity.Warehouse {
		if n := float32(float64((v.Production.Base.Iron+v.Production.Buildings.Iron)*v.Production.Multipliers.Iron) * hours); n <= v.Capacity.Warehouse {
			v.Resources.Iron += n
		}
	}
	fmt.Println(v)
}
