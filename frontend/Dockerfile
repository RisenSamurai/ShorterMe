# --- Stage 1: Build the SvelteKit app ---
FROM node:18-alpine3.21 AS build

WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# --- Stage 2: Run the Node.js server ---
FROM node:18-alpine3.21

WORKDIR /app


COPY --from=build /app/build ./build
COPY --from=build /app/package*.json ./
COPY --from=build /app/package-lock.json ./
 #Important for consistent installs


RUN npm install --production


EXPOSE 3000


CMD ["node", "build/index.js"]