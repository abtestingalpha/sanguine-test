import { Transition } from "@headlessui/react"
import { TRANSITION_PROPS } from "@/styles/transitions"


export function OverlayTransition({ show, children, ...props }) {
  return (
    <Transition show={show} {...TRANSITION_PROPS} {...props}>
      <div className='fixed z-50 w-full h-full bg-opacity-50 bg-slate-400/10'>
        {children}
      </div>
    </Transition>
  )
}