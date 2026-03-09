import type { ApiResponse } from '@/types/auth'
import type { NitroFetchOptions, NitroFetchRequest } from 'nitropack'

type FetchOptions = NitroFetchOptions<NitroFetchRequest> & {
  params?: Record<string, string>
}

export const useApi = () => {
  const config = useRuntimeConfig()
  const { getToken, deleteToken } = useAccessToken()

  const apiFetch = async <T>(
    path: string,
    options: FetchOptions = {},
  ): Promise<ApiResponse<T>> => {
    const token = getToken()
    const headers: Record<string, string> = {
      ...(options.headers as Record<string, string>),
    }

    if (token) {
      headers['Authorization'] = `Bearer ${token}`
    }

    try {
      return await $fetch<ApiResponse<T>>(`${config.public.apiBase}${path}`, {
        ...options,
        headers,
      })
    } catch (err: unknown) {
      const status = (err as { response?: { status?: number } })?.response?.status
      if (status === 401) {
        deleteToken()
        await navigateTo('/auth/login')
      }
      throw err
    }
  }

  return { apiFetch }
}
