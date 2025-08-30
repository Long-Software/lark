import { groups } from '@/configs/sidebar'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem
} from '@longslark/ui/components/sidebar'

const AppSidebar = () => {
  return (
    <Sidebar>
      <SidebarHeader>Everia</SidebarHeader>
      <SidebarContent>
        {groups.map(group => {
          return (
            <SidebarGroup key={group.title}>
              <SidebarGroupLabel className='text-lg font-bold'>{group.title}</SidebarGroupLabel>
              <SidebarGroupContent>
                <SidebarMenu>
                  {group.items.map(item => (
                    <SidebarMenuItem key={item.title}>
                      <SidebarMenuButton asChild>
                        <a href={item.url}>
                          <item.icon />
                          <span>{item.title}</span>
                        </a>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  ))}
                </SidebarMenu>
              </SidebarGroupContent>
            </SidebarGroup>
          )
        })}
      </SidebarContent>
    </Sidebar>
  )
}

export default AppSidebar
