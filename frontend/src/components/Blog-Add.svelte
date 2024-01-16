<script>
    import axios from "axios";
  
    let description = "";
    let title = "";
    let error = "";
  
    async function add() {
      const form = {
        title: title,
        description: description,
      };
  
      try {
        const response = await axios.post(
          "http://localhost:3000/api/v1/blog/add",
          form
        );
        if (response.data.status === "OK") {
            console.log(response.data);
            window.location.href = "/dashboard"
        } else {
          error = "Büyük ihtimalle boşluk var ";
        }
      } catch (error) {
        console.error("Hata oluştu:", error);
      }
    }
  </script>
  
  <main class="p-8 bg-gray-900">
    {#if error}
      <div class="mb-4">
        <p class="font-bold text-red-500">{error}</p>
      </div>
    {/if}
  
    <div>
      <form on:submit|preventDefault={add}>
        <label class="text-white" for="title">Title:</label>
        <input
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          type="word"
          bind:value={title}
          id="title"
        />
  
        <label for="description" class="text-white">Description:</label>
        <textarea
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          bind:value={description}
          id="description"
          style="width: 700px; height: 400px; resize: none;"
        ></textarea>
  
        <button
          type="submit"
          class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          Ekle
        </button>
      </form>
    </div>
  </main>
  
  