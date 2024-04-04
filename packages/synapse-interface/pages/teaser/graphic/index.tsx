import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import exampleImg from '@assets/example.png'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import Hero from '../Hero'
import ValueProps from '../ValueProps'

import Wrapper from '@/components/WipWrapperComponents/Wrapper'

import styles from './keyframes.module.css'
import FauxBridge from '../FauxBridge'

const LandingPage = () => {
  const { address: currentAddress } = useAccount()
  const router = useRouter()

  useEffect(() => {
    segmentAnalyticsEvent(`[Teaser] arrives`, {
      address: currentAddress,
      query: router.query,
      pathname: router.pathname,
    })
  }, [])

  const hslStr = (h, s, l, a = undefined) =>
    `hsl(${h}deg ${s}% ${l}%${a === undefined ? '' : ` / ${a}%`})`

  const ethereum = {
    stroke: hslStr(260, 77, 83), // lavender // mint: 176, 87, 85
    fill: hslStr(240, 78, 7), // l: 40
    text: hslStr(14, 61, 85), // peach
  }
  const arbitrum = {
    stroke: hslStr(204, 87, 55),
    fill: hslStr(216, 58, 9), // l: 19
    text: hslStr(204, 87, 65),
  }
  const avalanche = {
    stroke: hslStr(360, 78, 58),
    fill: hslStr(360, 78, 3),
    text: hslStr(360, 78, 58),
  }
  const blast = {
    stroke: hslStr(60, 88, 45), // 60 98 50
    fill: hslStr(60, 98, 3),
    text: hslStr(60, 98, 65),
  }

  const stroke = {
    inherit: 'inherit',
    synapse: hslStr(300, 100, 25),
    flash: hslStr(300, 100, 40),
    beam: hslStr(300, 80, 60),
    yellow: hslStr(60, 80, 60),
    orange: hslStr(25, 80, 60),
    blue: hslStr(195, 80, 60),
    green: hslStr(135, 80, 60),
    north: arbitrum.stroke,
    west: ethereum.stroke,
    east: avalanche.stroke,
    south: blast.stroke,
  }
  const fill = {
    inherit: 'inherit',
    synapse: hslStr(300, 100, 5),
    yellow: hslStr(60, 30, 3),
    orange: hslStr(25, 30, 3),
    blue: hslStr(195, 30, 3),
    green: hslStr(135, 30, 3),
    north: arbitrum.fill,
    west: ethereum.fill,
    east: avalanche.fill,
    south: blast.fill,
  }

  const text = {
    north: arbitrum.text,
    west: ethereum.text,
    east: avalanche.text,
    south: blast.text,
  }

  const animAttrs = (x1 = 0.5, x2 = 0.2, y1 = 0, y2 = 1) => {
    return {
      calcMode: 'spline',
      keyTimes: '0; 1',
      keySplines: `${x1} ${y1} ${x2} ${y2}`,
      fill: 'freeze',
    }
  }

  const flashAttrs = (from, to, dur = '.4s') => {
    return {
      values: `${stroke[from]}; ${stroke.flash}; ${stroke[to]}`,
      dur: dur,
      calcMode: 'spline',
      keyTimes: `0; .5; 1`,
      keySplines: '.5 0 .2 1; .5 0 .2 1',
      fill: 'freeze',
    }
  }

  const AnimateFlash = ({
    hasStroke = true,
    hasFill = false,
    from,
    to,
    begin,
    dur = '.4s',
  }) => {
    return (
      <>
        {hasStroke && (
          <animate
            attributeName="stroke"
            begin={begin}
            dur={dur}
            values={`${stroke[from]}; ${stroke.flash}; ${stroke[to]}`}
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        )}
        {hasFill && (
          <animate
            attributeName="fill"
            begin={begin}
            dur={dur}
            values={`${fill[from]}; ${fill.synapse}; ${fill[to]}`}
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        )}
      </>
    )
  }

  const paint = (color = 'synapse') => {
    return {
      stroke: stroke[color],
      fill: fill[color],
    }
  }

  const Platform = ({
    id,
    color,
    translate,
    begin = '0s',
  }: {
    id: string
    color?: string
    translate?: string
    begin?: string
  }) => {
    return (
      <path
        transform={translate ? `translate(${translate})` : null}
        {...paint(color)}
      >
        <animate
          id={id}
          attributeName="d"
          values="m0 0 0 0 0 0 0 0z; m0 -100 200 100 -200 100 -200 -100z"
          dur=".25s"
          begin={begin}
          {...animAttrs()}
        />
        <animateMotion
          path="M0 50 0 0"
          dur=".5s"
          begin={`${id}.begin`}
          {...animAttrs()}
        />
        <animate
          attributeName="opacity"
          values="0;1"
          repeatCount="3"
          dur=".1s"
          begin={`${id}.begin + .1s`}
        />
      </path>
    )
  }

  const Cube = ({
    color = 'synapse',
    translate,
    begin = 0,
    restart,
    children,
  }: {
    color?: string
    translate?: string
    begin?: number | string
    restart?: string
    children?: React.ReactNode
  }) => {
    if (typeof begin === 'number') begin = begin + 's'
    return (
      <g
        transform={translate ? `translate(${translate})` : undefined}
        stroke={stroke[color]}
      >
        <animate
          attributeName="opacity"
          begin={begin}
          dur=".375s"
          values=".25; 1"
          fill="freeze"
          restart={restart}
        />
        <animateMotion
          path="m0 0 v-6.25"
          additive="sum"
          begin={begin}
          dur=".375s"
          calcMode="spline"
          keyPoints="0; 1; 0"
          keyTimes="0; .5; 1"
          keySplines=".5 0 .2 1; .5 0 .2 1"
          fill="freeze"
          restart={restart}
        />
        {children}
        <path fill={fill[color]} vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m0,12.5 25,-12.5 0,0 -25,-12.5 -25,12.5 0,0 25,12.5; m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
            begin={begin}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
            restart={restart}
          />
        </path>
        <path fill="none" vectorEffect="non-scaling-stroke">
          <animate
            attributeName="d"
            values="m-25,0 25,12.5 25,-12.5 m-25,12.5 0,0; m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
            begin={begin}
            dur=".25s"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines=".5 0 .2 1"
            fill="freeze"
            restart={restart}
          />
        </path>
      </g>
    )
  }

  return (
    <Wrapper>
      <Hero />
      <svg
        id="hero-graphic"
        width="1200"
        height="675"
        viewBox="-700 -437.5 1400 875"
        // className={`border border-zinc-900 mx-auto my-8`}
        // stroke="#0ff"
        // stroke-width="1"
        fill="none"
      >
        <style>
          {`text { font-size: 90%; letter-spacing: 1px }`}
          {/* {`@keyframes circlePulse { from { r: 50; } to { r: 100; } }`} */}
        </style>
        <defs>
          <linearGradient id="a">
            <stop stop-color="#E54DE5" />
            <stop offset="1" stop-color="#B580FF" />
          </linearGradient>
          <marker id="b" viewBox="-.5 -.5 1 1">
            <circle r=".4" fill="url(#a)" />
          </marker>
        </defs>

        <path {...paint('synapse')}>
          <animate
            id="bridgeNw"
            attributeName="d"
            values="m-200,-100 0,0 0,0 0,0z; m-100,-150 0,0 -200,100 0,0z; m-120,-160 40,20 -200,100 -40,-20z"
            dur=".5s"
            begin="platformYellow.end + 1s"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path transform="translate(400 200)" {...paint('synapse')}>
          <animate
            id="bridgeSe"
            attributeName="d"
            values="m-200,-100 0,0 0,0 0,0z; m-100,-150 0,0 -200,100 0,0z; m-120,-160 40,20 -200,100 -40,-20z"
            dur=".5s"
            begin="airDropOut.end + 1s"
            restart="never"
            calcMode="spline"
            keyTimes="0; .5; 1"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
        </path>
        <path transform="translate(-40 -240)" {...paint('synapse')}>
          <animate
            id="dockN"
            attributeName="d"
            values="m0 0 0 0 0 0 0 0z; m-40 -60 40 20 -120 60 -40 -20z"
            dur=".5s"
            begin="bridgeNw.end + 3.5s"
            {...animAttrs()}
          />
        </path>
        <path transform="translate(-440 -40)" {...paint('synapse')}>
          <animate
            id="dockW"
            attributeName="d"
            values="m0 0 0 0 0 0 0 0z; m-40 -60 40 20 -120 60 -40 -20z"
            dur=".5s"
            begin="dockN.begin + 1s"
            {...animAttrs()}
          />
        </path>

        <path transform="translate(200 -200)" {...paint('synapse')}>
          <animate
            id="airpadN"
            attributeName="d"
            dur=".5s"
            values="m0 0 0 0 0 0 0 0z; m0 -30 60 30 -60 30 -60 -30z"
            begin="bargeCross.end"
            restart="never"
            {...animAttrs()}
          />
          <animateMotion
            path="m0 30 0 0"
            dur=".5s"
            begin="airpadN.begin"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpadN.begin + .1s"
          />
        </path>
        <path transform="translate(200 -60)" {...paint('synapse')}>
          <animate
            id="airpadE"
            attributeName="d"
            dur=".5s"
            values="m0 0 0 0 0 0 0 0z; m0 -30 60 30 -60 30 -60 -30z"
            begin="airpadN.begin + .5s"
            {...animAttrs()}
          />

          <animateMotion
            path="m0 30 0 0"
            dur=".5s"
            begin="airpadE.begin"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            values="0;1"
            repeatCount="3"
            dur=".1s"
            begin="airpadE.begin + .1s"
          />
        </path>

        <Platform
          id="platformBlue"
          translate="0 -200"
          color="north"
          begin="0s"
        />
        <Platform
          id="platformGreen"
          translate="400 0"
          color="east"
          begin=".1s"
        />
        <Platform
          id="platformOrange"
          translate="0 200"
          color="south"
          begin=".2s"
        />
        <Platform
          id="platformYellow"
          translate="-400 0"
          color="west"
          begin=".3s"
        />

        <g transform="translate(125 -290) scale(1.2)">
          <rect width="100" height="52" rx="4" fill="#111" stroke="#333" />
          <text x="10" y="20" fill={text.north}>
            Arbitrum
          </text>
          <text x="10" y="42" fill="white" opacity=".8">
            $2.55B
          </text>
        </g>

        <g transform="translate(-380 -140) scale(1.2)">
          <rect width="100" height="52" rx="4" fill="#111" stroke="#333" />
          <text x="10" y="20" fill={text.west}>
            Ethereum
          </text>
          <text x="10" y="42" fill="white" opacity=".8">
            $1.72B
          </text>
        </g>

        <g transform="translate(425 52) scale(1.2)">
          <rect width="100" height="52" rx="4" fill="#111" stroke="#333" />
          <text x="10" y="20" fill={text.east}>
            Avalanche
          </text>
          <text x="10" y="42" fill="white" opacity=".8">
            $885.41M
          </text>
        </g>

        <g transform="translate(50 240) scale(1.2)">
          <rect width="110" height="52" rx="4" fill="#111" stroke="#333" />
          <text x="10" y="20" fill={text.south}>
            New! Blast
          </text>
          <text x="10" y="42" fill="white" opacity=".8">
            $452.12M
          </text>
        </g>

        {/* <g transform="translate(-600 -400) scale(1.2)">
            <rect width="120" height="52" rx="4" fill="#111" stroke="#333" />
            <text x="10" y="20" fill={text.west}>
              60d volume
            </text>
            <text x="10" y="42" fill="white" opacity=".8">
              $7.01B
            </text>
          </g> */}

        <path
          id="barge"
          d="m50,-75 100,50 -200,100 -100,-50z"
          transform="translate(250 -515)"
          {...paint('synapse')}
        >
          <animateMotion
            id="bargeOut"
            dur="2s"
            begin="dockN.end; bargeIn.end"
            path="m0 0 -440 220"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeCross"
            dur="2s"
            begin="bargeOut.end + 2s"
            path="m0 0 -400 200"
            additive="sum"
            {...animAttrs()}
          />
          <animateMotion
            id="bargeIn"
            dur="2s"
            begin="bargeCross.end + 2s"
            path="m0 0 -400 200"
            additive="sum"
            {...animAttrs()}
          />
        </path>

        <Cube
          color="east"
          translate="400 0"
          begin="teleportE.begin + 4s"
          restart="never"
        />

        <Cube
          color="east"
          translate="375 12.5"
          begin="airDropOut.end + .5s"
          restart="never"
        />

        <Cube color="south" translate="0 150" begin=".7s" />
        <Cube color="south" translate="25 162.5" begin={20} />
        <Cube color="south" translate="-25 162.5" begin={21} />
        <Cube color="south" translate="0 122.05" begin={22} />

        <Cube
          translate="100 -225"
          color="north"
          begin="airpadN.end; airDropOut.end + 1ms"
        >
          <animateMotion
            id="airCubeOut"
            path="m0 0 100 50"
            begin="airpadE.end + 2s; airDropOut.end + 2s"
            dur="1s"
            {...animAttrs()}
          />
          <animateMotion
            id="airLift"
            additive="sum"
            path="m0 0 v-50 150"
            begin="airCubeOut.end"
            keyPoints="0; .25; 1"
            keyTimes="0; .33; 1"
            calcMode="spline"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            dur="6s"
            fill="freeze"
          />
          <animateMotion
            id="airDrop"
            additive="sum"
            path="m0 0 v40"
            begin="airLift.end"
            dur=".25s"
            {...animAttrs(0.33, 1)}
          />
          <animate
            attributeName="stroke"
            begin="airDrop.begin"
            dur=".33s"
            {...flashAttrs('north', 'east')}
          />
          <animateMotion
            id="airDropOut"
            additive="sum"
            path="m0 0 137.5 68.75 12.5 -8.75"
            begin="airDrop.end"
            dur=".75s"
            {...animAttrs()}
          />
          <set attributeName="opacity" to="0" begin="airDropOut.end" />
          <set
            attributeName="stroke"
            to={stroke.north}
            begin="airDropOut.end"
          />
          <animateMotion path="m0 0" begin="airDropOut.end" />
        </Cube>

        <g id="balloon" transform="translate(200 -200)">
          <animateMotion
            dur="6s"
            begin="airLift.begin"
            path="m0 0 v-50 150"
            keyPoints="0; .25; 1"
            keyTimes="0; .33; 1"
            calcMode="spline"
            keySplines=".5 0 .2 1; .5 0 .2 1"
            fill="freeze"
          />
          <animateMotion
            dur="1s"
            begin="airDrop.begin"
            additive="sum"
            path="m0 0 v-400"
            calcMode="spline"
            keyTimes="0; 1"
            keySplines="1 0 1 1"
            fill="freeze"
          />
          <path stroke={stroke.synapse}>
            <animate
              attributeName="d"
              values="m0 -2 v0; m0 -50 v48"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </path>
          <circle {...paint('synapse')}>
            <animate
              attributeName="r"
              values="0; 36"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
            <animateMotion
              path="m0 0 v-86"
              begin="airLift.begin"
              dur="2s"
              {...animAttrs()}
            />
          </circle>
        </g>

        <g transform="translate(350 25)" stroke={stroke.east} opacity="0">
          <set
            id="airBridge"
            attributeName="stroke"
            to={stroke.east}
            begin="airDropOut.end"
          />
          <set
            id="airBridge"
            attributeName="opacity"
            to="1"
            begin="airDropOut.end"
          />
          <animateMotion begin="airBridge.begin" path="m0 0" />
          <animateMotion
            dur="1s"
            begin="airBridge.begin + 2s"
            path="m0 0 -325 162.5 -25 -12.5"
            {...animAttrs()}
          />
          <AnimateFlash from="east" to="south" begin="airBridge.begin + 2.3s" />
          <animate
            attributeName="opacity"
            values="1; 0"
            dur="4s"
            begin="airBridge.begin + 4s"
            {...animAttrs()}
          />
          <path
            fill={fill.green}
            d="m0,12.5 25,-12.5 0,-27.95 -25,-12.5 -25,12.5 0,27.95 25,12.5"
            vectorEffect="non-scaling-stroke"
          />
          <path
            d="m-25,-27.95 25,12.5 25,-12.5 m-25,12.5 0,27.95"
            vectorEffect="non-scaling-stroke"
          />
        </g>

        <Cube color="north" translate="-25 -212.5" begin="bargeOut.begin">
          <set
            attributeName="stroke"
            to={stroke.north}
            begin="bargeOut.begin"
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeOut.end"
            dur="1s"
            by="-162.5 -81.25"
            {...animAttrs()}
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeCross.begin"
            dur="2s"
            by="-400 200"
            {...animAttrs()}
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeCross.end"
            dur="1s"
            by="137.5 68.75"
            {...animAttrs()}
          />
          <animate
            attributeName="opacity"
            begin="bargeCross.end + 1s"
            to="0"
            dur="2s"
            fill="freeze"
          />
          <animate
            attributeName="opacity"
            begin="bargeOut.begin"
            to="1"
            dur="2s"
            fill="freeze"
          />
          <animateTransform
            attributeName="transform"
            type="translate"
            begin="bargeOut.begin"
            dur="1ms"
            to="-25 -212.5"
            fill="freeze"
          />
          <animate
            attributeName="stroke"
            begin="bargeCross.begin + .5s"
            dur=".5s"
            {...flashAttrs('north', 'west')}
          />
        </Cube>

        <Cube color="north" translate="0 -200" begin=".5s" />

        {/* Simple Bridge Blue/Yellow Swap */}

        <Cube color="north" translate="-25 -187.5" begin="bridgeNw.end + .5s">
          <animateMotion
            id="bridgeCubeOut"
            dur="4s"
            begin="bridgeNw.end + 2.5s; bridgeCubeOut.end + 2s"
            path="M0 0 -350 175"
            calcMode="spline"
            keyPoints="0; 1; 1; 0"
            keyTimes="0; .25; .75; 1"
            keySplines=".5 0 .2 1; 0 0 1 1; .5 0 .2 1"
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${stroke.north}; ${stroke.synapse}; ${stroke.west}; ${stroke.west}; ${stroke.synapse}; ${stroke.north}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
          <animate
            attributeName="fill"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${fill.blue}; ${fill.synapse}; ${fill.yellow}; ${fill.yellow}; ${fill.synapse}; ${fill.blue}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
        </Cube>

        <Cube color="west" translate="-375 -12.5" begin="bridgeNw.end + 1s">
          <animateMotion
            dur="4s"
            begin="bridgeNw.end + 2.5s; bridgeCubeOut.end + 2s"
            path="M0 0 350 -175"
            calcMode="spline"
            keyPoints="0; 1; 1; 0"
            keyTimes="0; .25; .75; 1"
            keySplines=".5 0 .2 1; 0 0 1 1; .5 0 .2 1"
          />
          <animate
            attributeName="stroke"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${stroke.west}; ${stroke.synapse}; ${stroke.north}; ${stroke.north}; ${stroke.synapse}; ${stroke.west}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
          <animate
            attributeName="fill"
            begin="bridgeCubeOut.begin + .3s;"
            values={`${fill.yellow}; ${fill.synapse}; ${fill.blue}; ${fill.blue}; ${fill.synapse}; ${fill.yellow}`}
            keyTimes="0; .06; .12; .88; .94; 1"
            dur="3.4s"
          />
        </Cube>

        <g stroke={stroke.east} transform="translate(520,0)">
          <ellipse {...paint('synapse')}>
            <animate
              begin="teleportE.begin"
              attributeName="rx"
              by="30"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin"
              attributeName="ry"
              by="15"
              dur=".5s"
              {...animAttrs()}
            />
          </ellipse>
          <Cube
            color="east"
            translate="-120"
            begin="airLift.begin"
            restart="never"
          >
            <animateTransform
              attributeName="transform"
              type="translate"
              by="120 -12.5"
              dur="1s"
              begin="teleportE.begin"
              {...animAttrs()}
            />
            <animateMotion
              dur="2s"
              begin="teleportE.begin + .33s"
              path="m0 0 v-12.5"
              calcMode="spline"
              keyPoints="0; 1; 0"
              keyTimes="0; .5; 1"
              keySplines=".33 0 .67 1; .33 0 .67 1"
              repeatCount="indefinite"
            />
            <animate
              attributeName="stroke"
              begin="teleportBeams.begin + 1s"
              {...flashAttrs('east', 'south', '3s')}
            />
            <animate
              attributeName="stroke"
              begin="teleportBeams.begin + 8s"
              {...flashAttrs('south', 'east', '3s')}
            />
          </Cube>
          <path
            d="m-20 0 v-56.25 m10 3.5 v56.25 m10 3 v-62.5 m10 3.25 v56.25 m10 -3.5 v-56.25"
            opacity="0"
            strokeWidth="3"
            stroke={stroke.flash}
            strokeDasharray="8 6 6 8"
          >
            <animate
              id="teleportBeams"
              attributeName="opacity"
              values="0; 1; 0; 0; 1; 0"
              keyTimes="0; .2; .4; .6; .8; 1"
              begin="teleportE.begin + 2s; teleportBeams.end + 2s"
              dur="12s"
            />
            <animate
              attributeName="stroke-dashoffset"
              values="0; 28"
              dur="1s"
              repeatCount="indefinite"
            />
          </path>
          <ellipse {...paint('synapse')}>
            <animate
              id="teleportE"
              begin="airDropOut.end"
              restart="never"
              attributeName="cy"
              by="-75"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin"
              attributeName="rx"
              by="30"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin"
              attributeName="ry"
              by="15"
              dur=".5s"
              {...animAttrs()}
            />
          </ellipse>
        </g>

        <g stroke={stroke.south} transform="translate(0,260)">
          <ellipse {...paint('synapse')}>
            <animate
              begin="teleportE.begin"
              attributeName="rx"
              by="30"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin"
              attributeName="ry"
              by="15"
              dur=".5s"
              {...animAttrs()}
            />
          </ellipse>
          <Cube color="south" begin="bargeCross.begin" restart="never">
            <AnimateFlash
              from="south"
              to="east"
              begin="teleportBeams.begin + 1s"
              dur="3s"
            />
            <AnimateFlash
              from="east"
              to="south"
              begin="teleportBeams.begin + 8s"
              dur="3s"
            />
            <animateMotion
              dur="2s"
              path="m0 -25 v12.5"
              calcMode="spline"
              keyPoints="0; 1; 0"
              keyTimes="0; .5; 1"
              keySplines=".33 0 .67 1; .33 0 .67 1"
              repeatCount="indefinite"
            />
          </Cube>
          <path
            d="m-20 0 v-56.25 m10 3.5 v56.25 m10 3 v-62.5 m10 3.25 v56.25 m10 -3.5 v-56.25"
            opacity="0"
            strokeWidth="3"
            stroke={stroke.flash}
            strokeDasharray="8 6 6 8"
          >
            <animate
              attributeName="opacity"
              values="0; 1; 0; 0; 1; 0"
              keyTimes="0; .2; .4; .6; .8; 1"
              begin="teleportBeams.begin"
              dur="12s"
            />
            <animate
              attributeName="stroke-dashoffset"
              values="0; 28"
              dur="1s"
              repeatCount="indefinite"
            />
          </path>
          <ellipse {...paint('synapse')}>
            <animate
              begin="teleportE.begin + .5s"
              attributeName="cy"
              by="-75"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin + .5s"
              attributeName="rx"
              by="30"
              dur=".5s"
              {...animAttrs()}
            />
            <animate
              begin="teleportE.begin + .5s"
              attributeName="ry"
              by="15"
              dur=".5s"
              {...animAttrs()}
            />
          </ellipse>
        </g>

        <Cube color="west" translate="-400 0" begin=".8s" />

        <Cube color="west" translate="-450 0" begin={15} />
        <Cube color="west" translate="-425 12.5" begin={16} />
        <Cube color="west" translate="-425 -27.95" begin={18} />

        <Cube color="north" translate="25 -187.5" begin={17} />
        <Cube color="north" translate="0 -175" begin={19} />

        {/* <g {...paint('synapse')}>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="M0,12.5 25,0 25,-27.95 0,-40.45 -25,-27.95 -25,0z; M17.7,8.85 17.7,-21.15 17.7,-38.85 -17.7,-38.85 -17.7,-21.15 -17.7,8.85z; M25,0 25,-27.95 0,-40.45 -25,-27.95 -25,0 0,12.5z"
              dur="1s"
              repeatCount="indefinite"
              fill="freeze"
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="M0,12.5 0,-15.45 25,-27.95 M0,-15.45 -25,-27.95; M17.7,8.85 17.7,-21.15 17.7,-38.85 M17.7,-21.15 -17.7,-21.15; M0 0 0 0 0 0 M0 0 0 0"
              dur="1s"
              repeatCount="indefinite"
            />
            <animate
              attributeName="visibility"
              values="visible; hidden"
              dur="1s"
              repeatCount="indefinite"
            />
          </path>
          <path vectorEffect="non-scaling-stroke">
            <animate
              attributeName="d"
              values="M0 0 0 0 0 0 M0 0 0 0; M-17.7,8.85 -17.7,-21.15 -17.7,-38.85 M-17.7,-21.15 17.7,-21.15; M0,12.5 0,-15.45 -25,-27.95 M0,-15.45 25,-27.95"
              dur="1s"
              repeatCount="indefinite"
            />
            <animate
              attributeName="visibility"
              values="hidden; visible"
              dur="1s"
              repeatCount="indefinite"
            />
          </path>
        </g> */}
        {/* <path
          d="M0 18 18 0 -18 0 0 -18"
          stroke-width="5"
          stroke-linejoin="bevel"
          stroke="url(#a)"
          stroke-opacity=".5"
          marker-start="url(#b)"
          marker-mid="url(#b)"
          marker-end="url(#b)"
          // transform="matrix(1 .5 -1 .5 0 0) rotate(-45)"
          transform="scale(1.25)"
        >
          <animateTransform
            attributeName="transform"
            attributeType="XML"
            type="rotate"
            by="360"
            // dur="2s"
            repeatCount="indefinite"
          />
        </path> */}
      </svg>
      <section>
        <ul className="w-fit md:w-max grid md:flex text-xl md:text-lg text-center items-center place-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 gap-x-8 -mt-8 shadow-sm mx-auto mb-16 cursor-default">
          <li className="-mt-1 p-3">
            $<data className="mx-0.5">45.3B</data> Bridge volume
          </li>
          <li className="-mt-1 p-3">
            <data className="mx-0.5">10.6M</data> transactions
          </li>
          <li className="-mt-1 p-3">
            $<data className="mx-0.5">116.7M</data> Total value locked
          </li>
        </ul>
      </section>
      <section>
        <h2 className="text-4xl font-medium">The best brands use Synapse</h2>
      </section>
      <section className="flex odd:flex-col even:flex-col-reverse md:grid grid-cols-2 gap-12 items-center p-4 max-w-4xl">
        <div>
          <h2 className="text-4xl font-medium mb-4">Interchain apps</h2>
          <p className="text-lg leading-relaxed mb-4">
            Synapse Bridge is built on top of the cross-chain infrastructure
            enabling users to seamlessly transfer assets across all blockchains.
            The Bridge has become the most widely-used method to move assets
            cross-chain, offering low cost, fast, and secure bridging.
          </p>
        </div>
        <div className="grid justify-center">
          <FauxBridge />
        </div>
      </section>
    </Wrapper>
  )
}

export default LandingPage