import {
  LayoutDashboard,
  Users,
  AppWindow,
  HardDrive,
} from "lucide-vue-next"
import type { NavItem } from "@/types/nav"

export const useNavigation = () => {
  const items: NavItem[] = [
    {
      title: "Dashboard",
      url: "/dashboard",
      icon: LayoutDashboard,
      type: "item",
      group: "Overview",
    },
    {
      title: "Users",
      url: "/users",
      icon: Users,
      type: "item",
      group: "Management",
    },
    {
      title: "Apps",
      url: "/apps",
      icon: AppWindow,
      type: "item",
      group: "Management",
    },
    {
      title: "Storage",
      url: "/buckets",
      icon: HardDrive,
      type: "item",
      group: "Management",
      disabled: true,
      badge: { label: "API Key", color: "secondary" },
    },
  ]

  const getNavigationItems = (): NavItem[] => items

  return {
    getNavigationItems,
    items,
  }
}

/**
 * Example usage in components:
 * 
 * <script setup lang="ts">
 * const { getNavigationItems, getCurrentUserRole, getDashboardUrl } = useNavigation()
 * 
 * const navItems = computed(() => getNavigationItems())
 * const userRole = computed(() => getCurrentUserRole())
 * const dashboardUrl = computed(() => getDashboardUrl())
 * </script>
 * 
 * <template>
 *   <nav>
 *     <div>Role: {{ userRole }}</div>
 *     <NuxtLink :to="dashboardUrl">Go to Dashboard</NuxtLink>
 *     <ul>
 *       <li v-for="item in navItems" :key="item.url">
 *         <NuxtLink :to="item.url">
 *           <component :is="item.icon" v-if="item.icon" class="w-4 h-4" />
 *           {{ item.title }}
 *         </NuxtLink>
 *       </li>
 *     </ul>
 *   </nav>
 * </template>
 * 
 * Role Hierarchy:
 * - superadmin: Full system access, can manage everything
 * - admin: User/transaction management, reports, provider config
 * - seller: Process manual transfers in assigned regions
 * - api_user: Own dashboard, API keys, transactions
 * - viewer: Read-only access to system
 */
