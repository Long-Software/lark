'use client'

import { useRef } from 'react'
import dynamic from 'next/dynamic'
import FloatingNoteTitle from '@/components/note/FloatingNoteTitle'
import NotePreviewList from '@/components/note/NotePreviewList'

const MarkdownEditor = dynamic(() => import('@/components/note/NoteMarkdownEditor'), {
  ssr: false
})
const NotePage = () => {
  const ref = useRef<HTMLDivElement>(null)
  const resetScroll = () => {
    ref.current?.scrollTo(0, 0)
  }

  return (
    <div className='flex px-2'>
      <div ref={ref} className='h-full w-full'>
        <FloatingNoteTitle className='pt-2' />
        <MarkdownEditor />
      </div>
      <NotePreviewList onSelect={resetScroll} />
    </div>
  )
}

export default NotePage
