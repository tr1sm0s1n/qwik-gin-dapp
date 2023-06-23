import { component$, Slot } from '@builder.io/qwik';
import { NavBar } from '~/components/nav-bar';

export default component$(() => {
  return (
    <>
      <NavBar />
      <Slot />
    </>
  );
});
