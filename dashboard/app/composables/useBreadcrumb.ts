import type { BreadcrumbItems } from '@/types/breadcrumb.ts'

const breadcrumbItems = ref<BreadcrumbItems>([])

export const useBreadcrumb = () => {
  const setBreadcrumbItems = (items: BreadcrumbItems) => {
    breadcrumbItems.value = items
  }

  const addBreadcrumbItem = (item: BreadcrumbItems[0]) => {
    breadcrumbItems.value.push(item)
  }

  const clearBreadcrumb = () => {
    breadcrumbItems.value = []
  }

  const updateLastItem = (label: string, isActive = true) => {
    if (breadcrumbItems.value.length > 0) {
      const lastIndex = breadcrumbItems.value.length - 1
      breadcrumbItems.value[lastIndex] = {
        ...breadcrumbItems.value[lastIndex],
        label,
        isActive
      }
    }
  }

  return {
    breadcrumbItems: readonly(breadcrumbItems),
    setBreadcrumbItems,
    addBreadcrumbItem,
    clearBreadcrumb,
    updateLastItem
  }
}
