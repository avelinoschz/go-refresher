import { renderHook, waitFor } from '@testing-library/react'
import { beforeEach, vi } from 'vitest'
import { useCatalog } from './useCatalog'

const mockFetch = vi.fn()

beforeEach(() => {
  vi.stubGlobal('fetch', mockFetch)
  mockFetch.mockReset()
})

describe('useCatalog', () => {
  it('returns initial idle state for empty sku', () => {
    const { result } = renderHook(() => useCatalog(''))
    expect(result.current.loading).toBe(false)
    expect(result.current.data).toBeNull()
    expect(result.current.error).toBeNull()
  })

  it('returns product data on successful fetch', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ sku: 'HAMMER-001', name: 'Hammer', price: 25 }),
    })

    const { result } = renderHook(() => useCatalog('HAMMER-001'))

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
      expect(result.current.data?.name).toBe('Hammer')
      expect(result.current.error).toBeNull()
    })
  })

  it('returns error on failed fetch', async () => {
    mockFetch.mockResolvedValueOnce({
      ok: false,
      statusText: 'Not Found',
    })

    const { result } = renderHook(() => useCatalog('UNKNOWN-999'))

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
      expect(result.current.data).toBeNull()
      expect(result.current.error).toBe('Not Found')
    })
  })
})
