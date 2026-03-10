<script setup lang="ts">
import { ChevronRight, MoreHorizontal } from "lucide-vue-next"
import type { NavItem } from "@/types/nav"

const props = defineProps<{
  item: NavItem
  showDropdown?: boolean
}>()

const route = useRoute()

const isActive = computed(() => {
  if (!props.item.url) return false
  // exact match for root-level routes, prefix match for nested
  return route.path === props.item.url || route.path.startsWith(props.item.url + '/')
})
</script>

<template>
  <template>
    <!-- Link Type -->
    <SidebarMenuItem v-if="item.type === 'link'">
      <SidebarMenuButton 
        :tooltip="item.disabled ? `${item.title} (disabled)` : item.title" 
        :as-child="!item.disabled"
        :disabled="item.disabled"
        :is-active="isActive"
      >
        <NuxtLink 
          v-if="!item.disabled"
          :to="item.url" 
          class="flex items-center gap-2"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </NuxtLink>
        <div 
          v-else
          class="flex items-center gap-2 cursor-not-allowed opacity-50"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </div>
      </SidebarMenuButton>
    </SidebarMenuItem>

    <!-- Item Type (Simple) -->
    <SidebarMenuItem v-else-if="item.type === 'item'">
      <SidebarMenuButton 
        :tooltip="item.disabled ? `${item.title} (disabled)` : item.title" 
        :as-child="!item.disabled"
        :disabled="item.disabled"
        :is-active="isActive"
      >
        <NuxtLink 
          v-if="!item.disabled"
          :to="item.url" 
          class="flex items-center gap-2"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </NuxtLink>
        <div 
          v-else
          class="flex items-center gap-2 cursor-not-allowed opacity-50"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </div>
      </SidebarMenuButton>
    </SidebarMenuItem>

    <!-- Collapse Type -->
    <Collapsible
      v-else-if="item.type === 'collapse' && item.children?.length"
      as-child
      :default-open="item.isActive"
      class="group/collapsible"
    >
      <SidebarMenuItem>
        <CollapsibleTrigger as-child>
          <SidebarMenuButton :tooltip="item.title">
            <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
            <span>{{ item.title }}</span>
            <div v-if="item.badge" class="ml-auto mr-2">
              <Badge :variant="item.badge.color as any">
                {{ item.badge.label }}
              </Badge>
            </div>
            <ChevronRight class="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
          </SidebarMenuButton>
        </CollapsibleTrigger>
        <CollapsibleContent>
          <SidebarMenuSub>
            <SidebarMenuSubItem v-for="subItem in item.children" :key="subItem.title">
              <template>
                <SidebarMenuSubButton as-child>
                  <NuxtLink :to="subItem.url" class="flex items-center gap-2">
                    <component :is="subItem.icon" v-if="subItem.icon" class="h-4 w-4" />
                    <span>{{ subItem.title }}</span>
                    <div v-if="subItem.badge" class="ml-auto">
                      <Badge :variant="subItem.badge.color as any">
                        {{ subItem.badge.label }}
                      </Badge>
                    </div>
                  </NuxtLink>
                </SidebarMenuSubButton>
              </template>
            </SidebarMenuSubItem>
          </SidebarMenuSub>
        </CollapsibleContent>
      </SidebarMenuItem>
    </Collapsible>

    <!-- Dropdown Type -->
    <SidebarMenuItem v-else-if="item.type === 'dropdown'">
      <SidebarMenuButton 
        :tooltip="item.disabled ? `${item.title} (disabled)` : item.title" 
        :as-child="!item.disabled"
        :disabled="item.disabled"
      >
        <NuxtLink 
          v-if="!item.disabled"
          :to="item.url" 
          class="flex items-center gap-2"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </NuxtLink>
        <div 
          v-else
          class="flex items-center gap-2 cursor-not-allowed opacity-50"
        >
          <component :is="item.icon" v-if="item.icon" class="h-4 w-4" />
          <span>{{ item.title }}</span>
          <div v-if="item.badge" class="ml-auto">
            <Badge :variant="item.badge.color as any">
              {{ item.badge.label }}
            </Badge>
          </div>
        </div>
      </SidebarMenuButton>
      
      <!-- Dropdown Menu (hidden when disabled) -->
      <DropdownMenu v-if="!item.disabled && showDropdown && item.children?.length">
        <DropdownMenuTrigger as-child>
          <SidebarMenuAction show-on-hover>
            <MoreHorizontal class="h-4 w-4" />
            <span class="sr-only">More actions</span>
          </SidebarMenuAction>
        </DropdownMenuTrigger>
        <DropdownMenuContent class="w-48 rounded-lg" side="right" align="start">
          <template v-for="subItem in item.children" :key="subItem.title">
            <template>
              <DropdownMenuItem as-child :disabled="subItem.disabled">
                <NuxtLink 
                  v-if="!subItem.disabled"
                  :to="subItem.url" 
                  class="flex items-center gap-2"
                >
                  <component :is="subItem.icon" v-if="subItem.icon" class="h-4 w-4 text-muted-foreground" />
                  <span>{{ subItem.title }}</span>
                  <div v-if="subItem.badge" class="ml-auto">
                    <Badge :variant="subItem.badge.color as any" class="text-xs">
                      {{ subItem.badge.label }}
                    </Badge>
                  </div>
                </NuxtLink>
                <div 
                  v-else
                  class="flex items-center gap-2 cursor-not-allowed opacity-50"
                >
                  <component :is="subItem.icon" v-if="subItem.icon" class="h-4 w-4 text-muted-foreground" />
                  <span>{{ subItem.title }}</span>
                  <div v-if="subItem.badge" class="ml-auto">
                    <Badge :variant="subItem.badge.color as any" class="text-xs">
                      {{ subItem.badge.label }}
                    </Badge>
                  </div>
                </div>
              </DropdownMenuItem>
            </template>
          </template>
        </DropdownMenuContent>
      </DropdownMenu>
    </SidebarMenuItem>
  </template>
</template>
