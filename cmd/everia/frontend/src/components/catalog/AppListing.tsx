import { models } from '@wailsjs/go/models'

import CatalogItem from './CatalogItem'

type AppListingProps = {
  apps: models.AppCatalog[]
}

const AppListing = ({ apps }: AppListingProps) => {
  // TODO: list all the app and add an insert button to insert new app
  return (
    <div className='px-2'>
      AppListing
      <div className='grid grid-cols-2 gap-y-3'>
        {apps.map(app => (
          <CatalogItem key={app.name} {...app} description={app.version} />
        ))}
      </div>
    </div>
  )
}

export default AppListing
