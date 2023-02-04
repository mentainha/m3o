import type { ReactElement } from 'react'
import Link from 'next/link'
import { Subscription } from './Subscription'
import { SubscriptionPlans } from '@/lib/constants'

export function Subscriptions(): ReactElement {
  return (
    <div className="grid xl:grid-cols-3 mt-10 text-left gap-4">
      <Subscription
        cost="Start for Free"
        plan={SubscriptionPlans.Free}
        description="For individuals"
        features={[
          'Pay as you grow',
          '1 app deployment',
          '10 User accounts',
          '100 Image uploads',
          '1000 DB records',
          '10k API requests',
          '10 request/second',
          '1GB Space storage',
        ]}
        button={
          <Link href="/register">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-dev-start-button">
              Get Started
            </a>
          </Link>
        }
      />
      <Subscription
        cost="$35/month"
        plan={SubscriptionPlans.Pro}
        description="For professionals"
        features={[
          'Unlock Premium APIs',
          '10 app deployments',
          '20 lambda functions',
          '1k User accounts',
          '10k Image uploads',
          '100k DB records',
	  '1M API Requests',
          '100 request/second',
          '10GB Space storage',
        ]}
        button={
          <Link href="/register?subscription=pro">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-pro-start-button">
              Choose Pro
            </a>
          </Link>
        }
      />
      <Subscription
        cost=""
        plan={SubscriptionPlans.Business}
        description="For enterprises"
        features={[
          '100 app deployments',
          '1000 lambda functions',
          '100k User accounts',
          '1m Image uploads',
          '10M DB records',
          '100M API requests',
          '1000 request/second',
          '100GB Space storage',
        ]}
        button={
          <Link href="mailto:contact@m3o.com">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-business-start-button">
              Contact Us
            </a>
          </Link>
        }
      />
    </div>
  )
}
