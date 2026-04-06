import React, { useEffect, useState } from 'react'
import { Product } from '../types'
import { useCatalog } from '../ex04/useCatalog'

const CatalogList: React.FC = () => {
  const [query, setQuery] = useState<string>('')
  const [submittedSku, setSubmittedSku] = useState<string>('')
  const [results, setResults] = useState<Product[]>([])

  const { data, error } = useCatalog(submittedSku)

  useEffect(() => {
    if (!data) return
    setResults((prev) => {
      const exists = prev.some((p) => p.sku === data.sku)
      return exists ? prev : [...prev, data]
    })
    setQuery('')
    setSubmittedSku('')
  }, [data])

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (query.trim()) {
      setSubmittedSku(query.trim())
    }
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="Enter SKU"
        />
        <button type="submit">Search</button>
      </form>
      {error && <p>Error: {error}</p>}
      <ul>
        {results.map((p) => (
          <li key={p.sku}>
            {p.sku} — {p.name}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default CatalogList
