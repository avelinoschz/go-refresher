import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { beforeEach, vi } from 'vitest'
import CatalogList from './CatalogList'

const mockFetch = vi.fn()

beforeEach(() => {
  vi.stubGlobal('fetch', mockFetch)
  mockFetch.mockReset()
})

describe('CatalogList', () => {
  it('renders the search input and button', () => {
    render(<CatalogList />)
    expect(screen.getByPlaceholderText(/enter sku/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /search/i })).toBeInTheDocument()
  })

  it('adds product to list after successful search', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ sku: 'HAMMER-001', name: 'Hammer', price: 25 }),
    })

    render(<CatalogList />)
    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'HAMMER-001')
    await userEvent.click(screen.getByRole('button', { name: /search/i }))

    await waitFor(() => {
      expect(screen.getByText(/hammer-001/i)).toBeInTheDocument()
      expect(screen.getByText(/hammer/i)).toBeInTheDocument()
    })
  })

  it('shows error message when product is not found', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: false,
      statusText: 'Not Found',
    })

    render(<CatalogList />)
    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'UNKNOWN-999')
    await userEvent.click(screen.getByRole('button', { name: /search/i }))

    await waitFor(() => {
      expect(screen.getByText(/error/i)).toBeInTheDocument()
    })
  })

  it('does not add duplicate SKUs to the list', async () => {
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => ({ sku: 'HAMMER-001', name: 'Hammer', price: 25 }),
    })

    render(<CatalogList />)

    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'HAMMER-001')
    await userEvent.click(screen.getByRole('button', { name: /search/i }))
    await waitFor(() => expect(screen.getAllByText(/hammer-001/i)).toHaveLength(1))

    await userEvent.type(screen.getByPlaceholderText(/enter sku/i), 'HAMMER-001')
    await userEvent.click(screen.getByRole('button', { name: /search/i }))
    await waitFor(() => expect(screen.getAllByText(/hammer-001/i)).toHaveLength(1))
  })
})
