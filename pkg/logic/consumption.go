package logic

type ConsumptionHistory []int64

const (
	//Energi connection kind names
	monofasico string = "monofasico"
	bifasico string = "bifasico"
	trifasico string = "trifasico"
	
	//Energi connection kind in kWh
	monofasicaKWH int64 = 400
	bifasicaKWH   int64 = 500
	trifasicaKWH  int64 = 750

	YearOfConsumptions int = 12

	// Co2 percentege per 1000kwh
	co2Average float64 = 0.084
)

func Co2Savings(c ConsumptionHistory) float64 {
	var result float64

	for i, kWh := range c {
		result += (float64(kWh) * co2Average)
		if YearOfConsumptions == i+1 {
			break
		}
	}

	return result
}

func average(c ConsumptionHistory) int64 {
	var total int64 
	var i int
	var value int64

	for i, value = range c {
		total += value 
		if i+1 == YearOfConsumptions {
			break
		}
	} 
	return total / int64(i+1)
}

func ConectionIsValid(cunsumptionKind string, c ConsumptionHistory) bool {
	ammount := average(c);

	switch {
	case RegexpString(cunsumptionKind, monofasico):
		return ammount >= monofasicaKWH
	case RegexpString(cunsumptionKind, bifasico):
		return ammount >= bifasicaKWH
	case RegexpString(cunsumptionKind, trifasico):
		return ammount >= trifasicaKWH
	default:
		return false
	}

}
