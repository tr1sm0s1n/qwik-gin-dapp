import { component$ } from '@builder.io/qwik';
import '../icons/metamask.svg';

export const NavBar = component$(() => {
  return (
    <>
      <nav class="flex items-center justify-between flex-wrap bg-rose-800 p-6">
        <div class="flex items-center flex-shrink-0 text-white mr-6">
          <img class="h-8 w-8 mx-2" src="/logo.svg" alt="logo" />
          <span class="font-semibold text-xl tracking-tight">
            Certificate DApp
          </span>
        </div>
        <div class="flex justify-end w-auto">
          <div>
            <button class="flex items-center text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-rose-600 hover:bg-white">
              <span>Connect</span>
              <img class="h-6 w-6 ml-2" src="/metamask.svg" alt="metamask" />
            </button>
          </div>
        </div>
      </nav>
    </>
  );
});
