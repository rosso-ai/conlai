# Bringing AI to the world with ConLAi
Server for Ledger type Federated Learning

## What's ConLAi?
Con(sensus)L(erning) Ai is server module for Ledger type federated learning.  
Ledger type federated learning achieves federated learning in a way that feels like Git.  

![features](https://github.com/rosso-ai/conlai/blob/main/docs/images/conlai_features.png?raw=true)

## How to Start
Docker makes it easy to start a server.

```shell
docker pull ghcr.io/rosso-ai/conlai:latest
docker run -d -p 9200:9200 ghcr.io/rosso-ai/conlai
```

Connect to the client learning app using Websocket communication on port 9200.  
Please use the following drivers for client-side learning apps.  
https://github.com/rosso-ai/pyConLAi

## License
This software is dual licensed under AGPL-3.0 and commercial license.  
If you would like to use a commercial license, please contact [Rosso inc](https://www.rosso-tokyo.co.jp/contact/).

### including OSS
This software including to the following OSS:
* [gorilla](https://github.com/gorilla/websocket) : [BSD-2-Clause](https://github.com/gorilla/websocket/blob/main/LICENSE)
* [protobuf](https://github.com/protocolbuffers/protobuf) : [BSD-3-Clause](https://github.com/protocolbuffers/protobuf/blob/main/LICENSE)

## Authors
ConLAi is developed by [Rosso inc](https://www.rosso-tokyo.co.jp/).
