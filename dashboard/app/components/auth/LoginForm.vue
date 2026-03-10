<script setup lang="ts">
import { computed, type HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'
import { vAutoAnimate } from '@formkit/auto-animate/vue'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'
import { toast } from 'vue-sonner'
import { Loader2 } from 'lucide-vue-next'

import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'

import { useAuthStore } from '@/stores/useAuth'

const authStore = useAuthStore()


const route = useRoute()
const successMessage = computed(() => (route.query.message as string) || null)

const loginSchema = toTypedSchema(
  z.object({
    email: z.string().min(1, 'Email is required').email('Please enter a valid email'),
    password: z.string().min(1, 'Password is required'),
  }),
)

const { isFieldDirty, handleSubmit } = useForm({
  validationSchema: loginSchema,
})

const isLoading = computed(() => authStore.loading)

const onSubmit = handleSubmit(async (values) => {
  try {
    const redirect = route.query.redirect as string | undefined
    const success = await authStore.login(
      {
        email: values.email,
        password: values.password,
      },
      redirect,
    )

    if (success) {
      toast.success('Login successful!', {
        description: 'Welcome back to XStorage',
      })
    }
  } catch (err) {
    console.error('Login form error:', err)

    if (authStore.error) {
      toast.error('Login failed', {
        description: authStore.error,
      })
    } else {
      toast.error('Login failed', {
        description: 'Something went wrong. Please try again.',
      })
    }
  }
})

const props = defineProps<{
  class?: HTMLAttributes['class']
}>()
</script>

<template>
  <div :class="cn('flex flex-col gap-6', props.class)">
    <Card class="overflow-hidden p-0">
      <CardContent class="flex flex-col p-0">
        <form class="p-6 md:p-8" @submit="onSubmit">
          <div class="flex flex-col gap-6">
            <div class="flex flex-col items-center text-center">
              <h1 class="text-2xl font-bold">
                Welcome back
              </h1>
              <p class="text-balance text-muted-foreground">
                Sign in to your XStorage account
              </p>

              <div
                v-if="successMessage"
                class="mt-4 rounded-md border border-green-200 bg-green-50 p-3 text-sm text-green-700"
              >
                {{ successMessage }}
              </div>
            </div>

            <div class="grid gap-2">
              <FormField
                v-slot="{ componentField }"
                name="email"
                :validate-on-blur="!isFieldDirty"
              >
                <FormItem v-auto-animate>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      type="email"
                      placeholder="example@example.com"
                      autocomplete="email"
                      :disabled="isLoading"
                      v-bind="componentField"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>

            <div class="grid gap-2">
              <FormField
                v-slot="{ componentField }"
                name="password"
                :validate-on-blur="!isFieldDirty"
              >
                <FormItem v-auto-animate>
                  <div class="flex items-center justify-between">
                    <FormLabel>Password</FormLabel>
                    <NuxtLink
                      to="/auth/forgot-password"
                      class="text-xs text-muted-foreground underline underline-offset-4 hover:text-primary"
                    >
                      Forgot password?
                    </NuxtLink>
                  </div>
                  <FormControl>
                    <Input
                      type="password"
                      placeholder="********"
                      autocomplete="current-password"
                      :disabled="isLoading"
                      v-bind="componentField"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              </FormField>
            </div>

            <Button type="submit" :disabled="isLoading" class="w-full">
              <Loader2 v-if="isLoading" class="mr-2 h-4 w-4 animate-spin" />
              {{ isLoading ? 'Signing in...' : 'Sign in' }}
            </Button>

            <div v-if="authStore.error" class="mt-2 text-center text-sm text-red-600">
              {{ authStore.error }}
            </div>

            <div class="text-center text-sm text-muted-foreground">
              Contact your administrator to get access.
            </div>
          </div>
        </form>
      </CardContent>
    </Card>

    <div class="text-muted-foreground text-center text-xs text-balance">
      XStorage — Self-hosted object storage platform
    </div>
  </div>
</template>
