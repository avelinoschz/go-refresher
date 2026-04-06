import React from 'react'
import { Product } from '../types'

interface Props {
  product: Product
  badge?: string
}

const ProductCard: React.FC<Props> = ({ product, badge }) => {
  return (
    <div>
      <p>SKU: {product.sku}</p>
      <p>
        Name: {product.name}
        {badge && <span> [{badge}]</span>}
      </p>
      <p>Price: ${product.price}</p>
    </div>
  )
}

export default ProductCard
