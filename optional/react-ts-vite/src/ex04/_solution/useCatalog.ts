import { useState, useEffect } from 'react'
import { Product } from '../types'

export interface UseCatalogResult {
  loading: boolean
  data: Product | null
  error: string | null
}

export function useCatalog(sku: string): UseCatalogResult {
  const [loading, setLoading] = useState<boolean>(false)
  const [data, setData] = useState<Product | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (!sku) {
      setData(null)
      setError(null)
      return
    }

    const controller = new AbortController()
    setLoading(true)
    setError(null)

    fetch(`/api/catalog?sku=${encodeURIComponent(sku)}`, { signal: controller.signal })
      .then((res) => {
        if (!res.ok) throw new Error(res.statusText)
        return res.json() as Promise<Product>
      })
      .then((product) => {
        setData(product)
        setError(null)
      })
      .catch((err: Error) => {
        if (err.name === 'AbortError') return
        setError(err.message)
        setData(null)
      })
      .finally(() => setLoading(false))

    return () => controller.abort()
  }, [sku])

  return { loading, data, error }
}
