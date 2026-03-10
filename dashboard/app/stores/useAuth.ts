import { nextTick } from 'vue'
import { defineStore } from 'pinia'
import type { User, LoginCredentials, LoginData } from '@/types/auth'

export const useAuthStore = defineStore('auth', () => {
  const { tokenCookie, setToken, deleteToken } = useAccessToken()
  const { apiFetch } = useApi()

  const user = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const isInitialized = ref(false)

  const token = computed(() => tokenCookie.value)
  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const initialize = async () => {
    if (token.value) {
      try {
        await fetchUser()
      } catch {
        deleteToken()
        user.value = null
      }
    }
    isInitialized.value = true
  }

  const fetchUser = async () => {
    const res = await apiFetch<User>('/api/v1/auth/me')
    if (res.success && res.data) {
      user.value = res.data
    } else {
      throw new Error(res.error?.message ?? 'Failed to fetch user')
    }
  }

  const login = async (credentials: LoginCredentials, redirectTo?: string) => {
    loading.value = true
    error.value = null

    try {
      const res = await apiFetch<LoginData>('/api/v1/auth/login', {
        method: 'POST',
        body: credentials,
      })

      if (!res.success || !res.data) {
        error.value = res.error?.message ?? 'Login failed'
        return false
      }

      setToken(res.data.token)
      user.value = res.data.user

      await nextTick()
      await navigateTo(redirectTo ?? '/dashboard', { replace: true })
      return true
    } catch {
      error.value = 'Network error. Please try again.'
      return false
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    deleteToken()
    user.value = null
    error.value = null
    await navigateTo('/auth/login', { replace: true })
  }

  return {
    user,
    token,
    loading,
    error,
    isInitialized,
    isAuthenticated,
    initialize,
    fetchUser,
    login,
    logout,
  }
})
