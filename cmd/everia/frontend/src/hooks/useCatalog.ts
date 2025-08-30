'use client'

import { useAtom } from 'jotai'

import { apps_catalog_atom, websites_catalog_atom } from './catalogs'

export const useCatalog = () => {
  const authors = ['Long']
  const [apps, setApps] = useAtom(apps_catalog_atom)
  const [websites, setWebsites] = useAtom(websites_catalog_atom)

  return {
    authors,
    apps,
    setApps,
    websites,
    setWebsites
  }
}
