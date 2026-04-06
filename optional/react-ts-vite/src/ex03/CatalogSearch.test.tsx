import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { beforeEach, vi } from 'vitest'
import CatalogSearch from './CatalogSearch'

const mockFetch = vi.fn()

beforeEach(() => {
  vi.stubGlobal('fetch', mockFetch)
  mockFetch.mockReset()
})

describe('CatalogSearch', () => {
  it('renders an input', () => {
    render(<CatalogSearch />)
    expect(screen.getByPlaceholderText(/enter sku/i)).toBeInTheDocument()
  })

  it('shows product data on successful fetch', async () => {
    // Use mockResolvedValue (not Once) — userEvent.type fires per-character,
    // each triggering a new fetch call via useEffect.
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => ({ sku: 'HAMMER-001', name: 'Hammer', price: 25 }),
    })

    render(<CatalogSearch />)
    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'HAMMER-001')

    await waitFor(() => {
      expect(screen.getByText(/hammer — \$25/i)).toBeInTheDocument()
    })
  })

  it('shows error message on failed fetch', async () => {
    mockFetch.mockResolvedValue({
      ok: false,
      statusText: 'Not Found',
    })

    render(<CatalogSearch />)
    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'UNKNOWN-999')

    await waitFor(() => {
      expect(screen.getByText(/error: not found/i)).toBeInTheDocument()
    })
  })
})
