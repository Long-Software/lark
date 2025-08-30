'use client'

import AppListing from '@/components/catalog/AppListing'
import WebsiteListing from '@/components/catalog/WebsiteListing'
import { useCatalog } from '@/hooks/useCatalog'

const CatalogPage = () => {
  const { apps, websites } = useCatalog()

  return (
    <div className=''>
      <AppListing apps={apps} />
      <WebsiteListing websites={websites} />
    </div>
  )
}

export default CatalogPage
