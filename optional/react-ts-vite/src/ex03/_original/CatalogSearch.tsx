import React, { useState } from 'react'
import { Product } from '../types'

// TODO: implement the fetch logic.
//
// When the input changes:
//  - if the value is empty, clear data and error
//  - otherwise fetch /api/catalog?sku=<value>
//  - set loading=true before the request, loading=false after
//  - on success: set data to the parsed Product, clear error
//  - on failure: set error to the response status text or thrown message, clear data
//
// Use an AbortController to cancel in-flight requests when sku changes.

const CatalogSearch: React.FC = () => {
  const [sku, setSku] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(false)
  const [data, setData] = useState<Product | null>(null)
  const [error, setError] = useState<string | null>(null)

  // TODO: add useEffect that triggers fetch when sku changes

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
