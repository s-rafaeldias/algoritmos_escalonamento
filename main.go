// Rafael Dias Silveira
// RA 21908018
// rafael.silveira@sempreceub.com

package main

func main() {
	p1 := newProcess(1, 4, 5)
	p2 := newProcess(2, 6, 4)
	p3 := newProcess(3, 0, 3)
	p4 := newProcess(4, 6, 2)
	p5 := newProcess(5, 5, 4)
	processes := []*Process{p1, p2, p3, p4, p5}

	// FCFS(processes)

	p1 = newProcess(1, 0, 12)
	p2 = newProcess(2, 2, 4)
	p3 = newProcess(3, 3, 6)
	p4 = newProcess(4, 8, 5)
	processes = []*Process{p1, p2, p3, p4}
	// SJF(processes)

	p1 = newProcessEDF(0, 3, 7, 20)
	p2 = newProcessEDF(1, 2, 4, 5)
	p3 = newProcessEDF(2, 2, 8, 10)
	processes = []*Process{p1, p2, p3}
	// EDF(processes)

	p1 = newProcess(1, 0, 12)
	p2 = newProcess(2, 2, 4)
	p3 = newProcess(3, 3, 6)
	processes = []*Process{p1, p2, p3}
	RoundRobin(processes, 1)
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
