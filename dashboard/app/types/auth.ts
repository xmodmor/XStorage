export interface User {
  id: number
  email: string
  created_at: string
}

export interface LoginCredentials {
  email: string
  password: string
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: {
    code: string
    message: string
  }
}

export interface LoginData {
  token: string
  user: User
}
