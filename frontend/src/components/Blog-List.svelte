<script>
// @ts-nocheck

    export let perm; 
    import axios from "axios" 
    let error = "";
    let data = [];
    async function list() {
      try {
        const response = await axios.post("http://localhost:3000/api/v1/blog/list");
        if (response.data.status === "OK") {
            console.log(response.data);
            data = response.data.data
        } else {
          error = "Email Yada Parola hatalı";
        }
      } catch (error) {
        console.error("Hata oluştu:", error);
      }
    }
    list()

    async function deleteb(token) {
        try {
        const form = {
        token: token,
        };   
        const response = await axios.post("http://localhost:3000/api/v1/blog/del", form);
        if (response.data.status === "OK") {
            console.log(response.data);
            error = response.data.message
        } else {
          error = "Silinemedi";
        }
      } catch (error) {
        console.error("Hata oluştu:", error);
      }
    }
</script>

<main class="p-8">
    <div class="mb-4 text-red-500">{error}</div>
    {#if perm == "admin"}
      <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" class="px-4 py-2">Başlık</th>
            <th scope="col" class="px-4 py-2">Link</th>
            <th scope="col" class="px-4 py-2">Sil</th>
          </tr>
        </thead>
        <tbody>
          {#each data as datas}
            <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
              <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                {datas.title}
              </th>
              <td class="px-6 py-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded"
                  on:click={() => window.location.href = `/blogs/${datas.title}`}
                >
                  🔗
                </button>
              </td>
              <td class="px-6 py-4">
                <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded"
                  on:click={() => deleteb(datas.token)}
                >
                  🚯
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    {/if}
    {#if perm == "user"}
      <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" class="px-4 py-2">Başlık</th>
            <th scope="col" class="px-4 py-2">Link</th>
          </tr>
        </thead>
        <tbody>
          {#each data as datas}
            <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
              <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                {datas.title}
              </th>
              <td class="px-6 py-4">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded"
                  on:click={() => window.location.href = `/blogs/${datas.title}`}
                >
                  🔗
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    {/if}
  </main>
