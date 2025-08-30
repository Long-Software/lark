import React, { useState } from 'react'

function UnlockScreen({ onUnlock }: { onUnlock: (pw: string) => void }) {
  const [password, setPassword] = useState('')

  const handleSubmit = () => {
    onUnlock(password)
    setPassword('')
  }

  return (
    <div className='mx-auto mt-20 max-w-md rounded bg-white p-4 shadow'>
      <h2 className='mb-4 text-xl font-bold'>🔒 Enter Master Password</h2>
      <input
        type='password'
        value={password}
        onChange={e => setPassword(e.target.value)}
        className='mb-2 w-full rounded border p-2'
        placeholder='Master password'
      />
      <button onClick={handleSubmit} className='w-full rounded bg-blue-500 p-2 text-white'>
        Unlock
      </button>
    </div>
  )
}

export default UnlockScreen
