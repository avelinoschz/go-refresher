import React, { useEffect, useState } from 'react'
import { Product } from '../types'

const CatalogSearch: React.FC = () => {
  const [sku, setSku] = useState<string>('')
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

  return (
    <div>
      <input
        value={sku}
        onChange={(e) => setSku(e.target.value)}
        placeholder="Enter SKU (e.g. HAMMER-001)"
      />
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error}</p>}
      {data && (
        <p>
          {data.name} — ${data.price}
        </p>
      )}
    </div>
  )
}

export default CatalogSearch
