FROM node:22-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm install

# Copy source code (will be overridden by volume mount in dev)
COPY . .

EXPOSE 5173

# Run Vite dev server with a pinned port (don't auto-pick another port)
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173", "--strictPort"]
