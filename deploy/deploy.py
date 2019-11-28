#!/usr/bin/env python
# -*- coding: utf-8 -*-

from fabric import Connection
from invoke.exceptions import UnexpectedExit


print('Code For Curitiba - Sistema de deploy do projeto onibus-io-backend\n')

connect_kwargs = {
    "key_filename": ["id_rsa"]
}

with Connection(host='jarvis.preludian.com.br', 
                user='preludian', 
                port=40022,
                connect_kwargs=connect_kwargs) as c:
    try:
        print ('Stopping containers and destroying images...')
        c.run('docker stop code4cwb-onibus-io-backend-container && \
               docker rm code4cwb-onibus-io-backend-container && \
               docker rmi code4cwb/onibus-io-backend')
    except UnexpectedExit:
        print ('Delete deleting container not possible... Keep going')

    print('Running container...')
    c.run('docker run --name code4cwb-onibus-io-backend-container -d --restart=always --network=code4cwb-onibusio -m 512m -e ONIBUSIO_DB_HIST=onibus-historico -e ONIBUSIO_DB_URL=mongodb://code4cwb-mongodb/onibus-historico -e PORT=3000 -p 9002:3000 code4cwb/onibus-io-backend')

print('\nTerminado com sucesso')
