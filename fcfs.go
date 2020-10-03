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
	totalTime := 0

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
		if p.arrivalTime > totalTime {
			totalTime += p.arrivalTime - totalTime
		}

		// Tempo total eh somado ao burstTime (tempo de execucao em CPU)
		totalTime += p.burstTime
		// Turn Around Time eh o tempo de conclusao do processo (totalTime)
		// menos em qual instante o processo chegou
		p.turnAroundTime = totalTime - p.arrivalTime
		// Tempo efetivo de espera, que eh a diferenca entre o
		p.waitingTime = p.turnAroundTime - p.burstTime

		completionTime := p.turnAroundTime + p.arrivalTime

		fmt.Printf("Processo %d\tCMPT %d\tTAT %d\tWT %d\n", p.pid, completionTime, p.turnAroundTime, p.waitingTime)
	}

	fmt.Printf("Tempo total: %d\n", totalTime)
	fmt.Printf("Tempo medio TA: %f\n", AvgTurnAroundTime(processes))
	fmt.Printf("Tempo medio WT: %f\n", AvgWaitingTime(processes))
}
