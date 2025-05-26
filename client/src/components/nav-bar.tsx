import { component$, $, useStore, useContext } from '@builder.io/qwik'
import { Link, useNavigate } from '@builder.io/qwik-city'
import ImgLogo from '~/media/logo.svg?jsx'
import ImgMetamask from '~/media/metamask.svg?jsx'
import { AuthContext } from '~/routes/layout'

export const NavBar = component$(() => {
  const authorized = useContext(AuthContext)
  const nav = useNavigate()
  const connection = useStore({
    account: '',
    status: false,
    admin: false,
  })

  const connectToMetaMask = $(async () => {
    const ethereum = window.ethereum
    if (ethereum) {
      const wallet = await ethereum.request({ method: 'eth_requestAccounts' })
      const sign = await window.ethereum.request({
        method: 'personal_sign',
        params: [
          '0x57656c636f6d6520746f20436572746966696361746520444170702e204b696e646c79207369676e2074686973206d65737361676520746f2070726f636565642e20546869732070726f63656475726520646f6573206e6f74207265717569726520616e792045544820746f2070726f636573732e',
          wallet[0],
        ],
      })
      console.log(sign)

      connection.account = wallet[0]
      connection.status = true
      connection.admin = sign === import.meta.env.PUBLIC_ADMIN_SIGN

      if (connection.admin) {
        console.log('Welcome Admin')
        authorized.value = true
        await nav('/issue')
      } else {
        console.log('Welcome Guest')
      }
    }
  })

  return (
    <>
      <nav class="flex items-center justify-between flex-wrap bg-rose-800 p-6">
        <Link href="/" class="flex items-center flex-shrink-0 text-white mr-6">
          <div class="h-8 w-8 mx-2">
            <ImgLogo />
          </div>
          <span class="font-semibold text-xl tracking-tight">
            Certificate DApp
          </span>
        </Link>
        <div class="flex justify-end w-auto">
          <div>
            <button
              class="flex items-center text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-rose-600 hover:bg-white"
              onClick$={connectToMetaMask}
            >
              {connection.status ? (
                <>
                  <div class="h-6 w-6 mr-3">
                    <ImgMetamask />
                  </div>
                  <span>
                    {connection.account.slice(0, 5)}...
                    {connection.account.slice(-4)}
                  </span>
                </>
              ) : (
                <>
                  <span>Connect</span>
                  <div class="h-6 w-6 ml-3">
                    <ImgMetamask />
                  </div>
                </>
              )}
            </button>
          </div>
        </div>
      </nav>
    </>
  )
})
