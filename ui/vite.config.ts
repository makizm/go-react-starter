import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api/docs': {
        target: process.env.VITE_DOCS_URL || 'http://localhost:8081',
        changeOrigin: true,
      },
      '/api': {
        // Use 'server' hostname in Docker, 'localhost' for local dev
        target: process.env.VITE_API_URL || 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
