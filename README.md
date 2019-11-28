# Onibus-io-backend
[![Build Status](https://travis-ci.org/CodeForCuritiba/onibus-io-backend.svg?branch=master)](https://travis-ci.org/CodeForCuritiba/onibus-io-backend)
Sistema em GO que cuida do fornecimento dos dados para a aplicação frontent do Onibus-io.

## Desenvolvendo
Foi criado um arquivo do docker compose que cuida do desenvolvimento da aplicação; Pressupõe-se que você possua o docker e o docker-compose instalados.
Para trabalhar no projeto, execute:

    $ docker-compose up

Esse comando executará três containers. O primeiro será o servidor mongodb, a aplicação em si, onde ele fará o build e gerará o executável e o terceiro é um arquivo de importação.

O container de importação carrega uns dados estáticos na sua base de dados mongodb de maneira automática conforme arquivo `load_data.sh`. Para evitar que haja mais de um carregamento de dados, ele cria um arquivo de lock na pasta padrão chamado `data_imported.lock`.
Para reimportar os dados novamente, basta remover esse arquivo e executar o docker compose novamente.

## API

* **/versao** : Versão da API;
* **/api/linhas** : Lista de linhas (sem tabela e sem pontos);
* **/api/linhas/{codigoLinha}** : Retorna a linha de acordo com o código;
* **/api/veiculos** : Retorna uma lista de todos os veículos em circulação;
* **/api/veiculos/{codigoVeiculo}** : Retorna uma lista das últimas posições de um veículo;
* **/api/veiculos/linha/{codigoLinha}** : Retorna uma lista das últimas posições dos veículos de uma linha;

## FAQ

### Por que são retornados as últimas posições de um veículo e não apenas a última?

Isso é feito para que possamos calcular a direção para onde o veículo está indo.
