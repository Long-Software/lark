import {
  FileLock,
  HomeIcon,
  IconType,
  LayoutGrid,
  LucideMessageCircleWarning,
  NotebookIcon,
  SearchCode,
  Timer
} from '@longslark/ui/components'

type SidebarItem = {
  title: string
  url: string
  icon: IconType
}
type SidebarGroup = {
  title: string
  items: SidebarItem[]
}
export const groups: SidebarGroup[] = [
  {
    title: 'Application',
    items: [
      { title: 'Home', url: '/', icon: HomeIcon },
      { title: 'Note', url: '/note/', icon: NotebookIcon },
      { title: 'Password', url: '/password/', icon: FileLock },
      { title: 'Catalog', url: '/catalog/', icon: LayoutGrid },
      { title: 'Fetcher', url: '/fetcher/', icon: SearchCode },
      { title: 'Timer', url: '/timer/', icon: Timer }
    ]
  },
  {
    title: 'Logs',
    items: [{ title: 'Error Logger', url: '/logs/errors', icon: LucideMessageCircleWarning }]
  }
]
