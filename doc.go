// Package pncpdados e a API publica da biblioteca pncp-dados: harvest do
// indice offline de contratos assinados da API de Consulta v1 do PNCP, com
// saida configuravel (Postgres por padrao, ou planilha XLSX).
//
// Uso minimo (Postgres, configuracao 100% por codigo):
//
//	cliente, err := pncpdados.New(ctx, pncpdados.Options{
//		Postgres: &pncpdados.PostgresConfig{
//			Host: "localhost", Port: "5432", User: "app", Password: "s3cret", Database: "pncp_dados",
//		},
//	})
//	resumo, err := cliente.Harvest(ctx, types.Scope{States: []string{"GO"}})
//
// O pacote nao le variaveis de ambiente: toda a configuracao e explicita em
// Options. Os comandos CLI em cmd/ sao os responsaveis por ler o ambiente.
package pncpdados
