import type { ReactElement } from 'react'
import Link from 'next/link'
import { Subscription } from './Subscription'
import { SubscriptionPlans } from '@/lib/constants'

export function Subscriptions(): ReactElement {
  return (
    <div className="grid md:grid-cols-3 mt-10 text-left gap-4">
      <Subscription
        cost="Free to start"
        plan={SubscriptionPlans.Free}
        description="For hackers and hobbyists"
        features={[
          'Access to 50+ public APIs',
          '10,000 requests per month',
          'Top-up to pay as you grow',
          'Community support',
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
        plan={SubscriptionPlans.Pro}
        description="For professionals and teams"
        features={[
          'Everything in Free plus',
          '1 million requests per month',
          'Unlock paid APIs and features',
          'Higher SLAs & response times',
          'Increased usage limits',
          'Email support',
        ]}
        button={
          <Link href="/register?subscription=pro">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-pro-start-button">
              Buy Pro
            </a>
          </Link>
        }
      />
      <Subscription
        cost="Scale as you grow"
        plan={SubscriptionPlans.Business}
        description="For growing businesses"
        features={[
          'Everything in Pro plus',
          'Unlimited requests per month',
          'Dedicated infrastructure',
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
