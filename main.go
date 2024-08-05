package main

import (
	"goargcli/arguments"
	"goargcli/utils"
	"log"
	"math/rand/V2"

	"github.com/alexflint/go-arg"
)

type StoredArguments struct {
	Arg1 string
	Arg2 string
	Arg3 string
}

type ResponseBody struct {
	ID   int
	text string
}

type ProposalResponseBody struct {
	ID           int     `json:"id"`
	SimulationId int     `json:"simulationId"`
	Text         string  `json:"text"`
	Amount       float64 `json:"amount"`
}

func createSimulation(s string, headers map[string]string) (int, map[string]string) {
	body := &ResponseBody{
		rand.IntN(100),
		s,
	}

	return body.ID, headers
}

func createProposal(id int) []byte {
	log.Println("Recebendo id da simulação: ", id)
	body := &ProposalResponseBody{
		ID:           rand.IntN(128937),
		SimulationId: id,
		Text:         "Texto qualquer",
		Amount:       rand.Float64(),
	}

	parsed := utils.BodyParser(body)

	return parsed
}

var id int
var res []byte

func main() {
	var args arguments.ArgList
	arg.MustParse(&args)

	log.Println(args)

	for {
		if &args.Simulation != nil {
			log.Println("Criando simulação")
			m := make(map[string]string)
			m["teste"] = "valor"
			responseId, headers := createSimulation("Simuation teste", m)
			id = responseId

			log.Printf("simulationId: %d, headers: %v\n", id, headers)
		}

		if &id == nil || id == 0 {
			panic("simulationId não gerado")
		}

		if &args.Proposal != nil {
			proposalResponse := createProposal(id)
			res = proposalResponse
		}

		if args.ParseResonse {
      var t ProposalResponseBody
      parsedResponse := utils.ResponseBodyParser(res, t)
      log.Println("Resposta da proposta: ", string(parsedResponse))
		}


    break
	}
}
