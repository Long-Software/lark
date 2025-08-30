import typograhpy from '@tailwindcss/typography'

import sharedConfig, { Config } from '@longslark/tailwind/tailwind.config'

const config: Config = {
  presets: [sharedConfig],
  content: [
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './../../../../packages/ui/src/components/ui/*.{js,ts,jsx,tsx,mdx}'
  ],
  plugins: [typograhpy]
}
export default config
