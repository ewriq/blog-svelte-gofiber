<script>
    // @ts-ignore
    import Cookies from "js-cookie";
    import axios from "axios";
  
    let email = "";
    let password = "";
    let error = "";
  
    async function login() {
      const form = {
        email: email,
        password: password,
      };
  
      try {
        const response = await axios.post(
          "http://localhost:3000/api/v1/login",
          form
        );
        if (response.data.status === "OK") {
          const data = response.data;
          const token = data.token;
          Cookies.set("token", token);
          // @ts-ignore
          error = "Başarılı işlem birazdan yönlendiriliyorsunuz !";
          setTimeout(function () {
            window.location.href = "/dashboard";
          }, 3000);
        } else {
          error = "Email Yada Parola hatalı";
        }
      } catch (error) {
        console.error("Hata oluştu:", error);
      }
    }
  </script>
  
  <main>
    <div class="flex justify-center items-center h-screen p-4 mb-1.5 text-center">
      <form class="h-screen roundex-3xl bg-gray-800 py-5 px-5 w-80">
        <h1 class="text-lg">Giriş Yap</h1>
        {#if error}
        <div
          class="bg-green-100 text-green-700 p-4"
          role="alert"
        >
          <p class="font-bold">{error}</p>
        </div>
      {/if}
        <label for="email" class="text-blue-600 pt-4 ">E-posta:</label>
        <input
          type="email"
          bind:value={email}
          id="email"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        />
        <br />
        <label for="password" class="text-blue-600 pt-4">Şifre:</label>
        <input
          type="password"
          bind:value={password}
          id="password"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        />
        <br />
        <button
          class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
          on:click|preventDefault={login}>Giriş Yap</button>      
          <button
          class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"><a href="/register">Kayıt ol</a></button>
      </form>
    </div>
  </main>