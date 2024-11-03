# API de Gerenciamento de Cursos

## Descrição
Esta é uma API RESTful para gerenciar usuários, cursos e progresso em um sistema de aprendizado. A API permite que os usuários se registrem, façam login e gerenciem seus cursos e progresso.

## Funcionalidades
- **Autenticação**:
  - Registro de usuários com nome, email e senha (armazenamento seguro da senha).
  - Login de usuários com geração de token JWT.
  
- **Gerenciamento de Usuários**:
  - Obter lista de usuários (protegido).
  - Atualizar informações do usuário.
  - Deletar conta do usuário.

- **Gerenciamento de Cursos**:
  - Criar novos cursos (protegido).
  - Obter lista de cursos (protegido).
  - Atualizar informações de cursos (protegido).
  - Deletar cursos (protegido).

- **Gerenciamento de Progresso**:
  - Criar registros de progresso (protegido).
  - Obter progresso do usuário (protegido).
  - Atualizar registros de progresso (protegido).
  - Deletar registros de progresso (protegido).

## Tecnologias Utilizadas
- Go (Golang)
- GORM (ORM para Go)
- JWT (JSON Web Tokens) para autenticação
- Gorilla Mux para roteamento
- Bcrypt para hashing de senhas

## Como Configurar

### Pré-requisitos
- Go (versão 1.16 ou superior)
- Banco de dados (PostgreSQL, MySQL, etc.)

### Instalação
1. Clone o repositório:
   ```bash
   git clone https://github.com/Trickqz/course-management-golang
   cd course-management-golang
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure o banco de dados no arquivo `database/database.go`.

4. Crie um arquivo `.env` na raiz do projeto e adicione suas variáveis de ambiente, como a chave secreta para o JWT.

### Executar a Aplicação
Para iniciar o servidor, execute:

```
go run main.go
```

### Conclusão

Esse README fornece uma visão geral clara do seu projeto, como configurá-lo e como usá-lo. Se precisar de mais ajustes ou informações adicionais, é só avisar!

### Rotas

todas as rotas estão no arquivo routes.go ou caso queira testar as rotas no postman, segue o link do postman: [Link do Postman](https://workhub-0304.postman.co/workspace/f425429d-ec68-4fcc-a249-bcd28f0ef221)