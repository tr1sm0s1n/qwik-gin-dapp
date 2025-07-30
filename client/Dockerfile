FROM node:22.17-alpine AS builder

ARG API_URL
ARG ADMIN_SIGN
ENV PUBLIC_API=$API_URL
ENV PUBLIC_ADMIN_SIGN=$ADMIN_SIGN

RUN npm install -g pnpm

WORKDIR /app
COPY package.json ./
RUN pnpm install

COPY . .
RUN pnpm run build
