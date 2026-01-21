# Cotação Frete

`cotacao-frete` é um serviço desenvolvido em Go, para consultar cotações de frete via API da Frete Rápido `https://dev.freterapido.com/`.

---

## Requisitos

- Go 1.23+
---

## Ferramentas acopladas ao projeto
- Uber Fx (Injeção de dependências)
- OpenTelemetry (Observabilidade)
- JWT (Dependencia criada porém não utilizada para facilitar a usuabilidade em desenvolvimento)
- GORM (ORM para banco de dados)
- Fiber (framework usado para expor rotas de apis focado em performance extrema e facilidade de uso)
- slog (gerar logs estruturados)
- clientcredentials (Requisições com geração de token e autenticação, não utilizada para facilitar a usuabilidade em desenvolvimento)
---

## Instalação e Execução

### 1. Clone o repositório:
   ```bash

    $ git clone git@github.com:jader-pinheiro/cotacao-frete.git


  ```
### 2. Execute o comando abaixo na raiz do projeto para construir a imagem e subir os containers:
   ```bash

    $ docker-compose up --build

  ```

# Swagger

O Swagger inicia automaticamente com a API na porta 9000 ao executar a aplicação, permitindo a visualização e teste dos endpoints. Acesse a documentação no Swagger UI:

[http://localhost:9000/docs/index.html](http://localhost:9000/docs/index.html)


# Rotas da aplicação

GET http://localhost:9000/v1/quote/metrics <br>
POST http://localhost:9000/v1/quote