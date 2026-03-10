<script setup lang="ts">

import { useBreadcrumb } from '@/composables/useBreadcrumb'

const { breadcrumbItems } = useBreadcrumb()



</script>

<template>
  <div
    class="flex min-h-screen flex-col">

    <SidebarProvider>
      <AppSidebar />
      <SidebarInset
        class="transition-[padding] duration-300 ease-out"
      >
        <header class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12 border-b">
          <div class="flex items-center justify-between gap-2 px-4 w-full">
            <!-- Left side: Sidebar trigger and breadcrumb -->
            <div class="flex items-center gap-2">
              <SidebarTrigger class="-ml-1" />
              <Separator orientation="vertical" class="mr-2 h-4" />
              <Breadcrumb v-if="breadcrumbItems.length > 0">
                <BreadcrumbList>
                  <template v-for="(item, index) in breadcrumbItems" :key="index">
                    <BreadcrumbItem>
                      <BreadcrumbLink v-if="item.href && !item.isActive" :href="item.href">
                        {{ item.label }}
                      </BreadcrumbLink>
                      <BreadcrumbPage v-else>
                        {{ item.label }}
                      </BreadcrumbPage>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator v-if="index < breadcrumbItems.length - 1" />
                  </template>
                </BreadcrumbList>
              </Breadcrumb>
            </div>
          </div>
        </header>
        <div class="@container/main flex flex-1 flex-col gap-2 p-4">
          <slot />
        </div>
      </SidebarInset>
    </SidebarProvider>
  </div>
</template>

<style scoped>
/* Smooth transition for sidebar when mode bar shows/hides */
:deep([data-slot="sidebar"] > div.fixed) {
  transition: top 0.3s ease-out, height 0.25s ease-in;
}
/* When mode bar is visible, push fixed sidebar below it so the bar covers full width including above sidebar */
.has-mode-bar :deep([data-slot="sidebar"] > div.fixed) {
  top: 2.5rem;
  height: calc(100svh - 2.5rem);
}
</style>
