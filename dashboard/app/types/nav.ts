import type { LucideIcon } from "lucide-vue-next"

export interface AccessRequirements {
    permissions?: string[]
    roles?: string[]
    services?: string[]
    requireAll?: boolean
    allowSuperAdmin?: boolean
  }
  
  export interface NavItem {
    group?: string
    title: string
    url?: string
    icon?: LucideIcon
    isActive?: boolean
    permissions?: string[]
    roles?: string[]
    services?: string[]
    disabled?: boolean
    accessRequirements?: AccessRequirements
    type: 'link' | 'collapse' | 'item' | 'dropdown'
    badge?: {
      label: string
      color: 'default' | 'primary' | 'secondary' | 'destructive' | 'outline' | 'warning' | 'success' | 'info' | 'error'
    }
    children?: NavItem[]
  }