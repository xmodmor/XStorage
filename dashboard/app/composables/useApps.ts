import type { App, CreateAppPayload, APIKey } from '@/types/app'

export const useApps = () => {
  const { apiFetch } = useApi()

  const fetchApps = () => {
    return apiFetch<App[]>('/api/v1/apps')
  }

  const fetchApp = (id: number) => {
    return apiFetch<App>(`/api/v1/apps/${id}`)
  }

  const createApp = (payload: CreateAppPayload) => {
    return apiFetch<App>('/api/v1/apps', {
      method: 'POST',
      body: payload,
    })
  }

  const deleteApp = (id: number) => {
    return apiFetch<null>(`/api/v1/apps/${id}`, {
      method: 'DELETE',
    })
  }

  const fetchApiKeys = (appId: number) => {
    return apiFetch<APIKey[]>(`/api/v1/apps/${appId}/keys`)
  }

  const createApiKey = (appId: number) => {
    return apiFetch<APIKey>(`/api/v1/apps/${appId}/keys`, {
      method: 'POST',
    })
  }

  const deleteApiKey = (appId: number, keyId: number) => {
    return apiFetch<null>(`/api/v1/apps/${appId}/keys/${keyId}`, {
      method: 'DELETE',
    })
  }

  return { fetchApps, fetchApp, createApp, deleteApp, fetchApiKeys, createApiKey, deleteApiKey }
}
