<script setup lang="ts">
import { useNavigation } from "@/composables/useNavigation"

const props = withDefaults(defineProps<{
  collapsible?: "offcanvas" | "icon" | "none"
}>(), {
  collapsible: "icon",
})

// Navigation composable for role-based navigation
const { getNavigationItems } = useNavigation()

// Get navigation items based on user role
const navItems = computed(() => {
  return getNavigationItems()
})

</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <!-- OMXE Logo - Responsive to sidebar state -->
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton 
            size="lg" 
            as-child 
            class="hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
          >
            <NuxtLink 
              to="/" 
              class="flex items-center justify-start w-full"
            >
              <!-- Logo Full - Shows in expanded state -->
              <div class="flex items-center justify-center transition-all duration-200 ease-in-out group-data-[collapsible=icon]:hidden">
                <AppLogo 
                  mode="full"
                  size="md"
                  :link-to-home="false"
                />
              </div>
              <!-- Logo Icon - Shows in collapsed state -->
              <div class="hidden items-center justify-center transition-all duration-200 ease-in-out group-data-[collapsible=icon]:flex">
                <AppLogo 
                  mode="icon"
                  size="md"
                  :link-to-home="false"
                />
              </div>
            </NuxtLink>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
      
      <!-- Divider for visual separation -->
      <div class="border-b my-2" />

    </SidebarHeader>
    <SidebarContent>
      <NavGroup :items="navItems" :show-dropdowns="true" />
    </SidebarContent>
    <SidebarFooter>
      <NavUser />
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
