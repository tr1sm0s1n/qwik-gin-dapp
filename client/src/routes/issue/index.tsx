import { component$, useContext } from '@builder.io/qwik';
import { Form, type DocumentHead, routeAction$ } from '@builder.io/qwik-city';
import { AuthContext } from '../layout';

export const useIssue = routeAction$(async (candidate) => {
  console.log(candidate);
  const res = await fetch(`${import.meta.env.PUBLIC_API}/issue`, {
    method: 'POST',
    body: JSON.stringify(candidate),
    headers: {
      'Content-Type': 'application/json',
    },
  });
  console.log(await res.json());

  return {
    status: res.status,
    candidate,
  };
});

export default component$(() => {
  const authorized = useContext(AuthContext);
  const action = useIssue();

  return authorized.value ? (
    <section class="flex justify-center m-8">
      <div class="w-full max-w-4xl p-4 bg-white border border-gray-200 rounded-lg shadow">
        {!action.value?.status ? (
          <>
            <h3 class="text-2xl font-medium text-gray-900 text-center mb-4">
              Issue Certificate
            </h3>
            <Form action={action}>
              <div class="relative z-0 w-full mb-6 group">
                <input
                  type="number"
                  min={1}
                  name="candidateID"
                  id="candidateID"
                  class="block py-2.5 px-0 w-full text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                  placeholder=" "
                  required
                />
                <label
                  for="candidateID"
                  class="peer-focus:font-medium absolute text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                >
                  Candidate ID
                </label>
              </div>
              <div class="relative z-0 w-full mb-6 group">
                <input
                  type="text"
                  name="candidateName"
                  id="candidateName"
                  class="block py-2.5 px-0 w-full text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                  placeholder=" "
                  required
                />
                <label
                  for="candidateName"
                  class="peer-focus:font-medium absolute text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                >
                  Candidate Name
                </label>
              </div>
              <div class="relative z-0 w-full mb-6 group">
                <select
                  id="courseName"
                  name="courseName"
                  class="block py-2.5 px-0 w-full text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                  required
                >
                  <option selected disabled hidden>
                    Choose a course
                  </option>
                  <option value="CBA">CBA</option>
                  <option value="DEB">DEB</option>
                  <option value="CED">CED</option>
                  <option value="WEB3">WEB3</option>
                </select>
                <label
                  for="courseName"
                  class="peer-focus:font-medium absolute text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                >
                  Course
                </label>
              </div>
              <div class="grid md:grid-cols-2 md:gap-6">
                <div class="relative z-0 w-full mb-6 group">
                  <select
                    id="courseGrade"
                    name="courseGrade"
                    class="block py-2.5 px-0 w-full text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                    required
                  >
                    <option selected disabled hidden>
                      Choose a grade
                    </option>
                    <option value="S">S</option>
                    <option value="A">A</option>
                    <option value="B">B</option>
                    <option value="C">C</option>
                  </select>
                  <label
                    for="courseGrade"
                    class="peer-focus:font-medium absolute text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                  >
                    Grade
                  </label>
                </div>
                <div class="relative z-0 w-full mb-6 group">
                  <input
                    type="date"
                    name="courseDate"
                    id="courseDate"
                    class="block py-2.5 px-0 w-full text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer"
                    required
                  />
                  <label
                    for="courseDate"
                    class="peer-focus:font-medium absolute text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
                  >
                    Date
                  </label>
                </div>
              </div>
              <button
                type="submit"
                class="text-white bg-blue-700 hover:bg-blue-800 font-medium rounded-lg w-full sm:w-auto px-5 py-2.5 text-center"
              >
                Submit
              </button>
            </Form>
          </>
        ) : (
          <>
            <p class="text-xl text-green-700">
              Certificate issued successfully!
            </p>
          </>
        )}
      </div>
    </section>
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
