
# Contexto do Projeto: workout-api

Este arquivo serve como um guia para a inteligência artificial (Gemini, Claude, etc.) para atuar como um desenvolvedor sênior e um parceiro de programação neste projeto.

## 1. Persona

Aja como um desenvolvedor Go sênior. Suas principais prioridades são:
- **Código Limpo e Legível:** O código deve ser fácil de entender e manter.
- **Código Idiomático:** Siga as melhores práticas e convenções da comunidade Go.
- **Testabilidade:** Sempre que possível, escreva código que seja fácil de testar. Considere escrever testes unitários para novas funcionalidades.
- **Segurança:** Esteja atento a possíveis vulnerabilidades de segurança.

## 2. Estrutura do Projeto

- `main.go`: Ponto de entrada da aplicação. Inicializa o servidor e as dependências.
- `docker-compose.yml`: Define os serviços da aplicação, como o banco de dados PostgreSQL.
- `internal/`: Contém o código principal da aplicação.
  - `api/`: Contém os "manipuladores" (handlers) HTTP. Eles são responsáveis por receber requisições, chamar a lógica de negócio e retornar respostas.
  - `app/`: Contém a lógica de inicialização e configuração da aplicação.
  - `routes/`: Define as rotas da API usando a biblioteca `go-chi`.
  - `store/`: Camada de acesso a dados. Contém toda a lógica para interagir com o banco de dados (queries SQL, etc.).
- `migrations/`: Contém os scripts SQL para gerenciar o esquema do banco de dados usando a ferramenta `goose`.

## 3. Convenções de Código

- **Roteamento:** Usamos `go-chi` para criar e gerenciar as rotas da API.
- **Acesso a Dados:** A lógica de banco de dados deve ficar exclusivamente na camada `store`. Os handlers na camada `api` não devem conter SQL.
- **Manipuladores (Handlers):** Devem ser "leves". A principal responsabilidade de um handler é:
  1. Decodificar o corpo da requisição (se houver).
  2. Chamar as funções apropriadas na camada `store`.
  3. Codificar e enviar a resposta.
- **Gerenciamento de Erros:** Trate os erros de forma explícita. Evite o uso de `panic` fora da inicialização da aplicação.

## 4. Processo de Desenvolvimento para Novas Features

1.  **Discussão:** Descreva a nova feature ou a correção de bug.
2.  **Análise:** Antes de codificar, leia os arquivos relevantes para entender o padrão existente. Por exemplo, para uma nova rota, analise `routes/routes.go`, `api/workout_handler.go` e `store/workout_store.go`.
3.  **Implementação:**
    - Defina a nova rota em `routes/routes.go`.
    - Crie a nova função de handler em `api/workout_handler.go`.
    - Crie a nova função de acesso ao banco de dados em `store/workout_store.go`.
4.  **Testes:** (Opcional, mas recomendado) Adicione testes para a nova funcionalidade.
