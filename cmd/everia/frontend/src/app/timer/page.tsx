'use client'

import { useEffect } from 'react'
import TimerContainer from '@/components/timer/TimerContainer'
import { useTimer } from '@/hooks/useTimer'

const TimerPage = () => {
  const { seconds, isRunning, startTimer, pauseTimer, resetTimer } = useTimer()
  // const startKeys = ['Space']
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.code === 'Space') {
        e.preventDefault()
        isRunning ? pauseTimer() : startTimer()
      }
    }

    const handleMouseDown = (e: MouseEvent) => {
      if (e.button === 0) {
        // Left click
        isRunning ? pauseTimer() : startTimer()
      } else if (e.button === 2) {
        // Right click
        resetTimer()
      }
    }

    window.addEventListener('keydown', handleKeyDown)
    window.addEventListener('mousedown', handleMouseDown)
    window.addEventListener('contextmenu', e => e.preventDefault()) // Prevent context menu

    return () => {
      window.removeEventListener('keydown', handleKeyDown)
      window.removeEventListener('mousedown', handleMouseDown)
      window.removeEventListener('contextmenu', e => e.preventDefault())
    }
  }, [isRunning, startTimer, pauseTimer, resetTimer])

  return (
    <div className='w-full'>
      <p>
        <strong>Controls:</strong>
      </p>
      <ul>
        <li>
          <kbd>Space</kbd> — Toggle start/pause
        </li>
        <li>
          <kbd>Left click</kbd> — Start
        </li>
        <li>
          <kbd>Right click</kbd> — Reset
        </li>
      </ul>
      <div className='grid grid-cols-3 gap-4 px-4'>
        <TimerContainer title='Hours' time={Math.floor(seconds / 3600)} />
        <TimerContainer title='Minutes' time={Math.floor(seconds / 60)} />
        <TimerContainer title='Seconds' time={seconds % 60} />
      </div>
    </div>
  )
}

export default TimerPage
