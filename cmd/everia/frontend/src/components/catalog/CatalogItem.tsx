import Image from 'next/image'

type CatalogItemProps = {
  name: string
  image_url: string
  version: string
  description: string
}

const CatalogItem = ({ name, image_url, version, description }: CatalogItemProps) => {
  return (
    <div className='flex h-[84px] w-[250px] flex-row justify-start space-x-4 rounded-md border border-foreground p-4'>
      <Image src={image_url} width={36} height={36} alt={name} className='' />
      <div className=''>
        <p className='text-lg font-bold'>
          {name} : {version}
        </p>
        <p className='text-base'>{description}</p>
      </div>
    </div>
  )
}

export default CatalogItem
