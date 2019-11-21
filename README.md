# onibus-io-backend


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
