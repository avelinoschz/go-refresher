import { useState, useEffect } from 'react'
import { Product } from '../types'

export interface UseCatalogResult {
  loading: boolean
  data: Product | null
  error: string | null
}

// TODO: implement useCatalog.
//
// - fetch /api/catalog?sku=<sku> whenever sku changes
// - skip the fetch if sku is empty
// - manage loading, data, and error state
// - cancel the in-flight request with AbortController on cleanup

export function useCatalog(_sku: string): UseCatalogResult {
  const [loading, setLoading] = useState<boolean>(false)
  const [data, setData] = useState<Product | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    // TODO: implement
  }, [_sku])

  return { loading, data, error }
}
