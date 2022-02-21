import type { FC } from 'react'
import { Outlet, Link } from 'react-router-dom'
import { AppLinks } from './AppLinks'
import {
  SupportIcon
} from '@heroicons/react/outline'

export const Layout: FC = () => {
  return (
    <div className="pl-52">
      <aside className="bg-zinc-900 h-screen fixed top-0 left-0 w-52 border-r border-gray-700">
        <div className="p-6">
          <a
            href="https://m3o.com"
            className="relative hover:no-underline w-14 flex items-center"
          >
            <img src="/logo-white.png" alt="m3o logo" className="w-14" />
          </a>
        </div>
        <nav className="m-2">
          <h1 className="pl-4 pb-4 font-bold">Cloud</h1>
          <AppLinks />
          <h1 className="pl-4 pb-4 pt-2 font-bold">Support</h1>
          <a
            href="mailto:contact@m3o.com?subject=M3O Cloud Feedback"
            className="h-12 flex items-center rounded-md text-sm p-4 text-gray-400"
          >
            <SupportIcon className="w-4 mr-2" /> Feedback
          </a>
        </nav>
      </aside>
      <main className="min-h-screen">
        <Outlet />
      </main>
    </div>
  )
}
