import React, { useState } from 'react'
import { Product } from '../types'
import { useCatalog } from '../ex04/useCatalog'

// TODO: implement CatalogList.
//
// - a controlled <input> for typing a SKU query
// - a <form> with onSubmit that prevents default and triggers a lookup
// - use useCatalog(submittedSku) to fetch the product
// - on success: add the product to a `results` list and clear the input
// - on error: show an inline error message
// - render the list of found products (sku + name)

const CatalogList: React.FC = () => {
  const [query, setQuery] = useState<string>('')
  const [submittedSku, setSubmittedSku] = useState<string>('')
  const [results, setResults] = useState<Product[]>([])

  const { data, error } = useCatalog(submittedSku)

  // TODO: handle the effect of data/error changes to update results

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // TODO: trigger the lookup
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
            {/* TODO: render sku and name */}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default CatalogList
