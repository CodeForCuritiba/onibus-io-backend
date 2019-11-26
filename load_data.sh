#!/bin/sh

LOCK_FILE="data_imported.lock";

if [ -f "${LOCK_FILE}" ]
then
  echo "Arquivo ja importado.";
  exit 0;
fi

echo "Importando base de dados inicial...";

apt-get update && apt-get install -y wget unzip;
wget https://github.com/CodeForCuritiba/onibus-io-backend/releases/download/data/onibus_io_data.zip;
unzip onibus_io_data.zip;

cd onibus_io_data;
mongoimport -v --host=$MONGODB_SERVER --port 27017 --db=$MONGODB_DATABASE --collection=veiculos --file=veiculos.json;
mongoimport -v --host=$MONGODB_SERVER --port 27017 --db=$MONGODB_DATABASE --collection=linhas --file=linhas.json;

cd ..;

rm -Rf onibus_io_data.zip;
rm -Rf onibus_io_data;

date > $LOCK_FILE;

echo "Dados carregados!";
