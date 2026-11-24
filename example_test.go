package pncpdados_test

import (
	"context"
	"fmt"

	pncpdados "github.com/br-dados/pncp-dados"
	"github.com/br-dados/pncp-dados/types"
)

// ExampleNew mostra a construcao do cliente e um harvest minimo para a saida
// XLSX. O exemplo nao tem Output e, por isso, nao e executado pelo go test.
func ExampleNew() {
	ctx := context.Background()

	cliente, err := pncpdados.New(ctx, pncpdados.Options{
		Storage:  pncpdados.StorageXLSX,
		XLSXPath: "contratacoes-go.xlsx",
	})
	if err != nil {
		fmt.Println("erro ao criar cliente:", err)
		return
	}
	defer cliente.Close()

	resumo, err := cliente.Harvest(ctx, types.Scope{States: []string{"GO"}})
	if err != nil {
		fmt.Println("erro no harvest:", err)
		return
	}

	fmt.Printf("contratacoes=%d\n", resumo.Inserted.Procurements)
}
