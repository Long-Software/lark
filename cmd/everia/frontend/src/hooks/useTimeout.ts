'use client'

import { useEffect, useRef } from 'react'

export function useTimeout(timeoutMs: number, onTimeout: () => void) {
  const timerRef = useRef<number | null>(null)

  const resetTimer = () => {
    if (timerRef.current) {
      clearTimeout(timerRef.current)
    }
    timerRef.current = window.setTimeout(onTimeout, timeoutMs)
  }

  useEffect(() => {
    const events = ['mousemove', 'keydown', 'mousedown', 'touchstart']

    const handleEvent = () => resetTimer()

    events.forEach(e => window.addEventListener(e, handleEvent))

    resetTimer() // Initialize timer on mount

    return () => {
      if (timerRef.current) clearTimeout(timerRef.current)
      events.forEach(e => window.removeEventListener(e, handleEvent))
    }
  }, [])
}
