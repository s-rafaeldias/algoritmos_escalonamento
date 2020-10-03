// Rafael Dias Silveira
// RA 21908018
// rafael.silveira@sempreceub.com

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
}

func main() {
	p1 := newProcess(1, 4, 5)
	p2 := newProcess(2, 6, 4)
	p3 := newProcess(3, 0, 3)
	p4 := newProcess(4, 6, 2)
	p5 := newProcess(5, 5, 4)
	processes := []*Process{p1, p2, p3, p4, p5}

	// p1 := newProcess(1, 0, 3)
	// p2 := newProcess(1, 1, 2)
	// p3 := newProcess(1, 2, 1)
	// p4 := newProcess(1, 3, 4)
	// p5 := newProcess(1, 4, 5)
	// p6 := newProcess(1, 5, 2)
	// processes := []*Process{p1, p2, p3, p4, p5, p6}

	FCFS(processes)
}

func newProcess(pid, arrivalTime, burstTime int) *Process {
	return &Process{
		pid:            pid,
		arrivalTime:    arrivalTime,
		burstTime:      burstTime,
		turnAroundTime: 0,
		waitingTime:    0,
	}
}

func AvgTurnAroundTime(processes []*Process) float64 {
	value := 0.0
	for _, p := range processes {
		value += float64(p.turnAroundTime)
	}

	return value / float64(len(processes))
}

func AvgWaitingTime(processes []*Process) float64 {
	value := 0.0
	for _, p := range processes {
		value += float64(p.waitingTime)
	}

	return value / float64(len(processes))
}
