package main

// Definicao de processo
type Process struct {
	// Identificador unico de um processo
	pid int
	// Tempo em que o processo foi colocado na fila de `ready`
	arrivalTime int
	// Tempo de execucao em CPU do processo
	burstTime int
	// Tempo de execucao total de um processo, desde o momento que
	// entra na fila de `ready` ate a conclusao de sua execucao total
	turnAroundTime int
	// Tempo total de espera de um processo na fila de `ready`
	waitingTime int
	// Indica se um processo ja finalizou sua execucao
	isFinished bool
	// Tempo de execucao do processo em CPU
	timeExecuted int
	//
	deadline int
	period   int
}

func newProcess(pid, arrivalTime, burstTime int) *Process {
	return &Process{
		pid:            pid,
		arrivalTime:    arrivalTime,
		burstTime:      burstTime,
		turnAroundTime: 0,
		waitingTime:    0,
		isFinished:     false,
		timeExecuted:   0,
	}
}

func newProcessEDF(pid, burstTime, deadline, period int) *Process {
	return &Process{
		pid:            pid,
		arrivalTime:    0,
		burstTime:      burstTime,
		turnAroundTime: 0,
		waitingTime:    0,
		isFinished:     false,
		timeExecuted:   0,
		deadline:       deadline,
		period:         period,
	}
}
