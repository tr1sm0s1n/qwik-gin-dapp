import {
  component$,
  createContextId,
  type Signal,
  Slot,
  useContextProvider,
  useSignal,
} from '@builder.io/qwik';
import { NavBar } from '~/components/nav-bar';

export const AuthContext = createContextId<Signal<boolean>>('authorized');

export default component$(() => {
  const authorized = useSignal(false);
  useContextProvider(AuthContext, authorized);
  return (
    <>
      <NavBar />
      <Slot />
    </>
  );
});
