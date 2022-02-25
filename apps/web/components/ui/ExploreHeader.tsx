import type { FC } from 'react'
import { TrendingButton } from './Buttons/TrendingButton'

interface Props {
  title: string
  subTitle?: string
}

export const ExploreHeader: FC<Props> = ({ title, subTitle }: Props) => {
  return (
    <div className="mb-6">
      <div className="flex items-center justify-between">
        <h1 className="gradient-text text-3xl md:text-5xl mb-6 pb-4 font-bold">
          {title}
        </h1>
      </div>
      {subTitle && (
        <h2 className="mt-4 max-w-xl md:text-lg text-zinc-500 dark:text-zinc-300">
          {subTitle}
        </h2>
      )}
    </div>
  )
}
