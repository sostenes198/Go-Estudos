package Enderecos

import "strings"

// TipoDeEndereco Verifica se o endereço tem um tipo válido e retorna
func TipoDeEndereco(endereco string) string{
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]
	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco{
			return strings.Title(primeiraPalavraDoEndereco)
		}
	}
	return "Tipo Inválido"
}
