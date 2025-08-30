'use client'

import { FormEvent, useRef, useState } from 'react'
import { useError } from '@/hooks/useError'
import {
  headingsPlugin,
  listsPlugin,
  markdownShortcutPlugin,
  MDXEditor,
  MDXEditorMethods,
  quotePlugin
} from '@mdxeditor/editor'

import { Button } from '@longslark/ui/components/button'
import { Input } from '@longslark/ui/components/input'
import { Label } from '@longslark/ui/components/label'
import { Fetch, ReadFile } from '@wailsjs/go/main/App'

const BrowserPage = () => {
  const editorRef = useRef<MDXEditorMethods>(null)
  const [url, setUrl] = useState('')
  const { handle } = useError()

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    try {
      const filePath = await Fetch(url)
      const content = await ReadFile(filePath)
      if (editorRef.current) {
        editorRef.current.setMarkdown(content)
      }
    } catch (error) {
      handle(error)
    }
  }

  return (
    <div className=''>
      <form onSubmit={handleSubmit} className='flex w-full max-w-sm items-center space-x-2'>
        <Label htmlFor='url'>URL</Label>
        <Input
          id='url'
          onChange={e => setUrl(e.target.value)}
          value={url}
          placeholder='www.example.com'
          className='border border-foreground'
          type='url'
        />
        <Button className='w-fit bg-foreground' type='submit'>
          Fetch
        </Button>
      </form>
      <MDXEditor
        ref={editorRef}
        markdown='# Enter the URL'
        plugins={[headingsPlugin(), listsPlugin(), quotePlugin(), markdownShortcutPlugin()]}
        contentEditableClassName='outline-none min-h-screen max-w-none text-lg px-8 py-5 caret-yellow-500 prose prose-invert prose-p:my-3 prose-p:leading-relaxed prose-headings:my-4 prose-blockquote:my-4 prose-ul:my-2 prose-li:my-0 prose-code:px-1 prose-code:text-red-500 prose-code:before:content-[""] prose-code:after:content-[""]'
      />
    </div>
  )
}

export default BrowserPage
