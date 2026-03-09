<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Field, FieldGroup, FieldLabel } from '~/components/ui/field'
import { Input } from '~/components/ui/input'
import { Alert, AlertDescription } from '~/components/ui/alert'
import { useAuthStore } from '@/stores/useAuth'

const props = defineProps<{
  class?: HTMLAttributes['class']
}>()

const auth = useAuthStore()
const route = useRoute()

const email = ref('')
const password = ref('')

const onSubmit = async () => {
  const redirect = route.query.redirect as string | undefined
  await auth.login({ email: email.value, password: password.value }, redirect)
}
</script>

<template>
  <form :class="cn('flex flex-col gap-6', props.class)" @submit.prevent="onSubmit">
    <FieldGroup>
      <div class="flex flex-col items-center gap-1 text-center">
        <h1 class="text-2xl font-bold">
          Login to your account
        </h1>
        <p class="text-muted-foreground text-sm text-balance">
          Enter your email below to login to your account
        </p>
      </div>

      <Alert v-if="auth.error" variant="destructive">
        <AlertDescription>{{ auth.error }}</AlertDescription>
      </Alert>

      <Field>
        <FieldLabel for="email">Email</FieldLabel>
        <Input id="email" v-model="email" type="email" placeholder="m@example.com" required />
      </Field>

      <Field>
        <div class="flex items-center">
          <FieldLabel for="password">Password</FieldLabel>
          <a href="#" class="ml-auto text-sm underline-offset-4 hover:underline">
            Forgot your password?
          </a>
        </div>
        <Input id="password" v-model="password" type="password" required />
      </Field>

      <Field>
        <Button type="submit" :disabled="auth.loading" class="w-full">
          <template v-if="auth.loading">Logging in...</template>
          <template v-else>Login</template>
        </Button>
      </Field>
    </FieldGroup>
  </form>
</template>
