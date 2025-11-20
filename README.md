# Go API - Clean Architecture com Chi Router

API REST em Go seguindo os princípios de Clean Architecture, utilizando Chi como router HTTP e respostas pré-prontas (mock data).

## 🏗️ Arquitetura

O projeto segue a Clean Architecture, organizando o código em camadas:

```
.
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/            # Camada de apresentação (HTTP handlers)
│   ├── services/            # Camada de casos de uso (lógica de negócio)
│   ├── repositories/        # Camada de dados (acesso a dados)
│   ├── models/              # Entidades e DTOs
│   └── middleware/          # Middlewares HTTP
└── go.mod                   # Dependências do projeto
```

## 🚀 Tecnologias

- **Go 1.21+**
- **Chi Router** - Router HTTP leve e rápido
- **Chi CORS** - Middleware para CORS
- **Chi Render** - Middleware para renderização JSON

## 📦 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/CristianSsousa/go-api-actions-ci-cd.git
cd go-api-actions-ci-cd
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute a aplicação:
```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## 📚 Endpoints

### Health Check
- `GET /health` - Verifica o status da API
- `GET /` - Health check alternativo

### Usuários
- `GET /api/users` - Lista todos os usuários
- `GET /api/users/{id}` - Busca usuário por ID
- `POST /api/users` - Cria um novo usuário
- `PUT /api/users/{id}` - Atualiza um usuário
- `DELETE /api/users/{id}` - Remove um usuário

### Produtos
- `GET /api/products` - Lista todos os produtos
- `GET /api/products/{id}` - Busca produto por ID
- `GET /api/products/category/{category}` - Busca produtos por categoria
- `POST /api/products` - Cria um novo produto
- `PUT /api/products/{id}` - Atualiza um produto
- `DELETE /api/products/{id}` - Remove um produto

## 📝 Exemplos de Uso

### Criar um usuário
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Usuário",
    "email": "novo@example.com",
    "role": "user"
  }'
```

### Listar todos os produtos
```bash
curl http://localhost:8080/api/products
```

### Buscar produto por categoria
```bash
curl http://localhost:8080/api/products/category/Periféricos
```

## 🧪 Testes

Para executar os testes:
```bash
go test ./...
```

## 🔧 Variáveis de Ambiente

- `PORT` - Porta onde o servidor irá rodar (padrão: 8080)

## 📋 Estrutura de Resposta

Todas as respostas seguem o padrão:

```json
{
  "success": true,
  "message": "Mensagem opcional",
  "data": { ... },
  "error": "Mensagem de erro (se houver)"
}
```

## 🏛️ Clean Architecture

### Camadas

1. **Handlers** - Recebem requisições HTTP e retornam respostas
2. **Services** - Contêm a lógica de negócio e validações
3. **Repositories** - Gerenciam o acesso aos dados (atualmente em memória)

### Princípios

- **Separação de responsabilidades**: Cada camada tem uma responsabilidade específica
- **Inversão de dependências**: Camadas superiores dependem de abstrações
- **Independência de frameworks**: A lógica de negócio não depende de frameworks HTTP
- **Testabilidade**: Cada camada pode ser testada independentemente

## 🚢 CI/CD

O projeto inclui GitHub Actions para CI/CD automatizado com deploy no Google Cloud Run.

### Workflows Disponíveis

- **`deploy-gcp-cloud-run.yml`** - Deploy usando Google Container Registry (GCR)
- **`deploy-gcp-cloud-run-artifact-registry.yml`** - Deploy usando Artifact Registry (recomendado)

### Configuração

Para configurar o deploy automático, siga as instruções detalhadas em [DEPLOY.md](./DEPLOY.md).

**Resumo rápido:**
1. Configure um projeto no GCP
2. Crie uma Service Account com as permissões necessárias
3. Adicione os secrets `GCP_PROJECT_ID` e `GCP_SA_KEY` no GitHub
4. Faça push para a branch `main` - o deploy será automático!

Veja `.github/workflows/` para mais detalhes sobre os workflows.

## 📄 Licença

Este projeto está sob a licença MIT.

