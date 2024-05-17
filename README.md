# Eulabs

API para controle de produtos.

## Como rodar?
### Dependências
- Go 1.22
- Docker
- Docker Compose

### Rodando o projeto
1. Clone o repositório
2. Entre na pasta do projeto
3. Execute o comando `make run`
4. A API estará disponível em `http://localhost:8001`

Um arquivo `requests.http` está disponível na raiz do projeto com exemplos de requisições para testar a API.

## Estrutura do projeto
```
cmd -> Arquivos de inicialização da aplicação
deploy -> Arquivos de deploy
pkg -> Pacotes da aplicação
    - api -> Pacote com a implementação da API
    - entity -> Pacote com as entidades do domínio
    - repository -> Pacote com a implementação dos repositórios
    - service -> Pacote com a implementação dos serviços
test -> Arquivos para teste
```

## Decisões de projeto
- Utilização do framework Echo para criação e controle das rotas
- Utilização do Docker Compose para facilitar a execução do projeto
- Utilização do banco de dados MySQL
- Utilização de poucas dependências de terceiros

## Decisões de lógica
- Criação de uma tabela exclusiva para preço dos produtos, a fim de suportar diversas moedas
- Atualização de produtos funciona como um "sync", onde é necessário enviar todos os dados do produto

## Faltou tempo :(
- Implementar testes (unitários e de integração)
- Implementar documentação da API (Swagger)
- Implementar busca elástica de produtos
- Implementar camada de cache