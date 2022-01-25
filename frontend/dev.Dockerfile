FROM node:16-alpine
COPY frontend/tsconfig.json frontend/package.json frontend/package-lock.json /src/
COPY frontend/generated /src/generated
COPY frontend/public /src/public
WORKDIR /src
RUN npm install
ENTRYPOINT [ "npm", "start" ]
