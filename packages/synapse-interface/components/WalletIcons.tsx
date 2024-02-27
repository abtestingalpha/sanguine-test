import { ReactNode } from 'react'

const SvgIconWrapper = ({ width, height, children }): ReactNode => (
  <svg
    width={width}
    height={height}
    viewBox="0 0 48 48"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    className="rounded-sm"
  >
    {children}
  </svg>
)

export const CoinbaseWallet = ({ width, height }): ReactNode => (
  <SvgIconWrapper width={width} height={height}>
    <g clip-path="url(#clip0_1101_3533)">
      <rect width="48" height="48" fill="url(#paint0_linear_1101_3533)" />
      <circle cx="24" cy="24" r="24" fill="url(#paint1_linear_1101_3533)" />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M24 38C31.732 38 38 31.732 38 24C38 16.268 31.732 10 24 10C16.268 10 10 16.268 10 24C10 31.732 16.268 38 24 38ZM20.5 19.5C19.9477 19.5 19.5 19.9477 19.5 20.5V27.5C19.5 28.0523 19.9477 28.5 20.5 28.5H27.5C28.0523 28.5 28.5 28.0523 28.5 27.5V20.5C28.5 19.9477 28.0523 19.5 27.5 19.5H20.5Z"
        fill="white"
      />
    </g>
    <defs>
      <linearGradient
        id="paint0_linear_1101_3533"
        x1="24"
        y1="0"
        x2="24"
        y2="48"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="#345FF7" />
        <stop offset="1" stopColor="#1D43DB" />
      </linearGradient>
      <linearGradient
        id="paint1_linear_1101_3533"
        x1="24"
        y1="0"
        x2="24"
        y2="48"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="#345FF7" />
        <stop offset="1" stopColor="#1D43DB" />
      </linearGradient>
      <clipPath id="clip0_1101_3533">
        <rect width="48" height="48" fill="white" />
      </clipPath>
    </defs>
  </SvgIconWrapper>
)

export const MetaMask = ({ width, height }): ReactNode => (
  <SvgIconWrapper width={width} height={height}>
    <path
      d="M41.2956 5.34839L26.3051 16.4821L29.0772 9.91336L41.2956 5.34839Z"
      fill="#E2761B"
      stroke="#E2761B"
    />
    <path
      d="M6.68922 5.34839L21.5593 16.5876L18.9227 9.91336L6.68922 5.34839ZM35.902 31.1563L31.9096 37.2731L40.4519 39.6233L42.9077 31.2919L35.902 31.1563ZM5.1073 31.2919L7.54798 39.6233L16.0904 37.2731L12.0979 31.1563L5.1073 31.2919Z"
      fill="#E4761B"
      stroke="#E4761B"
    />
    <path
      d="M15.6083 20.8211L13.2279 24.4218L21.71 24.7985L21.4087 15.6836L15.6083 20.8211ZM32.3767 20.8211L26.501 15.5781L26.3051 24.7985L34.7722 24.4218L32.3767 20.8211ZM16.0904 37.273L21.1827 34.7872L16.7835 31.3521L16.0904 37.273ZM26.8023 34.7872L31.9096 37.273L31.2015 31.3521L26.8023 34.7872Z"
      fill="#E4761B"
      stroke="#E4761B"
    />
    <path
      d="M31.9097 37.273L26.8023 34.7871L27.2091 38.1167L27.1639 39.5178L31.9097 37.273ZM16.0905 37.273L20.8362 39.5178L20.8061 38.1167L21.1827 34.7871L16.0905 37.273Z"
      fill="#D7C1B3"
      stroke="#D7C1B3"
    />
    <path
      d="M20.9114 29.1525L16.6628 27.902L19.661 26.531L20.9114 29.1525ZM27.0734 29.1525L28.3239 26.531L31.337 27.902L27.0734 29.1525Z"
      fill="#233447"
      stroke="#233447"
    />
    <path
      d="M16.0904 37.2731L16.8135 31.1563L12.0979 31.2919L16.0904 37.2731ZM31.1864 31.1563L31.9096 37.2731L35.902 31.2919L31.1864 31.1563ZM34.7721 24.4219L26.3051 24.7985L27.0885 29.1526L28.339 26.5311L31.3521 27.9021L34.7721 24.4219ZM16.6629 27.9021L19.6761 26.5311L20.9115 29.1526L21.71 24.7985L13.2278 24.4219L16.6629 27.9021Z"
      fill="#CD6116"
      stroke="#CD6116"
    />
    <path
      d="M13.2279 24.4219L16.7835 31.3522L16.6629 27.9021L13.2279 24.4219ZM31.3522 27.9021L31.2015 31.3522L34.7722 24.4219L31.3522 27.9021ZM21.71 24.7985L20.9115 29.1526L21.9059 34.29L22.1319 27.5255L21.71 24.7985ZM26.3051 24.7985L25.8983 27.5104L26.0791 34.29L27.0885 29.1526L26.3051 24.7985Z"
      fill="#E4751F"
      stroke="#E4751F"
    />
    <path
      d="M27.0885 29.1526L26.079 34.29L26.8022 34.7872L31.2014 31.3522L31.3521 27.9021L27.0885 29.1526ZM16.6628 27.9021L16.7834 31.3522L21.1826 34.7872L21.9058 34.29L20.9114 29.1526L16.6628 27.9021Z"
      fill="#F6851B"
      stroke="#F6851B"
    />
    <path
      d="M27.1639 39.5178L27.2091 38.1166L26.8325 37.7852H21.1526L20.8061 38.1166L20.8362 39.5178L16.0905 37.2729L17.7477 38.6289L21.1074 40.9641H26.8776L30.2524 38.6289L31.9097 37.2729L27.1639 39.5178Z"
      fill="#C0AD9E"
      stroke="#C0AD9E"
    />
    <path
      d="M26.8023 34.7872L26.0791 34.29H21.9058L21.1827 34.7872L20.806 38.1168L21.1525 37.7853H26.8324L27.209 38.1168L26.8023 34.7872Z"
      fill="#161616"
      stroke="#161616"
    />
    <path
      d="M41.9285 17.2053L43.2091 11.0584L41.2957 5.34839L26.8023 16.1055L32.3767 20.8211L40.2562 23.1262L42.0038 21.0923L41.2505 20.5499L42.4558 19.4501L41.5217 18.7269L42.727 17.8079L41.9285 17.2053ZM4.79102 11.0584L6.07162 17.2053L5.25806 17.8079L6.46333 18.7269L5.54431 19.4501L6.74958 20.5499L5.99629 21.0923L7.72887 23.1262L15.6083 20.8211L21.1827 16.1055L6.68932 5.34839L4.79102 11.0584Z"
      fill="#763D16"
      stroke="#763D16"
    />
    <path
      d="M40.2561 23.1261L32.3766 20.821L34.7721 24.4218L31.2015 31.3521L35.902 31.2919H42.9077L40.2561 23.1261ZM15.6082 20.821L7.72877 23.1261L5.1073 31.2919H12.0979L16.7834 31.3521L13.2278 24.4218L15.6082 20.821ZM26.305 24.7985L26.8022 16.1054L29.0922 9.91333H18.9227L21.1826 16.1054L21.7099 24.7985L21.8907 27.5404L21.9058 34.29H26.0791L26.1092 27.5404L26.305 24.7985Z"
      fill="#F6851B"
      stroke="#F6851B"
    />
  </SvgIconWrapper>
)

export const RabbyWallet = ({ width, height }): ReactNode => (
  <SvgIconWrapper width={width} height={height}>
    <rect width="48" height="48" fill="#8697FF" />
    <path
      d="M41.1066 26.1652C42.4643 23.1221 35.7524 14.62 29.3404 11.0779C25.2987 8.3338 21.0872 8.71079 20.2343 9.91564C18.3624 12.5598 26.4326 14.8003 31.8298 17.4148C30.6697 17.9204 29.5764 18.8277 28.9334 19.9881C26.9213 17.784 22.5049 15.8859 17.3229 17.4148C13.8308 18.4451 10.9286 20.8741 9.80683 24.5428C9.5343 24.4213 9.23251 24.3537 8.91503 24.3537C7.70098 24.3537 6.7168 25.3413 6.7168 26.5594C6.7168 27.7776 7.70098 28.765 8.91503 28.765C9.14006 28.765 9.84368 28.6136 9.84368 28.6136L21.0872 28.6954C16.5907 35.8527 13.0372 36.899 13.0372 38.139C13.0372 39.3789 16.4373 39.0429 17.7139 38.5807C23.8255 36.3681 30.3896 29.4724 31.516 27.4873C36.2461 28.0794 40.2214 28.1495 41.1066 26.1652Z"
      fill="url(#paint0_linear_1560_2236)"
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M31.8294 17.4154C31.8297 17.4155 31.83 17.4156 31.8302 17.4157C32.0805 17.3168 32.04 16.9461 31.9712 16.655C31.8134 15.9859 29.0892 13.2867 26.531 12.0778C23.0452 10.4304 20.4783 10.5154 20.0991 11.2746C20.8092 12.7348 24.101 14.1058 27.5389 15.5376C29.0057 16.1484 30.4991 16.7704 31.83 17.4151C31.8298 17.4152 31.8296 17.4153 31.8294 17.4154Z"
      fill="url(#paint1_linear_1560_2236)"
    />
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M27.4059 32.1116C26.7009 31.8414 25.9046 31.5933 24.9992 31.3683C25.9646 29.6352 26.1671 27.0695 25.2555 25.4473C23.976 23.1707 22.3699 21.959 18.6377 21.959C16.5851 21.959 11.0583 22.6527 10.9603 27.2819C10.95 27.7676 10.96 28.2128 10.995 28.6221L21.0872 28.6955C19.7266 30.8612 18.4524 32.4674 17.3369 33.6887C18.6763 34.0331 19.7816 34.3223 20.7964 34.5877C21.7592 34.8395 22.6406 35.07 23.5631 35.3062C24.9546 34.289 26.2627 33.1798 27.4059 32.1116Z"
      fill="url(#paint2_linear_1560_2236)"
    />
    <path
      d="M9.67209 28.1464C10.0844 31.6629 12.0762 33.041 16.1463 33.4489C20.2163 33.8567 22.551 33.5831 25.6592 33.8668C28.2552 34.1038 30.5731 35.4311 31.433 34.9725C32.2068 34.5596 31.7738 33.0683 30.7384 32.1114C29.3961 30.871 27.5383 30.0087 24.2695 29.7026C24.921 27.913 24.7385 25.4037 23.7267 24.0385C22.2638 22.0645 19.5638 21.1721 16.1463 21.562C12.5759 21.9694 9.15473 23.7331 9.67209 28.1464Z"
      fill="url(#paint3_linear_1560_2236)"
    />
    <defs>
      <linearGradient
        id="paint0_linear_1560_2236"
        x1="16.9162"
        y1="23.4703"
        x2="40.8262"
        y2="30.2279"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="white" />
        <stop offset="1" stopColor="white" />
      </linearGradient>
      <linearGradient
        id="paint1_linear_1560_2236"
        x1="36.7901"
        y1="23.0147"
        x2="19.4885"
        y2="5.72897"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="#8697FF" />
        <stop offset="1" stopColor="#8697FF" stop-opacity="0" />
      </linearGradient>
      <linearGradient
        id="paint2_linear_1560_2236"
        x1="27.8856"
        y1="32.7158"
        x2="11.2937"
        y2="23.2086"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="#8697FF" />
        <stop offset="1" stopColor="#8697FF" stop-opacity="0" />
      </linearGradient>
      <linearGradient
        id="paint3_linear_1560_2236"
        x1="18.3428"
        y1="23.2915"
        x2="29.5857"
        y2="37.5288"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="white" />
        <stop offset="0.983895" stopColor="#D1D8FF" />
      </linearGradient>
    </defs>
  </SvgIconWrapper>
)

export const RainbowWallet = ({ width, height }): ReactNode => (
  <SvgIconWrapper width={width} height={height}>
    <rect width="48" height="48" fill="url(#paint0_linear_1559_2195)" />
    <mask
      id="mask0_1559_2195"
      //   style="mask-type:alpha"
      maskUnits="userSpaceOnUse"
      x="7"
      y="7"
      width="34"
      height="34"
    >
      <path
        d="M10.4145 7.875C26.7533 8.02944 39.9692 21.2453 40.1236 37.5842C40.0785 38.9943 38.9211 40.1236 37.5 40.1236H23.4375C21.9878 40.1236 20.8125 38.9484 20.8125 37.4986C20.8125 31.8032 16.1954 27.1861 10.5 27.1861C9.05025 27.1861 7.875 26.0109 7.875 24.5611V10.4986C7.875 9.07753 9.00431 7.92015 10.4145 7.875Z"
        fill="#D9D9D9"
      />
    </mask>
    <g mask="url(#mask0_1559_2195)">
      <circle
        cx="10.125"
        cy="37.875"
        r="30"
        fill="url(#paint1_radial_1559_2195)"
      />
      <circle
        cx="10.125"
        cy="37.875"
        r="23.5312"
        fill="url(#paint2_radial_1559_2195)"
      />
      <circle
        cx="10.125"
        cy="37.875"
        r="17.1562"
        fill="url(#paint3_radial_1559_2195)"
      />
    </g>
    <defs>
      <linearGradient
        id="paint0_linear_1559_2195"
        x1="24"
        y1="0"
        x2="24"
        y2="48"
        gradientUnits="userSpaceOnUse"
      >
        <stop stopColor="#17429A" />
        <stop offset="1" stopColor="#011E5A" />
      </linearGradient>
      <radialGradient
        id="paint1_radial_1559_2195"
        cx="0"
        cy="0"
        r="1"
        gradientUnits="userSpaceOnUse"
        gradientTransform="translate(10.125 37.875) rotate(-90) scale(30)"
      >
        <stop offset="0.8" stopColor="#ED441E" />
        <stop offset="1" stopColor="#8953C2" />
      </radialGradient>
      <radialGradient
        id="paint2_radial_1559_2195"
        cx="0"
        cy="0"
        r="1"
        gradientUnits="userSpaceOnUse"
        gradientTransform="translate(10.125 37.875) rotate(-90) scale(23.5312)"
      >
        <stop offset="0.729084" stopColor="#FFED05" />
        <stop offset="1" stopColor="#FF980A" />
      </radialGradient>
      <radialGradient
        id="paint3_radial_1559_2195"
        cx="0"
        cy="0"
        r="1"
        gradientUnits="userSpaceOnUse"
        gradientTransform="translate(10.125 37.875) rotate(-90) scale(17.1562)"
      >
        <stop offset="0.622951" stopColor="#05B5E5" />
        <stop offset="1" stopColor="#00DC49" />
      </radialGradient>
    </defs>
  </SvgIconWrapper>
)

export const WalletConnect = ({ width, height }): ReactNode => (
  <SvgIconWrapper width={width} height={height}>
    <path
      d="M9.8303 14.801C17.6559 7.13892 30.3442 7.13892 38.17 14.801L39.1117 15.7232C39.2046 15.8133 39.2784 15.9212 39.3288 16.0403C39.3792 16.1594 39.4051 16.2875 39.4051 16.4168C39.4051 16.5462 39.3792 16.6743 39.3288 16.7934C39.2784 16.9125 39.2046 17.0203 39.1117 17.1104L35.89 20.265C35.7951 20.3571 35.668 20.4086 35.5357 20.4086C35.4035 20.4086 35.2764 20.3571 35.1815 20.265L33.8853 18.996C28.4258 13.6508 19.5743 13.6508 14.1148 18.996L12.7268 20.3551C12.6319 20.4472 12.5048 20.4987 12.3725 20.4987C12.2403 20.4987 12.1132 20.4472 12.0183 20.3551L8.79654 17.2005C8.70371 17.1104 8.62991 17.0026 8.57951 16.8835C8.52911 16.7643 8.50315 16.6363 8.50315 16.5069C8.50315 16.3776 8.52911 16.2495 8.57951 16.1304C8.62991 16.0112 8.70371 15.9034 8.79654 15.8133L9.83014 14.801H9.8303ZM44.833 21.3248L47.7005 24.1324C47.7934 24.2225 47.8672 24.3303 47.9176 24.4494C47.968 24.5686 47.9939 24.6966 47.9939 24.826C47.9939 24.9553 47.968 25.0834 47.9176 25.2025C47.8672 25.3217 47.7934 25.4295 47.7005 25.5196L34.7709 38.1791C34.3796 38.5621 33.7453 38.5621 33.354 38.1791L24.1773 29.1943C24.1299 29.1482 24.0663 29.1225 24.0002 29.1225C23.9341 29.1225 23.8706 29.1482 23.8231 29.1943L14.6465 38.1792C14.2551 38.5623 13.6207 38.5623 13.2293 38.1792L0.299576 25.5196C0.206722 25.4295 0.132901 25.3216 0.0824901 25.2025C0.032079 25.0833 0.00610352 24.9553 0.00610352 24.8259C0.00610352 24.6965 0.032079 24.5684 0.0824901 24.4493C0.132901 24.3301 0.206722 24.2223 0.299576 24.1322L3.1671 21.3247C3.5583 20.9416 4.1927 20.9416 4.58406 21.3247L13.7609 30.3095C13.8083 30.3555 13.8719 30.3813 13.938 30.3813C14.0041 30.3813 14.0676 30.3555 14.1151 30.3095L23.2913 21.3248C23.6826 20.9418 24.3169 20.9416 24.7082 21.3248L33.885 30.3096C33.9325 30.3557 33.996 30.3814 34.0621 30.3814C34.1283 30.3814 34.1918 30.3557 34.2393 30.3096L43.4161 21.3248C43.8074 20.9418 44.4417 20.9418 44.833 21.3248Z"
      fill="#3B99FC"
    />
  </SvgIconWrapper>
)
