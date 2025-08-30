import { ReactNode } from 'react'

import { Toaster } from '@longslark/ui/components/sonner'

type ToasterProviderProps = {
  children: ReactNode
}

const ToasterProvider = ({ children }: ToasterProviderProps) => {
  return (
    <>
      <Toaster position='top-center' />
      {children}
    </>
  )
}

export default ToasterProvider
