// This file can be used to add references for global types like `vite/client`.

// Add global `vite/client` types. For more info, see: https://vitejs.dev/guide/features#client-types
/// <reference types="vite/client" />

declare global {
  interface Window {
    ethereum: {
      request(args: { method: 'eth_requestAccounts' }): Promise<string[]>
      request(args: {
        method: 'personal_sign'
        params: [string | null, string | null]
      }): Promise<string>
    }
  }
}

export {}
