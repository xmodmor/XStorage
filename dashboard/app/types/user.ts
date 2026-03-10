export interface User {
  id: number
  email: string
  created_at: string
}

export interface CreateUserPayload {
  email: string
  password: string
}

export interface UpdateUserPayload {
  email: string
  password?: string
}

export interface UsersListResponse {
  users: User[]
  total: number
  page: number
  limit: number
}
