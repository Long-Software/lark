import { ComponentProps } from 'react'

import { cn } from '@longslark/tailwind/lib'
import { note } from '@wailsjs/go/models'

export type NotePrevieProps = note.Info & {
  isActive?: boolean
} & ComponentProps<'div'>
const NotePreview = ({ title, lastEditTime, content, isActive, className, ...props }: NotePrevieProps) => {
  return (
    <div
      className={cn(
        'cursor-pointer rounded-md px-2.5 py-3',
        isActive ? 'bg-zinc-400/75' : 'hover:bg-zinc-500/75',
        className
      )}
      {...props}>
      <p className='mb-1 truncate text-xl font-bold'>{title}</p>
      <p className='mb-2 inline-block w-full text-left text-xs font-light'>{lastEditTime}</p>
    </div>
  )
}

export default NotePreview
