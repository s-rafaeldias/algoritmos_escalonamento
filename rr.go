package main

import (
	"fmt"
)

// RoundRobin eh um algoritmo desenvolvido especialmente para sistemas
// com timesharing. A fila de `ready` eh tratado como uma fila circular,
// ou seja, cada processo sera executado em CPU por um intervalo de tempo
// pre-determinado, chamado de `quantum`. Quando o periodo de 1 quantum passar,
// o proximo processo sera alocado na CPU e ate 1 quantum de intervalo para
// uso e assim sucessivamente.
func RoundRobin(processes []*Process, quantum int) {
	// Cria-se um array de runtime. Como os processos sao passados como
	// ponteiros, para nao perder os valores originais eh criado esse
	// array de runtime que tera seus valores atualizados ao longo do tempo
	// e o array original mantem os valores de referencia sem alteracao
	runtimeProcesses := make([]*Process, len(processes))
	for i := range processes {
		// https://stackoverflow.com/questions/21011023/copy-pointer-values-a-b-in-golang
		tmp := &Process{}
		*tmp = *processes[i]
		runtimeProcesses[i] = tmp
	}

	// tempo de referencia
	clock := 0
	// tempo gasto na troca de contexto (valor hipotetico)
	overhead := 1
	// Flag com a quantidade de processos finalizados
	finishedJobs := 0

	for {
		for i, p := range runtimeProcesses {
			// Caso o processo ja esteja finalizado, basta
			// ir para o proximo processo da fila
			if p.isFinished {
				continue
			}

			// Simulacao do tempo gasto na troca de contexto
			// Esse overhead sempre ocorre antes de qualquer
			// acao pois eh preciso tirar o processo anterior
			// da CPU para colocar o novo processo a ser executado
			clock += overhead

			// O `burstTime` do processo atual eh maior que o quantum,
			// ou seja, nao ira conseguir finalizar seu processameto no
			// momento.
			if p.burstTime > quantum {
				clock += quantum
				p.burstTime -= quantum
			} else {
				// Nesse caso, o `burstTime` restante eh menor
				// que o quantum, signifcando que o processo ira terminar antes
				// de um quantum
				clock += p.burstTime
				// Tempo total de espera do processo, ou seja,
				// na fila de `ready`, ou na fila de `IO`
				p.waitingTime = clock - processes[i].burstTime
				p.burstTime = 0
				p.isFinished = true
				finishedJobs++
			}

		}
		// Todos os processos foram finalizados
		if finishedJobs == len(processes) {
			break
		}
	}

	// Realiza o calculo de `turnAroundTime`, que eh o tempo total de espera mais
	// o tempo gasto em CPU
	for i, p := range runtimeProcesses {
		p.turnAroundTime = p.waitingTime + processes[i].burstTime
		fmt.Printf("Processo %d\tCMPT %d\tTAT %d\tWT %d\n", p.pid, clock, p.turnAroundTime, p.waitingTime)
	}

	fmt.Printf("Tempo total: %d\n", clock)
	fmt.Printf("Tempo medio TA: %f\n", AvgTurnAroundTime(runtimeProcesses))
	fmt.Printf("Tempo medio WT: %f\n", AvgWaitingTime(runtimeProcesses))
}
