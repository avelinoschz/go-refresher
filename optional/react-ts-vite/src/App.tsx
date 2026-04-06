import Counter from './ex01/Counter'
import ProductCard from './ex02/ProductCard'
import CatalogSearch from './ex03/CatalogSearch'
import CatalogList from './ex05/CatalogList'

const sampleProduct = { sku: 'HAMMER-001', name: 'Hammer', price: 25 }

export default function App() {
  return (
    <div style={{ fontFamily: 'monospace', padding: '2rem', maxWidth: '600px' }}>
      <h1>React + TypeScript Refresher</h1>

      <section>
        <h2>ex01 — Counter</h2>
        <Counter />
      </section>

      <hr />

      <section>
        <h2>ex02 — ProductCard</h2>
        <ProductCard product={sampleProduct} />
        <ProductCard product={sampleProduct} badge="New" />
      </section>

      <hr />

      <section>
        <h2>ex03 — CatalogSearch (requires Go backend)</h2>
        <CatalogSearch />
      </section>

      <hr />

      <section>
        <h2>ex05 — CatalogList (requires Go backend)</h2>
        <CatalogList />
      </section>
    </div>
  )
}
