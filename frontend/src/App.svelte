<script>
// @ts-nocheck

  import { Router, Route } from "svelte-routing";
  import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
  // @ts-ignore
  import Dashboard from "./pages/Dashboard.svelte";
  import Cookies from "js-cookie";
  import axios from "axios";
  import Navbar from "./components/Navbar.svelte"
  import Blogs from "./pages/Blogs.svelte";
  import A4 from "./pages/404.svelte";
  import BlogList from "./components/Blog-List.svelte";
  let Users;
  let user = Cookies.get("token");
  let pathname = "";

  async function GetData() {
    pathname = window.location.pathname;
    if (user) {
      try {
        const response = await axios.post(
          `http://localhost:3000/api/v1/user/`, {
            token: user
          }
        );

        if (response.data.status === "OK") {
          Users = response.data.data;
          console.log(response.data.data.token);
        } else {
          console.error("Kullanıcı hatası ?");
        }
      } catch (err) {
        console.error(err);
      }
    } else {

    }
  }
  GetData();

  export let url = "";
</script>

<main>
  <Router {url}>
   <Navbar />
    <div>
      {#if Users}
      <Route path="/dashboard"><Dashboard users={Users} /></Route>
      {/if}
      <Route path="/"><Home /></Route>
      <Route path="/login"><Login /></Route>
      <Route path="/register"><Register /></Route>
      <Route path="/blogs/:title" let:params>
        <Blogs title={params.title}/>
      </Route>
      <Route path="/404"><A4 /></Route>
    </div>
  </Router>

     {#if pathname =!  "/404"}
     <div> 
      <h1 class="mt-5 ml-10 text-5xl"> Bloglar Tablosu</h1>
      <BlogList perm="user" />
     </div>
     {/if}

</main>