import type { ReactElement } from 'react'
import Link from 'next/link'
import { Subscription } from './Subscription'
import { SubscriptionPlans } from '@/lib/constants'

export function Subscriptions(): ReactElement {
  return (
    <div className="grid xl:grid-cols-4 mt-10 text-left gap-4">
      <Subscription
        cost="Start for Free"
        plan={SubscriptionPlans.Free}
        description="For small projects"
        features={[
          '1 app deployment',
          '5 lambda functions',
          '10 User accounts',
          '100 Image uploads',
          '1k DB records',
          '10k API requests',
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
        cost="£20/month"
        plan={SubscriptionPlans.Solo}
        description="For creators and individuals"
        features={[
          '5 app deployments',
          '10 lambda functions',
          '100 User accounts',
          '1000 Image uploads',
          '10k DB records',
          '100k API requests',
          '5GB Space storage',
          'Community support',
        ]}
        button={
          <Link href="/register?subscription=solo">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-solo-start-button">
              Choose Solo
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
          '1k User accounts',
          '10k Image uploads',
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
        cost="£250/month"
        plan={SubscriptionPlans.Business}
        description="For growing companies"
        features={[
          '20 app deployments',
          '100 lambda functions',
          '10k User accounts',
          '100k Image uploads',
          '1M DB records',
          '10M API requests',
          '25GB Space storage',
          'Premium support',
        ]}
        button={
          <Link href="/register?subscription=business">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-business-start-button">
              Choose Business
            </a>
          </Link>
        }
      />
    </div>
  )
}
