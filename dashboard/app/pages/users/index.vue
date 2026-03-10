<script setup lang="ts">
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { Search, Plus, Pencil, Trash2, ChevronLeft, ChevronRight, UserRound } from 'lucide-vue-next'
import { toast } from 'vue-sonner'
import type { User } from '@/types/user'

definePageMeta({ layout: 'default' })

const { setBreadcrumbItems } = useBreadcrumb()
setBreadcrumbItems([
  { label: 'Dashboard', href: '/dashboard' },
  { label: 'Users', isActive: true },
])

const { fetchUsers, createUser, updateUser, deleteUser } = useUsers()

// --- State ---
const users = ref<User[]>([])
const total = ref(0)
const page = ref(1)
const limit = 15
const search = ref('')
const loading = ref(false)

// Dialogs
const showCreateDialog = ref(false)
const showEditDialog = ref(false)
const showDeleteDialog = ref(false)
const selectedUser = ref<User | null>(null)
const deleteLoading = ref(false)

// --- Data Fetching ---
const load = async () => {
  loading.value = true
  try {
    const res = await fetchUsers(page.value, limit, search.value)
    if (res.success && res.data) {
      users.value = res.data.users
      total.value = res.data.total
    }
  } catch {
    toast.error('Failed to load users')
  } finally {
    loading.value = false
  }
}

const totalPages = computed(() => Math.ceil(total.value / limit))

// Debounced search
let searchTimer: ReturnType<typeof setTimeout>
watch(search, () => {
  clearTimeout(searchTimer)
  searchTimer = setTimeout(() => {
    page.value = 1
    load()
  }, 400)
})

watch(page, load)

onMounted(load)

// --- Create Form ---
const createSchema = toTypedSchema(z.object({
  email: z.string().email('Please enter a valid email'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
}))

const createForm = useForm({ validationSchema: createSchema })

const onCreateSubmit = createForm.handleSubmit(async (values) => {
  try {
    const res = await createUser(values)
    if (res.success) {
      toast.success('User created successfully')
      showCreateDialog.value = false
      createForm.resetForm()
      await load()
    } else {
      toast.error(res.error?.message ?? 'Failed to create user')
    }
  } catch {
    toast.error('Failed to create user')
  }
})

// --- Edit Form ---
const editSchema = toTypedSchema(z.object({
  email: z.string().email('Please enter a valid email'),
  password: z.string().min(6, 'Password must be at least 6 characters').optional().or(z.literal('')),
}))

const editForm = useForm({ validationSchema: editSchema })

const openEditDialog = (user: User) => {
  selectedUser.value = user
  editForm.setValues({ email: user.email, password: '' })
  showEditDialog.value = true
}

const onEditSubmit = editForm.handleSubmit(async (values) => {
  if (!selectedUser.value) return
  try {
    const payload: { email: string; password?: string } = { email: values.email }
    if (values.password) payload.password = values.password
    const res = await updateUser(selectedUser.value.id, payload)
    if (res.success) {
      toast.success('User updated successfully')
      showEditDialog.value = false
      await load()
    } else {
      toast.error(res.error?.message ?? 'Failed to update user')
    }
  } catch {
    toast.error('Failed to update user')
  }
})

// --- Delete ---
const openDeleteDialog = (user: User) => {
  selectedUser.value = user
  showDeleteDialog.value = true
}

const onDelete = async () => {
  if (!selectedUser.value) return
  deleteLoading.value = true
  try {
    const res = await deleteUser(selectedUser.value.id)
    if (res.success) {
      toast.success('User deleted successfully')
      showDeleteDialog.value = false
      if (users.value.length === 1 && page.value > 1) page.value--
      else await load()
    } else {
      toast.error(res.error?.message ?? 'Failed to delete user')
    }
  } catch {
    toast.error('Failed to delete user')
  } finally {
    deleteLoading.value = false
  }
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric',
  })
}

// Field components
const [CreateEmail, createEmailAttrs] = createForm.defineField('email')
const [CreatePassword, createPasswordAttrs] = createForm.defineField('password')
const [EditEmail, editEmailAttrs] = editForm.defineField('email')
const [EditPassword, editPasswordAttrs] = editForm.defineField('password')
</script>

<template>
  <div class="space-y-6">
    <!-- Page Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold tracking-tight">Users</h1>
        <p class="text-muted-foreground mt-1">
          Manage system users — {{ total }} total
        </p>
      </div>
      <Button @click="showCreateDialog = true">
        <Plus class="mr-2 h-4 w-4" />
        Add User
      </Button>
    </div>

    <!-- Search -->
    <div class="relative max-w-sm">
      <Search class="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
      <Input
        v-model="search"
        placeholder="Search by email..."
        class="pl-9"
      />
    </div>

    <!-- Table Card -->
    <Card>
      <CardContent class="p-0">
        <!-- Loading skeleton -->
        <div v-if="loading" class="p-4 space-y-3">
          <div v-for="i in 5" :key="i" class="flex items-center gap-4">
            <div class="h-9 w-9 rounded-full bg-muted animate-pulse" />
            <div class="flex-1 space-y-1.5">
              <div class="h-4 w-48 rounded bg-muted animate-pulse" />
              <div class="h-3 w-32 rounded bg-muted animate-pulse" />
            </div>
            <div class="h-8 w-20 rounded bg-muted animate-pulse" />
          </div>
        </div>

        <!-- Empty state -->
        <div v-else-if="users.length === 0" class="flex flex-col items-center justify-center py-16 text-muted-foreground">
          <UserRound class="h-10 w-10 mb-3 opacity-30" />
          <p class="text-sm font-medium">No users found</p>
          <p class="text-xs mt-1">
            <span v-if="search">Try a different search term</span>
            <span v-else>Get started by adding the first user</span>
          </p>
        </div>

        <!-- Data Table -->
        <Table v-else>
          <TableHeader>
            <TableRow>
              <TableHead class="w-12">#</TableHead>
              <TableHead>User</TableHead>
              <TableHead>Created</TableHead>
              <TableHead class="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="user in users" :key="user.id" class="group">
              <TableCell class="text-muted-foreground font-mono text-xs">
                {{ user.id }}
              </TableCell>
              <TableCell>
                <div class="flex items-center gap-3">
                  <Avatar class="h-8 w-8">
                    <AvatarFallback class="text-xs bg-primary/10 text-primary font-semibold">
                      {{ user.email.slice(0, 2).toUpperCase() }}
                    </AvatarFallback>
                  </Avatar>
                  <span class="font-medium text-sm">{{ user.email }}</span>
                </div>
              </TableCell>
              <TableCell class="text-muted-foreground text-sm">
                {{ formatDate(user.created_at) }}
              </TableCell>
              <TableCell class="text-right">
                <div class="flex items-center justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                  <Button
                    variant="ghost"
                    size="icon"
                    class="h-8 w-8"
                    @click="openEditDialog(user)"
                  >
                    <Pencil class="h-3.5 w-3.5" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    class="h-8 w-8 text-destructive hover:text-destructive"
                    @click="openDeleteDialog(user)"
                  >
                    <Trash2 class="h-3.5 w-3.5" />
                  </Button>
                </div>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>

      <!-- Pagination -->
      <CardFooter v-if="!loading && totalPages > 1" class="border-t px-4 py-3 flex items-center justify-between">
        <p class="text-xs text-muted-foreground">
          Showing {{ (page - 1) * limit + 1 }}–{{ Math.min(page * limit, total) }} of {{ total }}
        </p>
        <div class="flex items-center gap-1">
          <Button
            variant="outline"
            size="icon"
            class="h-7 w-7"
            :disabled="page <= 1"
            @click="page--"
          >
            <ChevronLeft class="h-4 w-4" />
          </Button>
          <span class="text-xs px-2">{{ page }} / {{ totalPages }}</span>
          <Button
            variant="outline"
            size="icon"
            class="h-7 w-7"
            :disabled="page >= totalPages"
            @click="page++"
          >
            <ChevronRight class="h-4 w-4" />
          </Button>
        </div>
      </CardFooter>
    </Card>

    <!-- ─── Create Dialog ─── -->
    <Dialog v-model:open="showCreateDialog">
      <DialogContent class="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Add New User</DialogTitle>
          <DialogDescription>Create a new system user account.</DialogDescription>
        </DialogHeader>
        <form class="space-y-4 pt-2" @submit.prevent="onCreateSubmit">
          <FormField name="email">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input
                  v-model="CreateEmail"
                  v-bind="createEmailAttrs"
                  type="email"
                  placeholder="user@example.com"
                  autocomplete="off"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField name="password">
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input
                  v-model="CreatePassword"
                  v-bind="createPasswordAttrs"
                  type="password"
                  placeholder="Min. 6 characters"
                  autocomplete="new-password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <DialogFooter class="pt-2">
            <Button type="button" variant="outline" @click="showCreateDialog = false">Cancel</Button>
            <Button type="submit" :disabled="createForm.isSubmitting.value">
              <span v-if="createForm.isSubmitting.value">Creating…</span>
              <span v-else>Create User</span>
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>

    <!-- ─── Edit Dialog ─── -->
    <Dialog v-model:open="showEditDialog">
      <DialogContent class="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Edit User</DialogTitle>
          <DialogDescription>Update email or set a new password.</DialogDescription>
        </DialogHeader>
        <form class="space-y-4 pt-2" @submit.prevent="onEditSubmit">
          <FormField name="email">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input
                  v-model="EditEmail"
                  v-bind="editEmailAttrs"
                  type="email"
                  placeholder="user@example.com"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField name="password">
            <FormItem>
              <FormLabel>New Password <span class="text-muted-foreground font-normal">(optional)</span></FormLabel>
              <FormControl>
                <Input
                  v-model="EditPassword"
                  v-bind="editPasswordAttrs"
                  type="password"
                  placeholder="Leave blank to keep current"
                  autocomplete="new-password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <DialogFooter class="pt-2">
            <Button type="button" variant="outline" @click="showEditDialog = false">Cancel</Button>
            <Button type="submit" :disabled="editForm.isSubmitting.value">
              <span v-if="editForm.isSubmitting.value">Saving…</span>
              <span v-else>Save Changes</span>
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>

    <!-- ─── Delete Alert ─── -->
    <AlertDialog v-model:open="showDeleteDialog">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Delete User</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete
            <span class="font-semibold text-foreground">{{ selectedUser?.email }}</span>?
            This action cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            class="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            :disabled="deleteLoading"
            @click.prevent="onDelete"
          >
            <span v-if="deleteLoading">Deleting…</span>
            <span v-else>Delete</span>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  </div>
</template>
