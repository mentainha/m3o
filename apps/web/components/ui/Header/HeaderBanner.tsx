import React from 'react'
import { useState, useEffect } from 'react'
import classNames from 'classnames'
import { DiscordIcon } from './DiscordIcon'
import { XIcon } from '@heroicons/react/outline'
import { SessionStorageKeys } from '@/lib/constants'
import Link from 'next/link'

export function HeaderBanner() {
  const [show, setShow] = useState(false)

  useEffect(() => {
    const hideBanner = sessionStorage.getItem(SessionStorageKeys.HideBanner)

    if (!hideBanner) {
      setShow(true)
    }
  }, [])

  const hideBanner = () => {
    setShow(false)
    sessionStorage.setItem(SessionStorageKeys.HideBanner, 'true')
  }

  return (
    <div
      className={classNames(
        'bg-indigo-800 dark:bg-indigo-600 hover:bg-indigo-600 dark:hover:bg-indigo-700 relative transition ',
        {
          flex: show,
          hidden: !show,
        },
      )}>
      <a
        className="text-white text-xs md:text-sm w-full flex items-center px-6 py-2 dark:text-white"
        href="https://play.google.com/store/apps/details?id=com.m3o.mobile">
        Get early access to M3O on Android. Check it out on the Google Play store</a>
      <button onClick={hideBanner} className="mr-6">
        <XIcon className="w-4 text-white" />
      </button>
    </div>
  )
}
