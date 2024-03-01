import { Fragment } from 'react'
import { useRouter } from 'next/router'
import { Popover, Transition } from '@headlessui/react'
import { MenuIcon, XIcon } from '@heroicons/react/outline'
import Grid from '@tw/Grid'
import ForumIcon from '@icons/ForumIcon'
import TwitterIcon from '@icons/TwitterIcon'
import DiscordIcon from '@icons/DiscordIcon'
import TelegramIcon from '@icons/TelegramIcon'
import DocumentTextIcon from '@icons/DocsIcon'
import { Wallet } from '@components/Wallet'

import { SynapseLogoSvg, SynapseLogoWithTitleSvg } from './SynapseLogoSvg'
import { TopBarNavLink, checkIsRouteMatched } from './TopBarNavLink'
import {
  DISCORD_URL,
  SYNAPSE_DOCS_URL,
  FORUM_URL,
  LANDING_PATH,
  TELEGRAM_URL,
  TWITTER_URL,
} from '@/constants/urls'
import { NAVIGATION } from '@/constants/routes'
import { MoreButton } from './MoreButton'
import { PageFooter } from './PageFooter'

export function LandingPageWrapper({ nestedPage=false, children }) {
  return (
    <div
      className="from-black via-slate-950 to-black bg-gradient-to-br"
    >
      <div
          className="bg-[length:100vw_100vh] bg-opacity-50 transition-all bg-fixed"
          style={{
            backgroundImage: nestedPage ? `url('../landingBg.svg')` : `url('landingBg.svg')`
          }}
      >
        <LandingNav />
        <div>
          {children}
        </div>
        <PageFooter />
      </div>

    </div>
  )
}


function MobileNavButton({ label, IconComponent }) {
  return (
    <Popover.Button
      data-test-id="mobile-navbar-button"
      className="group flex items-center p-2 text-opacity-75 bg-bgBase/10 hover:bg-bgBase/20 ring-1 ring-white/10 hover:ring-white/30 text-secondaryTextColor hover:text-white rounded-md"
    >
      <span className="sr-only">{label}</span>
      <IconComponent className="w-6 h-6" aria-hidden="true" />
    </Popover.Button>
  )
}


export function LandingNav() {
  return (
    <Popover>
      <div className="flex gap-4 place-content-between p-8 max-w-[1440px] m-auto">
        <SynapseTitleLogo showText={true} />
        <div className="lg:hidden">
          <MobileNavButton label="Open menu" IconComponent={MenuIcon} />
        </div>
        <Popover.Group
          as="nav"
          className="flex-wrap justify-center hidden lg:flex"
          data-test-id="desktop-nav"
        >
          <TopBarButtons />
        </Popover.Group>
        <div className="hidden lg:flex h-fit">
          <div className="flex items-center space-x-2">
            <Wallet />
            <Popover className="relative">
              {({ open }) => (
                <>
                  <Popover.Button as="div" onMouseEnter={() => {}}>
                    <MoreButton open={open} />
                  </Popover.Button>
                  <PopoverPanelContainer className="-translate-x-full left-full">
                    <MoreInfoButtons />
                    <SocialButtons />
                  </PopoverPanelContainer>
                </>
              )}
            </Popover>
          </div>
        </div>
      </div>

      <Transition
        as={Fragment}
        enter="duration-100 ease-out"
        enterFrom=" opacity-0"
        enterTo=" opacity-100"
        leave="duration-75 ease-in"
        leaveFrom=" opacity-100"
        leaveTo=" opacity-0"
      >
        <Popover.Panel focus className="absolute top-0 z-10 w-screen">
          <div className="bg-bgBase/10 backdrop-blur-lg">
            <div className="flex items-center p-8 pb-4 place-content-between">
              <SynapseTitleLogo showText={true} />
              <MobileNavButton label="Close menu" IconComponent={XIcon} />
            </div>
            <div className="flex flex-col gap-2 py-4 pl-4" data-test-id="mobile-nav">
              <MobileBarButtons />
            </div>
            <div className="pr-2 pl-8 py-4 bg-white/10">
              <Wallet />
            </div>
          </div>
        </Popover.Panel>
      </Transition>
    </Popover>
  )
}

export function PopoverPanelContainer({
  children,
  className,
}: {
  children: any
  className?: string
}) {
  return (
    <Transition
      as={Fragment}
      enter="transition ease-out duration-200"
      enterFrom="opacity-0 translate-y-1"
      enterTo="opacity-100 translate-y-0"
      leave="transition ease-in duration-150"
      leaveFrom="opacity-100 translate-y-0"
      leaveTo="opacity-0 translate-y-1"
    >
      <Popover.Panel
        className={`
          absolute z-10 left-1/2 transform-gpu
          ${className ?? '-translate-x-1/2'}
          mt-3 w-screen max-w-xs sm:px-0
        `}
      >
        <div className="overflow-hidden rounded-md shadow-xl ring-1 ring-white/10">
          <div
            className={`
              relative grid gap-3 bg-bgBase/10 backdrop-blur-lg
              px-2.5 py-3 sm:p-2
            `}
          >
            {children}
          </div>
        </div>
      </Popover.Panel>
    </Transition>
  )
}

function TopBarButtons() {
  const topBarNavLinks = Object.entries(NAVIGATION).map(([key, value]) => (
    <TopBarNavLink
      key={key}
      to={value.path}
      labelText={value.text}
      match={value.match}
    />
  ))

  return <>{topBarNavLinks}</>
}

function MoreInfoButtons() {
  return (
    <>
      <MoreInfoItem
        className="mdl:hidden"
        to={NAVIGATION.Analytics.path}
        labelText={NAVIGATION.Analytics.text}
        description="See preliminary analytics of the bridge"
      />
      <MoreInfoItem
        to={NAVIGATION.Contracts.path}
        labelText={NAVIGATION.Contracts.text}
        description="View contract related information such as contract addresses"
      />
    </>
  )
}

function SocialButtons() {
  return (
    <Grid cols={{ xs: 2, sm: 1 }} gapY={'1'}>
      <MiniInfoItem
        href={SYNAPSE_DOCS_URL}
        labelText="Docs"
        IconComponent={DocumentTextIcon}
      />
      <MiniInfoItem
        href={DISCORD_URL}
        labelText="Discord"
        IconComponent={DiscordIcon}
      />
      <MiniInfoItem
        href={TELEGRAM_URL}
        labelText="Telegram"
        IconComponent={TelegramIcon}
      />
      <MiniInfoItem
        href={TWITTER_URL}
        labelText="Twitter"
        IconComponent={TwitterIcon}
      />
      <MiniInfoItem
        href={FORUM_URL}
        labelText="Forum"
        IconComponent={ForumIcon}
      />
    </Grid>
  )
}

function MobileBarButtons() {
  const mobileBarItems = Object.entries(NAVIGATION).map(([key, value]) => (
    <MobileBarItem
      key={key}
      to={value.path}
      labelText={value.text}
      match={value.match}
    />
  ))

  return <>{mobileBarItems}</>
}

function MobileBarItem({
  to,
  labelText,
  match,
}: {
  to: string
  labelText: string
  match?: string | { startsWith: string }
}) {
  const router = useRouter()

  const isRouteMatched = checkIsRouteMatched(router, match)

  const isInternal = to[0] === '/' || to[0] === '#'

  return (
    <a
      key={labelText}
      href={to}
      target={isInternal ? undefined : '_blank'}
      className={`
        px-4 py-2 text-2xl font-medium text-white
        ${isRouteMatched ? 'text-opacity-100' : 'text-opacity-30'}
      `}
    >
      {labelText}
    </a>
  )
}

function MoreInfoItem({
  to,
  labelText,
  description,
  className,
}: {
  to: string
  labelText: string
  description: string
  className?: string
}) {
  return (
    <a
      key={labelText}
      href={to}
      target={to[0] === '/' ? undefined : '_blank'}
      className={`block px-3 pt-2 pb-2 rounded-md hover:bg-white hover:bg-opacity-10 ${className}`}
    >
      <p className="text-base font-medium text-white">{labelText}</p>
      <p className="hidden mt-1 text-sm text-white text-opacity-60 md:block">
        {description}
      </p>
    </a>
  )
}

function MiniInfoItem({
  href,
  labelText,
  IconComponent,
}: {
  href: string
  labelText: string
  IconComponent: any
}) {
  return (
    <a
      key={labelText}
      href={href}
      className="block px-3 pt-1 pb-2 text-sm rounded-md group"
      target="_blank"
    >
      <div>
        <p className="text-base text-white text-opacity-40 group-hover:text-white">
          <IconComponent className="inline w-5 mr-2 -ml-1" />
          <span className="mt-1">{labelText}</span>
        </p>
      </div>
    </a>
  )
}

export function SynapseTitleLogo({ showText }: { showText: boolean }) {
  return (
    <a href={LANDING_PATH}>
      {showText ? <SynapseLogoWithTitleSvg /> : <SynapseLogoSvg />}
    </a>
  )
}
