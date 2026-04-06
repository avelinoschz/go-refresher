import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Counter from './Counter'

describe('Counter', () => {
  it('renders initial count of 0', () => {
    render(<Counter />)
    expect(screen.getByText(/count:/i)).toHaveTextContent('Count: 0')
  })

  it('increments count when Increment is clicked', async () => {
    render(<Counter />)
    await userEvent.click(screen.getByRole('button', { name: /increment/i }))
    expect(screen.getByText(/count:/i)).toHaveTextContent('Count: 1')
  })

  it('decrements count when Decrement is clicked', async () => {
    render(<Counter />)
    await userEvent.click(screen.getByRole('button', { name: /increment/i }))
    await userEvent.click(screen.getByRole('button', { name: /increment/i }))
    await userEvent.click(screen.getByRole('button', { name: /decrement/i }))
    expect(screen.getByText(/count:/i)).toHaveTextContent('Count: 1')
  })

  it('resets count to 0 when Reset is clicked', async () => {
    render(<Counter />)
    await userEvent.click(screen.getByRole('button', { name: /increment/i }))
    await userEvent.click(screen.getByRole('button', { name: /increment/i }))
    await userEvent.click(screen.getByRole('button', { name: /reset/i }))
    expect(screen.getByText(/count:/i)).toHaveTextContent('Count: 0')
  })
})
