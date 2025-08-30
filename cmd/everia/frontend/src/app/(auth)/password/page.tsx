'use client'

import { useState } from 'react'
import UnlockScreen from '@/components/password/UnlockScreen'
import VaultView from '@/components/password/VaultView'
import { useTimeout } from '@/hooks/useTimeout'

import { models } from '@wailsjs/go/models'

const TIMEOUT = 5 * 60 * 1000 // 5 minutes

function App() {
  const [unlocked, setUnlocked] = useState(false)
  const [entries, setEntries] = useState<models.PasswordEntry[]>([
    {
      site: 'www.example.com',
      username: 'long seng',
      password: '123',
      notes: 'slfj'
    },
    {
      site: 'www.example.com',
      username: 'long seng',
      password: '123',
      notes: 'slfj'
    },
    {
      site: 'www.example.com',
      username: 'long seng',
      password: '123',
      notes: 'slfj'
    }
  ])

  const handleUnlock = async (password: string) => {
    console.log(password)
    // UnlockVault(password)
    //   .then(res => {
    //     setEntries(res)
    //     setUnlocked(true)
    //   })
    //   .catch(err => {
    //     console.error(err)
    //     setUnlocked(false)
    //   })
  }

  const handleLock = () => {
    setUnlocked(false)
    setEntries([])
  }

  useTimeout(TIMEOUT, () => {
    if (unlocked) {
      alert('Session timed out. Auto-locking.')
      handleLock()
    }
  })

  return unlocked ? <VaultView entries={entries} /> : <UnlockScreen onUnlock={handleUnlock} />
}

export default App
