import { gql, useQuery } from 'urql'
import './App.css'

const PING_QUERY = gql`
  query GetPing {
    ping
  }
`

function App() {
  const [result] = useQuery({
    query: PING_QUERY,
  })

  const { fetching, error, data } = result

  console.log('App state:', { fetching, error, data })

  return (
    <div className="App" style={{ padding: '20px', textAlign: 'center' }}>
      <h1>friii - Connection Test</h1>
      <div className="card" style={{ border: '1px solid #ccc', padding: '20px', margin: '20px' }}>
        {fetching && <p>Loading data from backend...</p>}
        {error && (
          <div style={{ color: 'red' }}>
            <p>Error connecting to backend:</p>
            <pre>{error.message}</pre>
            <p>Make sure backend is running at http://localhost:8080/query</p>
          </div>
        )}
        {data && (
          <p style={{ fontSize: '1.5em' }}>
            Backend Response: <strong style={{ color: '#646cff' }}>{data.ping}</strong>
          </p>
        )}
      </div>
      <div style={{ marginTop: '20px', color: '#666' }}>
        <p>Backend URL: http://localhost:8080/query</p>
        <p>If you see "Loading" forever, check the browser console (F12) for CORS or Network errors.</p>
      </div>
    </div>
  )
}

export default App
