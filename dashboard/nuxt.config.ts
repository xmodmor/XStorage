import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: [
    'shadcn-nuxt',
    '@pinia/nuxt',
    '@vee-validate/nuxt',
  ],
  css: [
    '~/assets/css/tailwind.css',
    '~/assets/css/main.css',
  ],
  ssr: true,
  vite: {
    plugins: [tailwindcss() as []],
  },
  shadcn: {
    prefix: '',
  },
  pinia: {
      storesDirs: ['~/app/stores'],
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
      isDev: process.env.NODE_ENV === 'development',
    },
  },
})