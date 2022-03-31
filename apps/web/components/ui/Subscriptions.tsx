import type { ReactElement } from 'react'
import Link from 'next/link'
import { Subscription } from './Subscription'
import { SubscriptionPlans } from '@/lib/constants'

export function Subscriptions(): ReactElement {
  return (
    <div className="grid xl:grid-cols-4 mt-10 text-left gap-4">
      <Subscription
        cost="Free to start"
        plan={SubscriptionPlans.Free}
        description="For testing and hacking"
        features={[
          'Access to 50+ public APIs',
          '100,000 requests for free',
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
        cost="£10/month"
        plan={SubscriptionPlans.Solo}
        description="For creators and developers"
        features={[
          'Everything in Free plus',
          '1 million requests per month',
          'Unlock paid APIs and features',
          'Increased rate limits',
        ]}
        button={
          <Link href="/register?subscription=solo">
            <a
              className="btn block w-full text-center"
              data-testid="subscription-solo-start-button">
              Buy Solo
            </a>
          </Link>
        }
      />
      <Subscription
        cost="£20/month"
        plan={SubscriptionPlans.Pro}
        description="For professionals and teams"
        features={[
          'Everything in Solo plus',
          '10 million requests per month',
          'Higher SLAs & response times',
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
