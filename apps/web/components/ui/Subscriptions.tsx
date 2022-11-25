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
        description="For small projects"
        features={[
          '1 app deployment',
          '5 lambda functions',
          '10 Image uploads',
          '100 User accounts',
          '10k DB records',
          '100k API requests',
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
        cost="£100/month"
        plan={SubscriptionPlans.Pro}
        description="For professionals and teams"
        features={[
          '10 app deployments',
          '20 lambda functions',
          '1k Image uploads',
          '10k User accounts',
          '100k DB records',
	  '1M API Requests',
          '10GB Space storage',
          'Email support',
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
        cost="Custom pricing"
        plan={SubscriptionPlans.Business}
        description="For growing companies"
        features={[
          '20 app deployments',
          '100 lambda functions',
          '10k Image uploads',
          '100k User accounts',
          '1M DB records',
          '10M API requests',
          '100GB Space storage',
          'Premium support',
        ]}
        button={
          <a
            className="btn block w-full text-center"
            data-testid="subscription-business-start-button"
            href="mailto:contact@m3o.com?subject=M3O Business Plan">
            Contact Us
          </a>
        }
      />
    </div>
  )
}
