# Go Godog

Go é uma linguagem de programação criada pela Google e lançada em código livre. É uma linguagem compilada e focada em produtividade e programação concorrente.

Godog é o framework cucumber(BDD) oficial para Golang, ele mescla especificações e documentação de teste em geral, usando cenários formatados do Gherkin no formato de Dado, Quando, Então.

> Godog não interfere no comportamento padrão do comando go test. Você pode aproveitar ambas as estruturas para testar funcionalmente seu aplicativo, enquanto mantém todo o código fonte relacionado ao teste em arquivos _test.go

Godog age de forma semelhante em comparação com o comando go test , usando o compilador go e a ferramenta de vinculação para produzir o executável de teste.

> Saiba mais em https://github.com/cucumber/godog


## Instalar

Para instalar o framework Godog é necessário que a linguagem Go esteja instalada na máquina, portanto siga a sequência de passos abaixo:

1. Instale o GO

Para instalar a linguagem Go acesse https://golang.org/dl/ e escolha a opção que melhor condiz. 

> As novas versões já configuram o GOPATH

2. Configure as pastas

Para executar a aplicação GO é necessário que o projeto esteja dentro da pasta src(pasta GO/src), essa pasta é criada automáticamente após a instalação da linguagen GO, portanto criaremos nosso projeto dentro dessa pasta, recomenda-se configura-la da seguinte forma:

GO/src/github.com/lucasscaramelo

Sendo o github.com o local aonde está alocado o código seguido do usuário.

3. Instale o GODOG

Na pasta criada, executaremos o seguinte comando:

go get github.com/cucumber/godog/cmd/godog@v0.11.0

Assim, o framework será baixado e seus arquivos estarão alocados na pasta bin(GO/bin).

## Executar

Na pasta raiz do projeto, execute:

> Godog

Assim, as features serão executadas e os resultados serão apresentados.