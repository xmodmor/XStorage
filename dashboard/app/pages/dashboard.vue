<script setup lang="ts">
import { Users, AppWindow, HardDrive, TrendingUp, Clock } from 'lucide-vue-next'
import type { UsersListResponse } from '@/types/user'
import type { App } from '@/types/app'

definePageMeta({ layout: 'default' })

const { setBreadcrumbItems } = useBreadcrumb()

setBreadcrumbItems([
  { label: 'Dashboard', isActive: true },
])

const { apiFetch } = useApi()

const { data: usersData, pending: usersPending } = await useAsyncData('dashboard-users', async () => {
  const res = await apiFetch<UsersListResponse>('/api/v1/users?page=1&limit=1')
  return res.data
})

const { data: appsData, pending: appsPending } = await useAsyncData('dashboard-apps', async () => {
  const res = await apiFetch<App[]>('/api/v1/apps')
  return res.data
})

const totalUsers = computed(() => usersData.value?.total ?? 0)
const totalApps = computed(() => appsData.value?.length ?? 0)
const recentApps = computed(() => (appsData.value ?? []).slice(0, 5))

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric',
  })
}
</script>

<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Dashboard</h1>
      <p class="text-muted-foreground mt-1">Overview of your XStorage system</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid gap-4 md:grid-cols-3">
      <!-- Users -->
      <Card>
        <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle class="text-sm font-medium">Total Users</CardTitle>
          <div class="rounded-md bg-blue-500/10 p-2 text-blue-600 dark:text-blue-400">
            <Users class="h-4 w-4" />
          </div>
        </CardHeader>
        <CardContent>
          <div class="text-2xl font-bold">
            <span v-if="usersPending" class="text-muted-foreground animate-pulse">—</span>
            <span v-else>{{ totalUsers }}</span>
          </div>
          <p class="text-xs text-muted-foreground mt-1">Registered accounts</p>
        </CardContent>
      </Card>

      <!-- Apps -->
      <Card>
        <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle class="text-sm font-medium">Total Apps</CardTitle>
          <div class="rounded-md bg-violet-500/10 p-2 text-violet-600 dark:text-violet-400">
            <AppWindow class="h-4 w-4" />
          </div>
        </CardHeader>
        <CardContent>
          <div class="text-2xl font-bold">
            <span v-if="appsPending" class="text-muted-foreground animate-pulse">—</span>
            <span v-else>{{ totalApps }}</span>
          </div>
          <p class="text-xs text-muted-foreground mt-1">Connected applications</p>
        </CardContent>
      </Card>

      <!-- Storage -->
      <Card>
        <CardHeader class="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle class="text-sm font-medium">Storage</CardTitle>
          <div class="rounded-md bg-emerald-500/10 p-2 text-emerald-600 dark:text-emerald-400">
            <HardDrive class="h-4 w-4" />
          </div>
        </CardHeader>
        <CardContent>
          <div class="text-2xl font-bold">—</div>
          <p class="text-xs text-muted-foreground mt-1">Via API key authenticated routes</p>
        </CardContent>
      </Card>
    </div>

    <!-- Recent Apps -->
    <Card>
      <CardHeader class="flex flex-row items-center justify-between">
        <div>
          <CardTitle class="text-base">Recent Apps</CardTitle>
          <CardDescription>Your latest connected applications</CardDescription>
        </div>
        <NuxtLink to="/dashboard/apps">
          <Button variant="outline" size="sm">
            View all
          </Button>
        </NuxtLink>
      </CardHeader>
      <CardContent>
        <div v-if="appsPending" class="space-y-3">
          <div v-for="i in 3" :key="i" class="h-10 rounded-md bg-muted animate-pulse" />
        </div>

        <div v-else-if="recentApps.length === 0" class="py-8 text-center text-muted-foreground">
          <AppWindow class="mx-auto h-8 w-8 mb-2 opacity-40" />
          <p class="text-sm">No apps yet.</p>
          <NuxtLink to="/dashboard/apps">
            <Button variant="link" size="sm" class="mt-1">Create your first app</Button>
          </NuxtLink>
        </div>

        <div v-else class="divide-y">
          <div
            v-for="app in recentApps"
            :key="app.id"
            class="flex items-center justify-between py-3"
          >
            <div class="flex items-center gap-3">
              <div class="flex h-8 w-8 items-center justify-center rounded-md border bg-muted">
                <AppWindow class="h-4 w-4 text-muted-foreground" />
              </div>
              <div>
                <p class="text-sm font-medium leading-none">{{ app.name }}</p>
                <p class="text-xs text-muted-foreground mt-0.5 flex items-center gap-1">
                  <Clock class="h-3 w-3" />
                  {{ formatDate(app.created_at) }}
                </p>
              </div>
            </div>
            <NuxtLink :to="`/dashboard/apps`">
              <Button variant="ghost" size="sm">View</Button>
            </NuxtLink>
          </div>
        </div>
      </CardContent>
    </Card>

    <!-- Quick Links -->
    <div class="grid gap-4 md:grid-cols-2">
      <NuxtLink to="/dashboard/users">
        <Card class="group cursor-pointer transition-colors hover:border-primary/50">
          <CardContent class="flex items-center gap-4 pt-6">
            <div class="rounded-lg bg-blue-500/10 p-3 text-blue-600 dark:text-blue-400 group-hover:bg-blue-500/20 transition-colors">
              <Users class="h-5 w-5" />
            </div>
            <div>
              <p class="font-medium">Manage Users</p>
              <p class="text-sm text-muted-foreground">Create, edit and remove users</p>
            </div>
            <TrendingUp class="ml-auto h-4 w-4 text-muted-foreground" />
          </CardContent>
        </Card>
      </NuxtLink>

      <NuxtLink to="/dashboard/apps">
        <Card class="group cursor-pointer transition-colors hover:border-primary/50">
          <CardContent class="flex items-center gap-4 pt-6">
            <div class="rounded-lg bg-violet-500/10 p-3 text-violet-600 dark:text-violet-400 group-hover:bg-violet-500/20 transition-colors">
              <AppWindow class="h-5 w-5" />
            </div>
            <div>
              <p class="font-medium">Manage Apps</p>
              <p class="text-sm text-muted-foreground">Configure apps and API keys</p>
            </div>
            <TrendingUp class="ml-auto h-4 w-4 text-muted-foreground" />
          </CardContent>
        </Card>
      </NuxtLink>
    </div>
  </div>
</template>
