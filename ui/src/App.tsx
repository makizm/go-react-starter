import { useEffect, useState } from 'react'

function App() {
  const [health, setHealth] = useState('Check Health')
  const [info, setInfo] = useState({
    name: '',
    version: '',
    lastChecked: 'never',
  })
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    handleFetchInfo()
  }, [])

  const handleClick = () => {
    fetch('/api/health')
      .then((response) => response.json())
      .then((data) => {
        setHealth(data.status)
      })
      .catch((err) => {
        setError(err.message)
      })
  }

  const handleFetchInfo = () => {
    fetch('/api/info')
      .then((response) => response.json())
      .then((data) => {
        setInfo({
          name: data.name,
          version: data.version,
          lastChecked: new Date(data.time).toLocaleString(),
        })
      })
      .catch((err) => {
        setError(err.message)
      })
  }

  return (
    <div>

      <div style={{ display: 'flex', alignItems: 'baseline', gap: '0.5em', justifyContent: 'center' }}>
        <h1>{info.name}</h1>
        <span>v{info.version}</span>
      </div>

      {error && (
        <div style={{ color: 'red' }}>
          <strong>Error:</strong> {error}
        </div>
      )}

      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', gap: '1em' }}>
        <button name="check health" onClick={handleClick}>{health}</button>
        <button name="fetch info" onClick={handleFetchInfo}>
          Refresh
        </button>
        <span>Last refreshed: {info.lastChecked}</span>
      </div>
    </div>
  )
}

export default App
