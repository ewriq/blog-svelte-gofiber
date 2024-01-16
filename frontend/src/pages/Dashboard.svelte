<script>
  // @ts-nocheck
  import BlogAdd from "../components/Blog-Add.svelte";
  import BlogList from "../components/Blog-List.svelte";
  export let users;
  console.log(users);
  //...
  let showModal = false;
  let dialog;

  $: if (dialog && showModal) dialog.showModal();
</script>

<main class="bg-gray-800 text-white p-8">
  {#if users.perm == "admin"}
    <div class="mb-4">
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
      <dialog
        bind:this={dialog}
        class="fixed inset-0 z-10 overflow-y-auto"
        on:close={() => (showModal = false)}
        on:click|self={() => dialog.close()}
      >
        <div class="bg-gray-900 flex items-center justify-center min-h-screen">
          <div class="bg-gray-900 p-4 rounded shadow-lg">
            <BlogAdd />
          </div>
        </div>
      </dialog>
    </div>
  {/if}

  <div class="p-8 flex items-center justify-between">
    <h1 class="text-2xl font-bold">Paylaşılan Bloglar</h1>
    <button
      class="ml-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded"
      on:click={() => (showModal = true)}
    >
      Blog Ekle
    </button>
  </div>

  <BlogList perm={users.perm} />
</main>
