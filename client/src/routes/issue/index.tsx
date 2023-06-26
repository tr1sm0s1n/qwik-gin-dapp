import { component$, useContext } from '@builder.io/qwik';
import { type DocumentHead } from '@builder.io/qwik-city';
import { AuthContext } from '../layout';

export default component$(() => {
  const authorized = useContext(AuthContext);

  return authorized.value ? (
    <>
      <section class="my-8">
        <h1 class="text-4xl text-gray-700 my-4">Issue Certificate</h1>
        {/* <p class="text-lg text-cyan-700">Verification Portal</p> */}
      </section>
    </>
  ) : null;
});

export const head: DocumentHead = {
  title: 'Issue | Certificate DApp',
  meta: [
    {
      name: 'description',
      content: 'Qwik site description',
    },
  ],
};
