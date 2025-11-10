package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entrega struct {
	NomeEntregador string
	GasLevado      int
	GasVendido     int
	AguaLevado     int
}

func main() {
	var entregas []Entrega
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== SISTEMA DE CONTROLE DE GÁS E ÁGUA ===")
		fmt.Println("1 - Registrar entrega")
		fmt.Println("2 - Mostrar todas as entregas")
		fmt.Println("3 - Mostrar resumo do dia")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		input, _ := reader.ReadString('\n')
		opcao := strings.TrimSpace(input)

		switch opcao {
		case "1":
			fmt.Print("Nome do entregador: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Quantidade de gás levada: ")
			gasLevadoStr, _ := reader.ReadString('\n')
			gasLevado, _ := strconv.Atoi(strings.TrimSpace(gasLevadoStr))

			fmt.Print("Quantidade de gás vendida: ")
			gasVendidoStr, _ := reader.ReadString('\n')
			gasVendido, _ := strconv.Atoi(strings.TrimSpace(gasVendidoStr))

			fmt.Print("Quantidade de água levada (vendida toda): ")
			aguaLevadoStr, _ := reader.ReadString('\n')
			aguaLevado, _ := strconv.Atoi(strings.TrimSpace(aguaLevadoStr))

			entrega := Entrega{
				NomeEntregador: nome,
				GasLevado:      gasLevado,
				GasVendido:     gasVendido,
				AguaLevado:     aguaLevado,
			}
			entregas = append(entregas, entrega)
			fmt.Println("Entrega registrada com sucesso!")

		case "2":
			if len(entregas) == 0 {
				fmt.Println("Nenhuma entrega registrada ainda.")
			} else {
				fmt.Println("\n--- TODAS AS ENTREGAS ---")
				for i, e := range entregas {
					fmt.Printf("%d. %s - Gás levado: %d, vendido: %d, restante: %d, Água: %d\n",
						i+1, e.NomeEntregador, e.GasLevado, e.GasVendido, e.GasLevado-e.GasVendido, e.AguaLevado)
				}
			}

		case "3":
			if len(entregas) == 0 {
				fmt.Println("Nenhuma entrega para resumir.")
				continue
			}

			totalGasLevado := 0
			totalGasVendido := 0
			totalAgua := 0

			for _, e := range entregas {
				totalGasLevado += e.GasLevado
				totalGasVendido += e.GasVendido
				totalAgua += e.AguaLevado
			}

			fmt.Println("\n--- RESUMO DO DIA ---")
			fmt.Printf("Total de gás levado: %d\n", totalGasLevado)
			fmt.Printf("Total de gás vendido: %d\n", totalGasVendido)
			fmt.Printf("Gás restante: %d\n", totalGasLevado-totalGasVendido)
			fmt.Printf("Total de água vendida: %d\n", totalAgua)

		case "0":
			fmt.Println("Saindo do sistema...")
			return

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
