package pncpdados

import "github.com/br-dados/pncp-dados/types"

// Backends de saida aceitos por Options.Storage.
const (
	// StoragePostgres persiste contratos e contratacoes em Postgres (default).
	StoragePostgres = types.StoragePostgres

	// StorageXLSX persiste em planilha XLSX (sem buscas indexadas).
	StorageXLSX = types.StorageXLSX
)
