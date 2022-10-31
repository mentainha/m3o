import type { NextPage } from 'next'
import { NextSeo } from 'next-seo'
import { useState } from 'react'
import Link from 'next/link'
import { Routes } from '@/lib/constants'
import { ChevronRightIcon } from '@heroicons/react/outline'
import { DashboardLayout } from '@/components/layouts'
import { withAuth } from '@/lib/api/m3o/withAuth'
import { Modal, Button, SubscriptionPlanBubble } from '@/components/ui'
import {
  AddFunds,
  PaymentMethods,
  DeleteCardModal,
  History,
} from '@/components/pages/User'
import { BillingProvider } from '@/providers'
import seo from '@/lib/seo.json'
import { useBillingAccount, useDowngradeToFreeTier } from '@/hooks'
import { SubscriptionPlans } from '@/lib/constants'

interface Props {
  user: Account
}

export const getServerSideProps = withAuth(async context => {
  if (!context.req.user) {
    return {
      redirect: {
        destination: Routes.Home,
        permanent: false,
      },
    }
  }

  return {
    props: {
      user: context.req.user,
    },
  }
})

const Billing: NextPage<Props> = ({ user }) => {
  const [showCancelSubscriptionModal, setShowSubscriptionModal] =
    useState(false)
  const { subscriptionLevel } = useBillingAccount()
  const { downgrade, isLoading } = useDowngradeToFreeTier({
    onSuccess: () => {
      setShowSubscriptionModal(false)
    },
  })

  return (
    <>
      <NextSeo title={seo.account.billing.title} />
      <BillingProvider user={user}>
        <DashboardLayout>
          <div className="p-6 md:p-10">
            <h1 className="text-3xl mb-6 pb-4 font-medium">
              Billing
            </h1>
            <AddFunds />
            <PaymentMethods />
            <div className="mt-8 tbgc p-6 md:p-10 rounded-lg">
              <h3 className="font-bold text-2xl text-black mb-3 dark:text-white">
                Subscription
              </h3>
              {subscriptionLevel !== SubscriptionPlans.Free && (
              <p className="ttc mb-4">
                {subscriptionLevel && (
                  <SubscriptionPlanBubble plan={subscriptionLevel} />
                )}{' '}
              </p>
              )}
              {subscriptionLevel !== SubscriptionPlans.Free && (
                <div className="m-2 font-light">
                  <button
                    className="bg-zinc-200 dark:bg-zinc-700 py-2 px-4 rounded-full text-sm flex items-center"
                    onClick={() => setShowSubscriptionModal(true)}>
                    Cancel subscription{' '}
                    <ChevronRightIcon className="w-4 ml-2" />
                  </button>
                </div>
              )}
              {subscriptionLevel === SubscriptionPlans.Free && (
                <span className="font-light">
                  <Link href="/pricing">
                    <a className="bg-zinc-200 dark:bg-zinc-700 py-2 px-4 rounded-full text-sm inline-flex items-center">
                      Choose a Plan
                    </a>
                  </Link>
                </span>
              )}
            </div>
            <History />
            <DeleteCardModal />
            <Modal
              open={showCancelSubscriptionModal}
              closeModal={() => setShowSubscriptionModal(false)}>
              <p className="pr-6">
                Are you sure you would like to cancel your subscription?
              </p>
              <Button
                loading={isLoading}
                onClick={() => downgrade()}
                className="mt-6">
                Confirm
              </Button>
            </Modal>
          </div>
        </DashboardLayout>
      </BillingProvider>
    </>
  )
}

export default Billing
