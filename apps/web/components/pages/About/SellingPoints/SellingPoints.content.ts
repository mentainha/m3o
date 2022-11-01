import {
  LightningBoltIcon,
  CashIcon,
  EyeOffIcon,
  EmojiHappyIcon,
} from '@heroicons/react/outline'
import { SellingPointProps } from './SellingPoints.types'

export const SELLING_POINTS: SellingPointProps[] = [
  {
    Icon: EmojiHappyIcon,
    description:
      'Use one account for all your API needs. Access multiple services with a single token',
    title: 'Simple',
  },
  {
    Icon: LightningBoltIcon,
    description:
      "Using a new API is simple - no need to learn yet another API style, it's all the same Dev UX",
    title: 'Fast',
  },
  {
    Icon: CashIcon,
    description:
      "It's a simple pay as you grow model and everything is priced per request",
    title: 'Affordable',
  },
  {
    Icon: EyeOffIcon,
    description:
      "Don't get lost in infinite billing. We show you exactly what you use with no hidden costs",
    title: 'No Hidden Fees',
  },
]
