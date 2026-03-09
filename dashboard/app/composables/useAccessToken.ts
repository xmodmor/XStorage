export const useAccessToken = () => {
  const tokenCookie = useCookie<string | null>('access_token', {
    path: '/',
    sameSite: 'lax',
    secure: !import.meta.dev,
    maxAge: 60 * 60 * 24,
    default: () => null,
    watch: true,
  })

  const setToken = (token: string) => {
    tokenCookie.value = token
  }

  const getToken = (): string | null => {
    return tokenCookie.value
  }

  const deleteToken = () => {
    tokenCookie.value = null
  }

  return {
    tokenCookie,
    setToken,
    getToken,
    deleteToken,
  }
}
