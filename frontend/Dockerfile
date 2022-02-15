FROM node:16.13.2-alpine3.15
EXPOSE 3000
WORKDIR /app
COPY package*.json .
RUN npm ci --production
COPY . .
RUN npm run build
RUN npm install -g serve
CMD [ "serve","-s","-p","3000","build" ]