FROM node:20-slim 

RUN apt update && apt install -y openssl procps

RUN npm i -g @nestjs/cli@10.3.2

WORKDIR /home/node/app

USER node

CMD tail -f /dev/null
