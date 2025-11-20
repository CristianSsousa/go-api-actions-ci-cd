# Instruções para Deploy

## Configurar o repositório Git

1. Inicialize o repositório Git (se ainda não foi feito):
```bash
git init
```

2. Adicione o remote do GitHub:
```bash
git remote add origin https://github.com/CristianSsousa/go-api-actions-ci-cd.git
```

3. Adicione todos os arquivos:
```bash
git add .
```

4. Faça o commit inicial:
```bash
git commit -m "Initial commit: API Go com Clean Architecture e Chi Router"
```

5. Faça o push para o repositório:
```bash
git branch -M main
git push -u origin main
```

## Executar a API localmente

```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## Testar os endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### Listar usuários
```bash
curl http://localhost:8080/api/users
```

### Listar produtos
```bash
curl http://localhost:8080/api/products
```

### Criar um usuário
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Teste", "email": "teste@example.com", "role": "user"}'
```

