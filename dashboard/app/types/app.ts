export interface App {
  id: number
  name: string
  owner_id: number
  created_at: string
}

export interface CreateAppPayload {
  name: string
}

export interface APIKey {
  id: number
  app_id: number
  access_key: string
  created_at: string
}
