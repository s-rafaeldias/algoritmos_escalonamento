package main

import (
	"fmt"
	"sort"
)

// FCFS implementa o escalonamento usando a regra de First Come First Served.
// Eh o algoritmo mais simples de escalonamento, onde os processos sao alocados
// de acordo com sua ordem de chegada na fila de `ready`
func FCFS(processes []*Process) {
	// Tempo total de execucao dos processos
	completionTime := 0

	// Ordena os processos de acordo com arrivalTime
	sort.Slice(processes, func(a, b int) bool {
		// Se o `arrivalTime` for igual entre dois processos,
		// o processo com o menor pid tem prioridade
		if processes[a].arrivalTime == processes[b].arrivalTime {
			return processes[a].pid < processes[b].pid
		}
		return processes[a].arrivalTime < processes[b].arrivalTime
	})

	for _, p := range processes {
		// Se o arrivalTime for maior que o tempo total
		// de execucao ate o momento, significa que a CPU
		// ficou IDLE. O tempo que a CPU fica IDLE eh somado
		// ao tempo total de execucao
		if p.arrivalTime > completionTime {
			completionTime += p.arrivalTime - completionTime
		}

		// Tempo total eh somado ao burstTime (tempo de execucao em CPU)
		completionTime += p.burstTime
		// Turn Around Time eh a diferenca entre o tempo de conclusao do processo (completionTime)
		// e o instante o processo chegou, ou seja, eh o tempo decorrido entre o momento
		// que o processo entra na fila de `ready` ate o momento que finaliza sua completa
		// execucao
		p.turnAroundTime = completionTime - p.arrivalTime
		// Tempo efetivo de espera do processo desde o momento que entrou na fila de
		// `ready` ate o momento que o processo sera executado na CPU
		p.waitingTime = p.turnAroundTime - p.burstTime

		fmt.Printf("Processo %d\tCMPT %d\tTAT %d\tWT %d\n", p.pid, completionTime, p.turnAroundTime, p.waitingTime)
	}

	fmt.Printf("Tempo total: %d\n", completionTime)
	fmt.Printf("Tempo medio TA: %f\n", AvgTurnAroundTime(processes))
	fmt.Printf("Tempo medio WT: %f\n", AvgWaitingTime(processes))
}
