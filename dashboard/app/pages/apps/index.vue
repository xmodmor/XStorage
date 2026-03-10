<script setup lang="ts">
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { Plus, Trash2, AppWindow, Key, Copy, Eye, EyeOff, MoreVertical } from 'lucide-vue-next'
import { toast } from 'vue-sonner'
import type { App, APIKey } from '@/types/app'

definePageMeta({ layout: 'default' })

const { setBreadcrumbItems } = useBreadcrumb()
setBreadcrumbItems([
  { label: 'Dashboard', href: '/dashboard' },
  { label: 'Apps', isActive: true },
])

const { fetchApps, createApp, deleteApp, fetchApiKeys, createApiKey, deleteApiKey } = useApps()

// --- State ---
const apps = ref<App[]>([])
const loading = ref(false)

const showCreateDialog = ref(false)
const showDeleteDialog = ref(false)
const selectedApp = ref<App | null>(null)
const deleteLoading = ref(false)

// API Keys panel
const activeAppId = ref<number | null>(null)
const apiKeys = ref<Record<number, APIKey[]>>({})
const keysLoading = ref<Record<number, boolean>>({})
const newKeyVisible = ref<Record<number, boolean>>({})

// --- Load ---
const load = async () => {
  loading.value = true
  try {
    const res = await fetchApps()
    if (res.success && res.data) {
      apps.value = res.data
    }
  } catch {
    toast.error('Failed to load apps')
  } finally {
    loading.value = false
  }
}

onMounted(load)

// --- Create Form ---
const createSchema = toTypedSchema(z.object({
  name: z.string().min(1, 'App name is required').max(100, 'Name too long'),
}))
const createForm = useForm({ validationSchema: createSchema })
const [AppName, appNameAttrs] = createForm.defineField('name')

const onCreateSubmit = createForm.handleSubmit(async (values) => {
  try {
    const res = await createApp({ name: values.name })
    if (res.success) {
      toast.success('App created successfully')
      showCreateDialog.value = false
      createForm.resetForm()
      await load()
    } else {
      toast.error(res.error?.message ?? 'Failed to create app')
    }
  } catch {
    toast.error('Failed to create app')
  }
})

// --- Delete App ---
const openDeleteDialog = (app: App) => {
  selectedApp.value = app
  showDeleteDialog.value = true
}

const onDeleteApp = async () => {
  if (!selectedApp.value) return
  deleteLoading.value = true
  try {
    const res = await deleteApp(selectedApp.value.id)
    if (res.success) {
      toast.success('App deleted')
      showDeleteDialog.value = false
      await load()
    } else {
      toast.error(res.error?.message ?? 'Failed to delete app')
    }
  } catch {
    toast.error('Failed to delete app')
  } finally {
    deleteLoading.value = false
  }
}

// --- API Keys ---
const toggleApiKeys = async (appId: number) => {
  if (activeAppId.value === appId) {
    activeAppId.value = null
    return
  }
  activeAppId.value = appId
  if (!apiKeys.value[appId]) {
    await loadApiKeys(appId)
  }
}

const loadApiKeys = async (appId: number) => {
  keysLoading.value[appId] = true
  try {
    const res = await fetchApiKeys(appId)
    if (res.success && res.data) {
      apiKeys.value[appId] = res.data
    }
  } catch {
    toast.error('Failed to load API keys')
  } finally {
    keysLoading.value[appId] = false
  }
}

const onCreateApiKey = async (appId: number) => {
  try {
    const res = await createApiKey(appId)
    if (res.success && res.data) {
      if (!apiKeys.value[appId]) apiKeys.value[appId] = []
      apiKeys.value[appId].push(res.data)
      newKeyVisible.value[res.data.id] = true
      toast.success('API key created — copy it now, it won\'t be shown again')
    }
  } catch {
    toast.error('Failed to create API key')
  }
}

const onDeleteApiKey = async (appId: number, keyId: number) => {
  try {
    const res = await deleteApiKey(appId, keyId)
    if (res.success) {
      apiKeys.value[appId] = apiKeys.value[appId].filter(k => k.id !== keyId)
      toast.success('API key deleted')
    }
  } catch {
    toast.error('Failed to delete API key')
  }
}

const copyKey = (key: string) => {
  navigator.clipboard.writeText(key)
  toast.success('Copied to clipboard')
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric',
  })
}
</script>

<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold tracking-tight">Apps</h1>
        <p class="text-muted-foreground mt-1">
          Manage your applications and their API keys
        </p>
      </div>
      <Button @click="showCreateDialog = true">
        <Plus class="mr-2 h-4 w-4" />
        New App
      </Button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
      <div v-for="i in 3" :key="i" class="h-40 rounded-lg border bg-muted animate-pulse" />
    </div>

    <!-- Empty State -->
    <div v-else-if="apps.length === 0" class="flex flex-col items-center justify-center rounded-lg border border-dashed py-16 text-muted-foreground">
      <AppWindow class="h-12 w-12 mb-4 opacity-30" />
      <p class="font-medium">No apps yet</p>
      <p class="text-sm mt-1">Create your first app to get an API key for storage access</p>
      <Button class="mt-4" @click="showCreateDialog = true">
        <Plus class="mr-2 h-4 w-4" />
        Create App
      </Button>
    </div>

    <!-- Apps Grid -->
    <div v-else class="space-y-4">
      <Card
        v-for="app in apps"
        :key="app.id"
        class="overflow-hidden"
      >
        <!-- App Header -->
        <CardHeader class="pb-3">
          <div class="flex items-start justify-between">
            <div class="flex items-center gap-3">
              <div class="flex h-10 w-10 items-center justify-center rounded-lg border bg-muted">
                <AppWindow class="h-5 w-5 text-muted-foreground" />
              </div>
              <div>
                <CardTitle class="text-base">{{ app.name }}</CardTitle>
                <CardDescription class="text-xs mt-0.5">
                  ID: {{ app.id }} · Created {{ formatDate(app.created_at) }}
                </CardDescription>
              </div>
            </div>

            <DropdownMenu>
              <DropdownMenuTrigger as-child>
                <Button variant="ghost" size="icon" class="h-8 w-8 -mr-1">
                  <MoreVertical class="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end">
                <DropdownMenuItem
                  class="text-destructive focus:text-destructive"
                  @click="openDeleteDialog(app)"
                >
                  <Trash2 class="mr-2 h-4 w-4" />
                  Delete App
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </CardHeader>

        <!-- API Keys Toggle -->
        <CardContent class="pt-0 space-y-3">
          <Button
            variant="outline"
            size="sm"
            class="w-full justify-between"
            @click="toggleApiKeys(app.id)"
          >
            <span class="flex items-center gap-2">
              <Key class="h-3.5 w-3.5" />
              API Keys
              <Badge variant="secondary" class="h-4 px-1.5 text-xs">
                {{ apiKeys[app.id]?.length ?? '—' }}
              </Badge>
            </span>
            <span class="text-muted-foreground text-xs">
              {{ activeAppId === app.id ? 'Hide' : 'Manage' }}
            </span>
          </Button>

          <!-- Keys Panel -->
          <div v-if="activeAppId === app.id" class="space-y-2">
            <!-- Loading -->
            <div v-if="keysLoading[app.id]" class="space-y-2">
              <div v-for="i in 2" :key="i" class="h-10 rounded-md bg-muted animate-pulse" />
            </div>

            <!-- Empty keys -->
            <div v-else-if="!apiKeys[app.id]?.length" class="rounded-md border border-dashed p-4 text-center text-sm text-muted-foreground">
              No API keys — create one to access storage
            </div>

            <!-- Key list -->
            <div v-else class="space-y-2">
              <div
                v-for="key in apiKeys[app.id]"
                :key="key.id"
                class="flex items-center gap-2 rounded-md border bg-muted/50 px-3 py-2"
              >
                <Key class="h-3.5 w-3.5 shrink-0 text-muted-foreground" />
                <code class="flex-1 truncate font-mono text-xs">
                  <span v-if="newKeyVisible[key.id]">{{ key.access_key }}</span>
                  <span v-else>{{ key.access_key.slice(0, 8) }}••••••••••••</span>
                </code>
                <Button
                  variant="ghost"
                  size="icon"
                  class="h-6 w-6 shrink-0"
                  :title="newKeyVisible[key.id] ? 'Hide' : 'Show'"
                  @click="newKeyVisible[key.id] = !newKeyVisible[key.id]"
                >
                  <Eye v-if="!newKeyVisible[key.id]" class="h-3 w-3" />
                  <EyeOff v-else class="h-3 w-3" />
                </Button>
                <Button
                  v-if="newKeyVisible[key.id]"
                  variant="ghost"
                  size="icon"
                  class="h-6 w-6 shrink-0"
                  title="Copy key"
                  @click="copyKey(key.access_key)"
                >
                  <Copy class="h-3 w-3" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon"
                  class="h-6 w-6 shrink-0 text-destructive hover:text-destructive"
                  title="Delete key"
                  @click="onDeleteApiKey(app.id, key.id)"
                >
                  <Trash2 class="h-3 w-3" />
                </Button>
              </div>
            </div>

            <!-- Add Key Button -->
            <Button
              variant="outline"
              size="sm"
              class="w-full"
              @click="onCreateApiKey(app.id)"
            >
              <Plus class="mr-2 h-3.5 w-3.5" />
              Generate API Key
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>

    <!-- ─── Create App Dialog ─── -->
    <Dialog v-model:open="showCreateDialog">
      <DialogContent class="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Create New App</DialogTitle>
          <DialogDescription>
            Give your app a name. You'll get an API key to access XStorage buckets.
          </DialogDescription>
        </DialogHeader>
        <form class="space-y-4 pt-2" @submit.prevent="onCreateSubmit">
          <FormField name="name">
            <FormItem>
              <FormLabel>App Name</FormLabel>
              <FormControl>
                <Input
                  v-model="AppName"
                  v-bind="appNameAttrs"
                  placeholder="e.g. My Web App"
                  autocomplete="off"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <DialogFooter class="pt-2">
            <Button type="button" variant="outline" @click="showCreateDialog = false">Cancel</Button>
            <Button type="submit" :disabled="createForm.isSubmitting.value">
              <span v-if="createForm.isSubmitting.value">Creating…</span>
              <span v-else>Create App</span>
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>

    <!-- ─── Delete App Alert ─── -->
    <AlertDialog v-model:open="showDeleteDialog">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Delete App</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete
            <span class="font-semibold text-foreground">{{ selectedApp?.name }}</span>?
            All associated API keys and buckets will be removed. This cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            class="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            :disabled="deleteLoading"
            @click.prevent="onDeleteApp"
          >
            <span v-if="deleteLoading">Deleting…</span>
            <span v-else>Delete App</span>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  </div>
</template>
