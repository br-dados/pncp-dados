# Changelog

Todas as mudanças relevantes desta biblioteca são documentadas neste arquivo.

O formato segue [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/)
e o projeto adere ao [Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [Não publicado]

## [0.1.0] - 2026-11-24

### Adicionado

- Primeira versão pública da biblioteca como módulo
  `github.com/br-dados/pncp-dados`.
- Harvest do índice offline de contratos e contratações da API de Consulta v1
  do PNCP.
- Backends de persistência Postgres (padrão) e XLSX, além de suporte a backend
  próprio via `storage.Storage`.
- Buscas indexadas por fornecedor e filtros de contratos/contratações (Postgres).
- `pncpdados.ErrEmptyScope` e constantes `pncpdados.StoragePostgres` /
  `pncpdados.StorageXLSX` na API pública.
- Comandos CLI `cmd/harvest` e `cmd/migrate`.

[0.1.0]: https://github.com/br-dados/pncp-dados/releases/tag/v0.1.0
