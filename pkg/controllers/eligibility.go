package controllers

import (
	"github.com/Fr0zenBoy/eligibility/pkg/logic"
)

const (
	messageConsumeClass string = "Classe de consumo não aceita"
	messageTariffKind string = "Modalidade tarifária não aceita"
	messageConsumeConnection string = "Consumo muito baixo para tipo de conexão"
)

type Eligibible struct {
	NumeroDoDocumento string `json:"numeroDoDocumento" binding:"required,numeric,gt=11,lt=14"`
	TipoDeConexao string `json:"tipoDeConexao" binding:"required,alpha,endswith=fasico"`
	ClaseDeConsumo string `json:"claseDeConsumo" binding:"required,alpha,lt=15"`
	ModalidadeTarifaria string `json:"modalidadeTarifaria" binding:"required,alpha,lt=15"`
	HistoricoDeConsumo []int64 `json:"historicoDeConsumo" binding:"required"` //dive,gt=3,lt=12
}

type EligibleOutput struct {
	Elegivel bool `json:"elegivel" binding:"required"`
	EconomiaAnaulDeCo2 float64 `json:"economiaAnaulDeCO2"`
	RazoesInelegibilidade []string `json:"razoesInelegibilidade"`
}

func (e Eligibible) validConnection() bool {
	return	logic.ConectionIsValid(e.TipoDeConexao, e.HistoricoDeConsumo)
}

func (e Eligibible) custumerClass() bool {
	return logic.CustomerClasses().Allow(e.ClaseDeConsumo)
}

func (e Eligibible) tariffModality() bool {
	return logic.TariffModality().Allow(e.ModalidadeTarifaria)
}

func (e Eligibible) Co2Saving() float64 {
	return logic.Co2Savings(e.HistoricoDeConsumo)
}

func (e Eligibible) DenyReasons() []string {
	reasons := map[string]bool{
		messageConsumeClass: e.custumerClass(),
		messageTariffKind: e.tariffModality(),
		messageConsumeConnection: e.validConnection(),
	}

	deny := []string{}

	for k,v := range reasons {
		if !v {
			deny = append(deny, k)
		}
	}

	return deny
}

func (e Eligibible) Resp(r EligibleOutput) EligibleOutput {
	denyReasions := e.DenyReasons()

	if noErrors := len(denyReasions); noErrors == 0 {
		r.Elegivel = true
		r.EconomiaAnaulDeCo2 = e.Co2Saving()
		return r
	} else {
		r.Elegivel = false
		r.RazoesInelegibilidade = denyReasions
		return r
	}
}
