import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'

describe('Simple Test', () => {
  it('should render', () => {
    render(<div>Test</div>)
    expect(screen.getByText('Test')).toBeInTheDocument()
  })
})
