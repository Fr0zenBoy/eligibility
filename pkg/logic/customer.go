package logic

type customerClasses struct {

	Possible []string
	//Comercial, Residencial, Industrial, Poder Público, e Rural.
	Eligible []string
	//Comercial, Residencial e Industrial
}

type tariffModality struct {
	Possible []string
	//branca, azul, Verde and Convencional

	Eligible []string
	//Convencional, Branca
}

type allowed interface {
	Allow (value string) bool
}

func CustomerClasses() customerClasses {
	cc := customerClasses{
		Possible: []string{"comercial", "residencial", "industrial", "poder publico", "rural"},
		Eligible: []string{"comercial", "residencial", "industrial"},
	}
	return cc
}

func TariffModality() tariffModality {
	tm := tariffModality{
		Possible: []string{"branca", "azul", "verde", "convencional"},
		Eligible: []string{"convencional", "branca"},
	}
	return tm
}

func exist(s string, values []string) bool {
	for _, classes := range values {
		if RegexpString(s, classes){
			return true
		}
	}
	return false
}

func (c customerClasses) Allow(kind string) bool {
	return exist(kind, c.Eligible)
}

func (t tariffModality) Allow(kind string) bool {
	return exist(kind, t.Eligible)
}
