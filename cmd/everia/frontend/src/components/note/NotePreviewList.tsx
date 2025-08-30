'use client'

import { useEffect } from 'react'
import { useNote } from '@/hooks/useNote'
import { useNoteList } from '@/hooks/useNoteList'
import { mapWithSeparator } from '@/utils'

import { LucideFileSignature, LucideTrash } from '@longslark/ui/components'
import { Button } from '@longslark/ui/components/button'
import { ScrollArea } from '@longslark/ui/components/scroll-area'
import { Separator } from '@longslark/ui/components/separator'
import { ListNotes } from '@wailsjs/go/main/App'

import NotePreview from './NotePreview'

type NotePreviewListProps = {
  onSelect?: () => void
}
const NotePreviewList = ({ onSelect }: NotePreviewListProps) => {
  const { notes, setNotes, selectedNoteIndex, handleNoteSelect } = useNoteList({ onSelect })
  const { handleCreateNewNote, handleDeleteNote } = useNote()

  useEffect(() => {
    ListNotes().then(res => setNotes(res ?? []))
  }, [])

  if (!notes || notes.length === 0) {
    return (
      <div className='flex h-full w-[250px] flex-col rounded-md bg-muted px-2 py-2'>
        <div className='flex w-full flex-row justify-between'>
          <Button variant='outline' className='bg-transparent' onClick={handleCreateNewNote}>
            <LucideFileSignature className='size-4 text-zinc-300' />
          </Button>
          <Button variant='outline' className='bg-transparent' onClick={handleDeleteNote}>
            <LucideTrash className='size-4 text-zinc-300' />
          </Button>
        </div>
        <div>No Notes Yet!</div>
      </div>
    )
  }

  return (
    <div className='flex h-full w-[250px] flex-col rounded-md bg-muted px-2 py-2'>
      <div className='flex w-full flex-row justify-between'>
        <Button variant='outline' className='bg-transparent' onClick={handleCreateNewNote}>
          <LucideFileSignature className='size-4 text-zinc-300' />
        </Button>
        <Button variant='outline' className='bg-transparent' onClick={handleDeleteNote}>
          <LucideTrash className='size-4 text-zinc-300' />
        </Button>
      </div>
      <ScrollArea className='h-[75vh]'>
        {mapWithSeparator(
          notes.map((note, index) => (
            <NotePreview
              key={index}
              {...note}
              isActive={selectedNoteIndex === index}
              onClick={handleNoteSelect(index)}
            />
          )),
          <Separator className='border border-background' />
        )}
      </ScrollArea>
    </div>
  )
}

export default NotePreviewList
