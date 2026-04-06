import React from 'react'

// TODO: add useState and implement the handlers below.
// The component should have three buttons: Increment, Decrement, and Reset.

const Counter: React.FC = () => {
  // TODO: declare count state with useState<number>

  return (
    <div>
      <p>Count: {/* TODO: render count */}</p>
      <button onClick={/* TODO */undefined}>Increment</button>
      <button onClick={/* TODO */undefined}>Decrement</button>
      <button onClick={/* TODO */undefined}>Reset</button>
    </div>
  )
}

export default Counter
