// Rafael Dias Silveira
// RA 21908018
// rafael.silveira@sempreceub.com
package main

import (
	"fmt"
	"sort"
)

// SJF eh o algoritmo Shortest Job First, onde o processo com o menor
// tempo de execucao restante eh alocado no CPU. Nessa versao, a implementacao
// usa uma abordagem preemptiva, ou seja, o escalonador consegue desalocar um
// processo da CPU para colocar outro com menor `burstTime` restante.
func SJF(processes []*Process) {
	// Ordena os processos em ordem de chegada
	sort.Slice(processes, func(a, b int) bool {
		if processes[a].arrivalTime == processes[b].arrivalTime {
			return processes[a].pid < processes[b].pid
		}
		return processes[a].arrivalTime < processes[b].arrivalTime
	})

	// Tempo de referencia
	clock := 0

	// Inicio da simulacao do algoritmo.
	for {
		// Define qual o processo sera executado em CPU
		currentProcessIndex, err := chooseProcess(clock, processes)
		// Caso retorne um erro, significa que nao existem no momento
		// processo na fila de `ready`
		if err != nil {
			// Encerra a simulacao caso todos os processos ja tenham sido
			// finalizados
			if allProcessed(processes) {
				break
			}
			// Incrementa o tempo e continua a simulacao
			clock++
			continue
		}

		// Realiza um loop para analisar todos os processos e identificar
		// qual o processo da vez e quais estao na fila de `ready`, com seu
		// `waitingTime` ja contando
		for i, p := range processes {
			// Ignora os processos ja finalizados
			if p.isFinished {
				continue
			}

			// Realiza a execucao em CPU do processo da vez
			if i == currentProcessIndex {
				fmt.Printf("Clock: %d\tProcess: %d\n", clock, p.pid)
				p.burstTime--
				clock++

				// Caso o `burstTime` apos o processamento chegue em zero,
				// o processo foi encerrado e nao precisa mais ser considerado
				if p.burstTime == 0 {
					p.isFinished = true
				}
			} else {
				// Caso o processo ja esteja na fila de `ready` e nao esteja
				// sendo executado, o seu `waitingTime` eh incrementado.
				// Se o processo acabou de chegar na fila de `ready`, ou seja,
				// seu `arrivalTime` == `clock`, seu `waitingTime` so vai comecar
				// a contar a partir da proxima unidade de tempo
				if p.arrivalTime < clock {
					p.waitingTime++
				}
			}
		}

	}

	fmt.Printf("Tempo total: %d\n", clock)
	fmt.Printf("Tempo medio TA: %f\n", AvgTurnAroundTime(processes))
	fmt.Printf("Tempo medio WT: %f\n", AvgWaitingTime(processes))
}

// chooseProcess escolhe o processo com o menor `burstTime` para
// um determinado tempo, ou seja, somente os processos que ja estao
// na fila de `ready` poderam ser escolhidos
func chooseProcess(clock int, processes []*Process) (int, error) {
	burstTime := 1000
	nextProcess := -1
	for i, p := range processes {
		// O processo para ser escolhido precisa ainda esta sendo executado,
		// `isFinished` == false, seu `arrivalTime` ser menor ou igual ao `clock` e
		// possuir o menor `burstTime` restante.
		if !p.isFinished && p.arrivalTime <= clock && p.burstTime < burstTime {
			nextProcess = i
			burstTime = p.burstTime
		}
	}

	if nextProcess == -1 {
		return -1, fmt.Errorf("Processo na fila de `ready`")
	}
	return nextProcess, nil
}

// allProcessed verifica que se todos os processos ja foram finalizados
func allProcessed(processes []*Process) bool {
	for _, p := range processes {
		if p.isFinished == false {
			return false
		}
	}

	return true
}
