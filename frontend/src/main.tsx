import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Client, Provider, cacheExchange, fetchExchange } from 'urql'
import './index.css'
import App from './App.tsx'

const client = new Client({
  url: 'http://localhost:8080/query',
  exchanges: [cacheExchange, fetchExchange],
})

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Provider value={client}>
      <App />
    </Provider>
  </StrictMode>,
)
