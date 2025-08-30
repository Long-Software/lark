'use client'

import { AUTO_SAVING_TIME } from '@/constants/note'
import { useError } from '@/hooks/useError'
import { useNote } from '@/hooks/useNote'
import {
  headingsPlugin,
  listsPlugin,
  markdownShortcutPlugin,
  MDXEditor,
  MDXEditorMethods,
  quotePlugin
} from '@mdxeditor/editor'
import { throttle } from 'lodash'
import { useEffect, useRef } from 'react'

import { WriteNote } from '@wailsjs/go/main/App'

const NoteMarkdownEditor = () => {
  const editorRef = useRef<MDXEditorMethods>(null)
  const { selectedNote } = useNote()
  const { handle } = useError()

  useEffect(() => {
    if (editorRef.current && selectedNote) {
      editorRef.current.setMarkdown(selectedNote.content)
    }
  }, [selectedNote])

  if (!selectedNote) return null

  const handleAutoSaving = throttle(
    async () => {
      try {
        if (editorRef.current && selectedNote.title) {
          await WriteNote(selectedNote.title, editorRef.current.getMarkdown())
        }
      } catch (error) {
        handle(error)
      }
    },
    AUTO_SAVING_TIME,
    {
      leading: false,
      trailing: true
    }
  )

  // handler for when a new note is selected and the previous note have not been save
  const handleBlur = async () => {
    if (!selectedNote) return
    handleAutoSaving.cancel()
    const content = editorRef.current?.getMarkdown()
    if (content != null && selectedNote.title) await WriteNote(selectedNote.title, content)
  }

  return (
    <MDXEditor
      ref={editorRef}
      markdown={selectedNote.content}
      onChange={handleAutoSaving}
      onBlur={handleBlur}
      plugins={[headingsPlugin(), listsPlugin(), quotePlugin(), markdownShortcutPlugin()]}
      contentEditableClassName='outline-none min-h-screen max-w-none text-lg px-8 py-5 caret-yellow-500 prose prose-invert prose-p:my-3 prose-p:leading-relaxed prose-headings:my-4 prose-blockquote:my-4 prose-ul:my-2 prose-li:my-0 prose-code:px-1 prose-code:text-red-500 prose-code:before:content-[""] prose-code:after:content-[""]'
    />
  )
}

export default NoteMarkdownEditor
