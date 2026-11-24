# Changelog

Todas as mudanças relevantes desta biblioteca são documentadas neste arquivo.

O formato segue [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/)
e o projeto adere ao [Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [Não publicado]

## [0.1.1] - 2026-11-24

### Alterado

- Requer Go 1.25+ (mínimo elevado para receber as correções de segurança).
- Atualiza `github.com/jackc/pgx/v5` para v5.9.2 e `golang.org/x/text` para
  v0.41.0.
- CI migrado para `golangci-lint-action@v9` (golangci-lint v2) e matriz Go
  1.25/stable.

### Corrigido

- Erros de lint (`errcheck` e `unused`) no cliente HTTP, migrations e testes.
- Vulnerabilidades GO-2026-5004, GO-2026-5970 e GO-2026-6629.

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

[0.1.1]: https://github.com/br-dados/pncp-dados/releases/tag/v0.1.1
[0.1.0]: https://github.com/br-dados/pncp-dados/releases/tag/v0.1.0
