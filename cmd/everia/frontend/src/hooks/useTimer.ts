'use client'

import { useCallback, useRef, useState } from 'react'

export const useTimer = () => {
  const [seconds, setSeconds] = useState(0)
  const [isRunning, setIsRunning] = useState(false)
  const ref = useRef<NodeJS.Timeout | null>(null)

  const startTimer = useCallback(() => {
    if (ref.current !== null) return
    setIsRunning(true)
    ref.current = setInterval(() => {
      setSeconds(prev => prev + 1)
    }, 1000)
  }, [])
  const pauseTimer = useCallback(() => {
    if (ref.current) {
      clearInterval(ref.current)
      ref.current = null
      setIsRunning(false)
    }
  }, [])
  const resetTimer = useCallback(() => {
    pauseTimer()
    setSeconds(0)
  }, [pauseTimer])

  return { seconds, isRunning, startTimer, pauseTimer, resetTimer }
}
