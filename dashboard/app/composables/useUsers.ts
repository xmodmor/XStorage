import type { ApiResponse } from '@/types/auth'
import type { User, CreateUserPayload, UpdateUserPayload, UsersListResponse } from '@/types/user'

export const useUsers = () => {
  const { apiFetch } = useApi()

  const fetchUsers = (page = 1, limit = 20, search = '') => {
    const params: Record<string, string> = {
      page: String(page),
      limit: String(limit),
    }
    if (search) params.search = search

    return apiFetch<UsersListResponse>('/api/v1/users', { params })
  }

  const fetchUser = (id: number) => {
    return apiFetch<User>(`/api/v1/users/${id}`)
  }

  const createUser = (payload: CreateUserPayload) => {
    return apiFetch<User>('/api/v1/users', {
      method: 'POST',
      body: payload,
    })
  }

  const updateUser = (id: number, payload: UpdateUserPayload) => {
    return apiFetch<User>(`/api/v1/users/${id}`, {
      method: 'PUT',
      body: payload,
    })
  }

  const deleteUser = (id: number) => {
    return apiFetch<null>(`/api/v1/users/${id}`, {
      method: 'DELETE',
    })
  }

  return { fetchUsers, fetchUser, createUser, updateUser, deleteUser }
}
