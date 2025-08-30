import { atom } from 'jotai'

import { note } from '@wailsjs/go/models'

export const notesAtom = atom<note.Info[]>([])

export const selectedNoteIndexAtom = atom<number | null>(null)

export const selectedNoteContentAtom = atom<string>('')

export const selectedNoteAtom = atom(get => {
  const notes = get(notesAtom)
  const selectedNoteIndex = get(selectedNoteIndexAtom)

  if (selectedNoteIndex == null) return null
  const selectedNote = notes[selectedNoteIndex]
  const selectedNoteContent = get(selectedNoteContentAtom)

  return {
    ...selectedNote,
    content: selectedNoteContent
  }
})


