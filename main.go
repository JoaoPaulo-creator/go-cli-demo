package main

import (
	"goargcli/arguments"
	"goargcli/utils"
	"log"
	"math/rand/V2"
	"sync"

	"github.com/alexflint/go-arg"
)

type StoredArguments struct {
	Arg1 string
	Arg2 string
	Arg3 string
}

type ResponseBody struct {
	ID   int
	Text string
}

type ProposalResponseBody struct {
	ID           int     `json:"id"`
	SimulationId int     `json:"simulationId"`
	Text         string  `json:"text"`
	Amount       float64 `json:"amount"`
}

func createSimulation(s string, headers map[string]string) (int, map[string]string) {
	body := &ResponseBody{
		ID:   rand.IntN(100),
		Text: s,
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

func createSimulationAsync(s string, headers map[string]string, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	body := &ResponseBody{
		ID:   rand.IntN(100),
		Text: s,
	}

	result <- body.ID
}

func createProposalAsync(id int, wg *sync.WaitGroup, result chan<- []byte) {
	defer wg.Done()
	log.Println("Recebendo id da simulação async: ", id)
	body := &ProposalResponseBody{
		ID:           rand.IntN(128937),
		SimulationId: id,
		Text:         "Texto qualquer",
		Amount:       rand.Float64(),
	}

	parsed := utils.BodyParser(body)
	result <- parsed
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

		if args.Async != nil && args.Async.Run {
			responseChannel1 := make(chan int)
			responseChannel2 := make(chan []byte)

			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer close(responseChannel1)
				log.Println("Criando simulação async")
				m := make(map[string]string)
				m["teste"] = "valor"

				createSimulationAsync("Simuation teste", m, &wg, responseChannel1)
				log.Println("Id da simulação async: ", <-responseChannel1)
			}()

      id = <- responseChannel1

			// segunda função
			go func() {
				log.Println("Criando proposta async")
				createProposalAsync(id, &wg, responseChannel2)
			}()

			res = <-responseChannel2
			log.Println("Resposta da proposta async: ", string(res))

      wg.Wait()
		}

		break
	}
}
