import { component$ } from '@builder.io/qwik';
import { Form, type DocumentHead, routeAction$ } from '@builder.io/qwik-city';

type Certificate = {
  id: number;
  name: string;
  course: string;
  date: string;
  grade: string;
};

export const useFetch = routeAction$(async (input) => {
  const res = await fetch(
    `${import.meta.env.PUBLIC_API}/fetch?id=${input.certificateID}`
  );
  const certData: Certificate = await res.json();

  return {
    success: res.status === 200,
    data: certData,
  };
});

export default component$(() => {
  const action = useFetch();

  return (
    <>
      <section class="my-8 text-center">
        <h1 class="text-4xl text-gray-700 my-4">Certificate DApp</h1>
        <p class="text-lg text-cyan-700">Verification Portal</p>
      </section>
      <section>
        <Form class="mx-auto flex items-center max-w-xs" action={action}>
          <div class="relative w-full">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
              <svg
                aria-hidden="true"
                class="w-5 h-5 text-gray-500"
                fill="currentColor"
                viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg"
                data-darkreader-inline-fill=""
              >
                <path
                  fill-rule="evenodd"
                  d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                  clip-rule="evenodd"
                ></path>
              </svg>
            </div>
            <input
              type="text"
              name="certificateID"
              id="certificateID"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5"
              placeholder="Enter the certificate ID"
              required
            />
          </div>
        </Form>
      </section>
      {action.value?.success && (
        <section class="flex justify-center m-8">
          <div class="w-full max-w-4xl p-4 bg-white border border-gray-200 rounded-lg shadow">
            <h3 class="text-2xl font-medium text-gray-900 text-center my-6">
              Certificate of Completion
            </h3>
            <p class="text-gray-700 text-center italic">
              This certificate is awarded to
            </p>
            <p class="text-2xl text-center my-3">{action.value?.data?.name}</p>
            <p class="text-gray-700 text-center italic">
              in recognition of their successful completion of
            </p>
            <p class="text-xl text-center my-3">{action.value?.data?.course}</p>
            <p class="text-gray-700 text-center italic">
              with a{' '}
              <span class="text-red-800">{action.value?.data?.grade}</span>{' '}
              Grade.
            </p>
            <div class="flex justify-between mx-4 my-4">
              <div class="text-start">
                <h2 class="text-lg font-semibold">Certificate ID:</h2>
                <p class="text-gray-700">{action.value?.data?.id}</p>
              </div>
              <div class="text-end">
                <h2 class="text-lg font-semibold">Date:</h2>
                <p class="text-gray-700">{action.value?.data?.date}</p>
              </div>
            </div>
          </div>
        </section>
      )}
    </>
  );
});

export const head: DocumentHead = {
  title: 'Home | Certificate DApp',
  meta: [
    {
      name: 'description',
      content: 'Qwik site description',
    },
  ],
};
