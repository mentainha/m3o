import type { FC } from 'react'
import Link from 'next/link'
import { useBillingAccount } from '@/hooks'
import { SubscriptionPlanBubble } from '@/components/ui'
import { SubscriptionPlans } from '@/lib/constants'

export const Subscription: FC = () => {
  const { subscriptionLevel } = useBillingAccount()

  return (
    <div className="text-white">
      <h5 className="font-bold text-xl mb-4">Subscription Plan</h5>
      {subscriptionLevel && (
        <SubscriptionPlanBubble
          plan={subscriptionLevel}></SubscriptionPlanBubble>
      )}
     {subscriptionLevel === SubscriptionPlans.Free && (
	<div className="mt-4 font-light">
	  <Link href="/subscriptions/pro/card-details">
	    <a className="text-sm inline-flex items-center">
	      Upgrade to Pro
	    </a>
	  </Link>
	</div>
      )}
    </div>
  )
}
