'use client'

import { useAtom, useSetAtom } from 'jotai'

import { ReadNote } from '@wailsjs/go/main/App'

import { notesAtom, selectedNoteContentAtom, selectedNoteIndexAtom } from './notes'

export const useNoteList = ({ onSelect }: { onSelect?: () => void }) => {
  const [notes, setNotes] = useAtom(notesAtom)
  const [selectedNoteIndex, setSelectedNoteIndex] = useAtom(selectedNoteIndexAtom)
  const setNoteContent = useSetAtom(selectedNoteContentAtom)

  const handleNoteSelect = (index: number) => () => {
    ReadNote(notes[index]!.title).then(setNoteContent)
    setSelectedNoteIndex(index)

    if (onSelect) {
      onSelect()
    }
  }
  return {
    notes,
    setNotes,
    selectedNoteIndex,
    handleNoteSelect
  }
}
