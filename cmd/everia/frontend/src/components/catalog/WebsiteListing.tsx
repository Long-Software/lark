import { models } from '@wailsjs/go/models'

import CatalogItem from './CatalogItem'

type WebsiteListingProps = {
  websites: models.WebsiteCatalog[]
}

const WebsiteListing = ({ websites }: WebsiteListingProps) => {
  // TODO: list all the website and add an insert button to insert new website
  return (
    <div className='px-2'>
      WebsiteListing
      <div className='grid grid-cols-2 gap-y-3'>
        {websites.map(website => (
          <CatalogItem key={website.name} {...website} description={website.url} />
        ))}
      </div>
    </div>
  )
}

export default WebsiteListing
