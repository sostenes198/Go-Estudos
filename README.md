Comandos cmd GO

**go mod init NOME_MODULO** (cria um modulo)

**go run** (executa o projeto)

**go build** (builda o projeto e gera um .exe)

**go install** (builda o projeto e gera um .exe na pasta raiz do go)

**go get NOME_PACOTE** (Baixa pacotes de terceiros)

**go clean -i NOME_PACOTE** (Remove pacote)

**go mod tidy** (Remove todas dependências que não estão sendo usadas no modulo)

**go test** (Executa os testes do projeto)

**go test ./...** (Executar todos os testes do projeto)

**go test -v** (Exibe mais detalhes da execução)

**go test --cover** (Realiza a cobertura de código)

**go test --coverprofile Nome_Do_Arquivo** (Realiza a cobertura de código e exibe em um arquivo)

**go tool cover --func=Nome_Do_Arquivo** (Exibe no terminal detalhes do arquivo)

**go tool cover --html=Nome_Do_Arquivo** (Exibe um html mostrando o código do que esta coberto e oque não esta coberto)
