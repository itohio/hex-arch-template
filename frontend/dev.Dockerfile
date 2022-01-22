FROM node:alpine
COPY .env frontend/tsconfig.json frontend/package.json frontend/package-lock.json /src/
COPY frontend/public /src/public
WORKDIR /src
RUN npm install
COPY schema /schema
ENTRYPOINT [ "npm", "start" ]
