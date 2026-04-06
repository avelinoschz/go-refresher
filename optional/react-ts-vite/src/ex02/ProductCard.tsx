import React from 'react'
import { Product } from '../types'

interface Props {
  product: Product
  badge?: string
}

// TODO: render the product's sku, name, and price.
// If the badge prop is provided, render it alongside the product name.

const ProductCard: React.FC<Props> = (_props) => {
  return (
    <div>
      {/* TODO: implement */}
    </div>
  )
}

export default ProductCard
