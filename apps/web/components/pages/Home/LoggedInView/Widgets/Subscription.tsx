import type { FC } from 'react'
import Link from 'next/link'
import { useBillingAccount } from '@/hooks'
import { SubscriptionPlanBubble } from '@/components/ui'
import { SubscriptionPlans } from '@/lib/constants'

type SubscriptionUpgradeLinksProps = {
  level: SubscriptionPlans
}

const LINK_CLASSES = 'text-sm ml-2 mb-2 last:mb-0 inline-block capitalize rounded-full bg-zinc-200 dark:bg-zinc-700 p-1 px-4'

type Link = {
  tier: SubscriptionPlans
  show?: (tier: SubscriptionPlans) => boolean
}

const LINKS: Link[] = [
  {
    tier: SubscriptionPlans.Solo,
    show: tier => tier !== SubscriptionPlans.Solo,
  },
  {
    tier: SubscriptionPlans.Pro,
  },
]

function SubscriptionUpgradeLinks({ level }: SubscriptionUpgradeLinksProps) {
  return (
    <div className="mt-4">
      Upgrade to
      {LINKS.map(link => {
        if (!link.show || link.show(level)) {
          return (
            <Link href={`/subscriptions?tier=${link.tier}`}>
              <a className={LINK_CLASSES}>{link.tier}</a>
            </Link>
          )
        }

        return null
      })}
    </div>
  )
}

export const Subscription: FC = () => {
  const { subscriptionLevel } = useBillingAccount()

  return (
    <div className="text-white">
      <h5 className="font-bold text-xl mb-4">Subscription Plan</h5>
      {subscriptionLevel && (
        <span className="mb-2 inline-block">
          <SubscriptionPlanBubble
            plan={subscriptionLevel}></SubscriptionPlanBubble>
        </span>
      )}
      {(subscriptionLevel === SubscriptionPlans.Free ||
        subscriptionLevel === SubscriptionPlans.Solo) && (
        <SubscriptionUpgradeLinks
          level={subscriptionLevel as SubscriptionPlans}
        />
      )}
    </div>
  )
}
