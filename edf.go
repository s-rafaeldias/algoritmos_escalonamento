package main

import "fmt"

// EDF eh um algoritmo de escalonamento baseado em prioridades. A prioridade eh
// definida de acordo com o seu deadline. Aqueles mais proximos de seu deadline,
// tem maior prioridade na fila de execucao. As prioridades e deadline sao atualizados
// de maneira dinamica. No EDF as tarefas sao periodicas e independentes e o seus `deadlines`
// coincidem com seus periodos.
func EDF(processes []*Process) {
	// Cria-se um array de runtime. Como os processos sao passados como
	// ponteiros, para nao perder os valores originais de `deadline` e
	// `period` eh criado esse array de runtime que tera seus valores atualizados
	// durante o tempo e o array original mantem os valores de referencia sem
	// alteracao
	runtimeProcesses := make([]*Process, len(processes))
	for i, _ := range processes {
		// https://stackoverflow.com/questions/21011023/copy-pointer-values-a-b-in-golang
		tmp := &Process{}
		*tmp = *processes[i]
		runtimeProcesses[i] = tmp
	}

	// Variavel de referencia do tempo
	clock := 0

	for {
		// A cada `burstTime` de um processo, seus valores de
		// `deadline` e `period` sao atualizados. Como o algoritmo
		// busca pelos processos mais proximos de seu deadline,
		// eh necessario buscar o index do processo com o menor `deadline`
		// antes de iniciar sua alocacao no CPU
		currentProcessIndex := lowestDeadline(runtimeProcesses)
		// Salva o processo a ser executado na variavel `rp`
		rp := runtimeProcesses[currentProcessIndex]
		// Busca os valores de referencia do processo a ser executado
		// salvado na variavel `p`
		p := processes[currentProcessIndex]

		fmt.Printf("PID processo escolhido: %d\n", p.pid)

		// O clock eh somado ao `burstTime` do processo,
		// simulando sua execucao em CPU
		clock += p.burstTime
		fmt.Printf("Clock: %d\n", clock)
		fmt.Printf("Burst Time: %d\n", p.burstTime)

		fmt.Printf("Deadline anterior: %d\n", rp.deadline)
		// Atualiza dinamicamente o deadline,
		// considerando o numero de periodos anterior a sua
		// ativacao: `d = kP`
		rp.deadline += p.period
		fmt.Printf("Deadline atualizado: %d\n", rp.deadline)

		fmt.Printf("Periodo anterior: %d\n", rp.period)
		// Atualiza dinamicamente o periodo
		// Isso eh equivalente a formula `kP`, ou seja,
		// a quantidade de vezes que o periodo foi ativado
		rp.period += p.period
		fmt.Printf("Periodo atualizado: %d\n", rp.period)

		// Incrementa a quantidade de vezes que o processo
		// escolhido foi chamado
		rp.timeExecuted++
		fmt.Println()

		// Esse eh simplesmente um limite de clock
		// para definir quando finalizar a simulacao do algoritmo
		if clock >= 20 {
			break
		}
	}

	for _, p := range runtimeProcesses {
		fmt.Printf("Processo %d executou %d vezes\n", p.pid, p.timeExecuted)
	}
}

// lowestDeadline retorna o index do processo com o menor `deadline`
func lowestDeadline(processors []*Process) (index int) {
	index = 0
	ld := 10000

	for i, p := range processors {
		if p.deadline < ld {
			ld = p.deadline
			index = i
		}
	}

	return index
}
