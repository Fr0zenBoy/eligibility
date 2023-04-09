package controllers

import (
	"github.com/Fr0zenBoy/eligibility/pkg/logic"
)

const (
	messageConsumerClass string = "Classe de consumo não aceita"
	messageTariffKind string = "Modalidade tarifária não aceita"
	messageConsumerConnection string = "Consumo muito baixo para tipo de conexão"
)

type Eligibible struct {
	NumeroDoDocumento string `json:"numeroDoDocumento" binding:"required,numeric,min=11,max=14"`
	TipoDeConexao string `json:"tipoDeConexao" binding:"required,alpha,endswith=fasico"`
	ClasseDeConsumo string `json:"classeDeConsumo" binding:"required,ascii,lt=15"`
	ModalidadeTarifaria string `json:"modalidadeTarifaria" binding:"required,alpha,lt=15"`
	HistoricoDeConsumo []int64 `json:"historicoDeConsumo" binding:"required,min=3,max=12,dive,numeric"`
}

type eligibleOutput struct {
	Elegivel bool `json:"elegivel" binding:"required"`
	EconomiaAnaulDeCo2 float64 `json:"economiaAnaulDeCO2"`
}

type noteligibleOutput struct {
	Elegivel bool `json:"elegivel" binding:"required"`
	RazoesInelegibilidade []string `json:"razoesInelegibilidade"`
}

func (e Eligibible) isValidConnection() bool {
	return	logic.ConectionIsValid(e.TipoDeConexao, e.HistoricoDeConsumo)
}

func (e Eligibible) isValidcustumerClass() bool {
	return logic.CustomerClasses().Allow(e.ClasseDeConsumo)
}

func (e Eligibible) isValidTariffModality() bool {
	return logic.TariffModality().Allow(e.ModalidadeTarifaria)
}

func (e Eligibible) Co2Saving() float64 {
	return logic.Co2Savings(e.HistoricoDeConsumo)
}

func (e Eligibible) DenyReasons() []string {
	reasons := map[string]bool{
		messageConsumerClass: e.isValidcustumerClass(),
		messageTariffKind: e.isValidTariffModality(),
		messageConsumerConnection: e.isValidConnection(),
	}

	deny := []string{}

	for k,v := range reasons {
		if !v {
			deny = append(deny, k)
		}
	}

	return deny
}

func (e Eligibible) Resp() interface{} {
	output := eligibleOutput{}
	errOutput := noteligibleOutput{}
	if errors := e.DenyReasons(); len(errors) == 0 {
		output.Elegivel = true
		output.EconomiaAnaulDeCo2 = e.Co2Saving()
		return output
	} else {
		errOutput.RazoesInelegibilidade = errors
		return errOutput
	}
}
