FROM node:22 AS builder
WORKDIR /app

ARG VITE_AUTH_URL
ARG VITE_HUB_URL
ENV VITE_AUTH_URL=${VITE_AUTH_URL}
ENV VITE_HUB_URL=${VITE_HUB_URL}

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:22
WORKDIR /app
COPY --from=builder /app/build ./build
COPY --from=builder /app/node_modules ./node_modules
CMD ["node", "build"]