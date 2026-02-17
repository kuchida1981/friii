import { gql, useQuery } from '@apollo/client'
import './App.css'

const PING_QUERY = gql`
  query GetPing {
    ping
  }
`

function App() {
  const { loading, error, data } = useQuery(PING_QUERY)

  return (
    <div className="App">
      <h1>friii - Connection Test</h1>
      <div className="card">
        {loading && <p>Loading...</p>}
        {error && <p style={{ color: 'red' }}>Error: {error.message}</p>}
        {data && <p>Backend Response: <strong>{data.ping}</strong></p>}
      </div>
      <p className="read-the-docs">
        Check if backend is running on http://localhost:8080
      </p>
    </div>
  )
}

export default App
