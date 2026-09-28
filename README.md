# 💸 type_finance

> Planilha de gastos **offline**, **sem GUI**, rodando direto no terminal. Feita para organizar as minhas finanças.

![Go](https://img.shields.io/badge/Go-1.x-00ADD8?logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow)
![Offline](https://img.shields.io/badge/100%25-offline-brightgreen)

---

## 📌 Sobre

O **type_finance** é uma aplicação de linha de comando escrita em **Go** para registrar e acompanhar gastos pessoais sem depender de internet, conta em nuvem ou interface gráfica. Seus dados ficam na sua máquina.

## ✨ Funcionalidades

- [x] Execução 100% no terminal
- [x] Funciona offline
- [ ] Registrar gastos e receitas
- [ ] Categorizar lançamentos
- [ ] Resumo mensal
- [ ] Exportar para CSV

> Ajuste a lista acima conforme o que já está implementado.

## 🛠️ Tecnologias

- [Go](https://go.dev/)

## 📂 Estrutura do projeto

```
type_finance/
├── ctr/         # controle / lógica de negócio
├── spt/         # suporte / utilitários
├── main.go      # ponto de entrada
├── go.mod       # módulo e dependências
└── README.md
```

## 🚀 Como executar

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado (verifique com `go version`)

### Instalação

```bash
# Clonar o repositório
git clone https://github.com/CarlosGabellini/type_finance.git

# Entrar na pasta
cd type_finance

# Executar
go run main.go
```

### Gerar um binário

```bash
go build -o type_finance
./type_finance
```

## 🖥️ Exemplo de uso

```text
$ ./type_finance

=== type_finance ===
1) Adicionar gasto
2) Listar gastos
3) Resumo do mês
0) Sair

Escolha uma opção: 1
Descrição: Mercado
Valor: 152,90
Categoria: Alimentação

✔ Gasto registrado com sucesso!
```

> Exemplo ilustrativo. Troque pela saída real do programa.

## 🗺️ Roadmap

- [ ] Persistência local dos dados (arquivo)
- [ ] Filtros por período e categoria
- [ ] Relatórios mensais
- [ ] Testes automatizados

## 🤝 Contribuindo

Este é um projeto pessoal, mas sugestões são bem-vindas:

1. Faça um fork
2. Crie uma branch: `git checkout -b minha-feature`
3. Commit: `git commit -m "Adiciona minha feature"`
4. Push: `git push origin minha-feature`
5. Abra um Pull Request

## 📄 Licença

Defina uma licença para o projeto (ex.: [MIT](https://choosealicense.com/licenses/mit/)).

## 👤 Autor

**Carlos Gabellini**
[GitHub](https://github.com/CarlosGabellini)