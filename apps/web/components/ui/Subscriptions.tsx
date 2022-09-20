import type { ReactElement } from 'react'
import Link from 'next/link'
import { Subscription } from './Subscription'
import { SubscriptionPlans } from '@/lib/constants'

export function Subscriptions(): ReactElement {
  return (
    <div className="grid xl:grid-cols-4 mt-10 text-left gap-4">
      <Subscription
        cost="Pay as you grow"
        plan={SubscriptionPlans.Dev}
        description="For small projects"
        features={[
          'Free to get started',
          '1 app deployment',
          '5 lambda functions',
          '100 User accounts',
          '1k DB records',
          '1M API requests',
          '1GB Space storage',
        ]}
        button={
          <Link href="/register">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-free-start-button">
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
          '1k User accounts',
          '10k DB records',
          '5M API requests',
          '10GB Space storage',
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
          '10k User accounts',
          '100k DB records',
          '25M API requests',
          '100GB Space storage',
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
        plan={SubscriptionPlans.Enterprise}
        description="For growing companies"
        features={[
          '20 app deployments',
          '100 lambda functions',
          '100k User accounts',
          '1M DB records',
          '100M API requests',
          '1TB Space storage',
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
