# LInks : 
[https://go.dev/doc/tutorial/create-module](https://go.dev/doc/tutorial/create-module)

---
# Criando modulo greetings
`mkdir greetings`

`cd greetings`

# Iniciando modulo greetings
`go mod init soso.estudos.com/greetings`

---
# Criando modulo hello
`mkdir hello`

`cd hello`

# Iniciando modulo hello
`go mod init soso.estudos.com/hello`

# Alterando módulo que aponta para repositório externo para dependência local:
`go mod edit -replace soso.estudos.com/greetings=../greetings`

# Sincronizando projeto de exemplo `module`
`go mod tidy`

# Executando código para validar que está tudo correto
`go run .`