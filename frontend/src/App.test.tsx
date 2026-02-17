import { render, screen } from '@testing-library/react'
import { describe, it, expect, vi } from 'vitest'
import App from './App'
import { useQuery } from 'urql'

// urql をモック化
vi.mock('urql', async (importOriginal) => {
  const actual = await importOriginal<typeof import('urql')>()
  return {
    ...actual,
    useQuery: vi.fn(),
  }
})

describe('App Component', () => {
  it('renders loading state', () => {
    vi.mocked(useQuery).mockReturnValue([
      { fetching: true, error: undefined, data: undefined },
      vi.fn(),
    ])
    render(<App />)
    expect(screen.getByText(/Loading data from backend/i)).toBeInTheDocument()
  })

  it('renders success state', () => {
    vi.mocked(useQuery).mockReturnValue([
      { fetching: false, error: undefined, data: { ping: 'pong' } },
      vi.fn(),
    ])
    render(<App />)
    expect(screen.getByText(/Backend Response:/i)).toBeInTheDocument()
    expect(screen.getByText(/pong/i)).toBeInTheDocument()
  })

  it('renders error state', () => {
    vi.mocked(useQuery).mockReturnValue([
      { fetching: false, error: new Error('Network error'), data: undefined },
      vi.fn(),
    ])
    render(<App />)
    expect(screen.getByText(/Error connecting to backend/i)).toBeInTheDocument()
    // pre タグ内のテキストを正確に指定
    expect(screen.getByText('Network error')).toBeInTheDocument()
  })
})
