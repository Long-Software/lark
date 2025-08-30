import { toast } from '@longslark/ui/components/sonner'

export const useError = () => {
  const handle = (error: unknown) => {
    if (error instanceof Error) {
      toast.error(error.message)
    } else {
      toast.error('An unknown error occurred.')
    }
  }
  return {
    handle
  }
}
