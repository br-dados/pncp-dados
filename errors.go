package pncpdados

import "github.com/br-dados/pncp-dados/internal/harvest"

// ErrEmptyScope e retornado por Harvest quando o escopo nao define nenhuma
// dimensao de recorte (sem UFs, sem municipios e sem CNPJs de orgao).
//
// Use errors.Is para compara-lo:
//
//	_, err := cliente.Harvest(ctx, types.Scope{})
//	if errors.Is(err, pncpdados.ErrEmptyScope) {
//		// escopo invalido
//	}
var ErrEmptyScope = harvest.ErrEmptyScope
