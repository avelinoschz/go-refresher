import { render, screen } from '@testing-library/react'
import ProductCard from './ProductCard'

const product = { sku: 'HAMMER-001', name: 'Hammer', price: 25 }

describe('ProductCard', () => {
  it('renders product sku, name, and price', () => {
    render(<ProductCard product={product} />)
    // Use text content of specific elements to avoid ambiguous matches.
    expect(screen.getByText(/sku:/i)).toHaveTextContent('HAMMER-001')
    expect(screen.getByText(/^name:/i)).toHaveTextContent('Hammer')
    expect(screen.getByText(/price:/i)).toHaveTextContent('25')
  })

  it('renders badge when provided', () => {
    render(<ProductCard product={product} badge="New" />)
    expect(screen.getByText(/\[new\]/i)).toBeInTheDocument()
  })

  it('does not render badge when not provided', () => {
    render(<ProductCard product={product} />)
    expect(screen.queryByText(/\[new\]/i)).not.toBeInTheDocument()
  })
})
