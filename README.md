# teste-tecnico-ilumeo

# ADR: Escolha da Linguagem Go, Processamento do Arquivo SQL e Otimização com Índices

## Contexto

Recebi um arquivo `.sql` contendo milhões de registros para popular o banco de dados. A aplicação backend precisava ser performática, escalável e capaz de processar esses dados volumosos de forma eficiente. A decisão da stack e da arquitetura impactaria diretamente no desempenho e na manutenção do projeto.

## Decisão

### Escolha da Linguagem Go

- Optei por **Go** em vez de Node.js devido à necessidade de processar um grande volume de dados (milhões de registros) de forma eficiente, com melhor desempenho e menor consumo de memória.
- Go é uma linguagem compilada, gerando binários nativos otimizados para alta performance.
- Possui um modelo de concorrência leve baseado em goroutines, o que facilita o processamento paralelo e eficiente.
- Node.js, por ser baseado em JavaScript interpretado e modelo single-threaded, apresentaria limitações em processamento intensivo e maior consumo de recursos.
- A facilidade de empacotamento do Go em um binário único simplifica o deploy em ambientes Docker e facilita a portabilidade da aplicação.

### Processamento do Arquivo SQL usando comandos Linux

Para preparar o arquivo `.sql` para importação e evitar problemas com dados mal formatados, executei uma sequência de comandos Linux para limpar e transformar o conteúdo:

1. **Remover linhas em branco**  
   As linhas vazias poderiam gerar erros na importação ou leitura do arquivo:
   ```bash
   sed '/^\s*$/d' arquivo.sql > arquivo_sem_linhas_vazias.sql
   ```

2. **Remover aspas simples `'`**  
   As aspas simples presentes no arquivo poderiam quebrar a importação:
   ```bash
   sed "s/'//g" arquivo_sem_linhas_vazias.sql > arquivo_limpo.sql
   ```

3. **Converter o arquivo para CSV**  
   Utilizei `awk` para extrair as colunas desejadas e gerar um CSV que facilitaria a carga no banco:
   ```bash
   awk -F',' '{print $1","$2","$3","$4}' arquivo_limpo.sql > dados.csv
   ```
   (A estrutura de campos foi ajustada conforme o schema do banco.)

   ```

### Criação de Índices no Banco de Dados

Para garantir performance nas consultas, principalmente para as buscas e filtros mais comuns, criei os seguintes índices:

- Índice simples na coluna `data`, para acelerar consultas filtrando por intervalo de datas.
- Índice na coluna `origin` (origem da taxa de conversão), pois ela é frequentemente usada em cláusulas WHERE.
- Índice composto em (`start`, `end`) para acelerar buscas com filtros combinados por datas inicial e final.
- Índices em colunas usadas em joins e filtros frequentes, como `user_id` ou `transaction_id`, conforme a modelagem.

Estes índices reduziram significativamente o tempo das consultas, mesmo com o banco contendo milhões de linhas.

---

## Resumo

Ao escolher Go, consegui unir alta performance, baixo consumo de recursos e facilidade no deploy.  
A preparação cuidadosa dos dados usando comandos Linux permitiu uma importação limpa e rápida.  
A criação estratégica dos índices garantiu que as buscas no banco fossem rápidas e escaláveis.

Este conjunto de decisões resultou em uma solução robusta e eficiente para o processamento do volume de dados requerido.

---
