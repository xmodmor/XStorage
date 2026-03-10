import { useAuthStore } from '@/stores/useAuth'

const PUBLIC_PATHS = ['/', '/auth/login']

export default defineNuxtRouteMiddleware(async (to) => {
  const isPublic = PUBLIC_PATHS.includes(to.path) || to.path.startsWith('/docs')
  const auth = useAuthStore()

  if (!auth.isInitialized || (auth.token && !auth.user)) {
    await auth.initialize()
  }

  if (isPublic) {
    if (auth.isAuthenticated && to.path.startsWith('/auth/')) {
      return navigateTo('/dashboard')
    }
    return
  }

  if (!auth.isAuthenticated) {
    return navigateTo({
      path: '/auth/login',
      query: { redirect: to.fullPath },
    })
  }
})
