'use client'

import { ComponentProps } from 'react'
import { useNote } from '@/hooks/useNote'

import { cn } from '@longslark/tailwind/lib'

type FloatingNoteTitleProps = ComponentProps<'div'>

const FloatingNoteTitle = ({ className, ...props }: FloatingNoteTitleProps) => {
  const { selectedNote } = useNote()
  if (!selectedNote) return null

  return (
    <div className={cn('flex justify-center', className)} {...props}>
      <p className='text-gray-400'>{selectedNote.title}</p>
    </div>
  )
}

export default FloatingNoteTitle
