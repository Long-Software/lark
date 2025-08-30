'use client'

import { useAtomValue, useSetAtom } from 'jotai'

import { DeleteNote, ListNotes, SaveNoteDialog } from '@wailsjs/go/main/App'

import { notesAtom, selectedNoteAtom, selectedNoteIndexAtom } from './notes'

export const useNote = () => {
  const setNotes = useSetAtom(notesAtom)
  const selectedNote = useAtomValue(selectedNoteAtom)
  const setSelectedNoteIndex = useSetAtom(selectedNoteIndexAtom)

  const handleCreateNewNote = async () => {
    await SaveNoteDialog().then(_ => ListNotes().then(setNotes))
  }
  const handleDeleteNote = async () => {
    if (selectedNote?.title) {
      setSelectedNoteIndex(null)
      await DeleteNote(selectedNote.title).then(_ => {
        ListNotes().then(setNotes)
      })
    }
  }
  return {
    selectedNote,
    handleCreateNewNote,
    handleDeleteNote
  }
}
